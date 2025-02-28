package pkg

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

type VideoCompilerConfig struct {
	CacheFolder     string
	CacheFileType   string // what extension to use when saving new clips, might not need this?
	MaxVideoQuality string
}

func DefaultConfig() VideoCompilerConfig {
	return VideoCompilerConfig{
		CacheFolder:     "./cached-videos",
		CacheFileType:   "mp4",
		MaxVideoQuality: "1080",
	}
}

type VideoSource struct {
	URL      string
	Interval string
}

type VideoCompiler struct {
	Config VideoCompilerConfig
	Clips  []VideoResource
	log    *log.Logger

	VideoResourceFactory
}

func NewVideoCompiler(config VideoCompilerConfig) *VideoCompiler {
	log.SetFlags(log.Llongfile)
	log := log.New(os.Stdout, "VIDEO_COMPILER: ", log.Llongfile)

	return &VideoCompiler{
		Config: config,
		log:    log,

		VideoResourceFactory: NewVideoResourceFactory(config),
	}
}

func NewDefaultVideoCompiler() *VideoCompiler {
	panic("TODO")
}

func (vc *VideoCompiler) AddClips(clip ...VideoResource) {
	vc.Clips = append(vc.Clips, clip...)
}

func (vc *VideoCompiler) AddVideoSources(videoSource ...VideoSource) error {
	for i := range videoSource {
		clip, err := vc.NewVideoResource(videoSource[i].URL, videoSource[i].Interval)
		if err != nil {
			return fmt.Errorf("failed to create video resource: %w", err)
		}

		vc.Clips = append(vc.Clips, clip)
	}

	return nil
}

func (vc *VideoCompiler) MergeVideos(finalVideoPath string) (VideoResource, error) {
	if finalVideoPath == "" {
		vc.log.Println("No file-name provided in MergeVideos arguments")
		return nil, fmt.Errorf("no file-name provided")
	}

	listPath, err := writeMergeList(vc.Clips)
	if err != nil {
		vc.log.Println("Failed to write merge list required by ffmpeg to concatenate the videos", err)
		return nil, err
	}

	args := []string{}
	args = append(args, "-f", "concat")
	args = append(args, "-safe", "0")
	args = append(args, "-y")
	args = append(args, "-i", listPath)

	if !reEncode {
		args = append(args, "-c", "copy")
	}

	args = append(args, finalVideoPath)

	cmd := exec.Command(
		"ffmpeg",
		args...,
	)

	stdOut, err := cmd.StdoutPipe()
	if err != nil {
		log.Println("Error creating StdoutPipe: ", err)
		return nil, err
	}

	stdErr, err := cmd.StderrPipe()
	if err != nil {
		log.Println("Error creating StderrPipe", err)
		return nil, err
	}

	if err = cmd.Start(); err != nil {
		log.Println("Error starting the execution of cmd: ", err)
		return nil, err
	}

	scanOut := bufio.NewScanner(stdOut)
	scanErr := bufio.NewScanner(stdErr)
	readCommandOutput(scanOut, scanErr)

	if err = cmd.Wait(); err != nil {
		log.Println("Error while waiting cmd result: ", err)
		return nil, err
	}

	finalVideoResource, err := vc.NewLocalVideoResource(finalVideoPath, "")
	if err != nil {
		log.Println("Error getting final video resource: ", err)
		return nil, err
	}

	vc.Clips = []VideoResource{}

	return finalVideoResource, nil
}

func (vc *VideoCompiler) MergeVideos2(finalVideoPath string) (VideoResource, error) {
	if finalVideoPath == "" {
		vc.log.Println("No file-name provided in MergeVideos arguments")
		return nil, fmt.Errorf("no file-name provided")
	}

	listPath, err := writeMergeList(vc.Clips)
	if err != nil {
		vc.log.Println("Failed to write merge list required by ffmpeg to concatenate the videos", err)
		return nil, err
	}

	args := []string{}
	args = append(args, "-f", "concat")
	args = append(args, "-safe", "0")
	args = append(args, "-y")
	args = append(args, "-i", listPath)

	// args = append(args, "-r", "30")
	// args = append(args, "-c:v", "h264")
	// args = append(args, "-crf", "23")
	// args = append(args, "-preset", "veryfast")
	// args = append(args, "-c:a", "aac")
	// args = append(args, "-b:a", "192k")

	if !reEncode {
		args = append(args, "-c", "copy")
	}

	args = append(args, finalVideoPath)

	cmd := exec.Command(
		"ffmpeg",
		args...,
	)

	stdOut, err := cmd.StdoutPipe()
	if err != nil {
		log.Println("Error creating StdoutPipe: ", err)
		return nil, err
	}

	stdErr, err := cmd.StderrPipe()
	if err != nil {
		log.Println("Error creating StderrPipe", err)
		return nil, err
	}

	if err = cmd.Start(); err != nil {
		log.Println("Error starting the execution of cmd: ", err)
		return nil, err
	}

	scanOut := bufio.NewScanner(stdOut)
	scanErr := bufio.NewScanner(stdErr)
	readCommandOutput(scanOut, scanErr)

	if err = cmd.Wait(); err != nil {
		log.Println("Error while waiting cmd result: ", err)
		return nil, err
	}

	finalVideoResource, err := vc.NewLocalVideoResource(finalVideoPath, "")
	if err != nil {
		log.Println("Error getting final video resource: ", err)
		return nil, err
	}

	vc.Clips = []VideoResource{}

	return finalVideoResource, nil
}
