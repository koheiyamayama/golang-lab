package main

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type CacheData struct {
	ID      string
	Message string
}

type CacheClient struct {
	redis *redis.Client
}

func (c *CacheClient) Set(key string, value interface{}) error {
	buf := bytes.NewBuffer(nil)
	if err := gob.NewEncoder(buf).Encode(value); err != nil {
		return err
	}

	return c.redis.Set(context.TODO(), key, buf.Bytes(), 1*time.Minute).Err()
}

func (c *CacheClient) Get(key string, value interface{}) error {
	buf, err := c.redis.Get(context.TODO(), key).Bytes()
	if err != nil {
		return err
	}

	return gob.NewDecoder(bytes.NewBuffer(buf)).Decode(value)
}

func main() {
	setData := &CacheData{
		ID:      "1",
		Message: "Hello World",
	}

	client := &CacheClient{
		redis: redis.NewClient(&redis.Options{}),
	}

	if err := client.Set("key", setData); err != nil {
		panic(err)
	}

	getData := &CacheData{}
	if err := client.Get("key", getData); err != nil {
		panic(err)
	}

	fmt.Println(getData)

}
