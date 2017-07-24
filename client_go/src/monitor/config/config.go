package config

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"gopkg.in/ini.v1"
)

//定义Config类
type Config struct {
	//MQTT
	MqttServer   string `ini:"mqtt_server"` //`ini:name`的写法和映射的实现方法理暂时未理解
	MqttUsername string `ini:"mqtt_username"`
	MqttPassword string `ini:"mqtt_password"`
	MqttSubTopic string `ini:"mqtt_subtopic"`
	MqttPubTopic string `ini:"mqtt_pubtopic"`
	//MAC
	Mac string `ini:"mac"`
	//LOG
	LogDirWin   string `ini:"log_dir_win"`
	LogDirLinux string `ini:"log_dir_linux"`
	LogPrefix   string `ini:"log_prefix"`
	LogToFile   bool   `ini:"log_tofile"`
}

//把读取的配置格式化成字符串
func (c Config) String() string {
	mqtt := fmt.Sprintf("MQTT:[%v]/[%v]/[%v]/[Sub:%v]/[Pub:%v]", c.MqttServer, c.MqttUsername, c.MqttPassword, c.MqttSubTopic, c.MqttPubTopic)

	log := fmt.Sprintf("LOG:[win:%v]/[linux:%v]:[prefix:%v]:[tofile:%v]", c.LogDirWin, c.LogDirLinux, c.LogPrefix, c.LogToFile)

	mac := c.Mac
	return mqtt + ", " + log + ", " + mac
}

//参考  http://www.liuhaihua.cn/archives/238877.html
//Read Server's Config Value from "path"
func ReadConfig(path string) (Config, error) {
	//实例化一个对象
	var config Config
	//加载数据源
	conf, err := ini.Load(path)
	if err != nil {
		log.Println("load config file fail!")
		return config, err
	}
	//在确定只有读操作的时候 可以将该字段置false 提高读的性能
	conf.BlockMode = false

	//把文件映射到结构体？
	//如何映射？ 内部实现原理暂时还没有理解。
	err = conf.MapTo(&config)
	if err != nil {
		log.Println("mapto config file fail!")
		return config, err
	}

	return config, nil
}
