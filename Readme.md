# <NAME_TBD>

<NAME_TBD> is a video compiler that will download and compiler a list of videos

You can use Youtube and Twitch links with or without an interval to download.

Example

```go

videoResource, err := pkg.NewVideoResource(videoSource.URL, videoSource.Interval)
if err != nil {
	log.Fatal("Failed to create video resource", err)
}

```