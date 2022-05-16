package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

const (
	addr string = "192.168.126.129:6379"
)

var ctx context.Context
var client redis.UniversalClient

func init() {

	client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	ctx = context.Background()

}

func main() {

	setValue(200000, "10_bytes", generateValue(10))
	setValue(100000, "20_bytes", generateValue(20))
	setValue(100000, "50_bytes", generateValue(50))

	setValue(50000, "100_bytes", generateValue(100))
	setValue(50000, "200_bytes", generateValue(200))

	setValue(20000, "1000_bytes", generateValue(1000))
	setValue(10000, "5000_bytes", generateValue(5000))

}

func setValue(num int, key, value string) {
	for i := 0; i < num; i++ {
		k := fmt.Sprintf("%s_%d", key, i)
		err := client.Set(ctx, k, value, 0).Err()
		if err != nil {
			panic(err)
		}
	}
}

func generateValue(size int) string {
	arr := make([]byte, size)
	for i := 0; i < size; i++ {
		arr[i] = 'x'
	}
	return string(arr)
}
