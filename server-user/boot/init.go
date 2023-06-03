// init 项目初始化

package boot

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jeffcail/ginframe/server-common/config"
	_const "github.com/jeffcail/ginframe/server-common/const"
	"github.com/jeffcail/ginframe/server-common/nacosRF"
	_process "github.com/jeffcail/ginframe/server-common/process"
	_ip "github.com/jeffcail/ginframe/server-common/utils/ip"
	"github.com/jeffcail/ginframe/server-user/cachedb"
	"github.com/jeffcail/ginframe/server-user/daos/serverDao"
	"github.com/jeffcail/ginframe/server-user/global"
	_rpc "github.com/jeffcail/ginframe/server-user/rpc"
	"github.com/jeffcail/ginframe/server-user/ulogger"
	"go.uber.org/zap"
	"io"
	"log"
	"os"
	"strings"
)

var (
	RootDir       string
	err           error
	cf            string
	MysqlDbDsn    string
	MaxOpenConn   int
	MaxIdleConn   int
	RedisAddr     string
	RedisPassword string
	RedisDb       int
	MongoAddr     string
	LeveldbPath   string
	LoggerPath    string
	EsUrl         string
)

var (
	ip    = flag.String("ip", "127.0.0.1", "The nacos of ip address")
	port  = flag.Int("p", 7848, "The nacos of port")
	cfg   = flag.String("c", "server-user.yml", "The nacos of Data ID")
	group = flag.String("g", "gin-frame", "The nacos of Group")
)

func init() {
	flag.Parse()
	loadRemoteConfig(*ip, *port, *cfg, *group, &global.Config)
	if global.Config.IsEnableGOMAXPROCS {
		_process.GroRuntimeMaxCpu()
	}
	assignment()
	initDB()
	nacosRF.InitNacos(global.Config.Nacos)
	setUserServer()
	registerUserServer()
}

func loadRemoteConfig(ip string, port int, cfg string, group string, configs interface{}) {
	config.LoadCoreConfig(ip, port, cfg, group, configs)
}

// assignment
func assignment() {
	MysqlDbDsn = global.Config.Mysql.DbDsn
	MaxOpenConn = global.Config.Mysql.MaxOpenConns
	MaxIdleConn = global.Config.Mysql.MaxIdleConns
	RedisAddr = global.Config.Redis.Addr
	RedisPassword = global.Config.Redis.Password
	RedisDb = global.Config.Redis.Db
	MongoAddr = global.Config.Mongo.Addr
	LeveldbPath = global.Config.LevelDb.Path
	LoggerPath = global.Config.Http.LogPath
	//EsUrl = global.Config.Elastic.Url
}

func initDB() {
	InitLog(LoggerPath)
	InitDb(MysqlDbDsn, MaxOpenConn, MaxIdleConn)
	InitRedis(RedisAddr, RedisPassword, RedisDb)
	InitMongo(MongoAddr)
	InitLevelDb(LeveldbPath)
	//go InitEs(EsUrl)
}

// 注册服务， 服务必须存在server_flag表中， 可通过后台可视化页面添加
// server_name 需保持一致
// 如: 用户服务为`user`，所有的用户服务server_name 都为`user`，依次类推
func setUserServer() {
	ip := _ip.GetIp()
	server, err := serverDao.CheckIpIsNewServer(ip, _const.ServerNameU)
	if err != nil {
		ulogger.UserLogger.Error("==========> user服务初始化,检测是否为新服务失败 <==========\n",
			zap.String("ip", ip), zap.Error(err))
		os.Exit(1)
	}
	if server.Id == 0 {
		ulogger.UserLogger.Error("==========> user服务初始化,库中服务未找到 <==========\n",
			zap.String("ip", ip), zap.Error(err))
		os.Exit(1)
	}
	// 新服务注册
	err = cachedb.SetUserServerFlag(fmt.Sprintf("%s%s%s", ip, ":", _const.UserServerName),
		fmt.Sprintf("%s%d", _const.UserServerName, server.Sum))
	if err != nil {
		os.Exit(1)
	}
}

func registerUserServer() {
	ip := _ip.GetIp()
	k := fmt.Sprintf("%s%s%s", ip, ":", _const.UserServerName)
	nacosRF.NacosInstance.Register(cachedb.GetUserServerFlag(k), global.Config.Nacos.Port)
}

// Boot 项目初始化
func Boot() {
	errs := make(chan error)
	go func() {
		err = HttpServe()
		if err != nil {
			errs <- err
		}
	}()

	go func() {
		err = RunGrpcServer()
		if err != nil {
			errs <- err
		}
	}()

	go func() {
		err = WebsocketServer()
		if err != nil {
			errs <- err
		}
	}()

	go func() {
		err = _rpc.RpcServer()
		if err != nil {
			errs <- err
		}
	}()

	debugPrint("*****************************************************************")
	debugPrint("* Listening and serving HTTP on %s *\n", global.Config.Http.BindPort)
	debugPrint("* Listening and serving GRPC on %s *\n", global.Config.Grpc.Port)
	debugPrint("* Listening and serving Websocket on %s *\n", global.Config.Websocket.BindPort)
	debugPrint("* Listening and serving rpc on %s *\n", global.Config.Rpc.Listener)
	debugPrint("*****************************************************************")

	select {
	case err = <-errs:
		log.Fatalf("Run server err: %v", err)
	}

}

var DefaultWriter io.Writer = os.Stdout

func debugPrint(format string, values ...any) {
	if gin.IsDebugging() {
		if !strings.HasSuffix(format, "\n") {
			format += "\n"
		}
		fmt.Fprintf(DefaultWriter, "[GIN-debug] "+format, values...)
	}
}
