package service

import (
	"github.com/lukanzx/DouServer/cmd/video/dal/db"
	"github.com/lukanzx/DouServer/kitex_gen/video"
)

func (s *VideoService) GetVideoIDByUid(req *video.GetVideoIDByUidRequset) (videoIDList []int64, err error) {
	return db.GetVideoIDByUid(s.ctx, req.UserId)
}
