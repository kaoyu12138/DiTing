package service

import (
	"DiTing/go/dao"
	"DiTing/go/entity"
	"fmt"
	"math/rand"
	"os/exec"
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
	var maxN int
	if err = dao.SqlSession.Where("tag IN (?)", DangerTags).Find(&theWebsiteList).Count(&maxN).Error; err != nil {
		return websiteList, err
	}
	if maxN > 1000 {
		maxN = 1000
	}
	// 获取最近1000条诈骗网址，从中随机15条
	if err = dao.SqlSession.Where("tag IN (?)", DangerTags).
		Order("record_date desc").Offset(offset).Limit(maxN).Find(&theWebsiteList).Error; err != nil {
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

func CmdPythonGetTag(url string) (tag string, err error) {
	path := "F:/Go/workspace/projecct/learn/test.py" //执行python文件路径
	cmd := exec.Command("python", path, url)
	output, err := cmd.CombinedOutput() //执行命令，并获取输出结果
	if err != nil {
		fmt.Println("执行脚本时出错:", err)
		return "", err
	} else {
		fmt.Println("脚本输出:", string(output))
	}
	_, e := fmt.Sscanf(string(output), "%s", &tag) //匹配tag的输出内容
	if e != nil {
		return "", e
	}
	fmt.Println("tag :", tag)
	return tag, nil
}
