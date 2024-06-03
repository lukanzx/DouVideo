// Code generated by Kitex v0.6.2. DO NOT EDIT.

package interactionservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	interaction "github.com/lukanzx/DouServer/kitex_gen/interaction"
)

func serviceInfo() *kitex.ServiceInfo {
	return interactionServiceServiceInfo
}

var interactionServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "InteractionService"
	handlerType := (*interaction.InteractionService)(nil)
	methods := map[string]kitex.MethodInfo{
		"FavoriteAction":      kitex.NewMethodInfo(favoriteActionHandler, newInteractionServiceFavoriteActionArgs, newInteractionServiceFavoriteActionResult, false),
		"FavoriteList":        kitex.NewMethodInfo(favoriteListHandler, newInteractionServiceFavoriteListArgs, newInteractionServiceFavoriteListResult, false),
		"VideoFavoritedCount": kitex.NewMethodInfo(videoFavoritedCountHandler, newInteractionServiceVideoFavoritedCountArgs, newInteractionServiceVideoFavoritedCountResult, false),
		"UserFavoriteCount":   kitex.NewMethodInfo(userFavoriteCountHandler, newInteractionServiceUserFavoriteCountArgs, newInteractionServiceUserFavoriteCountResult, false),
		"UserTotalFavorited":  kitex.NewMethodInfo(userTotalFavoritedHandler, newInteractionServiceUserTotalFavoritedArgs, newInteractionServiceUserTotalFavoritedResult, false),
		"IsFavorite":          kitex.NewMethodInfo(isFavoriteHandler, newInteractionServiceIsFavoriteArgs, newInteractionServiceIsFavoriteResult, false),
		"CommentAction":       kitex.NewMethodInfo(commentActionHandler, newInteractionServiceCommentActionArgs, newInteractionServiceCommentActionResult, false),
		"CommentList":         kitex.NewMethodInfo(commentListHandler, newInteractionServiceCommentListArgs, newInteractionServiceCommentListResult, false),
		"CommentCount":        kitex.NewMethodInfo(commentCountHandler, newInteractionServiceCommentCountArgs, newInteractionServiceCommentCountResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "interaction",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.6.2",
		Extra:           extra,
	}
	return svcInfo
}

func favoriteActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*interaction.InteractionServiceFavoriteActionArgs)
	realResult := result.(*interaction.InteractionServiceFavoriteActionResult)
	success, err := handler.(interaction.InteractionService).FavoriteAction(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newInteractionServiceFavoriteActionArgs() interface{} {
	return interaction.NewInteractionServiceFavoriteActionArgs()
}

func newInteractionServiceFavoriteActionResult() interface{} {
	return interaction.NewInteractionServiceFavoriteActionResult()
}

func favoriteListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*interaction.InteractionServiceFavoriteListArgs)
	realResult := result.(*interaction.InteractionServiceFavoriteListResult)
	success, err := handler.(interaction.InteractionService).FavoriteList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newInteractionServiceFavoriteListArgs() interface{} {
	return interaction.NewInteractionServiceFavoriteListArgs()
}

func newInteractionServiceFavoriteListResult() interface{} {
	return interaction.NewInteractionServiceFavoriteListResult()
}

func videoFavoritedCountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*interaction.InteractionServiceVideoFavoritedCountArgs)
	realResult := result.(*interaction.InteractionServiceVideoFavoritedCountResult)
	success, err := handler.(interaction.InteractionService).VideoFavoritedCount(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newInteractionServiceVideoFavoritedCountArgs() interface{} {
	return interaction.NewInteractionServiceVideoFavoritedCountArgs()
}

func newInteractionServiceVideoFavoritedCountResult() interface{} {
	return interaction.NewInteractionServiceVideoFavoritedCountResult()
}

func userFavoriteCountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*interaction.InteractionServiceUserFavoriteCountArgs)
	realResult := result.(*interaction.InteractionServiceUserFavoriteCountResult)
	success, err := handler.(interaction.InteractionService).UserFavoriteCount(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newInteractionServiceUserFavoriteCountArgs() interface{} {
	return interaction.NewInteractionServiceUserFavoriteCountArgs()
}

func newInteractionServiceUserFavoriteCountResult() interface{} {
	return interaction.NewInteractionServiceUserFavoriteCountResult()
}

func userTotalFavoritedHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*interaction.InteractionServiceUserTotalFavoritedArgs)
	realResult := result.(*interaction.InteractionServiceUserTotalFavoritedResult)
	success, err := handler.(interaction.InteractionService).UserTotalFavorited(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newInteractionServiceUserTotalFavoritedArgs() interface{} {
	return interaction.NewInteractionServiceUserTotalFavoritedArgs()
}

func newInteractionServiceUserTotalFavoritedResult() interface{} {
	return interaction.NewInteractionServiceUserTotalFavoritedResult()
}

func isFavoriteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*interaction.InteractionServiceIsFavoriteArgs)
	realResult := result.(*interaction.InteractionServiceIsFavoriteResult)
	success, err := handler.(interaction.InteractionService).IsFavorite(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newInteractionServiceIsFavoriteArgs() interface{} {
	return interaction.NewInteractionServiceIsFavoriteArgs()
}

func newInteractionServiceIsFavoriteResult() interface{} {
	return interaction.NewInteractionServiceIsFavoriteResult()
}

func commentActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*interaction.InteractionServiceCommentActionArgs)
	realResult := result.(*interaction.InteractionServiceCommentActionResult)
	success, err := handler.(interaction.InteractionService).CommentAction(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newInteractionServiceCommentActionArgs() interface{} {
	return interaction.NewInteractionServiceCommentActionArgs()
}

func newInteractionServiceCommentActionResult() interface{} {
	return interaction.NewInteractionServiceCommentActionResult()
}

func commentListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*interaction.InteractionServiceCommentListArgs)
	realResult := result.(*interaction.InteractionServiceCommentListResult)
	success, err := handler.(interaction.InteractionService).CommentList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newInteractionServiceCommentListArgs() interface{} {
	return interaction.NewInteractionServiceCommentListArgs()
}

func newInteractionServiceCommentListResult() interface{} {
	return interaction.NewInteractionServiceCommentListResult()
}

func commentCountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*interaction.InteractionServiceCommentCountArgs)
	realResult := result.(*interaction.InteractionServiceCommentCountResult)
	success, err := handler.(interaction.InteractionService).CommentCount(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newInteractionServiceCommentCountArgs() interface{} {
	return interaction.NewInteractionServiceCommentCountArgs()
}

func newInteractionServiceCommentCountResult() interface{} {
	return interaction.NewInteractionServiceCommentCountResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) FavoriteAction(ctx context.Context, req *interaction.FavoriteActionRequest) (r *interaction.FavoriteActionResponse, err error) {
	var _args interaction.InteractionServiceFavoriteActionArgs
	_args.Req = req
	var _result interaction.InteractionServiceFavoriteActionResult
	if err = p.c.Call(ctx, "FavoriteAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FavoriteList(ctx context.Context, req *interaction.FavoriteListRequest) (r *interaction.FavoriteListResponse, err error) {
	var _args interaction.InteractionServiceFavoriteListArgs
	_args.Req = req
	var _result interaction.InteractionServiceFavoriteListResult
	if err = p.c.Call(ctx, "FavoriteList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) VideoFavoritedCount(ctx context.Context, req *interaction.VideoFavoritedCountRequest) (r *interaction.VideoFavoritedCountResponse, err error) {
	var _args interaction.InteractionServiceVideoFavoritedCountArgs
	_args.Req = req
	var _result interaction.InteractionServiceVideoFavoritedCountResult
	if err = p.c.Call(ctx, "VideoFavoritedCount", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UserFavoriteCount(ctx context.Context, req *interaction.UserFavoriteCountRequest) (r *interaction.UserFavoriteCountResponse, err error) {
	var _args interaction.InteractionServiceUserFavoriteCountArgs
	_args.Req = req
	var _result interaction.InteractionServiceUserFavoriteCountResult
	if err = p.c.Call(ctx, "UserFavoriteCount", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UserTotalFavorited(ctx context.Context, req *interaction.UserTotalFavoritedRequest) (r *interaction.UserTotalFavoritedResponse, err error) {
	var _args interaction.InteractionServiceUserTotalFavoritedArgs
	_args.Req = req
	var _result interaction.InteractionServiceUserTotalFavoritedResult
	if err = p.c.Call(ctx, "UserTotalFavorited", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) IsFavorite(ctx context.Context, req *interaction.IsFavoriteRequest) (r *interaction.IsFavoriteResponse, err error) {
	var _args interaction.InteractionServiceIsFavoriteArgs
	_args.Req = req
	var _result interaction.InteractionServiceIsFavoriteResult
	if err = p.c.Call(ctx, "IsFavorite", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CommentAction(ctx context.Context, req *interaction.CommentActionRequest) (r *interaction.CommentActionResponse, err error) {
	var _args interaction.InteractionServiceCommentActionArgs
	_args.Req = req
	var _result interaction.InteractionServiceCommentActionResult
	if err = p.c.Call(ctx, "CommentAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CommentList(ctx context.Context, req *interaction.CommentListRequest) (r *interaction.CommentListResponse, err error) {
	var _args interaction.InteractionServiceCommentListArgs
	_args.Req = req
	var _result interaction.InteractionServiceCommentListResult
	if err = p.c.Call(ctx, "CommentList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CommentCount(ctx context.Context, req *interaction.CommentCountRequest) (r *interaction.CommentCountResponse, err error) {
	var _args interaction.InteractionServiceCommentCountArgs
	_args.Req = req
	var _result interaction.InteractionServiceCommentCountResult
	if err = p.c.Call(ctx, "CommentCount", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
