# how to setup
```
$ cd /path/to/golang-gob
$ ./run-redis.sh
$ go run .
&{1 Hello World}
```

```
1708606994.130213 [0 172.17.0.1:58816] "client" "setinfo" "LIB-NAME" "go-redis(,go1.22.0)"
1708606994.130263 [0 172.17.0.1:58816] "client" "setinfo" "LIB-VER" "9.5.1"
1708606994.134215 [0 172.17.0.1:58816] "set" "key" ")\x7f\x03\x01\x01\tCacheData\x01\xff\x80\x00\x01\x02\x01\x02ID\x01\x0c\x00\x01\aMessage\x01\x0c\x00\x00\x00\x13\xff\x80\x01\x011\x01\x0bHello World\x00" "ex" "60"
1708606994.137371 [0 172.17.0.1:58816] "get" "key"
```

# purpose
I want to share how to implement cache client to hide a specific encode/decode methods with my teammate.
