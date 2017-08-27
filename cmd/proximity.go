package cmd

import (
	"github.com/catherinetcai/proximity/proximity"
	"github.com/shopspring/decimal"
	"github.com/spf13/cobra"
)

var SeedDBCmd = &cobra.Command{
	Use:   "seed",
	Short: "Create and seed db",
	Run:   seedAndCreateDB,
}

var ProximityCmd = &cobra.Command{
	Use:   "proximity",
	Short: "Find closest to latitude, longitude",
}

var haversineCmd = &cobra.Command{
	Use:   "haversine",
	Short: "Find with haversine formula",
	Run:   findHaversine,
}

var optimizedCmd = &cobra.Command{
	Use:   "optimized",
	Short: "Find with optimized haversine",
	Run:   findOptimized,
}

func init() {
	RootCmd.AddCommand(SeedDBCmd)
	RootCmd.AddCommand(ProximityCmd)
	ProximityCmd.AddCommand(haversineCmd)
	ProximityCmd.AddCommand(optimizedCmd)
}

func seedAndCreateDB(cmd *cobra.Command, args []string) {
	db := proximity.New()
	db.CreateDB()
	db.CreateSchema()
	db.SeedDB()
}

func findHaversine(cmd *cobra.Command, args []string) {
	db := proximity.New()
	// These are the coordinates of Dublin
	lat := decimal.NewFromFloat(53.349805)
	lon := decimal.NewFromFloat(-6.260310)
	db.FindClosestHaversine(lat, lon)
}

func findOptimized(cmd *cobra.Command, args []string) {
	db := proximity.New()
	// These are the coordinates of Dublin
	lat := decimal.NewFromFloat(53.349805)
	lon := decimal.NewFromFloat(-6.260310)
	db.FindClosestOptimized(lat, lon)
}
