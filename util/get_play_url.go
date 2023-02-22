package util

import (
	"fmt"

	"github.com/linzijie1998/bytedance_camp_douyin/global"
)

func GetPlayURLByFilename(filename string) string {
	fmt.Println(filename)
	return fmt.Sprintf("http://%s:%s/videos/%s",
		global.DOUYIN_CONFIG.Upload.PublicHost, global.DOUYIN_CONFIG.Upload.PublicPort, filename)
}

func GetCoverURLByFilename(filename string) string {
	return fmt.Sprintf("http://%s:%s/covers/%s",
		global.DOUYIN_CONFIG.Upload.PublicHost, global.DOUYIN_CONFIG.Upload.PublicPort, filename)
}
