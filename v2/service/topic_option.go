package service

import (
	"strconv"
	"vote/v2/dao"
	"vote/v2/model"
	"vote/v2/pkg"
)

func TopicOptionQueryByTopicId(topicId string) (os []*model.TopicOption, total int, err error) {
	return dao.TopicOptionQueryByTopicId(topicId)
}

// SingleVote 单选投票
func SingleVote(r *model.VoteRecord, i int32, topicId, token string) error {
	bearer, err := pkg.ParseTokenWithBearer(token)
	if err != nil {
		return err
	}
	uid := bearer.Id
	r.Uid, _ = strconv.Atoi(uid)

	// 检查是否投过票
	if err := IsVoted(uid, topicId); err != nil {
		return err
	}

	// 插入数据到投票记录表
	//if _, err := dao.RecordInsert(r); err != nil {
	//	return err
	//}
	//
	//// 该选项的票数 + i
	//if err := dao.TopicOptionAddNumber(i, strconv.Itoa(r.OptionId)); err != nil {
	//	return err
	//}

	if err := dao.TopicOptionSingleVote(r, i, topicId); err != nil {
		return err
	}

	return nil
}

// MultipleVote 多选投票
func MultipleVote(rs []*model.VoteRecord, i int32, topicId, token string) error {
	bearer, err := pkg.ParseTokenWithBearer(token)
	if err != nil {
		return err
	}
	// 从 token 中得到用户 id
	uid := bearer.Id

	// 检查是否投过票
	if err := IsVoted(uid, topicId); err != nil {
		return err
	}

	for _, r := range rs {
		// 将 userid 添加到每个 model.VoteRecord 对象中
		r.Uid, _ = strconv.Atoi(uid)
	}

	if err := dao.TopicOptionMultipleVote(rs, i, topicId); err != nil {
		return err
	}

	// TODO
	return nil
}
