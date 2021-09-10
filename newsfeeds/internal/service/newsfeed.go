package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/timestamppb"
	"newsfeeds/internal/biz"

	pb "newsfeeds/api/newsfeed/v1"
)

type NewsFeedService struct {
	pb.UnimplementedNewsFeedServer

	log *log.Helper
	uc *biz.NewsfeedUsecase
}

func NewNewsFeedService(u *biz.NewsfeedUsecase,logger log.Logger) *NewsFeedService {
	return &NewsFeedService{
		uc: u,
		log: log.NewHelper(logger),
	}
}

func (s *NewsFeedService) CreateNewsFeed(ctx context.Context, req *pb.CreateRequest) (*pb.CreateReply, error) {
	s.log.Info("req:",req)

	us:=biz.NewsFeed{

		UserId:  req.GetUserId(),
		TweetId: req.GetTweetId(),
	}

	err:=s.uc.CreateNewsfeed(ctx,&us)
	if err!=nil{
		return nil, err
	}

	return &pb.CreateReply{
		NewFeedInfo: &pb.NewsFeedInfo{
			Id:       uint64(us.ID),
			UserId: us.UserId,
			TweetId: us.TweetId,

		},
	}, nil

}
func (s *NewsFeedService) ListFollowTweet(ctx context.Context, req *pb.ListRequest) (*pb.ListReply, error) {
	s.log.Info("req:",req)

	userId:=req.UserId

	allfeeds,err:=s.uc.ListFollowTweet(ctx,userId)
	if err!=nil{
		return nil, err
	}

	newsfeedinfos:=[]*pb.NewsFeedInfo{}

	for i:=0;i<len(allfeeds);i++{
		ui:=pb.NewsFeedInfo{
			Id:       uint64(allfeeds[i].Nf.ID),
			UserId:   allfeeds[i].Nf.UserId,
			TweetId: allfeeds[i].Nf.TweetId,
			TweetInfo: &pb.TweetInfo{
				Id:            allfeeds[i].T.Id,
				UserId:        allfeeds[i].T.UserId,
				Content:        allfeeds[i].T.Content,
				LikesCount:    allfeeds[i].T.LikesCount,
				CommentsCount: allfeeds[i].T.CommentsCount,
				CreateAt: timestamppb.New(allfeeds[i].T.CreateAt),
			},

		}
		newsfeedinfos=append(newsfeedinfos,&ui)

	}

	return &pb.ListReply{
		NewFeedInfo: newsfeedinfos,
	}, err

}
