package connections

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type MysqlConfig struct {
	Host   string `default:"localhost"`
	Port   int    `default:"3306"`
	DbName string `default:"sample_go"`
	User   string `default:"root"`
	Pass   string `default:"123456"`
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
	db.LogMode(true)
	return db
}
