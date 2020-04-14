package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/endaaman/api.endaaman.me/infras"
	_ "github.com/endaaman/api.endaaman.me/routers"
	"github.com/endaaman/api.endaaman.me/services"
)

func fillCredentials() {
	hash := beego.AppConfig.String("password_hash")
	if hash != "" {
		logs.Info("A password hash exists. No need to fill credeintails.")
		return
	}
	password := beego.AppConfig.String("password")
	if password == "" {
		panic("Password is empty.")
	}
	newHash, _ := services.GeneratePasswordHash(password)
	beego.AppConfig.Set("password_hash", newHash)
	suc := services.ValidatePassword(newHash, password)
	if !suc {
		panic("Somethins wrog?")
	}
	logs.Info("Credeintail was successfully filled.")
}

func main() {
	fillCredentials()
	// TODO: mkdir shared/articles and shared/private
	infras.PrepareDirs()
	services.ReadAllArticles()
	go infras.StartWatching()
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.DelStaticPath("/static")
	// beego.SetStaticPath("/static", "shared/")

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		// AllowOrigins:     []string{"https://*.foo.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}), true)

	beego.Run()
}
