# Read-server

Server writes data from connected users to file.

## Installing

Required Go version: 1.17

Clone repository and run:
```shell
$ make server
$ make client
```

## Usage

Client and server are placed in ./bin folder.
To check arguments run with `-help`. Example:
```shell
$ ./bin/read-server --help
  -file string
    	file where will be written clients data (default "file.txt")
  -help
    	show help
  -max int
    	max connections to server (default 5)
  -s string
    	address of the read-server (default "localhost:8080")
```

With server you can use client:
```shell
$ ./bin/client -s localhost:8080
Enter your id
10
Enter message:
Привет
Enter message:
Как дела
Enter message:
Hello
Enter message:
How are you
Enter message:
Closing connect
```

You can use netcat (`nc $host $port`) as client,
but as first line you need to put your id (number):
```shell
$ nc localhost 8080
10
Привет
Как дела
Hello
How are you
```

All data sent by clients is written to file. Example:
```
10: `Привет`
10: `Как дела`
10: `Hello`
10: `How are you`
```