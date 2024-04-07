package utils

import (
	"fmt"
	"github.com/go-redis/redis"
	"go-blog/config"
	"strconv"
)

func NewClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Cfg.Redis.Addr,
		Password: config.Cfg.Redis.Password,
		DB:       config.Cfg.Redis.DataBase,
	})

	// 通过 client.Ping() 来检查是否成功连接到了 redis 服务器
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	return client
}

func AddViews(client *redis.Client, articleId int) {
	id := strconv.Itoa(articleId)
	_, err := client.Incr(id).Result()
	if err != nil {
		panic(err)
	}
}

func GetViews(client *redis.Client, articleId int) int {
	id := strconv.Itoa(articleId)
	result, _ := client.Exists(id).Result()
	if result != 0 {
		val, err := client.Get(id).Result()
		if err != nil {
			panic(err)
		}

		views, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		return views
	} else {
		return 0
	}

}

func CountString(client *redis.Client, userip string) error {
	val, err := client.Exists(userip).Result()
	if err != nil {
		return err
	}

	if val == 0 { // no exists
		_, err := client.Incr(userip).Result()
		if err != nil {
			return err
		}

		_, err = client.Expire(userip, 5*1000*1000*1000*10).Result() //设置键的过期时间  10s
		if err != nil {
			return err
		}

	} else { //exist
		count, err := client.Get(userip).Result()
		if err != nil {
			return err
		}

		temp, _ := strconv.Atoi(count)
		fmt.Println(temp)

		if temp < 100 {
			_, err = client.Incr(userip).Result()
			if err != nil {
				return err
			}
		} else {
			return fmt.Errorf("访问次数受限")
		}

	}
	return nil
}
