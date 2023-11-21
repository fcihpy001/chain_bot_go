package service

import (
	"fmt"
	"github.com/anypay/scanner/model"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var (
	config *model.Config
)

func getConfig() *model.Config {
	if config == nil {
		yamlFile, err := os.ReadFile("./config.yaml")
		if err != nil {
			log.Println(err)
		}
		err = yaml.Unmarshal(yamlFile, &config)
		if err != nil {
			fmt.Println("Error unmarshaling YAML:", err)
			return nil
		}
		log.Println("配置文件读取成功。。。")

		//	加载环境变量
		if err := godotenv.Load(); err != nil {
			log.Fatalf("无法加载 .env 文件: %v", err)
		}
		log.Println(".env 文件加载成功。。。")
	}
	return config
}

func LoadEnvFile() {
	//	加载环境变量
	if err := godotenv.Load(); err != nil {
		log.Fatalf("无法加载 .env 文件: %v", err)
	}
	log.Println(".env 文件加载成功。。。")
}
