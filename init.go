package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/streadway/amqp"
)

// Configuration File Opjects
type configuration struct {
	ServerName     string
	AppName        string
	AppVer         string
	Broker         string
	BrokerUser     string
	BrokerPwd      string
	BrokerExchange string
	BrokerVhost    string
	ChannelSize    int
	ChannelCount   int
	SrvPort        string
}

var (
	err           error
	arbitraryJSON interface{}
	conf          configuration
	messages      chan chanToRabbit
)

func init() {

	//Load Default Configuration Values
	conf.AppName = "Go - UDP to RMQ"
	conf.AppVer = "1.0"
	conf.ServerName, _ = os.Hostname()
	conf.ChannelSize = 2048
	conf.SrvPort = "8514"
	conf.Broker = "127.0.0.1"
	conf.BrokerUser = "guest"
	conf.BrokerPwd = "guest"
	conf.BrokerExchange = "amq.topic"
	conf.BrokerVhost = "/"
	conf.ChannelCount = 4

	//Load Configuration Data
	dat, _ := ioutil.ReadFile("conf.json")
	_ = json.Unmarshal(dat, &conf)

	fmt.Println(conf.AppName, " ver: ", conf.AppVer, " starting...")
	fmt.Println("On Port: ", conf.SrvPort)

	messages = make(chan chanToRabbit, conf.ChannelSize)

	// create the rabbitmq error channel
	rabbitCloseError = make(chan *amqp.Error)

	// run the callback in a separate thread
	go rabbitConnector(fmt.Sprint("amqp://" + conf.BrokerUser + ":" + conf.BrokerPwd + "@" + conf.Broker + conf.BrokerVhost))

	// establish the rabbitmq connection by sending
	// an error and thus calling the error callback
	rabbitCloseError <- amqp.ErrClosed

	for rabbitConn == nil {
		log.Println("Waiting for connection to rabbitmq...")
		time.Sleep(time.Second * 1)
	}

	for i := 0; i < conf.ChannelCount; i++ {
		go func() {
			for {
				chanPubToRabbit()
				time.Sleep(time.Second * 5)
			}
		}()
	}

}

func checkError(err error) {
	if err != nil {
		log.Println("Error: ", err)
	}
}
