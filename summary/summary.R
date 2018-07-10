if (!require("devtools")) {
  install.packages("devtools")
  library("devtools")
}
if (!require("jsonlite")) {
  install.packages("jsonlite", repos="https://cran.rstudio.com/")
  library("jsonlite")
}
if (!require("skimr")) {
  devtools::install_github("ropenscilabs/skimr")
  library(skimr)
}

json_file <- 'datapackage.json'
json_data <- fromJSON(paste(readLines(json_file), collapse=""))

sink('summary/SUMMARY.md')

for(i in 1:length(json_data$resources$mediatype)){
  file = json_data$resources$path[i]
  sep = json_data$resources$dialect$delimiter[i]
  header = json_data$resources$schema$fields[i][[1]][['name']]
  data <- read.csv(file=file, header=FALSE, sep=sep)
  names(data) <- header

  title = paste('##', file)
  summary = skim(data)
  space = ''
  code = '```'

  writeLines(title)
  writeLines(space)
  writeLines(code)
  print(summary)
  writeLines(code)
  writeLines(space)
}

sink()
