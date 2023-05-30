package types

type Post struct {
	ID          string `json:"id" db:"id"`
	AuthorId    string `json:"author_id" db:"author_id"`
	Body        string `json:"body" db:"body"`
	DateCreated string `json:"date_created" db:"date_created"`
	DateUpdated string `json:"date_updated" db:"date_updated"`
}

type Rating struct {
	ID          int64  `json:"id" db:"id"`
	PostId      string `json:"post_id" db:"post_id"`
	Rating      int8   `json:"rating" db:"rating"` // 1 or -1
	DateCreated string `json:"date_created" db:"date_created"`
	DateUpdated string `json:"date_updated" db:"date_updated"`
}

type Comment struct {
	ID          string `json:"id" db:"id"`
	PostId      string `json:"post_id" db:"post_id"`
	Body        string `json:"body" db:"body"`
	DateCreated string `json:"date_created" db:"date_created"`
	DateUpdated string `json:"date_updated" db:"date_updated"`
}
