package application

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type Application struct {
	Server struct {
		Host string
		Port string
	}
	Mongodb struct {
		Host       string
		Port       string
		Username   string
		Password   string
		Database   string
		Collection string
	}
}

var App = &Application{}

// init 初始化配置文件
func init() {
	// os.Getenv() 获取环境变量, ioutil.ReadFile() 读取对应的 yaml 配置文件
	file, err := ioutil.ReadFile("./application-" + os.Getenv("ENVCONFIG") + ".yaml")
	if err != nil {
		log.Println(err)
	}

	// yaml.Unmarshal() 将 .yaml 配置文件中的配置解析到 Application 类型的变量中
	err = yaml.Unmarshal(file, App)
	if err != nil {
		log.Println(err)
	}
}
