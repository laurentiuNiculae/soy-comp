package main

import (
	"test/pkg"
)

func main() {
	videoSource := []pkg.VideoSource{
		{URL: "https://www.twitch.tv/videos/2338102814", Interval: "01:04:54-01:06:38"},
		{URL: "https://www.twitch.tv/videos/2338102814", Interval: "01:06:56-01:07:08"},
	}

	videoCompiler := pkg.NewVideoCompiler(pkg.VideoCompilerConfig{
		CacheFolder:     "./cached-videos2",
		CacheFileType:   "mp4",
		MaxVideoQuality: "1080",
	})

	err := videoCompiler.AddVideoSources(videoSource...)
	if err != nil {
		panic(err)
	}

	_, err = videoCompiler.MergeVideos("finalVideo.mp4")
	if err != nil {
		panic(err)
	}
}
