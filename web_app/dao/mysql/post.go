package mysql

import (
	"strings"
	"web_app/models"

	"github.com/jmoiron/sqlx"
)

func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post
	(post_id, title, content, author_id, community_id)
	values (?, ?, ?, ?, ?)
	`
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return
}

func GetPostById(pid int64) (data *models.ApiPostDetail, err error) {
	data = new(models.ApiPostDetail)
	sqlStr := `select c.community_name community_name,
	p.title title,p.content content,u.username username from community c 
inner join post p on c.community_id = p.community_id left join user u 
on p.author_id  = u.user_id where p.post_id = ?`
	err = db.Get(data, sqlStr, pid)
	return
}

func GetPostList(page, size int64) (posts []*models.Post, err error) {

	sqlStr := `select post_id, title, content, author_id, community_id, create_time
	from post 
	order by create_time desc 
	limit ? offset ?`
	posts = make([]*models.Post, 0, 2)
	err = db.Select(&posts, sqlStr, size, (page-1)*size)
	return
}

func GetPostListByIDs(ids []string) (postList []*models.Post, err error) {
	sqlStr := `select post_id, title, content, author_id, community_id, create_time
	from post where post_id in (?)
	order by FIND_IN_SET(post_id, ?)`

	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}

	query = db.Rebind(query)

	err = db.Select(&postList, query, args...)
	return
}
