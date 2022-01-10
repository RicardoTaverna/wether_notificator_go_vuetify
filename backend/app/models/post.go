package models

type Post struct {
	ID      int64  `db:"id"`
	Title   string `db:"title"`
	Content string `db:"content"`
	Author  string `db:"author"`
}

type JsonPost struct {
	ID      int64  `db:"id"`
	Title   string `db:"title"`
	Content string `db:"content"`
	Author  string `db:"author"`
}

type PostRequest struct {
	Title   string `db:"title"`
	Content string `db:"content"`
	Author  string `db:"author"`
}
