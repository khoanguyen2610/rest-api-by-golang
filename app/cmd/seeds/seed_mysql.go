package main

import (
	"fmt"
	"sync"
	"time"
	"user-service/configs"
	"user-service/connections"
	"user-service/models"
	"user-service/repositories/mysql"

	"github.com/icrowley/fake"
	log "github.com/sirupsen/logrus"
)

var wg sync.WaitGroup
func main() {
	appConf := configs.GetAppConfigFromEnv()

	// init DB Connection
	mysqlConfig := connections.MysqlConfig{
		Host:   appConf.DB.Host,
		Port:   appConf.DB.Port,
		DbName: appConf.DB.DbName,
		User:   appConf.DB.User,
		Pass:   appConf.DB.Pass,
	}
	db := connections.ConnectDB(mysqlConfig)
	repo := mysql.NewBaseRepo(db)



	defer timeTrack(time.Now(), "bulkDatabase")
	wg.Add(1)
	go bulkDatabase(repo)
	go bulkDatabase(repo)
	go bulkDatabase(repo)
	go bulkDatabase(repo)
	go bulkDatabase(repo)
	go bulkDatabase(repo)
	wg.Wait()
}

func timeTrack(start time.Time, funcName string) {
	elapsed := time.Since(start)
	fmt.Println(funcName, "took", elapsed)
}

func bulkDatabase(repo *mysql.BaseRepo) {
	defer wg.Done()
	var email, fullname, password string



	for i:=0; i<10000; i++ {
		email = fake.EmailAddress()
		fullname = fake.FullName()
		password = fake.SimplePassword()
		user := models.User{Email: email, Fullname: fullname, Password: password}
		err := repo.Create(&user)
		if err != nil {
			fmt.Println("Tạch")
		}
		log.WithFields(log.Fields{
			"email": email,
			"fullname": fullname,
			"password": password,
		}).Info("Seeding dummy data")
	}
}