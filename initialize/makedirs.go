package initialize

import (
	"github.com/linzijie1998/bytedance_camp_douyin/global"
	"github.com/linzijie1998/bytedance_camp_douyin/util"
	"path/filepath"
)

func MakeUploadRootDirs() {
	videoPath := filepath.Join(global.DOUYIN_CONFIG.Upload.UploadRoot, "videos")
	coverPath := filepath.Join(global.DOUYIN_CONFIG.Upload.UploadRoot, "covers")
	err := util.CreateDir(videoPath, coverPath)
	if err != nil {
		panic(err)
	}
}
