package service

import (
	"bytes"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/lukanzx/DouServer/config"
	"github.com/lukanzx/DouServer/kitex_gen/video"
)

func (s *VideoService) UploadVideo(req *video.PutVideoRequest, videoName string) (err error) {
	fileReader := bytes.NewReader(req.VideoFile)
	err = s.bucket.PutObject(config.OSS.MainDirectory+"/"+videoName, fileReader)
	if err != nil {
		klog.Errorf("upload file error: %v\n", err)
	}
	return
}
