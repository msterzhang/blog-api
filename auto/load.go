package auto

import (
	"database/sql"
	"log"
	"blog-api/api/database"
	"blog-api/api/models"
	"blog-api/config"
)

func init() {
	db, err := sql.Open(config.DBDRIVER, config.DBDATAURL)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_, err = db.Exec("CREATE DATABASE " + config.DBNAME)
	if err != nil {
		log.Println("数据库已存在!")
		InitDatabase()
		return
	}
	log.Println("数据库创建成功！",err)
	InitDatabase()
}

func InitDatabase()  {
	err := database.InitDb()
	if err != nil {
		log.Fatal("Gorm初始化数据库失败！报错：" + err.Error())
	}
}

func Load() {
    var err error
	db := database.NewDb()
	
	err = db.Debug().AutoMigrate(&models.User{},).Error
	if err != nil {
		log.Fatal(err)
	}
	
	err = db.Debug().AutoMigrate(&models.Post{},).Error
	if err != nil {
		log.Fatal(err)
	}
	
	err = db.Debug().AutoMigrate(&models.Tag{},).Error
	if err != nil {
		log.Fatal(err)
	}
	
	err = db.Debug().AutoMigrate(&models.Comment{},).Error
	if err != nil {
		log.Fatal(err)
	}
	
	err = db.Debug().AutoMigrate(&models.Link{},).Error
	if err != nil {
		log.Fatal(err)
	}
	
}
