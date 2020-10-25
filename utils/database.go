package utils

import (
	"fmt"
	"strconv"
	"sync"

	"autodesk-go/config"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
)

/**
 * Making sure database instance is created only once by using "sync.Once" package
 * remember the instance in this package and return the pointer reference
 */

 type DB struct {
	Db *gorm.DB
}

var onceDb sync.Once
var instanceDB *DB
var conf *config.Config

/**
 * get the DB instance as a pointer reference from original one in this file
 */
 func GetInstanceDB() *gorm.DB {
	onceDb.Do(func() {
		conf, _ = config.GetConfig()
		// user:password@(localhost)/dbname?charset=utf8&parseTime=True&loc=Local
		connUrl := conf.Database.User + ":" + conf.Database.Password + "@(" + conf.Database.Host + ":" + strconv.FormatInt(conf.Database.Port, 10) + ")/" + conf.Database.Name + "?charset=utf8&parseTime=True&loc=Local"
		dbConnection, _ := gorm.Open("mysql", connUrl)
		dbConnection.LogMode(true)
		instanceDB = &DB{Db: dbConnection}
		fmt.Println("Mysql")
		fmt.Println("Connected to local")
	})

	fmt.Println(instanceDB.Db)
	return instanceDB.Db
}