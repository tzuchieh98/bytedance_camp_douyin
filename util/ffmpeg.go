package util

import (
	"fmt"
	"os/exec"
)

func GetCover(videoPath string, imgPath string, time string) error {
	cmd := exec.Command("ffmpeg", "-i", videoPath, "-ss", time, "-frames:v", "1", imgPath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("运行出错")
		fmt.Printf("%s\n", string(output))
	}
	return err
}
