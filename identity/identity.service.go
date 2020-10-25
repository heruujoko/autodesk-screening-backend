package identity

import (
	"fmt"
	"autodesk-go/utils"
	"sync"
	"github.com/jinzhu/gorm"
)

var once sync.Once
var DB *gorm.DB

func init() {
	once.Do(func() {
		DB = utils.GetInstanceDB()
		fmt.Println("init identity service")
		DB.AutoMigrate(&Identity{})
	})
}