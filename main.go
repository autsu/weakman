package main

import (
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"vote/v2/router"
)


func init() {
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	logrus.SetReportCaller(true)

	log.SetFlags(log.Llongfile | log.Lmicroseconds)
}

func main() {
	//logs.InitLogger()
	//defer logs.ZapLogger.Sync()
	g := router.Router()
	if err := g.Run(":8080"); err != nil {
		log.Fatalln(err)
	}
}
