package data

import (
	"encoding/json"
	"io"
	"time"
)

type Post struct {
	ID int `json:"id"`
	PostedBy string `json:"postedby"`
	Timestamp time.Time `json:"timestamp"`
	Description string `json:"description"`
	Comments `json:"comments"`
	Location `json:"location"`
	Medias `json:"medias"`
	Likes []string `json:"likes"`
	Dislikes []string `json:"dislikes"`
}

func (p *Post) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func (p *Posts) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

type Posts []*Post

func GetPosts() Posts {
	return postList
}

var postList = []*Post{

	{
		ID: 1,
		PostedBy: "lucyxz",
		Timestamp: time.Now(),
		Description: "the description",
		Comments:
			[]*Comment{
				{
					ID: 1,
					PostedBy: "lucyxz",
					Text: "some text here",
					Timestamp: time.Now(),
				},
			},
		Location: Location{
			ID: 1,
			Country: "Serbia",
			City: "Novi Sad",
		},
		Medias: []*Media{
			{
				ID: 1,
				Type: image,
				Path: "some_path",
			},
		},
		Likes: []string{"like1"},
		Dislikes: []string{"dislike1"},
	},
}