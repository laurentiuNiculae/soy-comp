package pkg2

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path"
	"sync"
)

type RenderCtx struct {
	DownloadQualityOptions string
	CacheFolder            string
	CacheFileType          string

	OverideLocalVideos bool
}

func DownloadVideoResource(ctx RenderCtx, remoteResource RemoteVideoResource) (VideoResource, error) {
	args := []string{}

	if remoteResource.Interval != nil && !remoteResource.Interval.IsMax() {
		args = append(args, "--download-sections", "*"+remoteResource.Interval.String())
	}

	// This is important here, clips don't use interval but are still cut automatically from the main video
	// so this makes sure we get the correct cut. This is why it's not inside the if above
	args = append(args, "--force-keyframes-at-cuts")

	if ctx.DownloadQualityOptions != "" {
		args = append(args, "-f", ctx.DownloadQualityOptions)
	}

	outputFile := path.Join(ctx.CacheFolder, remoteResource.Name)
	args = append(args, "-o", outputFile)

	// TODO: this will be updated using the update codec for some reason this doesn't work
	args = append(args, "--recode-video", ctx.CacheFileType)
	args = append(args, "--postprocessor-args", "-vf fps=30 -c:v libx264 -c:a aac -b:a 128k")
	// TODO

	args = append(args, remoteResource.URL.String())

	cmd := exec.Command(
		"yt-dlp",
		args...,
	)

	err := startAndWatchCmd(cmd)
	if err != nil {
		return VideoResource{}, err
	}

	// TODO: Update the codec and rest stuff

	return nil
}

func CutVideo(ctx RenderCtx, sourceVideo, outputVideo, intervalStr string) error {
	args := []string{}

	args = append(args, "-i", sourceVideo)
	args = append(args, "-y")

	if intervalStr != "" {
		interval, err := ParseInterval(intervalStr)
		if err != nil {
			return err
		}

		args = append(args, "-ss", interval.Begin.String())
		if !interval.End.isInf {
			args = append(args, "-t", interval.Duration().String())
		}
	}

	args = append(args, "-c:v", "copy")
	args = append(args, "-c:a", "copy")

	args = append(args, outputVideo)

	cmd := exec.Command(
		"ffmpeg",
		args...,
	)

	err := startAndWatchCmd(cmd)
	if err != nil {
		return err
	}

	return nil
}

func UpdateCodec(ctx RenderCtx, videoResource VideoResource) error {
	tempOutputPath := videoResource.VideoPath + ".tmp" + "." + ctx.CacheFileType

	args := []string{}
	args = append(args, "-i", videoResource.VideoPath)
	args = append(args, "-filter:v", "fps=30") // TODO: set all of these in RenderCtx
	args = append(args, "-crf", "23")
	args = append(args, "-c:v", "libx264")
	args = append(args, "-c:a", "aac")
	args = append(args, "-b:a", "128k")
	args = append(args, tempOutputPath)

	cmd := exec.Command(
		"ffmpeg",
		args...,
	)

	err := startAndWatchCmd(cmd)
	if err != nil {
		return err
	}

	return os.Rename(tempOutputPath, videoResource.VideoPath)
}

func startAndWatchCmd(cmd *exec.Cmd) error {
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

	return cmd.Wait()
}

func readCommandOutput(scanOut, scanErr *bufio.Scanner) error {
	wg := &sync.WaitGroup{}

	wg.Add(2)

	go scanAsync(scanOut, wg)
	go scanAsync(scanErr, wg)

	wg.Wait()

	return nil
}

func scanAsync(s *bufio.Scanner, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()

	for s.Scan() {
		fmt.Println(s.Text())
	}

	if s.Err() != nil {
		fmt.Println("Error while scanning: ", s.Err())
	}
}
