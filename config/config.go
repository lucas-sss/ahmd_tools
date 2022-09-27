package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type AppConfig struct {
	Mysql       Mysql
	ProxyServer ProxyServer
	Accounts    map[string]Account
}

type Mysql struct {
	Username string
	Password string
	Host     string
}

type ProxyServer struct {
	Ip string
}

type Account struct {
	User string
	Pwd  string
}

var cfgPath *string = flag.String("c", "./conf.toml", "the file path of config for this service")
var AppConf *AppConfig = &AppConfig{}

func ParseConfig() *AppConfig {
	flag.Parse()
	if _, err := os.Stat(*cfgPath); err == nil || os.IsExist(err) {
		if _, err := toml.DecodeFile(*cfgPath, AppConf); err != nil {
			fmt.Printf("Parse config file error %v\n", err)
			os.Exit(1)
		}
		if AppConf.Mysql.Username == "" {
			fmt.Printf("mysql username is empty\n")
			os.Exit(1)
		}
		if AppConf.Mysql.Password == "" {
			fmt.Printf("mysql password is empty\n")
			os.Exit(1)
		}
		if AppConf.Mysql.Host == "" {
			fmt.Printf("mysql host is empty\n")
			os.Exit(1)
		}
		if AppConf.ProxyServer.Ip == "" {
			fmt.Printf("proxy server ip is empty\n")
			os.Exit(1)
		}
		return AppConf
	}
	fmt.Printf("The configuration file is no exist:%s\n", *cfgPath)
	os.Exit(1)
	return nil
}
