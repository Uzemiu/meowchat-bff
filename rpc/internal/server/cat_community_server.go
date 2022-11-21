// Code generated by goctl. DO NOT EDIT!
// Source: cat_community.proto

package server

import (
	"context"

	"github.com/xh-polaris/cat-community-svc/rpc/internal/logic"
	"github.com/xh-polaris/cat-community-svc/rpc/internal/svc"
	"github.com/xh-polaris/cat-community-svc/rpc/pb"
)

type CatCommunityServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedCatCommunityServer
}

func NewCatCommunityServer(svcCtx *svc.ServiceContext) *CatCommunityServer {
	return &CatCommunityServer{
		svcCtx: svcCtx,
	}
}

//  Cat
func (s *CatCommunityServer) GetCatDetail(ctx context.Context, in *pb.CatDetailReq) (*pb.CatDetailResp, error) {
	l := logic.NewGetCatDetailLogic(ctx, s.svcCtx)
	return l.GetCatDetail(in)
}

//  查询猫咪信息
func (s *CatCommunityServer) QueryCat(ctx context.Context, in *pb.QueryCatReq) (*pb.QueryCatResp, error) {
	l := logic.NewQueryCatLogic(ctx, s.svcCtx)
	return l.QueryCat(in)
}

//  上传或更新猫咪信息
func (s *CatCommunityServer) UploadCat(ctx context.Context, in *pb.UploadCatReq) (*pb.UploadCatResp, error) {
	l := logic.NewUploadCatLogic(ctx, s.svcCtx)
	return l.UploadCat(in)
}

//  User
func (s *CatCommunityServer) GetUserDetail(ctx context.Context, in *pb.GetUserDetailReq) (*pb.GetUserDetailResp, error) {
	l := logic.NewGetUserDetailLogic(ctx, s.svcCtx)
	return l.GetUserDetail(in)
}

//  插入或更新用户信息
func (s *CatCommunityServer) UpsertUserDetail(ctx context.Context, in *pb.UpsertUserReq) (*pb.UpsertUserDetailResp, error) {
	l := logic.NewUpsertUserDetailLogic(ctx, s.svcCtx)
	return l.UpsertUserDetail(in)
}

//  Moment
func (s *CatCommunityServer) GetMomentDetail(ctx context.Context, in *pb.GetMomentDetailReq) (*pb.GetMomentDetailResp, error) {
	l := logic.NewGetMomentDetailLogic(ctx, s.svcCtx)
	return l.GetMomentDetail(in)
}

//  查询动态列表
func (s *CatCommunityServer) QueryMoment(ctx context.Context, in *pb.QueryMomentReq) (*pb.QueryMomentResp, error) {
	l := logic.NewQueryMomentLogic(ctx, s.svcCtx)
	return l.QueryMoment(in)
}

//  上传或更新动态
func (s *CatCommunityServer) UpsertMoment(ctx context.Context, in *pb.UpsertMomentReq) (*pb.UpsertMomentResp, error) {
	l := logic.NewUpsertMomentLogic(ctx, s.svcCtx)
	return l.UpsertMoment(in)
}

//  Comment
func (s *CatCommunityServer) ListComment(ctx context.Context, in *pb.ListCommentReq) (*pb.QueryMomentResp, error) {
	l := logic.NewListCommentLogic(ctx, s.svcCtx)
	return l.ListComment(in)
}

//  发表评论
func (s *CatCommunityServer) PostComment(ctx context.Context, in *pb.PostCommentReq) (*pb.PostCommentResp, error) {
	l := logic.NewPostCommentLogic(ctx, s.svcCtx)
	return l.PostComment(in)
}

//  Admin
func (s *CatCommunityServer) GetAdminDetail(ctx context.Context, in *pb.GetAdminDetailReq) (*pb.GetAdminDetailResp, error) {
	l := logic.NewGetAdminDetailLogic(ctx, s.svcCtx)
	return l.GetAdminDetail(in)
}

//  获取管理员列表
func (s *CatCommunityServer) ListAdmin(ctx context.Context, in *pb.ListAdminReq) (*pb.ListAdminResp, error) {
	l := logic.NewListAdminLogic(ctx, s.svcCtx)
	return l.ListAdmin(in)
}

//  Carousel
func (s *CatCommunityServer) ListCarousel(ctx context.Context, in *pb.ListCarouselReq) (*pb.ListCarouselResp, error) {
	l := logic.NewListCarouselLogic(ctx, s.svcCtx)
	return l.ListCarousel(in)
}