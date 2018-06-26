package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/frictionlessdata/datapackage-go/datapackage"
	"github.com/frictionlessdata/tableschema-go/csv"
	"googlemaps.github.io/maps"
)

var geoHeaders = []string{
	"CD_PROCESSO_ELEITORAL",
	"CD_PLEITO",
	"SG_UF",
	"CD_LOCALIDADE_TSE",
	"NM_LOCALIDADE",
	"NR_ZONA",
	"NR_SECAO",
	"ST_SECAO_USA_LOCVOT_TEMP",
	"NR_LOCVOT",
	"NM_LOCVOT",
	"DS_ENDERECO",
	"DS_BAIRRO",
	"NR_CEP",
	"LATLNG",
}

const (
	PlaceIndex  = 9
	LatLngIndex = 13
	ZipIndex    = 12
	StateIndex  = 2
	CityIndex   = 4
)

func main() {
	if os.Getenv("GMAPS_API_KEY") == "" {
		log.Fatalf("GMAPS_API_KEY env var not set.")
	}
	pkgFName := os.Args[1]
	if pkgFName == "" {
		log.Fatalf("Usage: ./geocoder [package path]")
	}

	pkg, err := datapackage.Load(pkgFName)
	if err != nil {
		log.Fatalf("Error loading data package (%s): %q", pkgFName, err)
	}

	// 1. Read geocoded file and create a map from where we should proceed.
	geocodeMap := make(map[string]string)
	res2016Geo := pkg.GetResource("2016_geocoded")
	if res2016Geo == nil {
		log.Fatalf("Error fetching resource 2016_geocoded")
	}
	items2016Geo, err := res2016Geo.ReadAll(csv.SetHeaders(geoHeaders...))
	if err != nil {
		log.Fatalf("Error reading resource contents: %q", err)
	}
	for _, item := range items2016Geo {
		// Maybe someone improved the algorithm or the place has been picked up by google maps
		// and now it can be resolved.
		if item[LatLngIndex] != "" {
			geocodeMap[getGeoMapIndex(item)] = item[LatLngIndex]
		}
	}
	fmt.Printf("Reused %d coordinates.\n", len(geocodeMap))

	// 2. Sweep the old table and try to find more to geocode.
	// Must update the output file while sweeping.
	fName := "../../2016_geocoded.csv"
	f, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Error trying to open file (%s):%q", fName, err)
	}
	defer f.Close()
	res2016GeoWriter := csv.NewWriter(f)
	res2016GeoWriter.Comma = ';'

	client, err := maps.NewClient(maps.WithAPIKey(os.Getenv("GMAPS_API_KEY")))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	res2016 := pkg.GetResource("2016")
	items2016, err := res2016.ReadAll(csv.LoadHeaders())
	if err != nil {
		log.Fatalf("Error reading resource contents: %q", err)
	}
	apiCalls := 0
	for count, item := range items2016 {
		if _, ok := geocodeMap[getGeoMapIndex(item)]; !ok {
			addr := fmt.Sprintf("%s,%s,%s,%s", item[PlaceIndex], item[ZipIndex], item[CityIndex], item[StateIndex])
			// Special case: When the address is 0.
			if item[ZipIndex] == "0" {
				addr = fmt.Sprintf("%s,%s,%s", item[PlaceIndex], item[CityIndex], item[StateIndex])
			}
			req := &maps.GeocodingRequest{Address: addr}
			resp, err := client.Geocode(
				context.Background(),
				req,
			)
			apiCalls += 1
			if err != nil {
				log.Fatalf("Error geocoding item[%s]: %q", spew.Sdump(item), err)
			}
			var r *maps.GeocodingResult
			if len(resp) == 1 {
				r = &resp[0]
			} else {
				// When there is ambiguity, must explicity fall back.
				for _, v := range resp {
					for _, t := range v.Types {
						// TODO: rank the options based on type (from more to less specific).
						// Deliberate choose not to fall into city or country types. Better to not have it at all.
						if t == "postal_code" || t == "school" || t == "administrative_area_level_2" || t == "local_government_office" || t == "locality" || t == "point_of_interest" || t == "route" {
							r = &v
							break
						}
					}
					if r != nil {
						break
					}
				}
			}
			if r == nil {
				log.Printf("Could not determine the correct geo coordinate to use. Item:%s\nReq:%s\nResp:%s\n", spew.Sdump(item), spew.Sdump(req), spew.Sdump(resp))
				geocodeMap[getGeoMapIndex(item)] = ""
			} else {
				geocodeMap[getGeoMapIndex(item)] = fmt.Sprintf("%s,%s",
					strconv.FormatFloat(r.Geometry.Location.Lat, 'f', -1, 64),
					strconv.FormatFloat(r.Geometry.Location.Lng, 'f', -1, 64))
			}
			fmt.Printf("Geolocation API called. Record: %+v\n", item)
			time.Sleep(time.Second / 100)
		}
		if err := res2016GeoWriter.Write(append(item, geocodeMap[getGeoMapIndex(item)])); err != nil {
			log.Fatalf("Error writing record [%v]: %q", item, err)
		}
		res2016GeoWriter.Flush()
		if (count+1)%10 == 0 {
			fmt.Printf("Processed %d records, Made %d GMAPS API calls\n", count+1, apiCalls)
		}
	}
}

func getGeoMapIndex(item []string) string {
	// Places could have the same name: differentiating based on the postal code would be enough, but
	// there are cases where the postal code is 0.
	return item[PlaceIndex] + item[CityIndex] + item[ZipIndex]
}
