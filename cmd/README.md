# Queue Manager command line tools
Provides command line tools for Queue manager.

## Commands

| Command                     | Description                                                         |
|:----------------------------|:--------------------------------------------------------------------|
| [qmanager](qmanager/)       | Queue manager main service.                                         |
| [qmclient](qmclient/)       | Tool for adding and reading tasks in Queue.                         |
| [qmsimulator](qmsimulator/) | Simulator demonstration tool which adds and process tasks in Queue. |
| [qmworker](qmworker/)       | Worker tool which processes registered tasks in Queue.              |

## Installation

```bash
go get -u github.com/mitjaziv/qmanager/cmd/...
```

## Screenshots
qmanager:
![](../docs/qmanager.gif)
qmclient:
![](../docs/qmclient.gif)
qmsimulator:
![](../docs/qmsimulator.gif)
qmworker:
![](../docs/qmworker.gif)

## License

[WTFPL]()
