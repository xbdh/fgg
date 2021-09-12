package data

import (
	"context"
	"friendships/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type friendshipRepo struct {
	data *Data
	log  *log.Helper
}

func (fr *friendshipRepo) GetFollowings(ctx context.Context, fromUserId uint64) ([]*biz.Friendship, error) {
	friendships:=[]*biz.Friendship{}

	result:=fr.data.db.WithContext(ctx).Where("from_User_id = ?",fromUserId).Find(&friendships)
	//key:=fmt.Sprintf("followings:%d",fromUserId)
	//if result.Error==nil{
	//	for i:=0;i<len(friendships);i++{
	//		err:=fr.data.rdb.LPush(ctx,key,friendships[i].FromId).Err()
	//		if err!=nil{
	//			fr.log.Info("send to reids wrong")
	//		}
	//	}
	//}
	return friendships,result.Error
}

func (fr *friendshipRepo) GetFollowers(ctx context.Context, toUserId uint64) ([]*biz.Friendship, error) {
	friendships:=[]*biz.Friendship{}

	result:=fr.data.db.WithContext(ctx).Where("to_user_id = ?", toUserId).Find(&friendships)

	return friendships,result.Error
}

func (fr *friendshipRepo) CreateFriendship(ctx context.Context, friendship *biz.Friendship) error {
	result:=fr.data.db.WithContext(ctx).Create(friendship)

	return result.Error
}

// NewFriendshipRepo .
func NewFriendshipRepo(data *Data, logger log.Logger) biz.FriendshipRepo {
	return &friendshipRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}


