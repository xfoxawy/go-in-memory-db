# Go Simple In Memory DB

## Getting Started

Small In Memory DB for local development, no need to install Redis or whatsoever.
Just Run the DB server and use it to fiddle with your memory.
Note : this DB is not persistent, Don't use it in production.

### Prerequisites

(https://golang.org/doc/install) - How to get started with Go Lang

```
Give examples
```

### Installing

A step by step series of examples that tell you how to get a development env running

Say what the step will be

```
git clone 
```

And 

```
go run in-memory-db.go --port=8080 (default is 8080)
```
OR
```
go build in-memory-db.go
```
then 
```
./in-memory-db --port=8080 (default is 8080)
```


then through Telnet or any TCP connection
```
telnet 127.0.0.1 8080
```


### Usage
to SET data 
```
SET key value
```

to GET data
```
GET key
```

to check if key isset
```
ISSET key
```

to dump data to screen
```
DUMP
```

to Delete
```
DEL key
```

to clear memory
```
CLEAR
```

to close the connection
```
BYE
```
and Help
```
HELP
```

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
