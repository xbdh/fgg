package service

import (
	"context"
	pb "friendships/api/friendship/v1"
	"friendships/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type FriendshipService struct {
	pb.UnimplementedFriendshipServer

	log *log.Helper
	uc *biz.FriendshipUsecase
}

func NewFriendshipService(u *biz.FriendshipUsecase,logger log.Logger) *FriendshipService {
	return &FriendshipService{
		uc: u,
		log: log.NewHelper(logger),
	}
}

func (s *FriendshipService) CreateFriendship(ctx context.Context, req *pb.CreateRequest) (*pb.CreateReply, error) {
	s.log.Info("req=",req)

	friendship:=biz.Friendship{

		FromId: req.GetFromUserId(),
		ToId:   req.GetToUserId(),
	}

	err:=s.uc.CreateFriendship(ctx,&friendship)

	if err!=nil{
		return nil, err
	}

	return &pb.CreateReply{
		FriendshipInfo: &pb.FrienddhipInfo{
			Id:         uint64(friendship.ID),
			FromUserId: friendship.FromId,
			ToUserId:   friendship.ToId,
		},
	}, nil
}
func (s *FriendshipService) Followings(ctx context.Context, req *pb.FollowingRequest) (*pb.FollowingReply, error) {
	s.log.Info("req: ",req)
	fromId:=req.FromUserId

	friendships,err:=s.uc.GetFollowings(ctx,fromId)
	if err!=nil{
		return nil, err
	}

	friendshipInfo:=[]*pb.FrienddhipInfo{}

	for i:=0;i<len(friendships);i++{
		ui:=pb.FrienddhipInfo{
			Id:            uint64(friendships[i].ID),
			FromUserId:     friendships[i].FromId,
			ToUserId:      friendships[i].ToId,

		}
		friendshipInfo=append(friendshipInfo,&ui)

	}

	return &pb.FollowingReply{
		FriendshipInfo:friendshipInfo,
	}, nil
}
func (s *FriendshipService) Followers(ctx context.Context, req *pb.FollowerRequest) (*pb.FollowerReply, error) {
	s.log.Info("req: ",req)
	toId:=req.ToUserId

	friendships,err:=s.uc.GetFollowers(ctx,toId)
	if err!=nil{
		return nil, err
	}

	friendshipInfo:=[]*pb.FrienddhipInfo{}

	for i:=0;i<len(friendships);i++{
		ui:=pb.FrienddhipInfo{
			Id:            uint64(friendships[i].ID),
			FromUserId:     friendships[i].FromId,
			ToUserId:      friendships[i].ToId,

		}
		friendshipInfo=append(friendshipInfo,&ui)

	}

	return &pb.FollowerReply{
		FriendshipInfo:friendshipInfo,
	}, nil
}
