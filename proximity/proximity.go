package proximity

import (
	"fmt"
	"math"

	"github.com/davecgh/go-spew/spew"
	"github.com/golang/geo/s1"
	"github.com/shopspring/decimal"
)

type Location struct {
	Name      string          `db:"name",sql:"text"`
	Latitude  decimal.Decimal `db:"latitude",sql:"decimal(16,13)"`
	Longitude decimal.Decimal `db:"longitude",sql:"decimal(16,13)"`
	Distance  decimal.Decimal `db:"distance",sql:"decimal(16,13)"`
}

func (d *DB) FindClosestHaversine(lat, lon decimal.Decimal) {
	distance := 10
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

func (d *DB) FindClosestOptimized(lat, lon decimal.Decimal) {
	distance := 10
	// 1 degree of latitude is ~69 miles
	// 1 degree of longitude ~ cos(latitude)*69
	lon1 := lon - distance/math.Abs(math.Cos(s1.Angle(lat).Radians()*69))
	lon2 := lon + distance/math.Abs(math.Cos(s1.Angle(lat).Radians()*69))
	lat1 := lat - (distance / 69)
	lat2 := lat + (distance / 69)
	query := fmt.Sprintf(`
		SELECT *, 3956 * 2 * ASIN(SQRT(
			POWER(SIN((%v - locations.latitude) * pi()/180/2), 2) +
			COS(%v * pi()/180) * COS(locations.latitude * pi()/180) *
			POWER(SIN((%v - locations.longitude) * pi()/180/2), 2))) as distance
		FROM locations
		WHERE locations.longitude 
		between %v and %v
		and locations.latitude
		between %v and %v
		ORDER BY distance limit 20;
	`, lat, lat, lon, lon1, lon2, lat1, lat2)
}
