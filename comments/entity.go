package comments

import "time"

type CurrentUser struct {
	ID       int
	Username string
	Image    string
}

type Reply struct {
	ID         int
	Username   string
	Content    string
	CreatedAt  time.Time
	ReplyingTo string
	Score      int
	Image      string
}

type Comment struct {
	ID        int
	Content   string
	CreatedAt time.Time
	Score     int
	Username  string
	Image     string
}

type CtrCurrentUser struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Image    string `json:"image"`
}

type CtrUser struct {
	Image    string `json:"image"`
	Username string `json:"username"`
}

type CtrReply struct {
	ID         int       `json:"id"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
	ReplyingTo string    `json:"replying_to"`
	Score      int       `json:"score"`
	Username   string    `json:"username"`
	Image      string    `json:"image"`
}

type CtrComment struct {
	ID        int        `json:"id"`
	Content   string     `json:"content"`
	CreatedAt time.Time  `json:"created_at"`
	Score     int        `json:"score"`
	Username  string     `json:"username"`
	Image     string     `json:"image"`
	Replies   []CtrReply `json:"replies"`
}

type CtrData struct {
	CurrentUser CtrCurrentUser `json:"current_user"`
	Comments    []CtrComment   `json:"comments"`
}

type CtrPostComment struct {
	Content  string `json:"content" validate:"required"`
	Score    int    `json:"score"`
	Username string `json:"username" validate:"required"`
}

type CtrPostReply struct {
	Content    string `json:"content" validate:"required"`
	Score      int    `json:"score"`
	ReplyingTo string `json:"replying_to" validate:"required"`
	Username   string `json:"username" validate:"required"`
}

type CtrEditComment struct {
	Content string `json:"content" validate:"required"`
}

type CtrEditReply struct {
	Content string `json:"content" validate:"required"`
}

type CtrEditScore struct {
	Score int `json:"score" validate:"required"`
}
