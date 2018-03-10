package main

import (
	"encoding/gob"
	"flag"
	"fmt"
	"log"
	"net/rpc"
	"os"

	. "github.com/mitjaziv/qmanager/internal/structures"
	. "github.com/mitjaziv/qmanager/internal/uuid"
)

func main() {
	// Register RPC/gob data type.
	registerTypes()

	// Commands.
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)

	// List Add command flag pointers
	getTypePtr := addCmd.String(
		"type",
		"",
		"Task Type {fibonacci|arithmetic|reverse|encoder}. (Required)",
	)
	getInputPtr := addCmd.String(
		"input",
		"",
		"Task Input. (Required)",
	)
	getAddHostPtr := addCmd.String(
		"rpc",
		"0.0.0.0:8090",
		"RPC Server Host:Port",
	)

	// List Get command flag pointers
	getIdPtr := getCmd.String(
		"id",
		"",
		"Task ID. (Required)",
	)

	getGetHostPtr := getCmd.String(
		"rpc",
		"0.0.0.0:8090",
		"RPC Server Host:Port",
	)

	// Check for command.
	if len(os.Args) < 2 {
		fmt.Println("add or get is required")
		os.Exit(1)
	}

	// Check for command attributes.
	switch os.Args[1] {
	case "add":
		err := addCmd.Parse(os.Args[2:])
		if err != nil {
			log.Fatal(err)
		}
	case "get":
		err := getCmd.Parse(os.Args[2:])
		if err != nil {
			log.Fatal(err)
		}
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Add command
	if addCmd.Parsed() {

		// Task type
		if _, ok := TaskTypes[*getTypePtr]; !ok {
			addCmd.PrintDefaults()
			os.Exit(1)
		}

		// Task input
		if *getInputPtr == "" {
			addCmd.PrintDefaults()
			os.Exit(1)
		}

		// Dial server
		client, err := rpc.DialHTTP("tcp", *getAddHostPtr)
		if err != nil {
			log.Fatal("dialing:", err)
		}

		// args
		args := &Task{
			Type:  *getTypePtr,
			Input: *getInputPtr,
		}

		// Synchronous call
		var reply interface{}
		err = client.Call("RpcHandler.AddTask", args, &reply)
		if err != nil {
			log.Fatal("RPC error:", err)
		}
		id := reply.(UUID)

		// Print
		fmt.Printf("Task ID result: %s\n", id)
	}

	// Get command
	if getCmd.Parsed() {
		if *getIdPtr == "" {
			getCmd.PrintDefaults()
			os.Exit(1)
		}

		// Dial server
		client, err := rpc.DialHTTP("tcp", *getGetHostPtr)
		if err != nil {
			log.Fatal("dialing:", err)
		}

		// Synchronous call
		var reply interface{}
		err = client.Call("RpcHandler.GetTask", *getIdPtr, &reply)
		if err != nil {
			log.Fatal("RPC error:", err)
		}
		if reply != nil {
			t := reply.(Task)

			// Write to output
			fmt.Println("Task", t)
		}
	}
}

// registerTypes will register RPC/gob data types.
func registerTypes() {
	task := new(Task)
	uuid := new(UUID)

	gob.Register(*task)
	gob.Register(*uuid)
}
