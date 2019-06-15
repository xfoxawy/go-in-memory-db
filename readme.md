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
go run go-in-memory-db.go --port=8080 (default is 8080)
```
OR
```
go build go-in-memory-db.go
```
then 
```
./go-in-memory-db --port=8080 (default is 8080)
```

to run tests

```
go test -cover ./...
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

to set linked list
```
Lset key
```

to get data from linked list
```
LGET key
```

to delete linked list
```
LDEL key
```

to push in linked list
```
LPUSH key value1 value2 etc...
```

to remove last element in linked list 
```
LPOP key
```

to push in left of linked list 
```
LSHIFT key value
```

to remove from left of linked list 
```
LUNSHIFT key
```

to remove value in linked list 
```
LRM key value1 value2 etc...
OR
LREMOVE key value1 value2 etc...
```

to remove values in linked list using (index) int
```
LUNLINK key index1 index2 etc...
```

to get value using (index) int
```
LSEEK key index
```
to set hashtable
 ```
 HSET key 
 OR 
 HSET key value1 value2 etc...
 ```

 to get data from hashtable
 ```
 HGET key
 ```

 to delete hashtable
 ```
 HDEL key
 ```

 to push element in hashtable 
 ```
 HPUSH key value1 value2 etc...
 ```

 to remove element from hashtable
 ```
 HRM key value1 value2 etc...
 OR
 HREMOVE key value1 value2 etc...
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

supporting namespacing, you can switch databases by using
```
USE key
```
to switch back to `master` database simply run
```
USE master
```

to list all open databases
```
SHOW 
```

and to get which is the current database
```
WHICH
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
