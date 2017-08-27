package proximity

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/shopspring/decimal"
)

type Location struct {
	Name      string          `db:"name",sql:"text"`
	Latitude  decimal.Decimal `db:"latitude",sql:"decimal(16,13)"`
	Longitude decimal.Decimal `db:"longitude",sql:"decimal(16,13)"`
	Distance  decimal.Decimal `db:"distance",sql:"decimal(16,13)"`
}

func (d *DB) FindClosestHaversine(lat, lon decimal.Decimal) {
	distance := 500
	query := fmt.Sprintf(`
		SELECT *, 3956 * 2 * ASIN(SQRT(
			POWER(SIN((%v - abs(locations.latitude)) * pi()/180/2), 2) +
			COS(%v * pi()/180) * COS(abs(locations.latitude) *
			pi()/180) * POWER(SIN((%v - locations.longitude) *
			pi()/180/2), 2))) as distance
		FROM locations 
		having distance < %v
		ORDER BY distance limit 20;
	`, lat, lat, lon, distance)
	spew.Dump(query)
	rows, err := d.Queryx(query)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var l Location
		err = rows.StructScan(&l)
		spew.Dump(l)
	}
}
