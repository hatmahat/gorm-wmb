## livecode-gorm

## Add your files

<p align="center">
    Entity Relationship Diagram for WMB
</p> 
<p align="center" width="100%">
    <img width="100%" src="images/Gorm WMB.png">
</p>

# run with flag
```
go run app.go -host=localhost -port=5432 -username=postgres -password= -dbname=wmbgorm -env=migration
```

# main function
```
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
```
nomor 1 sampai 7    : sample.Nomor01sampai07(db)
nomor 8             : sample.Nomor08(db)
nomor 9             : sample.Nomor09(db)
nomor 10            : sample.Nomor10(db)

untuk nomor 10 menggunakan db dari nomor 9