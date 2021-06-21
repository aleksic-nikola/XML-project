package repository

import (
	"fmt"
	"xml/content-service/data"

	"gorm.io/gorm"
)

type StoryRepository struct {
	Database *gorm.DB
}

func (repo *StoryRepository) CreateStory(story *data.Story) error {
	result := repo.Database.Create(story)
	//TODO convert to logs
	fmt.Println(result.RowsAffected)
	return result.Error
}

func (repo *StoryRepository) StoryExists(id uint) bool {
	var count int64
	repo.Database.Where("id = ?", id).Find(&data.Story{}).Count(&count)
	return count != 0
}

func (repo *StoryRepository) GetAllStoriesForUser(username string) (data.Stories) {
	var stories data.Stories
	repo.Database.Where("posted_by = ?", username).Find(&stories)
	return stories
}

func (repo *StoryRepository) GetAllStoriesForFeed(usernames []string) []data.Story {
	var stories []data.Story
	//repo.Database.Where("posted_by IN ?", usernames).Find(&posts)

	//repo.Database.Preload("Medias", "posted_by IN ?", usernames).Find(&posts)
	repo.Database.Preload("Media").Where("posted_by IN ?", usernames).Find(&stories)
	//repo.Database.Preload("Followers").Find(&profile, id)

	//db.Preload("Orders").Preload("Profile").Preload("Role").Find(&users)

	fmt.Println("DOBILII POSTS: ")
	fmt.Println(stories)
	fmt.Println("\n\n******************************************************* ")


	return stories
}
