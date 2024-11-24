package services

import (
    "errors"
    "mini-social-media/models"
    "github.com/google/uuid"
)


func CreatePost(content string) (*models.Post, error) {
	models.PostsLock.Lock()
	defer models.PostsLock.Unlock()

	id := uuid.New().String()
	post := &models.Post{
		ID:       id,
		Content:  content,
		Likes:    0,
		Comments: []string{},
	}

	models.Posts[id] = post
	return post, nil
}

func UpdatePost(id, content string) (*models.Post, error) {
	models.PostsLock.Lock()
	defer models.PostsLock.Unlock()

	post, exists := models.Posts[id]
	if !exists {
		return nil, errors.New("post not found")
	}

	post.Content = content
	return post, nil
}

func GetAllPosts() []*models.Post {
	models.PostsLock.RLock()
	defer models.PostsLock.RUnlock()

	posts := []*models.Post{}
	for _, post := range models.Posts {
		posts = append(posts, post)
	}
	return posts
}

func LikePost(id string) (*models.Post, error) {
	models.PostsLock.Lock()
	defer models.PostsLock.Unlock()

	post, exists := models.Posts[id]
	if !exists {
		return nil, errors.New("post not found")
	}

	post.Likes++
	return post, nil
}

func AddComment(id, comment string) (*models.Post, error) {
	models.PostsLock.Lock()
	defer models.PostsLock.Unlock()

	post, exists := models.Posts[id]
	if !exists {
		return nil, errors.New("post not found")
	}

	post.Comments = append(post.Comments, comment)
	return post, nil
}

func GetPostDetails(id string) (*models.Post, error) {
	models.PostsLock.RLock()
	defer models.PostsLock.RUnlock()

	post, exists := models.Posts[id]
	if !exists {
		return nil, errors.New("post not found")
	}

	return post, nil
}
