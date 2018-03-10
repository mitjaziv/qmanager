# Queue Manager service
Queue manager main service.

## Usage
Command line arguments:

| Arguments | Description                                          |
|:----------|:-----------------------------------------------------|
| http      | define HTTP host:port (default "localhost:8080")     |
| rpc       | RPC Server Host:Port (default "localhost:8090")      |


Example of usage:
~~~
# ./qmanager -h
Usage of ./qmanager:
  -http string
    	define HTTP host:port (default "localhost:8080")
  -rpc string
    	define RPC host:port (default "localhost:8090")

# ./qmanager -http "localhost:8080" &
2018/03/11 15:12:44 rpc.Register: method "RegisterHandler" has 1 input parameters; needs exactly three
2018/03/11 15:12:44 Starting instance
2018/03/11 15:12:44 Registering HTTP Listener: localhost:8080
2018/03/11 15:12:44 Registering RPC Listener: localhost:8090
~~~

## Installation

```bash
go get -u github.com/mitjaziv/qmanager/cmd/qmanager
```

## Screenshots

![](../../docs/qmanager.gif)

## License

[WTFPL]()
