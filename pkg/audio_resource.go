package pkg

type AudioResource interface {
	CachePath() string
	IsCached() bool
}
