## 2016.csv

```
Skim summary statistics
 n obs: 432961 
 n variables: 13 

── Variable type:factor ────────────────────────────────────────────────────────
                 variable missing complete      n n_unique
        CD_LOCALIDADE_TSE       0   432961 432961     5569
                CD_PLEITO       0   432961 432961        2
    CD_PROCESSO_ELEITORAL       0   432961 432961        2
                DS_BAIRRO       0   432961 432961    30471
              DS_ENDERECO       0   432961 432961    83434
            NM_LOCALIDADE       0   432961 432961     5297
                NM_LOCVOT       0   432961 432961    85016
                   NR_CEP       0   432961 432961    26406
                NR_LOCVOT       0   432961 432961      287
                 NR_SECAO       0   432961 432961     1035
                  NR_ZONA       0   432961 432961      426
                    SG_UF       0   432961 432961       27
 ST_SECAO_USA_LOCVOT_TEMP       0   432961 432961        3
                                     top_counts ordered
   710: 25054, 600: 11609, 384: 4704, 138: 4492   FALSE
                     187: 432960, CD_: 1, NA: 0   FALSE
                     185: 432960, CD_: 1, NA: 0   FALSE
   CEN: 99231, ZON: 15203, CEN: 3366, INT: 1578   FALSE
        EST: 1118, RUA: 640, RUA: 277, RUA: 211   FALSE
   SÃO: 25054, RIO: 11609, SAL: 4704, FOR: 4492   FALSE
         PRE: 296, ESC: 189, ESC: 187, SEN: 172   FALSE
         0: 5160, 650: 1536, 689: 794, 580: 788   FALSE
 101: 31546, 102: 25584, 103: 21760, 104: 19266   FALSE
             1: 2583, 3: 2548, 5: 2530, 2: 2528   FALSE
             2: 6452, 1: 6001, 3: 4665, 4: 4574   FALSE
     SP: 92016, MG: 46773, RJ: 33415, BA: 31700   FALSE
              N: 431313, S: 1647, ST_: 1, NA: 0   FALSE
```

## 2016_geocoded.csv

```
Skim summary statistics
 n obs: 50052 
 n variables: 15 

── Variable type:factor ────────────────────────────────────────────────────────
                 variable missing complete     n n_unique
                DS_BAIRRO       0    50052 50052     5689
              DS_ENDERECO       0    50052 50052    12221
            NM_LOCALIDADE       0    50052 50052      648
                NM_LOCVOT       0    50052 50052    12927
                    SG_UF       0    50052 50052        6
 ST_SECAO_USA_LOCVOT_TEMP       0    50052 50052        2
                                                                 top_counts
  CEN: 14006, ZON: 1831, CID: 377, TAN: 220                                
                                         AV.: 66, RUA: 64, AVE: 61, AV : 51
 SAL: 4704, MAN: 3322, MAC: 1530, FEI: 1073                                
         COL: 90, COL: 74, UNI: 60, COL: 56                                
    BA: 31700, AM: 6693, AL: 6287, CE: 2037                                
                    N: 49690, S: 362, NA: 0                                
 ordered
   FALSE
   FALSE
   FALSE
   FALSE
   FALSE
   FALSE

── Variable type:integer ───────────────────────────────────────────────────────
              variable missing complete     n        mean          sd   p0
     CD_LOCALIDADE_TSE       0    50052 50052 28700.24    15011.74    1007
             CD_PLEITO       0    50052 50052   187           0        187
 CD_PROCESSO_ELEITORAL       0    50052 50052   185           0        185
                NR_CEP       0    50052 50052     5.2e+07     1.1e+07    0
             NR_LOCVOT       0    50052 50052  1294.15      310.94    1015
              NR_SECAO       0    50052 50052   171.98      176.6        1
               NR_ZONA       0    50052 50052    69.64       59.58       1
         p25         p50         p75  p100     hist
 27235       34673       38075       98515 ▃▁▇▅▁▁▁▁
   187         187         187         187 ▁▁▁▇▁▁▁▁
   185         185         185         185 ▁▁▁▇▁▁▁▁
     4.5e+07     4.8e+07     5.8e+07 7e+07 ▁▁▁▁▂▇▂▃
  1074        1171        1414        3581 ▇▂▁▁▁▁▁▁
    51         114         225        1034 ▇▃▂▁▁▁▁▁
    18          53         113         205 ▇▅▃▂▂▂▂▂

── Variable type:numeric ───────────────────────────────────────────────────────
 variable missing complete     n   mean    sd      p0    p25    p50    p75
      LAT     269    49783 50052 -10.24  5.61  -42.47 -12.97 -11.66  -9.37
      LNG     269    49783 50052 -43.77 10.17 -122.71 -43.31 -39.57 -38.47
   p100     hist
  56.55 ▁▁▇▂▁▁▁▁
 138.6  ▁▂▇▁▁▁▁▁
```

