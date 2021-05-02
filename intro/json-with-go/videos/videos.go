package main

import (
	"io/ioutil"
	"encoding/json"
)

type video struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	ImageUrl string `json: "imageUrl"`
	Url string `json:"url"`
}

func getVideos()(videos []video) {
	
	fileBytes, err := ioutil.ReadFile("./videos.json")

	if err != nil {
		panic(err) // something went wrong
	}

	err = json.Unmarshal(fileBytes, &videos)

	if err != nil {
		panic(err) // something went wrong
	}

	return videos
}

func saveVideos(videos []video) {
	videoBytes, err := json.Marshal(videos)

	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("./videos-updated.json", videoBytes, 0644)

	if err != nil {
		panic(err)
	}
}