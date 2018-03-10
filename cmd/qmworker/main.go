package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mitjaziv/qmanager/worker"
	"github.com/mitjaziv/qmanager/worker/operations"
)

type (
	arrayFlags []string
)

func main() {
	// Command line settings with flags.
	var types arrayFlags
	var host string
	var delay int64
	var duration time.Duration

	flag.StringVar(
		&host,
		"rpc",
		"0.0.0.0:8090",
		"RPC Server Host:Port",
	)
	flag.Var(
		&types,
		"type",
		"Task Type {all|fibonacci|arithmetic|reverse|encoder}",
	)
	flag.Int64Var(
		&delay,
		"delay",
		1,
		"Send/Receive delay in sec (Default: 1 sec)",
	)
	flag.Parse()

	// Calculate delay duration
	duration = time.Duration(delay)

	// Check flags
	if len(types) == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}
	all := types.Contains("all")

	// Create callback factory
	cbFactory := worker.NewCallbackFactory()

	// Register fibonacci
	if all || types.Contains("fibonacci") {
		cbFactory = cbFactory.Register(&operations.Fibonacci{}, "fibonacci")
	}

	// Register arithmetic's solver
	if all || types.Contains("arithmetic") {
		cbFactory = cbFactory.Register(&operations.ArithmeticSolver{}, "arithmetic")
	}

	// Register reverse text
	if all || types.Contains("reverse") {
		cbFactory = cbFactory.Register(&operations.ReverseText{}, "reverse")
	}

	// Register text encoder
	if all || types.Contains("encoder") {
		cbFactory = cbFactory.Register(&operations.BCrypt{}, "encoder")
	}

	// Create worker and start processing
	w, err := worker.NewWorker(
		cbFactory,
		worker.Host(host),
		worker.Delay(time.Second*duration),
	)
	if err != nil {
		log.Fatalln(err)
	}
	w.StartProcessing()

	// Run worker
	log.Println("Worker registered")
	if err := run(); err != nil {
		log.Println(err)
	}
}

func run() error {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)

	// wait on kill signal
	select {
	case <-ch:
	}

	return nil
}

func (i *arrayFlags) String() string {
	return "" // we don't need it
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func (i *arrayFlags) Contains(value string) bool {
	for _, v := range *i {
		if v == value {
			return true
		}
	}
	return false
}
