package models


const (
	OrderTime = "time"
	OrderScore = "score"
)

type SignupUser struct {
	Username        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
}

type ParamLogin struct {
	Username        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required"`
}

type ParamsVoteData struct {
	PostID int64 `json:"post_id,string" binding:"required"`
	Direction int8 `json:"direction,string" binding:"oneof=1 0 -1"`
}

type ParamPostList struct {
	Page int64 `json:"page" form:"page"`
	Size int64 `json:"size" form:"size"`
	Order string `json:"order" form:"order"`
}