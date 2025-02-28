package pkg2

import (
	"fmt"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type VideoResource struct {
	VideoPath string
	// TODO: add info/metadata about the video itself
}

func NewVideoResource(ctx RenderCtx, videoSource VideoSource) (VideoResource, error) {
	var videoResource VideoResource
	if fileInfo, err := os.Stat(videoSource.Source); err == nil {
		if fileInfo.IsDir() {
			return VideoResource{}, fmt.Errorf("video source can't be a directory")
		}

		videoResource, err = NewLocalVideoResource(ctx, videoSource.Source, videoSource.Interval)
		if err != nil {
			return VideoResource{}, err
		}
	} else if _, err := url.Parse(videoSource.Source); err == nil {
		videoResource, err = NewRemoteVideoResource(ctx, videoSource.Source, videoSource.Interval)
		if err != nil {
			return VideoResource{}, err
		}
	} else {
		return VideoResource{}, fmt.Errorf("video source is not a valid URL or file Path")
	}

	err := UpdateCodec(ctx, videoResource)
	if err != nil {
		return VideoResource{}, err
	}

	return videoResource, nil
}

func NewLocalVideoResource(ctx RenderCtx, videoSource string, intervalStr string) (VideoResource, error) {
	_, err := os.Stat(videoSource)
	if err != nil {
		return VideoResource{}, fmt.Errorf("video file was not found: %w", err)
	}

	if intervalStr == "" {
		return VideoResource{
			VideoPath: videoSource,
		}, nil
	}

	interval, err := ParseInterval(intervalStr)
	if err != nil {
		return VideoResource{}, err
	}

	cutVideoPath := filepath.Join(ctx.CacheFolder, filepath.Base(videoSource)+"_"+interval.PathFormat()+"."+ctx.CacheFileType)

	err = CutVideo(ctx, videoSource, cutVideoPath, intervalStr)
	if err != nil {
		return VideoResource{}, err
	}

	return VideoResource{
		VideoPath: cutVideoPath,
	}, nil
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

func NewRemoteVideoResource(ctx RenderCtx, videoURLStr string, intervalStr string) (VideoResource, error) {
	var video RemoteVideoResource

	videoURL, err := url.Parse(videoURLStr)
	if err != nil {
		return VideoResource{}, err
	}

	interval, err := ParseInterval(intervalStr)
	if err != nil {
		return VideoResource{}, err
	}

	switch GetVideoURLKind(videoURL) {
	case YoutubeNormal:
		video, err = VideoResourceFromFullYoutubeLink(ctx, videoURL, interval)
		if err != nil {
			return VideoResource{}, err
		}
	case YoutubeClip:
		video, err = VideoResourceFromYoutubeClipLink(ctx, videoURL, interval)
		if err != nil {
			return VideoResource{}, err
		}
	case YoutubeMinified:
		video, err = VideoResourceFromMiniYoutubeLink(ctx, videoURL, interval)
		if err != nil {
			return VideoResource{}, err
		}
	case TwitchClip:
		video, err = VideoResourceFromTwitchClipLink(ctx, videoURL, interval)
		if err != nil {
			return VideoResource{}, err
		}
	case TwitchVod:
		video, err = VideoResourceFromTwitchVodLink(ctx, videoURL, interval)
		if err != nil {
			return VideoResource{}, err
		}
	default:
		return VideoResource{}, fmt.Errorf("not supported url kind")
	}

	videoResource, err := LoadVideoResource(ctx, video.Name)
	if err != nil {
		videoResource, err = DownloadVideoResource(ctx, video)
		if err != nil {
			return VideoResource{}, err
		}
	}

	return videoResource, nil
}

func LoadVideoResource(ctx RenderCtx, path string) (VideoResource, error) {
	panic("unimplemented")
}

func VideoResourceFromFullYoutubeLink(ctx RenderCtx, videoURL *url.URL, interval *Interval) (RemoteVideoResource, error) {
	if videoURL == nil {
		return RemoteVideoResource{}, ErrURLIsNull
	}

	videoFileName := videoURL.Query().Get("v") + "_" + interval.PathFormat() + "." + ctx.CacheFileType
	localPath := path.Join(ctx.CacheFolder, videoFileName)

	return RemoteVideoResource{
		URL:       videoURL,
		Interval:  interval,
		Name:      videoFileName,
		LocalPath: localPath,
	}, nil
}

func VideoResourceFromYoutubeClipLink(ctx RenderCtx, videoURL *url.URL, interval *Interval,
) (RemoteVideoResource, error) {
	if videoURL == nil {
		return RemoteVideoResource{}, nil
	}

	videoFileName := path.Base(videoURL.Path) + "_" + interval.PathFormat() + "." + ctx.CacheFileType
	localPath := path.Join(ctx.CacheFolder, videoFileName)

	return RemoteVideoResource{
		URL:       videoURL,
		Interval:  interval,
		Name:      videoFileName,
		LocalPath: localPath,
	}, nil
}

func VideoResourceFromMiniYoutubeLink(ctx RenderCtx, videoURL *url.URL, interval *Interval,
) (RemoteVideoResource, error) {
	if videoURL == nil {
		return RemoteVideoResource{}, nil
	}

	videoFileName := path.Base(videoURL.Path) + "_" + interval.PathFormat() + "." + ctx.CacheFileType
	localPath := path.Join(ctx.CacheFolder, videoFileName)

	return RemoteVideoResource{
		URL:       videoURL,
		Interval:  interval,
		Name:      videoFileName,
		LocalPath: localPath,
	}, nil
}

func VideoResourceFromTwitchClipLink(ctx RenderCtx, videoURL *url.URL, interval *Interval,
) (RemoteVideoResource, error) {
	if videoURL == nil {
		return RemoteVideoResource{}, nil
	}

	videoFileName := path.Base(videoURL.Path) + "_" + interval.PathFormat() + "." + ctx.CacheFileType
	localPath := path.Join(ctx.CacheFolder, videoFileName)

	return RemoteVideoResource{
		URL:       videoURL,
		Interval:  interval,
		Name:      videoFileName,
		LocalPath: localPath,
	}, nil
}

func VideoResourceFromTwitchVodLink(ctx RenderCtx, videoURL *url.URL, interval *Interval,
) (RemoteVideoResource, error) {
	if videoURL == nil {
		return RemoteVideoResource{}, nil
	}

	videoFileName := path.Base(videoURL.Path) + "_" + interval.PathFormat() + "." + ctx.CacheFileType
	localPath := path.Join(ctx.CacheFolder, videoFileName)

	return RemoteVideoResource{
		URL:       videoURL,
		Interval:  interval,
		Name:      videoFileName,
		LocalPath: localPath,
	}, nil
}

// default name for output -> %(title)s.%(ext)s
// this means that the name is the video title

type RemoteVideoResource struct {
	URL       *url.URL
	Interval  *Interval
	Name      string
	LocalPath string
}

func (rvr RemoteVideoResource) IsCached() bool {
	_, err := os.Stat(rvr.LocalPath)
	return err == nil
}

type VideoSource struct {
	Source   string
	Interval string
}

func InitializeVideoResources(videoSources ...VideoSource) ([]VideoResource, error) {
	videoResources := make([]VideoResource, len(videoSources))

	for i := range videoSources {
		clip, err := NewVideoResource(videoSources[i])
		if err != nil {
			return nil, fmt.Errorf("failed to create video resource: %w", err)
		}

		videoResources[i] = clip
	}

	return videoResources, nil
}
