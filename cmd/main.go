package main

import (
	"fmt"
	"github.com/dataleodev/igridmq"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	mflog "github.com/mainflux/mainflux/logger"

	"github.com/mainflux/mproxy/pkg/mqtt"
	"github.com/mainflux/mproxy/pkg/session"
	"github.com/mainflux/mproxy/pkg/websocket"
)

const (

	//WS
	defWSPath         = "/mqtt"
	defWSPort         = "8080"
	defWSTargetScheme = "ws"
	defWSTargetHost   = "localhost"
	defWSTargetPort   = "8888"
	defWSTargetPath   = "/mqtt"

	// MQTT
	defMQTTHost       = "0.0.0.0"
	defMQTTPort       = "1883"
	defMQTTTargetHost = "mosquitto"
	defMQTTTargetPort = "1884"

)


func main() {

	logger, err := mflog.New(os.Stdout, "debug")
	if err != nil {
		log.Fatalf(err.Error())
	}

	h := igridmq.New(logger)

	errs := make(chan error, 3)

	// WS
	logger.Info(fmt.Sprintf("Starting WebSocket proxy on port %s ",defWSPort))
	go proxyWS(logger, h, errs)

	// MQTT
	logger.Info(fmt.Sprintf("Starting MQTT proxy on port %s ", defMQTTPort))
	go proxyMQTT(logger, h, errs)
	//}

	go func() {
		c := make(chan os.Signal, 2)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	err = <-errs
	logger.Error(fmt.Sprintf("igridmq terminated: %s", err))
}

func proxyWS(logger mflog.Logger, handler session.Handler, errs chan error) {
	target := fmt.Sprintf("%s:%s", defWSTargetHost, defWSTargetPort)
	wp := websocket.New(target, defWSTargetPath, defWSTargetScheme, handler, logger)
	http.Handle(defWSPath, wp.Handler())

	errs <- wp.Listen(defWSPort)
}

func proxyMQTT(logger mflog.Logger, handler session.Handler, errs chan error) {
	address := fmt.Sprintf("%s:%s", defMQTTHost, defMQTTPort)
	target := fmt.Sprintf("%s:%s", defMQTTTargetHost, defMQTTTargetPort)
	mp := mqtt.New(address, target, handler, logger)

	errs <- mp.Listen()
}

