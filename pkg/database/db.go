package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
 )

 const Dsn = `postgres://vijay:zmxmcmvbn@localhost/databasevj`

 var Db *gorm.DB
 func InitDB()*gorm.DB{
	 Db = connectDB()
	 return Db
 }

 func connectDB()(*gorm.DB) {
	 var err error
	 db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database!")
	}
	fmt.Printf("\n Database type = '%s'\nConnected to database Successfully !", db.Name())
	return db
}