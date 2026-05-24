package redis

import (
	"web_app/models"

	"github.com/go-redis/redis"
)

func GetPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}
	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1
	return client.ZRevRange(key, start, end).Result()
}

// GetPostVoteData 根据 ids 查询每篇帖子投赞成票的数据
func GetPostVoteData(ids []string) (data []int64, err error) {
	data = make([]int64, 0, len(ids))
	//for _, id := range ids {
	//	key := getRedisKey(KeyPostVotedZSetPrefix + id)
	//	// 查找 key 中分数是 1 的元素的数量 --> 统计每篇帖子的赞成票的数量
	//	v := client.ZCount(key, "1", "1").Val()
	//	data = append(data, v)
	//}
	// 使用 pipeline 一次发送多条命令，减少 RTT
	pipeline := client.TxPipeline()
	for _, id := range ids {
		key := getRedisKey(KeyPostVotedZSetPrefix + id)
		// 查找 key 中分数是 1 的元素的数量 --> 统计每篇帖子的赞成票的数量
		pipeline.ZCount(key, "1", "1")
	}
	cmderData, err := pipeline.Exec()
	if err != nil {
		return nil, err
	}
	for _, cmder := range cmderData {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}
	return
}
