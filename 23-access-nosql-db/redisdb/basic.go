package main

import (
	"23-access-nosql-db/session"
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
	"time"
)

func keyValue() {
	client := session.GetRedisClient()
	defer client.Close()

	client.Set("one", 1, 0)
	val, _ := client.Get("one").Result()
	fmt.Println(val)

	oldVal, _ := client.GetSet("one", "cos(0)").Result()
	fmt.Println(oldVal)
	val, _ = client.Get("one").Result()
	fmt.Println(val)

	client.Set("two", 2, 0)
	client.MSet("three", 3, "four", 4, "five", 5)
	vals, _ := client.MGet("one", "two", "three", "four", "five").Result()
	fmt.Println(vals)

	client.SetNX("once", "atomic", 0).Err()
	did, _ := client.SetNX("once", "atomic", 0).Result()
	fmt.Println(did)
}

func incrDecr() {
	client := session.GetRedisClient()
	defer client.Close()

	client.Set("counter", 0, 0)
	client.Incr("counter")
	client.IncrBy("counter", 2)

	counter, _ := client.Get("counter").Result()
	fmt.Println(counter)

	client.IncrByFloat("counter", 0.2)
	counter, _ = client.Get("counter").Result()
	fmt.Println(counter)

	client.Decr("counter")
	client.DecrBy("counter", 2)
	counter, _ = client.Get("counter").Result()
	fmt.Println(counter)

	client.Expire("counter", 0) // 立刻过期
	exists, _ := client.Exists("counter").Result()
	fmt.Println(exists)
}

func hash() {
	client := session.GetRedisClient()
	defer client.Close()

	client.HSet("user", "username", "zhangsan")
	client.HMSet("user", map[string]interface{}{
		"age": 20,
		"job": "driver",
	})

	age, _ := client.HGet("user", "age").Result()
	fmt.Println(age)
	fields, _ := client.HMGet("user", "username", "job").Result()
	fmt.Println(fields)
}

func list() {
	client := session.GetRedisClient()
	defer client.Close()

	client.RPush("digits", 1, 2, 3)
	number, _ := client.LPop("digits").Result()
	fmt.Println(number)
}

func set() {
	client := session.GetRedisClient()
	defer client.Close()

	client.SAdd("scores", 89, 98, 82)
	all, _ := client.SMembers("scores").Result()
	fmt.Println(all)
}

func zset() {
	client := session.GetRedisClient()
	defer client.Close()

	client.ZAdd("events", redis.Z{Score: float64(time.Now().Unix()), Member: "create"})

	time.Sleep(1 * time.Second)
	t1 := time.Now().Unix()

	client.ZAdd("events", redis.Z{Score: float64(time.Now().Unix()), Member: "login"})
	time.Sleep(1 * time.Second)
	client.ZAdd("events", redis.Z{Score: float64(time.Now().Unix()), Member: "buy"})
	time.Sleep(1 * time.Second)

	t2 := time.Now().Unix()
	time.Sleep(1 * time.Second)

	client.ZAdd("events", redis.Z{Score: float64(time.Now().Unix()), Member: "exit"})

	all, _ := client.ZRange("events", 0, time.Now().Unix()).Result()
	fmt.Println(all)

	fmt.Println(t1, t2)
	all, _ = client.ZRangeByScore("events", redis.ZRangeBy{Min: fmt.Sprintf("%d", t1), Max: fmt.Sprintf("%d", t2)}).Result()
	fmt.Println(all)

	count, _ := client.ZCard("events").Result()
	fmt.Println(count)

	count, _ = client.ZCount("events", "0", fmt.Sprintf("%d", t2)).Result()
	fmt.Println(count)

	rank, _ := client.ZRank("events", "buy").Result()
	fmt.Println(rank)
}

func pubsub() {
	client := session.GetRedisClient()
	defer client.Close()

	sub := client.Subscribe("channel")
	go func() {
		time.Sleep(1 * time.Second)
		client.Publish("channel", "a message")
		time.Sleep(1 * time.Second)
		sub.Close()
	}()

	for message := range sub.Channel() {
		fmt.Println(message.Channel, message.Payload)
	}
}

func pipe() {
	client := session.GetRedisClient()
	defer client.Close()

	// MULTI-DISCARD-EXEC 整体提交
	// MULTI
	pipe := client.TxPipeline()
	pipe.Set("counter", 1, 0)
	pipe.IncrBy("counter", 1)
	pipe.Expire("counter", 10*time.Second)
	_, err := pipe.Exec()
	// EXEC

	if err != nil {
		fmt.Println(err)
	}
}

func watch() {
	client := session.GetRedisClient()
	defer client.Close()

	// WATCH-CAS 监测如改动则 EXEC 放弃
	doFunc := func(tx *redis.Tx) error {
		// MULTI
		flag, _ := tx.Get("flag").Result()
		val, _ := strconv.ParseInt(flag, 10, 0)

		//pipe := tx.Pipeline()
		//pipe.Set("flag", val+1, 0)
		//pipe.Exec()

		_, err := tx.Pipelined(func(pipe redis.Pipeliner) error {
			pipe.Set("flag", val+1, 0)
			return nil
		})
		// EXEC

		return err
	}

	//client.Watch(doFunc, "flag1", "flag2", "flag3")
	client.Watch(doFunc, "flag")
}

func main() {
	keyValue()
	incrDecr()
	hash()
	list()
	set()
	zset()
	pubsub()
	pipe()
	watch()
}
