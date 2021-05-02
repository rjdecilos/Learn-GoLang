package main

import (
	"fmt"
)

func main() {
	videos := getVideos()

	fmt.Println(videos)

	for x, video := range videos {
		videos[x].Description = video.Description + "\n Remember to Like and Subscribe!"
	}

	fmt.Println(videos)

	saveVideos(videos)
}