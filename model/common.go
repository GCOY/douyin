package model

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	Id            int64  `json:"id,omitempty"`
	Author        User   `json:"author"`
	AuthorID      int64   `json:"author_id"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
	PublishTime   int64	 `json:"publish_time,omitempty"`
}

type Comment struct {
	Id         int64  `json:"id,omitempty"`
	User       User   `json:"user"`
	UserID	   int64  `json:"user_id"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
	VideoID    int64  `json:"video_id,omitempty"`
}

type User struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

type LoginInfo struct {
	ID int64
	UserName string
	PassWord string
}

type FavoriteVideo struct {
	Id            		int64 `json:"id,omitempty"`
	UserID          	int64 `json:"user_id,omitempty"`
	VideoID   			int64 `json:"video_id,omitempty"`
	FavoriteStatus      bool  `json:"favorite_status,omitempty"`
}
type Follow struct {
	Id  int64 `json:"id,omitempty"`
	FollowID int64 `json:"follow_user,omitempty"`
	FollowerID int64 `json:"follower_user,omitempty"`
	FollowStatus bool `json:"follow_status,omitempty"`
}

type Token struct {
	Id int64
	UserToken string
	UserID int64
}