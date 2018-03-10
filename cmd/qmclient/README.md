# Queue Manager client tool
Provides functionality for adding new tasks to queue and retrieving finished tasks from queue.

## Commands
Command add or get is required.

#### Add

| Arguments | Description                                                  |
|:----------|:-------------------------------------------------------------|
| rpc       | RPC Server Host:Port (default "localhost:8090")              |
| type      | Task Type (fibonacci,arithmetic,reverse,encoder). (Required) |
| input     | Task Input. (Required)                                       |

Example of usage:
~~~
# ./qmclient add -type fibonacci -input 15
Task ID result: 04b9b218-9cf8-4f50-b9c3-0345398cc1b4
~~~

#### Get

|Arguments | Description                                     |
|:---------|:------------------------------------------------|
| id       | Task ID. (Required)                             |
| rpc      | RPC Server Host:Port (default "localhost:8090") |

Example of usage:
~~~
# ./qmclient get -id 04b9b218-9cf8-4f50-b9c3-0345398cc1b4
Task ID: 04b9b218-9cf8-4f50-b9c3-0345398cc1b4	Type: fibonacci	Input: 15	Output: 610
~~~

## Installation

```bash
go get -u github.com/mitjaziv/qmanager/cmd/qmclient
```

## Screenshots

![](../../docs/qmclient.gif)

## License

[WTFPL]()
