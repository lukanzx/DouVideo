package main

import (
	"context"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/lukanzx/DouServer/cmd/follow/rpc"
	"github.com/lukanzx/DouServer/kitex_gen/follow"
	"github.com/lukanzx/DouServer/kitex_gen/user"
)

func testFollowList(t *testing.T) {
	monkey.Patch(rpc.GetUser, func(ctx context.Context, req *user.InfoRequest) (*user.User, error) {
		return &user.User{Id: touserid}, nil
	})

	defer monkey.UnpatchAll()

	_, err := followService.FollowList(&follow.FollowListRequest{
		UserId: id,
		Token:  token,
	})

	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}
}

func benchmarkFollowList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := followService.FollowList(&follow.FollowListRequest{
			UserId: id,
			Token:  token,
		})

		if err != nil {
			b.Errorf("err: [%v] \n", err)
		}
		time.Sleep(100 * time.Millisecond) // Add a sleep to simulate some processing time
	}
}
