### Installation
```sh
go get github.com/viniciosanacleto/gonimbus
```

### How to use
See `example.go`:
```go
// Create a new cache system
cacheSystem := make(cache.CacheSystem)

// Set the key "foo" in the cache system with the value "bar" expiring in 2 seconds
cacheSystem.Set("foo", "bar", 2)

// Get the value of key "foo" in cache system
value := cacheSystem.Get("foo")
fmt.Println(value) // "bar"

time.Sleep(3 * time.Second)

// Get again the value of the key "foo" after sleep for 3 seconds
value = cacheSystem.Get("foo")
fmt.Println(value) // <nil>
```
