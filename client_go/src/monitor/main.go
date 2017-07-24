package main

import (
	"fmt"
	"monitor/applog"
	//后端相关程序
	"monitor/backend"
	//配置文件读取相关
	"monitor/config"
	"monitor/dataserver"
	//获取一些平台相关的方法，以及一些goroutine的配置
	"runtime"

	//os包中实现了平台无关的接口，设计向Unix风格，但是错误处理是go风格，
	//当os包使用时，如果失败之后返回错误类型而不是错误数量．
	"os"
	//用于监听和取消监听消息
	"os/signal"
	//OS级别的系统调用
	"syscall"
	//结构化、可热插拔的 Go记录
	//log相当于给这个包起了一个 别名
	log "github.com/Sirupsen/logrus"
	//cil是一个帮助生成 命令app 开发的框架包
	//详见http://hao.jobbole.com/cli/
	"github.com/codegangsta/cli"
)

func run(c *cli.Context) error {
	//查找"conf"关键字 并返回其地址
	conf, err := config.ReadConfig(c.String("conf"))
	if err != nil {
		log.Error("read from conf fail!", c.String("conf"))
		return err
	}
	fmt.Println("conf =  ", conf)

	//go的运行系统
	fmt.Println("runtime.GOOS = ", runtime.GOOS)

	//打印程序运行的一些信息
	//疑问 调用这个包 比直接println等的优点在什么地方？
	var logger *applog.AutoDailyLoger
	if runtime.GOOS == "windows" {
		logger = applog.NewAutoDailyLoger(conf.LogDirWin, conf.LogPrefix)
	} else {
		logger = applog.NewAutoDailyLoger(conf.LogDirLinux, conf.LogPrefix)
	}
	logger.Start()
	defer logger.Stop()

	//Create MQTT Backend
	//利用读到的配置信息创建一个后端程序
	mqtt, err := backend.NewBackend(conf.MqttServer, conf.MqttUsername, conf.MqttPassword, conf.MqttPubTopic)
	if err != nil {
		log.Error("can not connect mqtt server")
		return err
	}
	defer mqtt.Close()

	// Subscribe Topic
	//订阅主题
	topic := conf.MqttSubTopic
	err = mqtt.SubscribeTopic(topic)
	if err != nil {
		log.Errorf("SubscribeTopic %v Error", topic)
		return err
	}
	defer mqtt.UnSubscribeTopic(topic)

	//start DataServer
	ds := dataserver.NewDataServer(mqtt, conf.Mac)
	go func() {
		ds.Start()
	}()
	defer ds.Stop()

	//start http server
	//	go func() {
	//		if runtime.GOOS == "windows" {
	//			httpserver.StartHttpServer(mqtt, db, conf.HttpServerWin)
	//		} else {
	//			httpserver.StartHttpServer(mqtt, db, conf.HttpServerLinux)
	//		}
	//	}()

	//quit when receive end signal
	//看到这里就迷糊了 感觉还是对程序的并发理解的不够。
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	log.Infof("signal received signal %v", <-sigChan)
	log.Warn("shutting down server")
	return nil
}

func main() {
	//实例化一个对象
	app := cli.NewApp()

	app.Name = "monitor"
	app.Usage = "monitor pi board and transmit temprature and control led from remote server!"
	app.Copyright = "monitor123456@gmail.com"
	app.Version = "0.0.3"
	app.Action = run
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "conf,c",
			Usage:  "Set conf path here",
			Value:  "appserver.conf",
			EnvVar: "APP_CONF",
		},
	}
	app.Run(os.Args)
}
