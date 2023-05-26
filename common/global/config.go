// 公共的全局配置 package
// GoAppConfig 框架应用配置
// Http http服务配置
// Mysql 数据库配置
// Redis redis配置

package global

import "sync"

// GoAppConfig 框架应用配置
type GoAppConfig struct {
	GinAppDebug string    `json:"gin_app_debug"`
	GinAppName  string    `json:"gin_app_name"`
	Http        Http      `json:"http"`
	Websocket   Websocket `json:"websocket"`
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

// Mysql 数据库配置
type Mysql struct {
	DbDsn string `json:"db_dsn"`
}

// Redis redis配置
type Redis struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var (
	Config    *GoAppConfig
	once      sync.Once
	AppConfig *Application
)

func NewGoAppConfig() {
	once.Do(new)
}

func new() {
	Config = &GoAppConfig{}
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
