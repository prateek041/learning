Go has package level variables. These are instantiated before anything has the change to get moving. Therefore you get singleton patter by default.

```go
package somepack

var (
connection = createConn()
)

func Connection() SomeConnection {
	return connection
}
```

`connection` will be created once and therefore `Connection()` will return the same instance of it safely.

The other case would be when resources are limited and we want singleton which are "lazy". That is when we use[[Once# | sync.Once]].

```go
var (
connection SomeConnection
connectionOnce sync.Once
)

func Connection()someConnection{
	connectionOnce.Do(func(){
		connection = createConn()	
	})
}
```

this makes sure that `createConn` is only called once per instance of connectionOnce.