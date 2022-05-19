package service

import (
	"errors"
	"github.com/ACking-you/byte_douyin_project/models"
)

type VideoList struct {
	Videos []*models.Video `json:"video_list,omitempty"`
}

func QueryVideoListByUserId(userId int64) (*VideoList, error) {
	return NewQueryVideoListByUserIdFlow(userId).Do()
}

func NewQueryVideoListByUserIdFlow(userId int64) *QueryVideoListByUserIdFlow {
	return &QueryVideoListByUserIdFlow{userId: userId}
}

type QueryVideoListByUserIdFlow struct {
	userId int64
	videos []*models.Video

	videoList *VideoList
}

func (q *QueryVideoListByUserIdFlow) Do() (*VideoList, error) {
	if err := q.checkNum(); err != nil {
		return nil, err
	}
	if err := q.packData(); err != nil {
		return nil, err
	}
	return q.videoList, nil
}

func (q *QueryVideoListByUserIdFlow) checkNum() error {
	//检查userId是否存在
	if !models.NewUserInfoDAO().IsUserExistById(q.userId) {
		return errors.New("用户不存在")
	}

	return nil
}

//注意：Video由于在数据库中没有存储作者信息，所以需要手动填充
func (q *QueryVideoListByUserIdFlow) packData() error {
	err := models.NewVideoDAO().QueryVideoListByUserId(q.userId, &q.videos)
	if err != nil {
		return err
	}
	//作者信息查询
	var userInfo models.UserInfo
	err = models.NewUserInfoDAO().QueryUserInfoById(q.userId, &userInfo)
	if err != nil {
		return err
	}
	//填充Author字段
	for i := range q.videos {
		q.videos[i].Author = userInfo
	}
	//手动填充作者信息
	q.videoList = &VideoList{Videos: q.videos}

	return nil
}
