package pkg

import (
	"net/url"
	"path"
)

type YoutubeVideoResource struct {
	*BaseRemoteResource
}

func (yt *YoutubeVideoResource) DownloadFormat() string {
	return "bestvideo[height<=" + maxVideoQuality + "][protocol=https]+bestaudio[protocol=https]"
}

func VideoResourceFromFullYoutubeLink(videoURL *url.URL, interval *Interval) (*YoutubeVideoResource, error) {
	if videoURL == nil {
		return nil, ErrURLIsNull
	}

	cachePath := path.Join(
		cacheBasePath,
		videoURL.Query().Get("v")+"_"+interval.PathFormat()+"."+cacheEncoding)

	return &YoutubeVideoResource{
		BaseRemoteResource: &BaseRemoteResource{
			videoURL:  videoURL,
			interval:  interval,
			cachePath: cachePath,
		},
	}, nil
}

func VideoResourceFromYoutubeClipLink(videoURL *url.URL, interval *Interval) (*YoutubeVideoResource, error) {
	if videoURL == nil {
		return &YoutubeVideoResource{}, nil
	}

	cachePath := path.Join(
		cacheBasePath,
		path.Base(videoURL.Path)+"_"+interval.PathFormat()+"."+cacheEncoding)

	return &YoutubeVideoResource{
		BaseRemoteResource: &BaseRemoteResource{
			videoURL:  videoURL,
			interval:  interval,
			cachePath: cachePath,
		},
	}, nil
}

func VideoResourceFromMiniYoutubeLink(videoURL *url.URL, interval *Interval) (*YoutubeVideoResource, error) {
	if videoURL == nil {
		return &YoutubeVideoResource{}, nil
	}

	cachePath := path.Join(
		cacheBasePath,
		path.Base(videoURL.Path)+"_"+interval.PathFormat()+"."+cacheEncoding)

	return &YoutubeVideoResource{
		BaseRemoteResource: &BaseRemoteResource{
			videoURL:  videoURL,
			interval:  interval,
			cachePath: cachePath,
		},
	}, nil
}
