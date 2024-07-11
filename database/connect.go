package database

import (
	"carrier/config"
	"carrier/model/carrier"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strconv"
)

// ConnectDB connect to db and auto make migration
func ConnectDB() *gorm.DB {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		panic("failed to parse database port")
	}

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Config("DB_HOST"),
		port,
		config.Config("DB_USER"),
		config.Config("DB_PASSWORD"),
		config.Config("DB_NAME"),
	)
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("connection opened to database...")
	err = DB.AutoMigrate(&carrier.Carrier{})
	if err != nil {
		panic(err.Error())
		return nil
	}
	fmt.Println("database migrated", DB)
	return DB
}
