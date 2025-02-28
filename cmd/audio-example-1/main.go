package main

import (
	"test/pkg"
)

func main() {
	videoSource := struct {
		URL      string
		Interval string
	}{
		URL:      "https://www.youtube.com/watch?v=OFdi_61HnFk",
		Interval: "00:00:27-00:00:41",
	}

	fact := pkg.NewVideoResourceFactory(pkg.VideoCompilerConfig{
		CacheFolder:     "./audio-example-cache",
		CacheFileType:   "mp4",
		MaxVideoQuality: "1080",
	})

	videoResource, err := fact.NewVideoResource(videoSource.URL, videoSource.Interval)
	if err != nil {
		panic(err)
	}

	_, err = pkg.ExtractAudioFromVideo(videoResource, "yes.mp3", "00:00:01-00:00:10")
	if err != nil {
		panic(err)
	}
}
