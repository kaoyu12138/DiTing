package service

import (
	"DiTing/go/dao"
	"DiTing/go/entity"
)

func CreateEssay(essay *entity.Essay) (err error) {
	if err = dao.SqlSession.Create(essay).Error; err != nil {
		return err
	}
	return
}

func GetEssayList(offset int) (essayList []entity.Essay, err error) {
	if err = dao.SqlSession.Order("publish_date desc").Offset(offset).Limit(15).Find(&essayList).Error; err != nil {
		return essayList, err
	}
	return essayList, nil
}

func GetUserEssayList(offset int, userName string) (essayList []entity.Essay, err error) {
	if err = dao.SqlSession.Where("user_name = ?", userName).
		Order("publish_date desc").Offset(offset).Limit(15).Find(&essayList).Error; err != nil {
		return essayList, err
	}
	return essayList, nil
}

func GetEssay(essayId int) (essayContent string, err error) {
	var theEssay entity.Essay
	if err = dao.SqlSession.Where("essay_id = ?", essayId).Find(&theEssay).Error; err != nil {
		return "", err
	}
	return theEssay.EssayContent, nil
}

func UpdateLikeCount(essayId int) (likeCount int, err error) {
	var theEssay entity.Essay
	if err = dao.SqlSession.Where("essay_id = ?", essayId).Find(&theEssay).Error; err != nil {
		return 0, err
	}
	likes := theEssay.LikeCount + 1
	dao.SqlSession.Model(&theEssay).Where("essay_id = ?", essayId).Update("like_count", likes)
	return likes, nil
}
