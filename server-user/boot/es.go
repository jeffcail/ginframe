package boot

import (
	"github.com/jeffcail/ginframe/server-common/driver"
	"github.com/jeffcail/ginframe/server-user/core"
	"github.com/jeffcail/ginframe/server-user/ulogger"
	"log"
)

// InitEs init elastic
func InitEs(url string) {
	es, err := driver.InitEs(url)
	if err != nil {
		log.Fatalln(err)
	}
	core.SetEsClient(es)

	ulogger.UserLogger.Info("🚀🚀🚀🚀🚀🚀ES success...🚀🚀🚀🚀🚀🚀")
}
