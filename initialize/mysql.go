package initialize

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go-demo/global"
	"go-demo/model"
	"os"
)

func Mysql() {
	config := global.GVA_CONFIG.Mysql
	link := config.Username+":"+config.Password+"@("+config.Path+")/"+config.Dbname+"?"+config.Config
	if db, err := gorm.Open("mysql", link); err != nil {
		fmt.Println("mysql connect failed", err.Error())
		os.Exit(0)
	} else {
		global.GVA_DB = db
		global.GVA_DB.DB().SetMaxIdleConns(config.MaxIdleConns)
		global.GVA_DB.DB().SetMaxOpenConns(config.MaxOpenConns)
		global.GVA_DB.SingularTable(true)
	}
}

func DBTables() {
	global.GVA_DB.AutoMigrate(model.User{})
}
