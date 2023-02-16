package util

import (
	"fmt"
	"testing"
)

func TestGetCover(t *testing.T) {
	videoPath := "../upload/videos/20230215164545-share_0d25b695bce99d046544684000aa69e9.mp4"
	coverPath := "../upload/covers/20230215164545-share_0d25b695bce99d046544684000aa69e9.mp4.jpg"
	err := GetCover(videoPath, coverPath, "00:00:02")
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}
