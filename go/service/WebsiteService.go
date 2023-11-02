package service

import (
	"DiTing/go/dao"
	"DiTing/go/entity"
	"math/rand"
)

// IsUrlExits 判断一个网址是否存在于数据库
func IsUrlExits(url string) (entity.Website, bool) {
	var website entity.Website
	dao.SqlSession.Where("url = ?", url).Find(&website)
	if website.UrlId != 0 {
		return website, true
	}
	return website, false
}

func GetDangerUrlList(offset, limit int) (websiteList []entity.Website, err error) {
	var DangerTags []string = []string{"danger"}
	var UsedN = make(map[int]bool)
	var numbers []int
	var theWebsiteList []entity.Website
	if err = dao.SqlSession.Where("tag IN (?)", DangerTags).Offset(offset).Find(&theWebsiteList).Error; err != nil {
		return websiteList, err
	}
	end := len(theWebsiteList)
	if end < limit {
		limit = end
	}
	for len(numbers) < limit {
		randomNumber := rand.Intn(end)
		if UsedN[randomNumber] {
			continue
		}
		UsedN[randomNumber] = true
		numbers = append(numbers, randomNumber)
	}
	for _, i := range numbers {
		websiteList = append(websiteList, theWebsiteList[i])
	}
	return websiteList, nil
}

func GetTag(url string) (tag string, err error) {
	var theWebsite entity.Website
	if err = dao.SqlSession.Where("Url = ?", url).Find(&theWebsite).Error; err != nil {
		return "", err
	}
	return theWebsite.Tag, nil
}

func CreateWebsite(website *entity.Website) (err error) {
	if err = dao.SqlSession.Create(website).Error; err != nil {
		return err
	}
	return
}
