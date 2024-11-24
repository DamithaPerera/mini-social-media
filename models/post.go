package models

import "sync"

type Post struct {
	ID       string   `json:"id"`
	Content  string   `json:"content"`
	Likes    int      `json:"likes"`
	Comments []string `json:"comments"`
}

var (
	Posts     = make(map[string]*Post) // In-memory data
	PostsLock = sync.RWMutex{}         // Concurrency handling
)
