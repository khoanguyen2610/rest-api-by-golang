package connections

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type MysqlConfig struct {
	Host   		string
	Port   		int
	DbName 		string
	User   		string
	Pass   		string
	LogEnable 	bool
}

func (mC MysqlConfig) GetGormMysqlUrl() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local",
		mC.User,
		mC.Pass,
		mC.Host,
		mC.Port,
		mC.DbName,
	)
}

func ConnectDB(mC MysqlConfig) *gorm.DB {
	db, err := gorm.Open("mysql", mC.GetGormMysqlUrl())
	if err != nil {
		panic(err)
	}
	db.LogMode(mC.LogEnable)
	return db
}
