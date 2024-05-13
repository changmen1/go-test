package main

// Hash表 哈希表是一种键值对存储结构，在redis中，它允许一个键关联多个键值对
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

	defer c.Close()
	// Hset 表示将指定哈希表中的字段设置为指定值
	// books 哈希表的名称，即要操作的键名
	// abc 字段的名称，既要设置的键名
	// 100 字段的值，既要设置的键值
	_, err = c.Do("HSet", "books", "abc", 100)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Hget 表示获取指定哈希表中的指定字段的值 books哈希表名称，即要操作的键名 abc字段的名称，既要获取值的键名
	r, err := redis.Int(c.Do("HGet", "books", "abc"))
	if err != nil {
		fmt.Println("get abc failed,", err)
		return
	}

	fmt.Println(r)
}
