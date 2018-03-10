package main

import (
	"encoding/gob"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/rpc"
	"os"
	"os/signal"
	"syscall"
	"time"

	. "github.com/mitjaziv/qmanager/internal/structures"
	. "github.com/mitjaziv/qmanager/internal/uuid"

	"github.com/mitjaziv/qmanager/worker"
	"github.com/mitjaziv/qmanager/worker/operations"
)

var (
	host     string
	delay    int64
	duration time.Duration

	letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")
	ops     = []string{"+", "-", "*", "/"}
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Register RPC/gob data type.
	registerTypes()

	// Command line settings with flags.
	flag.StringVar(
		&host,
		"rpc",
		"0.0.0.0:8090",
		"RPC Server Host:Port",
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

	// Start worker as go routine
	go processor()

	// Start simulator of requests
	go simulator()

	// Run as service until interrupt/kill
	log.Println("Simulator registered")
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

func processor() {
	// Create callback factory and register for all operations.
	cbFactory := worker.NewCallbackFactory().
		Register(&operations.Fibonacci{}, "fibonacci").
		Register(&operations.ArithmeticSolver{}, "arithmetic").
		Register(&operations.ReverseText{}, "reverse").
		Register(&operations.BCrypt{}, "encoder")

	// Create worker and start processing.
	w, err := worker.NewWorker(
		cbFactory,
		worker.Host(host),
		worker.Delay(time.Second*duration),
	)
	if err != nil {
		log.Fatalln(err)
	}
	w.StartProcessing()

	log.Println("Worker registered")
}

func simulator() {
	// Setup client
	client, err := rpc.DialHTTP("tcp", host)
	if err != nil {
		panic(err)
	}

	// Loop endless and send various requests to Queue manager.
	for {
		args := &Task{}

		// Randomize task type.
		n := rand.Intn(4)
		for k := range TaskTypes {
			if n == 0 {
				args.Type = k
			}
			n--
		}

		// Randomize input.
		args.Input = randomInput(args.Type)

		// Add task to Queue manager and log it.
		var reply interface{}
		err = client.Call("RpcHandler.AddTask", args, &reply)
		if err != nil {
			log.Fatal("RPC error:", err)
		}
		log.Printf("Task ID result: %s\n", reply.(UUID))

		// Wait for next request
		time.Sleep(time.Second * duration)
	}
}

// registerTypes will register RPC/gob data types.
func registerTypes() {
	task := new(Task)
	uuid := new(UUID)

	gob.Register(*task)
	gob.Register(*uuid)
}

// randomInput generates random input according to task type.
func randomInput(typ string) interface{} {
	switch typ {
	case "fibonacci":
		return rand.Intn(20)
	case "encoder":
		fallthrough
	case "reverse":
		text := make([]rune, rand.Intn(50))
		for i := range text {
			text[i] = letters[rand.Intn(len(letters))]
		}
		return string(text)
	case "arithmetic":
		text := ""
		for i := 0; i < 5; i++ {
			// Add random number and operations for arithmetic.
			text += fmt.Sprintf("%d", rand.Intn(50)) + " "
			text += ops[rand.Intn(len(ops))] + " "
			text += fmt.Sprintf("%d", rand.Intn(50))
		}
		return text
	}
	return nil
}
