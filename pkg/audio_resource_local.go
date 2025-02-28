package pkg

type LocalAudioResource struct {
}

func NewFromVideoResource(video VideoResource) *LocalAudioResource {
	videoPath := video.CachePath()
	_ = videoPath

	return &LocalAudioResource{}
}
