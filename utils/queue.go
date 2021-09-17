package utils

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/spf13/cast"
)

type RedisQueueDelay struct {
	client   *redis.Client
	interval uint64
	number   uint64
}
type QueueMsg struct {
	Topic   string `json:"topic"`
	Content []byte `json:"content"`
}

// NewRedisQueueDelay 创建一个新的队列
// client redis连接, interval 消费端间隔，单位毫秒, number 每次消费数目，结果小于等于此数字
func NewRedisQueueDelay(client *redis.Client, interval, number uint64) *RedisQueueDelay {
	return &RedisQueueDelay{
		client:   client,
		interval: interval,
		number:   number,
	}
}

func (q *RedisQueueDelay) Produce(msg QueueMsg, t time.Time) error {
	redisKey := fmt.Sprintf("queue:delay:%v", msg.Topic)
	member := redis.Z{
		Member: msg,
		Score:  cast.ToFloat64(t.UnixNano()),
	}
	return q.client.ZAdd(redisKey, member).Err()
}

func (q *RedisQueueDelay) Consume(topic string, f func(msg QueueMsg) error) {
	ticker := time.Tick(time.Millisecond * time.Duration(q.interval))
	redisKey := fmt.Sprintf("queue:delay:%v", topic)
	number := int64(q.number)
	if number > 0 {
		number = number - 1
	}
	for {
		select {
		case <-ticker:
			var members []redis.Z
			if number > 0 {
				members = q.client.ZRangeWithScores(redisKey, 0, number).Val()
			} else {
				members = q.client.ZRangeByScoreWithScores(redisKey, redis.ZRangeBy{
					Min: "0",
					Max: cast.ToString(time.Now().UnixNano()),
				}).Val()
			}

			for _, v := range members {
				if v.Score > cast.ToFloat64(time.Now().UnixNano()) {
					continue
				}
				msg, ok := v.Member.(QueueMsg)
				if !ok {
					continue
				}

				if err := f(msg); err == nil {
					q.client.ZRem(redisKey, v.Member)
				}
			}
		}
	}
}
