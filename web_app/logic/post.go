package logic

import (
	"web_app/dao/mysql"
	"web_app/dao/redis"
	"web_app/models"
	"web_app/pkg/snowflake"

	"go.uber.org/zap"
)

func CreatePost(p *models.Post) (err error) {
	p.ID = snowflake.GenID()
	err = mysql.CreatePost(p)
	if err != nil {
		return
	}
	err = redis.CreatePost(p.ID, p.CommunityID)
	return

}

func GetPostById(pid int64) (data *models.ApiPostDetail, err error) {
	return mysql.GetPostById(pid)
}

func GetPostList(page, size int64) (posts []*models.Post, err error) {
	return mysql.GetPostList(page, size)
}

func GetPostList2(p *models.ParamPostList) (posts []*models.Post, err error){
	ids, err := redis.GetPostIDsInOrder(p)
	if err != nil {
		return
	}

	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostIDsInOrder(p) return 0 data")
		return
	}
	zap.L().Debug("GetPostList2", zap.Any("ids", ids))
	posts, err = mysql.GetPostListByIDs(ids)
	if err != nil {
		zap.L().Error("mysql.GetPostListByIDs(ids) error", zap.Error(err))
		return
	}
	zap.L().Debug("GetPostList2", zap.Any("posts", posts))
	
	return
}
