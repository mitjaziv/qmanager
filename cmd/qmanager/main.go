package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/mitjaziv/qmanager/manager"
)

var (
	conf = manager.Config{}
)

func main() {
	// Read config variables from command line. Usually i would do this with some kind of config
	flag.StringVar(
		&conf.Http,
		"http",
		"0.0.0.0:8080",
		"define HTTP host:port",
	)
	flag.StringVar(
		&conf.RPC,
		"rpc",
		"0.0.0.0:8090",
		"define RPC host:port",
	)
	flag.Parse()

	// Create queue manager instance
	m := manager.NewManager()

	// Create and register HTTP handler
	httpHandler := manager.NewHttpHandler(m)
	httpHandler.RegisterHandler()

	// Create and register RPC handler
	rpcHandler := manager.NewRpcHandler(m)
	err := rpcHandler.RegisterHandler()
	if err != nil {
		log.Fatalln(err)
	}

	// Run service
	log.Println("Starting instance")
	if err := run(); err != nil {
		log.Println(err)
	}
}

func run() error {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)

	// Start HTTP listener
	go func() {
		log.Println("Registering HTTP Listener:", conf.Http)
		err := http.ListenAndServe(conf.Http, nil)
		if err != nil {
			log.Fatalln(err)
		}
	}()

	// Start RPC listener
	go func() {
		log.Println("Registering RPC Listener:", conf.RPC)
		addy, err := net.ResolveTCPAddr("tcp", conf.RPC)
		if err != nil {
			log.Fatal(err)
		}

		conn, err := net.ListenTCP("tcp", addy)
		if err != nil {
			log.Fatal(err)
		}
		err = http.Serve(conn, nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

	// wait on kill signal
	select {
	case <-ch:
	}

	return nil
}
