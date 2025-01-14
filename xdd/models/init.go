package models

import (
	"os"
	"path/filepath"

	"github.com/beego/beego/v2/core/logs"
)

func init() {
	killp()
	for _, arg := range os.Args {
		if arg == "-d" {
			Daemon()
		}
	}
	ExecPath, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	logs.Info("当前%s", ExecPath)
	initConfig()
	go initVersion()
	go initUserAgent()
	go initContainer()
	initDB()
	initHandle()
	initCron()
	go initTgBot()
	InitReplies()
}
