// Package config 这里的结构体不保存到数据库，从配置文件中获取
package config

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
)

type tomlConfig struct {
	Viewer Viewer
	System SystemConfig
	Redis  Redis
}

type Redis struct {
	Addr     string
	Password string
	DataBase int
}

type Viewer struct {
	Title       string
	Description string
	Logo        string
	Navigation  []string
	Bilibili    string
	Zhihu       string
	Avatar      string
	UserName    string
	UserDesc    string
}

type SystemConfig struct {
	AppName         string
	Version         float32
	CurrentDir      string
	CdnURL          string
	QiniuAccessKey  string
	QiniuSecretKey  string
	Valine          bool
	ValineAppid     string
	ValineAppkey    string
	ValineServerURL string
}

var Cfg *tomlConfig

func init() {
	// 程序启动的时候
	Cfg = new(tomlConfig)
	Cfg.System.AppName = "go-blog"
	Cfg.System.Version = 0.6
	Cfg.System.CurrentDir, _ = os.Getwd()
	_, err := toml.DecodeFile("config/config.toml", &Cfg)
	if err != nil {
		log.Fatal(err)
	}
}
