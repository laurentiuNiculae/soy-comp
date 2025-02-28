package pkg

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
)

const (
	reEncode = true
)

func DownloadVideo(videoResource RemoteVideoResource) error {
	args := []string{}

	if videoResource.Interval() != nil && !videoResource.Interval().IsMax() {
		args = append(args, "--download-sections", "*"+videoResource.Interval().String())
	}

	// This is important here, clips don't use interval but are still cut automatically from the main video
	// so this makes sure we get the correct cut. This is why it's not inside the if above
	args = append(args, "--force-keyframes-at-cuts")

	if videoResource.DownloadQualityOptions() != "" {
		args = append(args, "-f", videoResource.DownloadQualityOptions())
	}

	args = append(args, "-o", videoResource.CachePath())
	args = append(args, "--recode-video", videoResource.CacheFileType())
	args = append(args, "--postprocessor-args", "-vf fps=30 -c:v libx264 -c:a aac -b:a 128k")
	args = append(args, videoResource.URL().String())

	cmd := exec.Command(
		"yt-dlp",
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

	return nil
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

func CutVideo(sourceVideo, outputVideo, intervalStr string) error {
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

	fmt.Println(cmd.String())

	stdOut, err := cmd.StdoutPipe()
	if err != nil {
		log.Println(err)
		return err
	}

	stdErr, err := cmd.StderrPipe()
	if err != nil {
		log.Println(err)
		return err
	}

	if err = cmd.Start(); err != nil {
		log.Println(err)
		return err
	}

	scanOut := bufio.NewScanner(stdOut)
	scanErr := bufio.NewScanner(stdErr)
	readCommandOutput(scanOut, scanErr)

	if err = cmd.Wait(); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func MergeVideos(finalVideoPath string, videos ...VideoResource) error {
	listPath, err := writeMergeList(videos)
	if err != nil {
		log.Println(err)
		return err
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
		log.Println(err)
		return err
	}

	stdErr, err := cmd.StderrPipe()
	if err != nil {
		log.Println(err)
		return err
	}

	if err = cmd.Start(); err != nil {
		log.Println(err)
		return err
	}

	scanOut := bufio.NewScanner(stdOut)
	scanErr := bufio.NewScanner(stdErr)
	readCommandOutput(scanOut, scanErr)

	if err = cmd.Wait(); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func writeMergeList(videos []VideoResource) (string, error) {
	for i := range videos {
		if !videos[i].IsCached() {
			return "", fmt.Errorf("video has not been downloaded, I can't find the cached video file locally")
		}
	}

	mergeFilePath, err := filepath.Abs("merge-list")
	if err != nil {
		return "", err
	}

	file, err := os.Create(mergeFilePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	fileWriter := bufio.NewWriter(file)

	for _, video := range videos {
		_, err := fileWriter.WriteString("file '" + video.CachePath() + "'\n")
		if err != nil {
			return "", fmt.Errorf("error while writing merge-list:%w", err)
		}
	}

	err = fileWriter.Flush()
	if err != nil {
		return "", fmt.Errorf("error while writing merge-list:%w", err)
	}

	return mergeFilePath, nil
}
