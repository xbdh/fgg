package service

import (
	"comments/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"

	pb "comments/api/comment/v1"
)

type CommentService struct {
	pb.UnimplementedCommentServer
	log *log.Helper
	uc *biz.CommentUsecase
}

func NewCommentService(u *biz.CommentUsecase,logger log.Logger) *CommentService {
	return &CommentService{
		uc: u,
		log: log.NewHelper(logger),
	}
}

func (s *CommentService) CreateComment(ctx context.Context, req *pb.CreateRequest) (*pb.CreateReply, error) {

	s.log.Info("req:",req)

	us:=biz.Comment{

		UserId:  req.GetUserId(),
		TweetId: req.GetTweetId(),
		Content: req.GetContent(),
	}

	err:=s.uc.CreateComment(ctx,&us)
	if err!=nil{
		return nil, err
	}

	return &pb.CreateReply{
		CommentInfo: &pb.CommentInfo{
			Id:       uint64(us.ID),
			UserId: us.UserId,
			TweetId: us.TweetId,
			Content: us.Content,
		},
	}, nil


}
func (s *CommentService) ListTweetComment(ctx context.Context, req *pb.ListRequest) (*pb.ListReply, error) {
	s.log.Info("req:",req)

	tweetId:=req.TweetId

	comments,err:=s.uc.ListTweetComment(ctx,tweetId)
	if err!=nil{
		return nil, err
	}

	commentInfos :=[]*pb.CommentInfo{}

	for i:=0;i<len(comments);i++{
		ui:=pb.CommentInfo{
			Id:      uint64(comments[i].ID),
			UserId:  comments[i].UserId,
			TweetId: comments[i].TweetId,
			Content: comments[i].Content,
		}
		commentInfos =append(commentInfos,&ui)

	}

	return &pb.ListReply{
		CommentInfo: commentInfos,
	}, err


}
