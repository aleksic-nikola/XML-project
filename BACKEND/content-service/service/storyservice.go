package service

import (
	"time"
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

func (service *StoryService) GetAllStoriesForUser(username string) data.Stories {

	stories := service.Repo.GetAllStoriesForUser(username)
	today := time.Now()
	var retList data.Stories
	for _, oneStory := range stories{
		if oneStory.CreatedAt.Add(24*time.Hour).After(today){
			retList = append(retList, oneStory)
		}
	}

	return retList
}

func (service *StoryService) GetAllStoriesForFeed(usernames []string) []data.Story {

	stories:= service.Repo.GetAllStoriesForFeed(usernames)

	today := time.Now()
	var retList []data.Story
	for _, oneStory := range stories{
		if oneStory.CreatedAt.Add(24*time.Hour).After(today){
			retList = append(retList, oneStory)
		}
	}

	return retList
}

func (service *StoryService) SetStoryHighlightedOn(id int) error {

	err:= service.Repo.SetStoryHighlightedOn(id)
	return err

}

func (service *StoryService) SetStoryHighlightedOff(id int) error {
	err:= service.Repo.SetStoryHighlightedOff(id)
	return err
}

func (service *StoryService) GetAllArchiveStories(username string) data.Stories {
	stories := service.Repo.GetAllStoriesForUser(username)

	return stories
}
