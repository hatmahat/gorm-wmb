package main

import (
	"livecode3/config"
	"livecode3/sample"
)

func main() {
	cfg := config.NewConfigDB()
	db := cfg.DbConn()
	defer cfg.DbClose()

	sample.Nomor01sampai07(db)

	// sample.Nomor08(db)

	// sample.Nomor09(db)

	// nomor 10 menggunakan database dari nomor 09
	// sample.Nomor10(db)

}

// go run app.go -host=localhost -port=5432 -username=postgres -password= -dbname=wmbgorm -env=migration
