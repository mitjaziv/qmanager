package worker

import (
	"encoding/gob"
	"log"
	"net/rpc"
	"os"
	"os/signal"
	"syscall"
	"time"

	. "github.com/mitjaziv/qmanager/internal/structures"
	. "github.com/mitjaziv/qmanager/internal/uuid"
)

type (
	// Worker interface
	Worker interface {
		StartProcessing()
		Close()
	}

	worker struct {
		host  string
		delay time.Duration

		client  *rpc.Client
		factory *CallbackFactory
	}

	Option func(*worker)
)

// NewWorker function creates worker service connected to RPC server and with registered callbacks.
func NewWorker(factory *CallbackFactory, options ...Option) (Worker, error) {
	// Create worker.
	w := &worker{
		factory: factory,
	}

	// Apply configuration
	applyDefaults(w)
	for _, option := range options {
		option(w)
	}

	// Dial RPC server and add client to worker.
	if client, err := rpc.DialHTTP("tcp", w.host); err == nil {
		w.client = client
	} else {
		return nil, err
	}

	// Register RPC/gob data type.
	registerTypes()

	return w, nil
}

// StartProcessing function starts processing task from queue.
func (w *worker) StartProcessing() {
	go w.handleShutdown(make(chan os.Signal, 1))

	// handle messages from the cluster consumer
	go w.handleTasks()
}

// Close function closes connection to RPC server.
func (w *worker) Close() {
	if w.client != nil {
		err := w.client.Close()
		if err != nil {
			log.Println(err)
		}
	}
}

// handleShutdown function closes connection on system interrupt or kill
func (w *worker) handleShutdown(signals chan os.Signal) {
	defer func() {
		log.Println("Worker shutdown")
		w.Close()
	}()
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)
	<-signals
}

// handleTasks function is collecting tasks from RPC server.
func (w *worker) handleTasks() {

	// Get all registered types.
	types := w.factory.Types()

	for {
		// Sleep between checks
		w.sleep()

		// Take task from Queue manager.
		task, err := w.takeTask(types)
		if err == nil && task == (Task{}) {
			continue
		}
		if err != nil {
			log.Println("TakeTask error:", err)
			continue
		}
		log.Println("Received Task:", task)

		// Process task.
		err = w.factory.Call(&task)
		if err != nil {
			log.Println("ProcessTask error:", err)

			err = w.retryTask(task)
			if err != nil {
				log.Println("RetryTask error:", err)
			}
			continue
		}

		// Save task.
		err = w.saveTask(task)
		if err != nil {
			log.Println("FinishTask error:", err)
			continue
		}
		log.Println("Finished Task:", task)
	}
}

// saveTask function sends output of processed task to Queue manager.
func (w *worker) takeTask(types []string) (Task, error) {
	var reply interface{}

	// Call RpcHandler.TakeTask, which will return next task in queue.
	err := w.client.Call("RpcHandler.TakeTaskByType", types, &reply)
	if err != nil {
		return Task{}, err
	}

	// If there was no response from server return empty task.
	if reply == nil {
		return Task{}, nil
	}

	// Return task.
	return reply.(Task), nil
}

// saveTask function sends output of processed task to Queue manager.
func (w *worker) saveTask(t Task) error {
	var reply interface{}

	// Call RpcHandler.FinishTask, which will save output to queue.
	err := w.client.Call("RpcHandler.FinishTask", t, &reply)
	if err != nil {
		return err
	}
	return nil
}

// retryTask function moves task back to end of the wait queue.
func (w *worker) retryTask(t Task) error {
	var reply interface{}

	// Call RpcHandler.FinishTask, which will save output to queue.
	err := w.client.Call("RpcHandler.RetryTask", t.Id, &reply)
	if err != nil {
		return err
	}
	return nil
}

// sleep function.
func (w *worker) sleep() {
	time.Sleep(w.delay)
}

// registerTypes will register RPC/gob data types.
func registerTypes() {
	task := new(Task)
	uuid := new(UUID)

	gob.Register(*task)
	gob.Register(*uuid)
}

// applyDefaults
func applyDefaults(w *worker) {
	w.host = "0.0.0.0:8080"
	w.delay = time.Second * time.Duration(5)
}
