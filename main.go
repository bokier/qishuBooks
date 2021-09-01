package main

import (
	"fmt"
	"qishuBook/common"
	"qishuBook/dao/mysql"
	"qishuBook/engine"
	"qishuBook/setting"
	"time"
)

const (
	qiShuUrl = "https://www.kankezw.com/soft/sort010/"
)

func main() {
	// 1. 初始化配置文件
	err := setting.InitViper()
	common.HandlerError(err, "[error] setting.InitViper failed.")
	fmt.Println("[info] InitViper success...")
	// 2. 初始化数据库
	err = mysql.InitMysql(setting.Conf.MysqlConfig)
	common.HandlerError(err, "[error] mysql.InitMysql failed.")
	fmt.Println("[info] InitMysql success...")

	t := time.Now()
	engine.RunEngineV1(qiShuUrl)
	// engine.GetBookInfosTest("https://www.kankezw.com/soft/sort011/index_2.html")
	elapsed := time.Since(t)
	fmt.Println(elapsed)
}
