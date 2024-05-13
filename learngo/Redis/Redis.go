package main

// 连接Redis
import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}
	_, err = c.Do("Set", "zxxx", 123)
	r, err := redis.Int(c.Do("Get", "zxxx"))
	fmt.Println("redis", r)
	fmt.Println("redis conn success")
	defer c.Close()
}
