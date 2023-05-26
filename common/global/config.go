// 公共的全局配置 package
// GoAppConfig 框架应用配置
// Http http服务配置
// Websocket Websocket配置
// Grpc Grpc配置
// Mysql 数据库配置
// Redis redis配置

package global

import "sync"

// GlobalConfig 框架应用配置
type GlobalConfig struct {
	GinAppDebug string    `json:"gin_app_debug"`
	GinAppName  string    `json:"gin_app_name"`
	Http        Http      `json:"http"`
	Websocket   Websocket `json:"websocket"`
	Grpc        Grpc      `json:"grpc"`
	Mysql       Mysql     `json:"mysql"`
	Redis       Redis     `json:"redis"`
}

// Http http服务配置
type Http struct {
	BindPort string `json:"bind_port"`
	LogPath  string `json:"log_path"`
}

// Websocket Websocket服务配置
type Websocket struct {
	BindPort string `json:"bind_port"`
}

// Grpc Grpc配置
type Grpc struct {
	Port string `json:"port"`
}

// Mysql 数据库配置
type Mysql struct {
	DbDsn        string `json:"db_dsn"`
	MaxOpenConns int    `json:"max_open_conns"`
	MaxIdleConns int    `json:"max_idle_conns"`
}

// Redis redis配置
type Redis struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var (
	Config    *GlobalConfig
	once      sync.Once
	AppConfig *Application
)

func NewGoAppConfig() {
	once.Do(new)
}

func new() {
	Config = &GlobalConfig{}
}

type Application struct {
	ConfigRemote bool `json:"config_remote"`
}

func NewApplicationConfig() {
	once.Do(newApp)
}

func newApp() {
	AppConfig = &Application{}
}
