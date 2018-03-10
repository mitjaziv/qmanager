# Queue Manager simulation tool
Simulation tool is used for Queue manager demonstration.

## Usage

Command line arguments:

| Arguments | Description                                      |
|:----------|:-------------------------------------------------|
| rpc       | RPC Server Host:Port (default "localhost:8090")  |
| delay     | Send/Receive delay in sec (Default: 1 sec)       |

Example of usage:
~~~
# ./qmsimulator -delay 3
2018/03/11 16:13:15 Simulator registered
2018/03/11 16:13:15 registered callback for: fibonacci
2018/03/11 16:13:15 registered callback for: arithmetic
2018/03/11 16:13:15 registered callback for: reverse
2018/03/11 16:13:15 registered callback for: encoder
2018/03/11 16:13:15 Worker registered
2018/03/11 16:13:15 Task ID result: 938b025e-917e-4963-829d-616a3dfd1473
2018/03/11 16:13:16 Received Task: ID: 938b025e-917e-4963-829d-616a3dfd1473	Type: reverse	Input: eCZXv5JnvMCtfsL5B9aGMVopPQ	Output: <nil>
2018/03/11 16:13:16 Task ID result: 0f9761ae-5b6f-43d3-a260-334951aaff2b
2018/03/11 16:13:16 Finished Task: ID: 938b025e-917e-4963-829d-616a3dfd1473	Type: reverse	Input: eCZXv5JnvMCtfsL5B9aGMVopPQ	Output: QPpoVMGa9B5LsftCMvnJ5vXZCe
2018/03/11 16:13:17 Task ID result: 943fd810-44e2-45ef-9846-01ce6f53a68f
2018/03/11 16:13:17 Received Task: ID: 0f9761ae-5b6f-43d3-a260-334951aaff2b	Type: encoder	Input: LDxsOPQzzQhMQ8DYYr2A6MKvmVMqmab	Output: <nil>
2018/03/11 16:13:17 Finished Task: ID: 0f9761ae-5b6f-43d3-a260-334951aaff2b	Type: encoder	Input: LDxsOPQzzQhMQ8DYYr2A6MKvmVMqmab	Output: $2a$10$CQFS4FtmGFl2TdPs4r4eseJxhDnOOevc1B6yY5irPZBoZ.AY4Tnne
2018/03/11 16:13:18 Task ID result: f40d4811-6702-4463-90ce-9b9267159f44
~~~

## Installation

```bash
go get -u github.com/mitjaziv/qmanager/cmd/qmsimulator
```

## Screenshots

![](../../docs/qmsimulator.gif)

## License

[WTFPL]()
