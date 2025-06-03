# sync.Pool

## sync.Pool Method

**`pool.New`**

```go
// create pool instance
var pool sync.pool
// sync.pool constructor method
pool.New = func() any {
    return &instance{
        body: []byte(
            {
                "key": "value",
            }
        ),
    }
}
```