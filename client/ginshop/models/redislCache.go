package models

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"gopkg.in/ini.v1"
)

var (
	ctx         = context.Background()
	rdbClient   *redis.Client
	redisEnable bool
)

func init() {
	config, err := ini.Load("./conf/app.ini")
	if err != nil {
		fmt.Printf("无法读取配置文件: %v\n", err)
		os.Exit(1)
	}

	// 获取 Redis 是否启用
	redisSection := config.Section("redis")
	redisEnable, err = redisSection.Key("enable").Bool()
	if err != nil {
		fmt.Printf("无法解析 Redis 启用配置: %v\n", err)
		return
	}

	// 获取 Redis 密码
	password := redisSection.Key("password").String()

	// 获取哨兵配置
	sentinelSection := config.Section("redis-sentinel")
	masterName := sentinelSection.Key("masterName").String()
	sentinels := []string{
		sentinelSection.Key("sentinel1").String(),
		sentinelSection.Key("sentinel2").String(),
		sentinelSection.Key("sentinel3").String(),
	}

	if redisEnable {
		//初始化 Redis 哨兵客户端
		rdbClient = redis.NewFailoverClient(&redis.FailoverOptions{
			MasterName:    masterName,
			SentinelAddrs: sentinels,
			Password:      password,
			DB:            0,
		})
		if err := rdbClient.Ping(ctx).Err(); err != nil {
			fmt.Printf("Redis 哨兵模式连接失败: %v\n", err)
		} else {
			fmt.Println("Redis 哨兵模式连接成功")
		}
	}
}

type CacheDB struct{}

func (c *CacheDB) Set(key string, value interface{}, expiration time.Duration) error {
	if !redisEnable {
		return nil
	}
	v, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("JSON编码失败: %w", err)
	}
	return rdbClient.Set(ctx, key, string(v), expiration).Err()
}

func (c *CacheDB) Get(key string, obj interface{}) (bool, error) {
	if !redisEnable {
		return false, nil
	}
	valueStr, err := rdbClient.Get(ctx, key).Result()
	if err == redis.Nil {
		return false, nil // 键不存在
	} else if err != nil {
		return false, fmt.Errorf("获取缓存失败: %w", err)
	}
	if valueStr == "" || valueStr == "[]" {
		return false, nil
	}
	if err := json.Unmarshal([]byte(valueStr), obj); err != nil {
		return false, fmt.Errorf("JSON解码失败: %w", err)
	}
	return true, nil
}

func (c *CacheDB) FlushAll() error {
	if !redisEnable {
		return nil
	}
	return rdbClient.FlushAll(ctx).Err()
}

var CacheDb = &CacheDB{}
