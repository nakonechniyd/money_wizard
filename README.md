# Money Wizard

Is a toy project that demonstrate how hard to get wealth accumulation

![Lint Go Code](https://github.com/nakonechniyd/money_wizard/workflows/Lint%20Go%20Code/badge.svg)

![Tests Across Matrix](https://github.com/nakonechniyd/money_wizard/workflows/Test%20Across%20Matrix/badge.svg)

![Build and Publish Docker Image](https://github.com/nakonechniyd/money_wizard/workflows/Build%20and%20Publish%20Docker%20Image/badge.svg)

## build
```
make build
```
or
```
o build -o money_wizard main.go
```

## run
```
./money_wizard
```
it'll start API server on localhost:8080

## calculation by swagger ui
open in browser
```
http://localhost:8080/swaggerui/
```

## calculation

```
make curl
```
or
```
curl "http://localhost:8080/api/v1/calc_years/12?tax_percent=19.5&passive_percent=14.0&capital=10000&monthly_replenishment=5000&retire_year=10&life_percent=45.0"
```
where:
- **years** - Ammount of calculation years (*Required*) = calculate for 12 years
- **tax_percent** - Total percent of all taxes (*Required*) = shared percent of all taxes 19.5%
- **passive_percent** - Passive percent (*Required*) = 14.0% of interest for example deposit
- **capital** - Ammount of start capital (*Required* if no monthly_replenishment, Default value : 0) = start accumulation with 10000 of capital
- **monthly_replenishment** - Monthly replenishment (*Required* if no capital, Default value : 0) = put 5000 every month to the account
- **retire_year** - Retire year = get retire after 10 years
- **life_percent** - Life percent from net passive income = take for life 45.0% from passive monthly income

as result you will get json like this:
```
[
  {
    "year": 1,
    "capital_before": 10000,
    "total_monthly_replenishment": 60000,
    "total_passive_income": 5447.7402421265715,
    "total_net_passive_income": 4385.430894911891,
    "capital_after": 74385.4308949119,
    "months": [
      {
        "year": 1,
        "month": 1,
        "capital_before": 10000,
        "monthly_replenishment": 5000,
        "passive_income": 116.66666666666669,
        "net_passive_income": 93.91666666666669,
        "capital_after": 15093.916666666666
      },
      {
        "year": 1,
        "month": 2,
        "capital_before": 15093.916666666666,
        "monthly_replenishment": 5000,
        "passive_income": 176.09569444444446,
        "net_passive_income": 141.7570340277778,
        "capital_after": 20235.673700694446
      },
      {
        "year": 1,
        "month": 3,
        "capital_before": 20235.673700694446,
        "monthly_replenishment": 5000,
        "passive_income": 236.08285984143524,
        "net_passive_income": 190.04670217235537,
        "capital_after": 25425.720402866802
      },
      {
        "year": 1,
        "month": 4,
        "capital_before": 25425.720402866802,
        "monthly_replenishment": 5000,
        "passive_income": 296.6334047001127,
        "net_passive_income": 238.78989078359072,
        "capital_after": 30664.510293650394
      },
      {
        "year": 1,
        "month": 5,
        "capital_before": 30664.510293650394,
        "monthly_replenishment": 5000,
        "passive_income": 357.75262009258796,
        "net_passive_income": 287.9908591745333,
        "capital_after": 35952.50115282493
      },
      {
        "year": 1,
        "month": 6,
        "capital_before": 35952.50115282493,
        "monthly_replenishment": 5000,
        "passive_income": 419.4458467829576,
        "net_passive_income": 337.6539066602808,
        "capital_after": 41290.15505948521
      },
      {
        "year": 1,
        "month": 7,
        "capital_before": 41290.15505948521,
        "monthly_replenishment": 5000,
        "passive_income": 481.71847569399415,
        "net_passive_income": 387.7833729336653,
        "capital_after": 46677.93843241887
      },
      {
        "year": 1,
        "month": 8,
        "capital_before": 46677.93843241887,
        "monthly_replenishment": 5000,
        "passive_income": 544.5759483782202,
        "net_passive_income": 438.3836384444673,
        "capital_after": 52116.32207086334
      },
      {
        "year": 1,
        "month": 9,
        "capital_before": 52116.32207086334,
        "monthly_replenishment": 5000,
        "passive_income": 608.0237574934057,
        "net_passive_income": 489.45912478219157,
        "capital_after": 57605.781195645526
      },
      {
        "year": 1,
        "month": 10,
        "capital_before": 57605.781195645526,
        "monthly_replenishment": 5000,
        "passive_income": 672.0674472825312,
        "net_passive_income": 541.0142950624377,
        "capital_after": 63146.79549070796
      },
      {
        "year": 1,
        "month": 11,
        "capital_before": 63146.79549070796,
        "monthly_replenishment": 5000,
        "passive_income": 736.7126140582595,
        "net_passive_income": 593.0536543168989,
        "capital_after": 68739.84914502487
      },
      {
        "year": 1,
        "month": 12,
        "capital_before": 68739.84914502487,
        "monthly_replenishment": 5000,
        "passive_income": 801.9649066919568,
        "net_passive_income": 645.5817498870252,
        "capital_after": 74385.4308949119
      }
    ]
  },
  ....
]
```
