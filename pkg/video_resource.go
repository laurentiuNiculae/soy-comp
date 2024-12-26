package pkg

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

const cacheBasePath = "./cached-videos"
const cacheEncoding = "webm"
const maxVideoQuality = "360"

type VideoResource interface {
	CachePath() string
	IsCached() bool
}

type RemoteVideoResource interface {
	Interval() *Interval
	CachePath() string
	IsCached() bool
	URL() *url.URL
	DownloadFormat() string
}

type BaseRemoteResource struct {
	videoURL  *url.URL
	interval  *Interval
	cachePath string
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

func NewVideoResource(videoSource string, intervalStr string) (VideoResource, error) {
	interval, err := ParseInterval(intervalStr)
	if err != nil {
		return nil, fmt.Errorf("can't parse interval for '%v' interval='%v':\nCaused by: %w", videoSource, intervalStr, err)
	}

	if videoURL, err := url.Parse(videoSource); err == nil {
		return NewRemoteVideoResource(videoURL, interval)
	} else if fileInfo, err := os.Stat(videoSource); err == nil {
		if fileInfo.IsDir() {
			return nil, fmt.Errorf("video source can't be a directory")
		}

		return NewLocalVideoResource(videoSource, interval)
	} else {
		return nil, fmt.Errorf("video source is not a valid URL or file Path")
	}
}

func NewRemoteVideoResource(videoURL *url.URL, interval *Interval) (RemoteVideoResource, error) {
	var video RemoteVideoResource
	var err error

	switch GetVideoURLKind(videoURL) {
	case YoutubeNormal:
		video, err = VideoResourceFromFullYoutubeLink(videoURL, interval)
		if err != nil {
			return nil, err
		}
	case YoutubeClip:
		video, err = VideoResourceFromYoutubeClipLink(videoURL, interval)
		if err != nil {
			return nil, err
		}
	case YoutubeMinified:
		video, err = VideoResourceFromMiniYoutubeLink(videoURL, interval)
		if err != nil {
			return nil, err
		}
	case TwitchClip:
		video, err = VideoResourceFromTwitchClipLink(videoURL, interval)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("not supported url kind")
	}

	if !video.IsCached() {
		err = DownloadVideo(video)
		if err != nil {
			panic(err)
		}
	}

	return video, nil
}

type VideoURLKind int

const (
	InvalidLink = iota
	YoutubeNormal
	YoutubeClip
	YoutubeMinified
	TwitchClip
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
		if strings.Contains(videoURL.Path, "/clip") {
			return TwitchClip
		}
		return InvalidLink
	default:
		return InvalidLink
	}
}
