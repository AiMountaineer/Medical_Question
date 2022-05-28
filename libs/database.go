package libs

import (
	"fmt"
	"time"

	logger "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/plugin/dbresolver"

	"gorm.io/gorm"
)

var (
	Db *gorm.DB
)

func InitDb() {

	var err error
	var dialector gorm.Dialector

	conn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=True&loc=Local", Config.DB.User, Config.DB.Password, Config.DB.Host, Config.DB.Port, Config.DB.Name)

	fmt.Printf("%+v", "DBUSERです")
	fmt.Printf("%+v", Config.DB.User)
	dialector = mysql.Open(conn)

	Db, err = gorm.Open(dialector, &gorm.Config{})

	if err != nil {
		logger.Println(err)
	}

	_ = Db.Use(
		dbresolver.Register(
			dbresolver.Config{}).
			SetConnMaxIdleTime(time.Hour).
			SetConnMaxLifetime(24 * time.Hour).
			SetMaxIdleConns(100).
			SetMaxOpenConns(200),
	)
	Db.Session(&gorm.Session{FullSaveAssociations: true, AllowGlobalUpdate: false})
}
