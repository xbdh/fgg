package service

import (
	pb "accounts/api/user/v1"
	"accounts/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type UserService struct {
	pb.UnimplementedUserServer
	
	log *log.Helper
	uc *biz.UserUsecase
}



func NewUserService(u *biz.UserUsecase,logger log.Logger) *UserService {
	return &UserService{
		uc: u,
		log: log.NewHelper(logger),
	}
}


func (s *UserService) GetUserByName(ctx context.Context, req *pb.NameRequest) (*pb.NameReply, error) {
	s.log.Info("req:",req)


	u,err:=s.uc.GetByName(ctx,req.Name)
	if err!=nil{
		return nil, err
	}
	
	return &pb.NameReply{
		Userinfo: &pb.UserInfo{
			Id:       uint64(u.ID),
			Name:     u.Name,
			Password: u.Password,
			Email:    u.Email,
		},
	}, nil
}
func (s *UserService) GetUserById(ctx context.Context, req *pb.IdRequest) (*pb.IdReply, error) {
	s.log.Info("req:",req)


	u,err:=s.uc.GetById(ctx,req.Id)
	if err!=nil{
		return nil, err
	}

	return &pb.IdReply{
		Userinfo: &pb.UserInfo{
			Id:       uint64(u.ID),
			Name:     u.Name,
			Password: u.Password,
			Email:    u.Email,
		},
	}, nil
}
func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateRequest) (*pb.CreateReply, error) {
	s.log.Info("req:",req)

	us:=biz.User{
		Name:     req.Name,
		Password: req.Password,
		Email:    req.Email,
	}

	err:=s.uc.CreateUser(ctx,&us)
	if err!=nil{
		return nil, err
	}

	return &pb.CreateReply{
		Userinfo: &pb.UserInfo{
			Id:       uint64(us.ID),
			Name:     us.Name,
			Password: us.Password,
			Email:    us.Email,
		},
	}, nil

}
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateReply, error) {
	return &pb.UpdateReply{}, nil
}
func (s *UserService) ListUser(ctx context.Context, req *pb.ListRequest) (*pb.ListReply, error) {
	s.log.Info("req:",req)

	ids:=req.Id

	users,err:=s.uc.ListUser(ctx,ids)
	if err!=nil{
		return nil, err
	}

	userinfo:=[]*pb.UserInfo{}

	for i:=0;i<len(users);i++{
		ui:=pb.UserInfo{
			Id:       uint64(users[i].ID),
			Name:     users[i].Name,
			Password: users[i].Password,
			Email:    users[i].Email,
		}
		userinfo=append(userinfo,&ui)

	}
	//fmt.Println(users)


	return &pb.ListReply{
		Results: userinfo,
	}, err


}
