# api

need control over population per city.

* add families
* add names
* add kingdoms
* create kingdom (creates municipals, respective plots, resources, and npc's)
* view kingdom
* view world

## profile

```
{
  "name": "qin",
  "population": 235235,
  "land": 63462,
  "wealth": 430040,
  "trees": 36082,
  "rocks": 58230,
  "cows": 25239,
}
```

## kingdom

plots are the lowest granular level of land within a municipal.  the sum of all plot areas is equal to the total area of the municipal.  plots are randomly generated.

endpoint | params | auth | notes
---|---|---|---
POST /api/kingdoms | profile | admin | add kingdom
GET /api/kingdoms | | | get all kingdoms
POST /api/kingdoms/:id/municipals | name, population, land, wealth, trees, rocks, cows | admin | create municipal
GET /api/kingdoms/:id/municipals/:id | | admin | get municipal info
GET /api/kingdoms/:id/municipals | name, land, wealth, trees, rocks, cows | admin | get municipals
GET /api/kingdoms/:id/municipals/:id/plots | name, land, wealth, trees, rocks, cows | admin | get plots