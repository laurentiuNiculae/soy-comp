package pkg

import (
	"net/url"
	"os"
	"strings"
)

type VideoResource interface {
	CachePath() string
	IsCached() bool
}

type RemoteVideoResource interface {
	Interval() *Interval
	CachePath() string
	IsCached() bool
	URL() *url.URL
	CacheFileType() string
	DownloadQualityOptions() string
}

type BaseRemoteResource struct {
	videoURL        *url.URL
	interval        *Interval
	cachePath       string
	cacheFileType   string
	maxVideoQuality string
}

func (rr *BaseRemoteResource) URL() *url.URL {
	return rr.videoURL
}

func (rr *BaseRemoteResource) Interval() *Interval {
	return rr.interval
}

func (rr *BaseRemoteResource) CachePath() string {
	return rr.cachePath
}

func (rr *BaseRemoteResource) IsCached() bool {
	_, err := os.Stat(rr.cachePath)
	return err == nil
}

func (rr *BaseRemoteResource) CacheFileType() string {
	return rr.cacheFileType
}

type VideoURLKind int

const (
	InvalidLink = iota
	YoutubeNormal
	YoutubeClip
	YoutubeMinified
	TwitchClip
	TwitchVod
)

func GetVideoURLKind(videoURL *url.URL) VideoURLKind {
	switch {
	case strings.HasSuffix(videoURL.Host, "youtube.com"):
		switch {
		case strings.Contains(videoURL.Path, "/watch"):
			return YoutubeNormal
		case strings.Contains(videoURL.Path, "/clip"):
			return YoutubeClip
		}
		return InvalidLink
	case strings.HasSuffix(videoURL.Host, "youtu.be"):
		return YoutubeMinified
	case strings.HasSuffix(videoURL.Host, "twitch.tv"):
		switch {
		case strings.Contains(videoURL.Path, "/clip"):
			return TwitchClip
		case strings.Contains(videoURL.Path, "/video"):
			return TwitchVod
		}
		return InvalidLink
	default:
		return InvalidLink
	}
}

type YoutubeVideoResource struct {
	*BaseRemoteResource
}

func (yt *YoutubeVideoResource) DownloadQualityOptions() string {
	return "bestvideo[height<=" + yt.maxVideoQuality + "][protocol=https]+bestaudio[protocol=https]"
}

type TwitchVideoResource struct {
	*BaseRemoteResource
}

func (tw *TwitchVideoResource) DownloadQualityOptions() string {
	return "best[height<=" + tw.maxVideoQuality + "]"
}

type LocalVideoResource struct {
	VideoPath string
}

func (lv *LocalVideoResource) CachePath() string {
	return lv.VideoPath
}

func (lv *LocalVideoResource) IsCached() bool {
	_, err := os.Stat(lv.VideoPath)
	return err == nil
}
