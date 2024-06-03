package service_test

import (
	"context"
	"testing"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/lukanzx/DouServer/cmd/follow/dal"
	"github.com/lukanzx/DouServer/cmd/follow/rpc"
	"github.com/lukanzx/DouServer/cmd/follow/service"
	"github.com/lukanzx/DouServer/config"
	"github.com/lukanzx/DouServer/kitex_gen/follow"
	"github.com/lukanzx/DouServer/pkg/utils"
)

var friendListTests = []Test{
	{10001, 10002, "", 1},
	{11001, 10002, "", 1},
}

func TestFriendList(t *testing.T) {
	config.InitForTest()
	dal.Init()
	rpc.Init()
	followService := service.NewFollowService(context.Background())
	for i, test := range friendListTests {
		test.token, _ = utils.CreateToken(test.id)
		_, err := followService.FriendList(&follow.FriendListRequest{
			UserId: test.id,
			Token:  test.token,
		})

		if err != nil {
			klog.Infof("test num %v,err:%v", i, err)
			continue
		}
	}
}
