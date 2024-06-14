package conf

import (
	"bufio"
	"encoding/json"
	"os"
)

type config struct {
	Env              string                     `json:"env"`
	Port             string                     `json:"port"`
	AesKey           string                     `json:"aes_key"`
	Mysql            Mysql                      `json:"mysql"`
	Redis            Redis                      `json:"redis"`
	Kafka            Kafka                      `json:"kafka"`
	Log              Log                        `json:"log"`
	Minio            Minio                      `json:"minio"`
	Nacos            Nacos                      `json:"nacos"`
	APPAuth          map[string]Auth            // 存放认证用户信息
	ExcludeAuth      map[string]map[string]bool // 存放不校验的URL
	LoginExcludeAuth map[string]map[string]bool // 存放不校验的URL
}

type Auth struct {
	Name   string
	Secret string
}

type Nacos struct {
	Endpoints []NacosEndpoint `json:"endpoints"`
	Username  string          `json:"username"`
	Password  string          `json:"password"`
	Namespace string          `json:"namespace"`
}
type NacosEndpoint struct {
	Addr string `json:"addr"`
	Port int    `json:"port"`
}

type Minio struct {
	Endpoint  string `json:"endpoint"`
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
	Bucket    string `json:"bucket"`
	Secure    bool   `json:"secure"`
}

type Mysql struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"db_name"`
}

type Log struct {
	Level      string `json:"level"`
	Filename   string `json:"filename"`
	MaxSize    int    `json:"maxsize"`
	MaxAge     int    `json:"max_age"`
	MaxBackups int    `json:"max_backups"`
}

type Redis struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	DB       int    `json:"db"`
	Password string `json:"password"`
}

type Kafka struct {
	BootstrapServers []string `json:"bootstrap_servers"`
	Topic            string   `json:"topic"`
	Key              string   `json:"key"`
}
type RedisKey struct {
	SubnetList   string
	SubnetDetail string
}

var Config *config

type api struct {
	Method string
	Uri    string
}

// LoginExcludeAuth 登录后都有权限的接口
var LoginExcludeAuth = []api{
	{"GET", "/auth/*"},
}

// ExcludeAuth 都有权限的接口
var ExcludeAuth = []api{
	{"GET", "/auth/public_key"},
	{"GET", "/"},
	{"GET", ""},
	{"POST", "/auth/login"},
} // 存放不校验的URL

func InitConfig() {
	Config = &config{}
	file, err := os.Open("conf/config.json")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)
	if err = decoder.Decode(&Config); err != nil {
		panic(err)
	}
}
