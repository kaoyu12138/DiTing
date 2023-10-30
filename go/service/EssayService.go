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
	if err = dao.SqlSession.Order("created_at desc").Offset(offset).Limit(15).Find(&essayList).Error; err != nil {
		return essayList, err
	}
	return essayList, nil
}

func GetUserEssayList(offset int, userName string) (essayList []entity.Essay, err error) {
	if err = dao.SqlSession.Where("UserName = ?", userName).
		Order("created_at desc").Offset(offset).Limit(15).Find(&essayList).Error; err != nil {
		return essayList, err
	}
	return essayList, nil
}

func UpdateLikeCount(essayId int) (likeCount int, err error) {
	var theEssay entity.Essay
	if err = dao.SqlSession.Where("EssayId = ?", essayId).Find(&theEssay).Error; err != nil {
		return 0, err
	}
	theEssay.LikeCount += 1
	if err = dao.SqlSession.Save(theEssay).Error; err != nil {
		return theEssay.LikeCount, err
	}
	return theEssay.LikeCount, nil
}
