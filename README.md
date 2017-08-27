## Proximity

Playing around with latitude/longitude proximity searches.

### Getting Going
Start TiDB
```
make db
```

Seed DB (it'll show you some errors you can ignore, probably)
```
make db-seed
```

Test Haversine - runs this from the lat/lon of Dublin against the DB
```
go run main.go proximity
```
