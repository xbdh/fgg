package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/timestamppb"


	"tweets/internal/biz"

	pb "tweets/api/tweet/v1"
)

type TweetsService struct {
	pb.UnimplementedTweetsServer

	log *log.Helper
	tc *biz.TweetUsecase
}

func NewTweetsService(tc *biz.TweetUsecase,logger log.Logger) *TweetsService {
	return &TweetsService{
		tc: tc,
		log: log.NewHelper(logger),
	}
}

func (s *TweetsService) CreateTweets(ctx context.Context, req *pb.CreateRequest) (*pb.CreateReply, error) {
	s.log.Info("req:",req)

	ts:=biz.Tweet{
		UserId:    req.GetUserId(),
		Content: req.GetContent(),
		LikesCount:    req.GetLikesCount(),
		CommentsCount: req.GetCommentsCount(),
	}

	err:=s.tc.CreateTweet(ctx,&ts)
	if err!=nil{
		return nil, err
	}

	return &pb.CreateReply{
		TweetInfo: &pb.TweetInfo{
			Id:       uint64(ts.ID),
			UserId: ts.UserId,
			Content: ts.Content,
			LikesCount: ts.LikesCount,
			CommentsCount: ts.CommentsCount,
			CreateAt:timestamppb.New(ts.CreatedAt),
		},
	}, nil

}
func (s *TweetsService) GetTweetByUserId(ctx context.Context, req *pb.UserIdRequest) (*pb.UserIdReply, error) {
	s.log.Info("req:",req)


	userId:=req.GetUserId()

	tweets,err:=s.tc.GetTweetByUserId(ctx,userId)

	if err!=nil{
		return nil, err
	}

	tweetinfo:=[]*pb.TweetInfo{}

	for i:=0;i<len(tweets);i++{
		ti:=pb.TweetInfo{
			Id:       uint64(tweets[i].ID),
			UserId: tweets[i].UserId,
			Content: tweets[i].Content,
			LikesCount: tweets[i].LikesCount,
			CommentsCount: tweets[i].CommentsCount,
			CreateAt:timestamppb.New(tweets[i].CreatedAt),
		}
		tweetinfo=append(tweetinfo,&ti)

	}
	//fmt.Println(users)


	return &pb.UserIdReply{
		TweetInfo: tweetinfo,
	}, err

}
func (s *TweetsService) GetTweetByTweetId(ctx context.Context, req *pb.TweetIdRequest) (*pb.TweetIdReply, error) {
	s.log.Info("req:",req)


	userId:=req.GetTweetId()

	tweet,err:=s.tc.GetTweetByTweetId(ctx,userId)

	if err!=nil{
		return nil, err
	}


	ti:=pb.TweetInfo{
		Id:       uint64(tweet.ID),
		UserId: tweet.UserId,
		Content: tweet.Content,
		LikesCount: tweet.LikesCount,
		CommentsCount: tweet.CommentsCount,
		CreateAt:timestamppb.New(tweet.CreatedAt),

	}
	//fmt.Println(users)

	return &pb.TweetIdReply{
		TweetInfo: &ti,
	}, nil
}
