# soy-comp

**soy-comp** is a simple video compilation making library that will download and compile a list of videos

You can fully or partially download videos from Youtube and Twitch links.

# Requirements

You will need:
- `go` compiler v1.23 (I'm confident it works with lower versions too but you'll have to update the go.mod file)
- `yt-dlp` executable in PATH for downloading the videos
- `ffmpeg` for video processing

# Example

There is an example in `cmd/video-example-1/main.go` you can run with

> $ go run cmd/video-example-1/main.go

This is the gist of what it does

```go

func main() {
	url := "https://www.youtube.com/watch?v=PGSba51aRYU"
	interval := "00:23:23-00:23:30"
	// this downloads and cuts and caches your clip locally 	in ./cached-videos
	videoResource1, err := pkg.NewVideoResource(url, videoSource.Interval)
	if err != nil {
		log.Fatal("Failed to create video resource", err)
	}

	url = "https://www.youtube.com/clip/	UgkxL8ZYx8HKI4NissPiScboPm6d03NI-QME"
	interval = ""
	videoResource2, err := pkg.NewVideoResource(url, videoSource.Interval)
	if err != nil {
		log.Fatal("Failed to create video resource", err)
	}

	url = "https://www.twitch.tv/tsoding/clip/	BlightedAdorableFungusTBCheesePull-H_6ATRHzHnpQnfYa"
	interval = ""
	videoResource3, err := pkg.NewVideoResource(url, videoSource.Interval)
	if err != nil {
		log.Fatal("Failed to create video resource", err)
	}

	// This creates a video file with all the clips merged
	err := pkg.MergeVideos("./bin/finalVideo.mp4", videoResource1, videoResource2, videoResource3)
	if err != nil {
		log.Fatal("Failed to merge videos")
	}
}

```
