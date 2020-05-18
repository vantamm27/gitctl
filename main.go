package main

import (
	"fmt"
	"gitctl/models"
	_ "gitctl/routers"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	fmt.Println("main init")

	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "build.db")
	orm.Debug = true
	name := "default"
	force := true
	verbose := true

	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err.Error)
	}

	// orm.RegisterModel(new(models.Service))
	// orm.RegisterModel(new(models.Environment))
	// orm.RegisterModel(new(models.Build))

	serkakacam := models.Service{Name: "kakacam", Createdate: time.Now().Local().Format("2006-01-02T15:04:05")}
	serkakacamser := models.Service{Name: "kakacam-service", Createdate: time.Now().Local().Format("2006-01-02T15:04:05")}
	o := orm.NewOrm()
	o.Insert(&serkakacam)
	o.Insert(&serkakacamser)

}

func main() {
	o := orm.NewOrm()
	o.Using("default") // Using default, you can use other database

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
