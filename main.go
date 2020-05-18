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
	beego.SetLevel(beego.AppConfig.DefaultInt("loglevel", beego.LevelInformational))
	beego.SetLogger("file", `{"filename":"logs/`+beego.AppConfig.DefaultString("appname", "app")+`.log"}`)
	beego.SetLogFuncCall(true)
	initDatabase()
}
func initDatabase() {
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "build.db")
	orm.Debug = true
	name := "default"
	force := true
	verbose := true

	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err.Error)
		return
	}

	// orm.RegisterModel(new(models.Service))
	// orm.RegisterModel(new(models.Environment))
	// orm.RegisterModel(new(models.Build))

	env_master := models.Environment{Name: "stg", Createdate: time.Now().Local().Format("2006-01-02T15:04:05")}
	serkakacam := models.Service{Name: "kakacam", Createdate: time.Now().Local().Format("2006-01-02T15:04:05")}

	o := orm.NewOrm()
	o.Insert(&env_master)
	o.Insert(&serkakacam)
	build_kakacam := models.Build{
		Id:          1,
		EnvId:       env_master.Id,
		EnvName:     env_master.Name,
		ServiceId:   serkakacam.Id,
		ServiceName: serkakacam.Name,
		Createdate:  time.Now().Local().Format("2006-01-02T15:04:05")}
	o.Insert(&build_kakacam)
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
