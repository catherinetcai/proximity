package proximity

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	*sqlx.DB
}

// New DB
func New() *DB {
	cfg := &mysql.Config{
		User: "root",
		//Passwd: "test",
		Net:    "tcp",
		Addr:   "127.0.0.1:4000",
		DBName: "proximity",
	}
	spew.Dump(cfg.FormatDSN())
	db := sqlx.MustConnect("mysql", cfg.FormatDSN())
	return &DB{db}
}

func (d *DB) CreateDB() {
	query := "CREATE DATABASE proximity;"
	_, err := d.Exec(query)
	spew.Dump(err)
}

func (d *DB) CreateSchema() {
	schema := `
		CREATE TABLE locations (
			name text,
			latitude DECIMAL(16, 13) NOT NULL,
			longitude DECIMAL(16, 13) NOT NULL
		);
	`
	_, err := d.Exec(schema)
	spew.Dump(err)
}
func (d *DB) SeedDB() {
	query := `
		INSERT INTO locations 
			(name, latitude, longitude)
		VALUES
			('google', 37.419200897217, -122.05740356445),
			('facebook', 53.347198486328, -6.2438998222351),
			('apple', 37.30419921875, -122.09459686279),
			('reddit', 37.76969909668, -122.39330291748),
			('hotmail', 53.338901519775, -6.2595000267029),
			('baidu', 34.772499084473, 113.72660064697),
			('timewarner', 38.650001525879, -90.533401489258);
	`
	_, err := d.Exec(query)
	spew.Dump(err)
}
