package config

import (
	"flag"
	"fmt"
	"livecode3/model"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type dbConfig struct {
	dbHost     string
	dbPort     string
	dbUser     string
	dbPassword string
	dbName     string
}

type Config struct {
	DB *gorm.DB
}

func (c *Config) initDB() {
	dbCfg := new(dbConfig)
	var env string

	flag.StringVar(&dbCfg.dbHost, "host", "", "database host")
	flag.StringVar(&dbCfg.dbPort, "port", "", "database port")
	flag.StringVar(&dbCfg.dbUser, "username", "", "database username")
	flag.StringVar(&dbCfg.dbPassword, "password", "", "database password")
	flag.StringVar(&dbCfg.dbName, "dbname", "", "database name")
	flag.StringVar(&env, "env", "", "environment")
	flag.Parse()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbCfg.dbHost, dbCfg.dbUser, dbCfg.dbPassword, dbCfg.dbName, dbCfg.dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err == nil {
		log.Println("connected")
	} else {
		panic(err)
	}

	if env == "dev" {
		c.DB = db.Debug()
	} else if env == "migration" {
		c.DB = db.Debug()
		err := c.DB.AutoMigrate(
			&model.BillDetail{},
			&model.Bill{},
			&model.Customer{},
			&model.Discount{},
			&model.MenuPrice{},
			&model.Menu{},
			&model.Table{},
			&model.TransType{},
		)
		if err != nil {
			return
		}
	} else {
		c.DB = db
	}
}

func (c *Config) DbConn() *gorm.DB {
	return c.DB
}

func (c *Config) DbClose() {
	enigmaDb, err := c.DB.DB()

	if err != nil {
		panic(err)
	}
	err = enigmaDb.Close()
	if err != nil {
		panic(err)
	}
}

func NewConfigDB() Config {
	cfg := Config{}
	cfg.initDB()
	return cfg
}
