package config

import (
	"encoding/json"
	"os"
)

//Config ...
type Config struct {
	Mysql Mysql `json:"mysql"`
	Redis Redis `json:"redis"`
}

//Mysql ...
type Mysql struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Host         string `json:"host"`
	Port         string `json:"port"`
	Dbname       string `json:"dbname"`
	Config       string `json:"config"`
	MaxIdleConns int    `json:"maxIdleConns"`
	MaxOpenConns int    `json:"maxOpenConns"`
}

//Redis ...
type Redis struct {
	NetWork  string `json:"net_work"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Password string `json:"password"`
	Prefix   string `json:"prefix"`
}

//InitConfig ...
func InitConfig() *Config {
	var config *Config
	file, err := os.Open("config.json")
	if err != nil {
		panic("配置文件读取异常!")
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		panic(err.Error())
	}
	return config
}
