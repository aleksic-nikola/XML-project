package service

import (
	"fmt"
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
	fmt.Println(username)
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


func (service *PostService) LikePost(id string, username string) error {

	post := service.Repo.GetPostByID(id)

	for _, user := range post.Likes {
		if user.Username == username {
			//fmt.Println("Already likes this post")
			//return fmt.Errorf("Already likes this post")

			err := service.unlikePost(id, username)
			return err

		}
	}

	post.Likes = append(post.Likes, data.User{Username: username})
	err := service.Repo.SavePost(post)

	err = service.undislikePost(id, username)
	return err
}

func (service *PostService) unlikePost(id string, username string) error {

	post := service.Repo.GetPostByID(id)
	var new_likes []data.User
	for _, like := range post.Likes {
		if username != like.Username {
			new_likes = append(new_likes, like)
		}
	}

	err := service.Repo.ClearLikes(post)

	if err != nil {
		return err
	}
	post.Likes = new_likes
	err = service.Repo.SavePost(post)

	if err != nil {
		return err
	}

	return nil
}

func (service *PostService) DislikePost(id string, username string) error {

	post := service.Repo.GetPostByID(id)

	for _, user := range post.Dislikes {
		if user.Username == username {
			err := service.undislikePost(id, username)
			return err
		}
	}

	post.Dislikes = append(post.Dislikes, data.User{Username: username})
	err := service.Repo.SavePost(post)
	err = service.unlikePost(id, username)
	return err
}

func (service *PostService) undislikePost(id string, username string) error {

	post := service.Repo.GetPostByID(id)
	var new_dislikes []data.User
	for _, dislike := range post.Dislikes {
		if username != dislike.Username {
			new_dislikes = append(new_dislikes, dislike)
		}
	}

	err := service.Repo.ClearDislikes(post)

	if err != nil {
		return err
	}
	post.Dislikes = new_dislikes
	err = service.Repo.SavePost(post)

	if err != nil {
		return err
	}

	return nil
}

func (service *PostService) GetPostsByIds(ids dtos.PostIdsDto) *data.Posts {

	var posts data.Posts
	for _, i := range ids.Ids {
		post := service.Repo.GetPostById(i.Id)
		posts = append(posts, &post)
	}

	return &posts
}

func (service *PostService) PostComment(id string, text string, username string) error {

	post := service.Repo.GetPostByID(id)
	post.Comments = append(post.Comments, data.Comment{Text : text, PostedBy: username})


	err := service.Repo.SavePost(post)

	return err

}

func (service *PostService) GetUserByPost(postid string) dtos.UsernameDto {
	post := service.Repo.GetPostByID(postid)

	var user dtos.UsernameDto
	user.Username = post.PostedBy

	return user
}

