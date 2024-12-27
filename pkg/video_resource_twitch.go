package pkg

import (
	"net/url"
	"path"
)

type TwitchVideoResource struct {
	*BaseRemoteResource
}

func (tw *TwitchVideoResource) DownloadFormat() string {
	return "best[height<=" + maxVideoQuality + "]"
}

func VideoResourceFromTwitchClipLink(videoURL *url.URL, interval *Interval) (*TwitchVideoResource, error) {
	if videoURL == nil {
		return &TwitchVideoResource{}, nil
	}

	cachePath := path.Join(
		cacheBasePath,
		path.Base(videoURL.Path)+"_"+interval.PathFormat()+"."+cacheEncoding)

	return &TwitchVideoResource{
		BaseRemoteResource: &BaseRemoteResource{
			videoURL:  videoURL,
			interval:  interval,
			cachePath: cachePath,
		},
	}, nil
}

func VideoResourceFromTwitchVodLink(videoURL *url.URL, interval *Interval) (*TwitchVideoResource, error) {
	if videoURL == nil {
		return &TwitchVideoResource{}, nil
	}

	cachePath := path.Join(
		cacheBasePath,
		path.Base(videoURL.Path)+"_"+interval.PathFormat()+"."+cacheEncoding)

	return &TwitchVideoResource{
		BaseRemoteResource: &BaseRemoteResource{
			videoURL:  videoURL,
			interval:  interval,
			cachePath: cachePath,
		},
	}, nil
}
