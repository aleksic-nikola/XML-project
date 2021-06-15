package service

import (
	"xml/content-service/data"
	"xml/content-service/repository"
)

type StoryService struct {
	Repo *repository.StoryRepository
}

func (service *StoryService) CreateStory(story *data.Story) error {
	error := service.Repo.CreateStory(story)
	return error
}

func (service *StoryService) StoryExists(id uint) (bool, error) {

	exists := service.Repo.StoryExists(id)
	return exists, nil
}

func (service *StoryService) GetAllStoriesForUser(username string) (data.Stories) {

	stories := service.Repo.GetAllStoriesForUser(username)
	return stories
}
