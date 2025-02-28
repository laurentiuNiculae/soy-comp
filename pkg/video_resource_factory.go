package pkg

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

type VideoResourceFactory struct {
	VideoCompilerConfig
}

func NewVideoResourceFactory(config VideoCompilerConfig) VideoResourceFactory {
	return VideoResourceFactory{
		VideoCompilerConfig: config,
	}
}

func (vrf VideoResourceFactory) NewVideoResource(videoSource, intervalStr string) (VideoResource, error) {
	interval, err := ParseInterval(intervalStr)
	if err != nil {
		return nil, fmt.Errorf("can't parse interval for '%v' interval='%v':\nCaused by: %w", videoSource, intervalStr, err)
	}

	var videoResource VideoResource
	if fileInfo, err := os.Stat(videoSource); err == nil {
		if fileInfo.IsDir() {
			return nil, fmt.Errorf("video source can't be a directory")
		}

		videoResource, err = vrf.NewLocalVideoResource(videoSource, intervalStr)
		if err != nil {
			return nil, err
		}
	} else if videoURL, err := url.Parse(videoSource); err == nil {
		videoResource, err = vrf.NewRemoteVideoResource(videoURL, interval)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("video source is not a valid URL or file Path")
	}

	err = vrf.UpdateCodec(videoResource)
	if err != nil {
		return nil, err
	}

	return videoResource, nil
}

func (vrf VideoResourceFactory) UpdateCodec(videoResource VideoResource) error {
	tempOutputPath := videoResource.CachePath() + ".tmp" + "." + vrf.CacheFileType

	args := []string{}
	args = append(args, "-i", videoResource.CachePath())
	args = append(args, "-filter:v", "fps=30")
	args = append(args, "-crf", "23")
	args = append(args, "-c:v", "libx264")
	args = append(args, "-c:a", "aac")
	args = append(args, "-b:a", "128k")
	args = append(args, tempOutputPath)

	cmd := exec.Command(
		"ffmpeg",
		args...,
	)

	stdOut, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	stdErr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	if err = cmd.Start(); err != nil {
		return err
	}

	scanOut := bufio.NewScanner(stdOut)
	scanErr := bufio.NewScanner(stdErr)
	readCommandOutput(scanOut, scanErr)

	if err = cmd.Wait(); err != nil {
		return err
	}

	return os.Rename(tempOutputPath, videoResource.CachePath())
}

func (vrf VideoResourceFactory) NewRemoteVideoResource(videoURL *url.URL, interval *Interval,
) (RemoteVideoResource, error) {
	var video RemoteVideoResource
	var err error

	switch GetVideoURLKind(videoURL) {
	case YoutubeNormal:
		video, err = vrf.VideoResourceFromFullYoutubeLink(videoURL, interval)
		if err != nil {
			return nil, err
		}
	case YoutubeClip:
		video, err = vrf.VideoResourceFromYoutubeClipLink(videoURL, interval)
		if err != nil {
			return nil, err
		}
	case YoutubeMinified:
		video, err = vrf.VideoResourceFromMiniYoutubeLink(videoURL, interval)
		if err != nil {
			return nil, err
		}
	case TwitchClip:
		video, err = vrf.VideoResourceFromTwitchClipLink(videoURL, interval)
		if err != nil {
			return nil, err
		}
	case TwitchVod:
		video, err = vrf.VideoResourceFromTwitchVodLink(videoURL, interval)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("not supported url kind")
	}

	if !video.IsCached() {
		err = DownloadVideo(video)
		if err != nil {
			return nil, err
		}
	}

	return video, nil
}

func (vrf VideoResourceFactory) VideoResourceFromFullYoutubeLink(videoURL *url.URL, interval *Interval,
) (*YoutubeVideoResource, error) {
	if videoURL == nil {
		return nil, ErrURLIsNull
	}

	cachePath := path.Join(
		vrf.CacheFolder,
		videoURL.Query().Get("v")+"_"+interval.PathFormat()+"."+vrf.CacheFileType)

	return &YoutubeVideoResource{
		BaseRemoteResource: &BaseRemoteResource{
			videoURL:        videoURL,
			interval:        interval,
			cachePath:       cachePath,
			maxVideoQuality: vrf.MaxVideoQuality,
			cacheFileType:   vrf.CacheFileType,
		},
	}, nil
}

func (vrf VideoResourceFactory) VideoResourceFromYoutubeClipLink(videoURL *url.URL, interval *Interval,
) (*YoutubeVideoResource, error) {
	if videoURL == nil {
		return &YoutubeVideoResource{}, nil
	}

	cachePath := path.Join(
		vrf.CacheFolder,
		path.Base(videoURL.Path)+"_"+interval.PathFormat()+"."+vrf.CacheFileType)

	return &YoutubeVideoResource{
		BaseRemoteResource: &BaseRemoteResource{
			videoURL:        videoURL,
			interval:        interval,
			cachePath:       cachePath,
			maxVideoQuality: vrf.MaxVideoQuality,
			cacheFileType:   vrf.CacheFileType,
		},
	}, nil
}

func (vrf VideoResourceFactory) VideoResourceFromMiniYoutubeLink(videoURL *url.URL, interval *Interval,
) (*YoutubeVideoResource, error) {
	if videoURL == nil {
		return &YoutubeVideoResource{}, nil
	}

	cachePath := path.Join(
		vrf.CacheFolder,
		path.Base(videoURL.Path)+"_"+interval.PathFormat()+"."+vrf.CacheFileType)

	return &YoutubeVideoResource{
		BaseRemoteResource: &BaseRemoteResource{
			videoURL:        videoURL,
			interval:        interval,
			cachePath:       cachePath,
			maxVideoQuality: vrf.MaxVideoQuality,
			cacheFileType:   vrf.CacheFileType,
		},
	}, nil
}

func (vrf VideoResourceFactory) VideoResourceFromTwitchClipLink(videoURL *url.URL, interval *Interval,
) (*TwitchVideoResource, error) {
	if videoURL == nil {
		return &TwitchVideoResource{}, nil
	}

	cachePath := path.Join(
		vrf.CacheFolder,
		path.Base(videoURL.Path)+"_"+interval.PathFormat()+"."+vrf.CacheFileType)

	return &TwitchVideoResource{
		BaseRemoteResource: &BaseRemoteResource{
			videoURL:        videoURL,
			interval:        interval,
			cachePath:       cachePath,
			maxVideoQuality: vrf.MaxVideoQuality,
			cacheFileType:   vrf.CacheFileType,
		},
	}, nil
}

func (vrf VideoResourceFactory) VideoResourceFromTwitchVodLink(videoURL *url.URL, interval *Interval,
) (*TwitchVideoResource, error) {
	if videoURL == nil {
		return &TwitchVideoResource{}, nil
	}

	cachePath := path.Join(
		vrf.CacheFolder,
		path.Base(videoURL.Path)+"_"+interval.PathFormat()+"."+vrf.CacheFileType)

	return &TwitchVideoResource{
		BaseRemoteResource: &BaseRemoteResource{
			videoURL:        videoURL,
			interval:        interval,
			cachePath:       cachePath,
			maxVideoQuality: vrf.MaxVideoQuality,
			cacheFileType:   vrf.CacheFileType,
		},
	}, nil
}

func (vrf VideoResourceFactory) NewLocalVideoResource(videoSource string, intervalStr string) (VideoResource, error) {
	_, err := os.Stat(videoSource)
	if err != nil {
		return nil, fmt.Errorf("video file was not found: %w", err)
	}

	if intervalStr == "" {
		return &LocalVideoResource{
			VideoPath: videoSource,
		}, nil
	}

	interval, err := ParseInterval(intervalStr)
	if err != nil {
		return nil, err
	}

	cutVideoPath := filepath.Join(vrf.CacheFolder, filepath.Base(videoSource)+"_"+interval.PathFormat()+"."+vrf.CacheFileType)

	err = CutVideo(videoSource, cutVideoPath, intervalStr)
	if err != nil {
		return nil, err
	}

	return &LocalVideoResource{
		VideoPath: cutVideoPath,
	}, nil
}
