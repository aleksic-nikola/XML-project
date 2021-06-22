package service

import (
	"xml/content-service/data"
	"xml/content-service/data/dtos"
	"xml/content-service/repository"
)

type PostService struct {
	Repo *repository.PostRepository
}

func (service *PostService) CreatePost(post *data.Post) error {
	error := service.Repo.CreatePost(post)
	return error
}

func (service *PostService) PostExists(id uint) (bool, error) {

	exists := service.Repo.PostExists(id)
	return exists, nil
}

func (service *PostService) GetPostsByUser(username string) data.Posts {
	posts := service.Repo.GetPostsByUser(username)
	return posts
}

func (service *PostService) GetAllPostsForUser(username string) (data.Posts) {

	posts := service.Repo.GetAllPostsForUser(username)
	return posts
}

func (service *PostService) GetLikedPostsByUser(username string) data.Posts {
	posts := service.Repo.GetAllPosts()
	likedPosts := data.Posts{}
	for _,post := range posts {
		for _,likedby := range post.Likes {
			if likedby.Username == username {
				likedPosts = append(likedPosts, post)
			}
		}
	}
	return likedPosts
}

func (service *PostService) GetDislikedPostsByUser(username string) data.Posts {
	posts := service.Repo.GetAllPosts()
	dislikedPosts := data.Posts{}
	for _,post := range posts {
		for _,dislikedby := range post.Dislikes {
			if dislikedby.Username == username {
				dislikedPosts = append(dislikedPosts, post)
			}
		}
	}
	return dislikedPosts
}


func (service *PostService) GetAllPostsForFeed(usernames []string) []data.Post{
	posts := service.Repo.GetAllPostsForFeed(usernames)
	return posts
}

func (service *PostService) GetPostsByIds(ids dtos.PostIdsDto) *data.Posts {

	var posts data.Posts
	for _, i := range ids.Ids {
		post := service.Repo.GetPostById(i.Id)
		posts = append(posts, &post)
	}

	return &posts
}