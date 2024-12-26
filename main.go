package main

import (
	"log"
	"test/pkg"
)

func main() {
	log.SetFlags(log.Llongfile)

	videoSource := []struct {
		URL      string
		Interval string
	}{
		{URL: "https://youtube.com/clip/UgkxL8ZYx8HKI4NissPiScboPm6d03NI-QME", Interval: ""},
		{URL: "https://www.twitch.tv/tsoding/clip/AmusedBlueBunnyPupper-WoNMXC--rEeHJXz2", Interval: ""},
		{URL: "https://www.twitch.tv/tsoding/clip/PrettiestImportantAardvarkMVGame-1GNH-l7E9nAey2W0", Interval: ""},
	}

	videos := []pkg.VideoResource{}
	for i, videoSource := range videoSource {
		videoResource, err := pkg.NewVideoResource(videoSource.URL, videoSource.Interval)
		if err != nil {
			log.Fatal("Failed to create video resource with index: ", i)
		}

		videos = append(videos, videoResource)
	}

	err := pkg.MergeVideos("./bin/finalVideo.mp4", videos...)
	if err != nil {
		log.Fatal("Failed to merge videos")
	}
}
