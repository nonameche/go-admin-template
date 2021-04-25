package main

import (
	"bin/models"
	"bin/pages"
	"bin/tables"
	"io/ioutil"
	"log"
	"os"
	"os/signal"

	_ "github.com/GoAdminGroup/go-admin/adapter/gin" // 引入适配器
	"github.com/GoAdminGroup/go-admin/engine"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql" // 引入对应数据库引擎
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	_ "github.com/GoAdminGroup/themes/sword" // 引入主题
	"github.com/gin-gonic/gin"
)

func main() {
	startServer()
}

func startServer() {
	/*
	**初始化配置
	 */
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	r := gin.Default() //创建带有默认中间件的路由:日志与恢复中间件

	template.AddComp(chartjs.NewChart())

	eng := engine.Default() // 实例化一个GoAdmin引擎对象

	// 增加配置与插件，使用Use方法挂载到Web框架中
	if err := eng.AddConfigFromJSON("./config.json").
		// 这里引入你需要管理的业务表配置
		// 后面会介绍如何使用命令行根据你自己的业务表生成Generators
		AddGenerators(tables.Generators).
		Use(r); err != nil {
		panic(err)
	}

	r.Static("/uploads", "./uploads")

	eng.HTML("GET", "/admin", pages.GetDashBoard)
	eng.HTMLFile("GET", "/admin/hello", "./html/hello.tmpl", map[string]interface{}{
		"msg": "Hello world",
	})

	models.Init(eng.MysqlConnection())

	_ = r.Run(":80")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Print("closing database connection")
	eng.MysqlConnection().Close()

}
