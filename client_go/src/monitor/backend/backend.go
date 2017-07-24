package backend

import (
	//"encoding/base64"
	//SON(JavaScript Object Notation)是一种轻量级的数据交换格式。
	//可以去json.org 查看json标准的清晰定义。
	//json package 是GO 语言官方提供的包
	"encoding/json"
	//"errors"
	//"fmt"
	"monitor/packets"
	"sync"

	log "github.com/Sirupsen/logrus"

	//提供了连接到MQTT 发布和订阅主题 和接收消息的方法
	"github.com/eclipse/paho.mqtt.golang"
	//"github.com/satori/go.uuid"
)

// Backend implements a MQTT pub-sub backend.
type Backend struct {
	conn     mqtt.Client
	rxChan   chan packets.LedPacket
	mutex    sync.RWMutex
	pubtopic string //发布主题
	subtopic string //订阅主题
}

// NewBackend creates a new Backend.
func NewBackend(server, username, password string, pubtopic string) (*Backend, error) {
	b := Backend{
		rxChan:   make(chan packets.LedPacket),
		pubtopic: pubtopic,
		subtopic: "",
	}
	//创建对象 NewClientOptions 里面包含了一些MQTT 的设置
	opts := mqtt.NewClientOptions()
	opts.AddBroker(server)
	opts.SetUsername(username)
	opts.SetPassword(password)
	opts.SetOnConnectHandler(b.onConnected)
	opts.SetConnectionLostHandler(b.onConnectionLost)
	log.Infof("backend/mqttpubsub: connecting to mqtt broker %v", server)
	//用获取的配置信息来创建一个新的客户端
	b.conn = mqtt.NewClient(opts)

	if token := b.conn.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	return &b, nil
}

// Close closes the backend.
func (b *Backend) Close() {
	b.conn.Disconnect(250) // wait 250 milisec to complete pending actions
	log.Info("backend/mqttpubsub: Disconnect mqtt broker")
}

// RxDataChan returns the TabRxData channel.
func (b *Backend) RxDataChan() chan packets.LedPacket {
	return b.rxChan
}

// Subscribe RxData
func (b *Backend) SubscribeTopic(topic string) error {
	defer b.mutex.Unlock()
	b.mutex.Lock()

	log.Infof("backend/mqttpubsub: subscribing to topic %v", topic)
	if token := b.conn.Subscribe(topic, 2, b.rxDataHandler); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	b.subtopic = topic
	return nil
}

// UnSubscribe RxData
func (b *Backend) UnSubscribeTopic(topic string) error {
	defer b.mutex.Unlock()
	b.mutex.Lock()

	log.Infof("backend/mqttpubsub: unsubscribing from topic %v", topic)
	if token := b.conn.Unsubscribe(topic); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	b.subtopic = ""
	return nil
}

//发布温度
func (b *Backend) PublishTempPacket(pkt packets.TempPacket) error {
	bytes, err := json.Marshal(pkt)
	if err != nil {
		return err
	}

	log.WithFields(log.Fields{
		"topic": b.pubtopic,
	}).Info("publish:")
	if token := b.conn.Publish(b.pubtopic, 2, false, bytes); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

//处理收到的数据
func (b *Backend) rxDataHandler(c mqtt.Client, msg mqtt.Message) {
	log.Infof("backend/mqttpubsub: packet received from topic %v", msg.Topic())
	var rxData packets.LedPacket
	if err := json.Unmarshal(msg.Payload(), &rxData); err != nil {
		log.Errorf("backend/mqttpubsub: decode rxData error: %s", err)
		return
	}
	log.Infof("backend/mqttpubsub: packet received  %v", rxData)
	b.rxChan <- rxData
}

func (b *Backend) onConnected(c mqtt.Client) {
	defer b.mutex.RUnlock()
	b.mutex.RLock()
	if b.subtopic == "" {
		return
	}
	log.Info("backend/mqttpubsub: onConnected to mqtt broker")
	if token := b.conn.Subscribe(b.subtopic, 2, b.rxDataHandler); token.Wait() && token.Error() != nil {
		return
	}
}

func (b *Backend) onConnectionLost(c mqtt.Client, reason error) {
	log.Errorf("backend/mqttpubsub: mqtt onConnectionLost error: %s", reason)
}
