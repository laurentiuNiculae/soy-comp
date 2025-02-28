package pkg

import (
	"bufio"
	"fmt"
	"os/exec"
)

func ExtractAudioFromVideo(video VideoResource, output string, intervalStr string) (AudioResource, error) {
	args := []string{}

	args = append(args, "-i", video.CachePath())
	args = append(args, "-y")

	if intervalStr != "" {
		interval, err := ParseInterval(intervalStr)
		if err != nil {
			return nil, err
		}

		args = append(args, "-ss", interval.Begin.String())
		args = append(args, "-t", interval.Duration().String())
	}

	args = append(args, "-q:a", "0")
	args = append(args, "-map", "a")
	args = append(args, output)

	cmd := exec.Command(
		"ffmpeg",
		args...,
	)

	fmt.Println("EXECUTING: ", cmd.String())

	stdOut, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	stdErr, err := cmd.StderrPipe()
	if err != nil {
		return nil, err
	}

	if err = cmd.Start(); err != nil {
		return nil, err
	}

	scanOut := bufio.NewScanner(stdOut)
	scanErr := bufio.NewScanner(stdErr)
	readCommandOutput(scanOut, scanErr)

	if err = cmd.Wait(); err != nil {
		return nil, err
	}

	return nil, nil
}
