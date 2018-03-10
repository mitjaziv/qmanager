# Queue Manager worker tool
Worker tool which processes tasks in queue. It has possibility to register to all type of tasks or only to the specific ones.

## Usage
Command line arguments:

| Arguments | Description                                          |
|:----------|:-----------------------------------------------------|
| rpc       | RPC Server Host:Port (default "localhost:8090")      |
| delay     | Send/Receive delay in sec (Default: 1 sec)           |
| type      | Task Type (all,fibonacci,arithmetic,reverse,encoder) |

Example of usage:
~~~
# ./qmworker -type encoder -type reverse
2018/03/11 12:05:42 registered callback for: encoder
2018/03/11 12:05:42 registered callback for: reverse
2018/03/11 12:05:42 Worker registered

# ./qmworker -type all
2018/03/11 16:13:44 registered callback for: fibonacci
2018/03/11 16:13:44 registered callback for: arithmetic
2018/03/11 16:13:44 registered callback for: reverse
2018/03/11 16:13:44 registered callback for: encoder
2018/03/11 16:13:44 Worker registered
2018/03/11 16:13:45 Received Task: ID: 6721bfa3-b08b-4ceb-9af5-554accc7bbf3	Type: reverse	Input: GHRghlxL6pnVmJECYl54FtC1ItAu7YloJ	Output: <nil>
2018/03/11 16:13:45 Finished Task: ID: 6721bfa3-b08b-4ceb-9af5-554accc7bbf3	Type: reverse	Input: GHRghlxL6pnVmJECYl54FtC1ItAu7YloJ	Output: JolY7uAtI1CtF45lYCEJmVnp6LxlhgRHG
2018/03/11 16:13:46 Received Task: ID: 1e9becb9-14d4-43a8-9ea9-2fb089bb02db	Type: reverse	Input: 5cJAtXIPA2aV1MeOkY	Output: <nil>
2018/03/11 16:13:46 Finished Task: ID: 1e9becb9-14d4-43a8-9ea9-2fb089bb02db	Type: reverse	Input: 5cJAtXIPA2aV1MeOkY	Output: YkOeM1Va2APIXtAJc5
2018/03/11 16:13:47 Received Task: ID: 66acf134-442c-4bda-a294-0fa872968849	Type: arithmetic	Input: 36 - 3321 + 423 + 2242 - 382 * 0	Output: <nil>
2018/03/11 16:13:47 Finished Task: ID: 66acf134-442c-4bda-a294-0fa872968849	Type: arithmetic	Input: 36 - 3321 + 423 + 2242 - 382 * 0	Output: -620
2018/03/11 16:13:48 Received Task: ID: a17b16c2-a394-4589-8e9d-d8c2e44ba9a9	Type: fibonacci	Input: 4	Output: <nil>
2018/03/11 16:13:48 Finished Task: ID: a17b16c2-a394-4589-8e9d-d8c2e44ba9a9	Type: fibonacci	Input: 4	Output: 3
2018/03/11 16:13:49 Received Task: ID: c237be6a-10c4-4b0b-92b8-1a0f7d84f0a3	Type: reverse	Input: YXxznJRRqJxbAxgZlSWgmaNMb5sjVYJ7VzPH	Output: <nil>
2018/03/11 16:13:49 Finished Task: ID: c237be6a-10c4-4b0b-92b8-1a0f7d84f0a3	Type: reverse	Input: YXxznJRRqJxbAxgZlSWgmaNMb5sjVYJ7VzPH	Output: HPzV7JYVjs5bMNamgWSlZgxAbxJqRRJnzxXY
~~~

## Installation

```bash
go get -u github.com/mitjaziv/qmanager/cmd/qmworker
```

## Screenshots

![](../../docs/qmworker.gif)

## License

[WTFPL]()
