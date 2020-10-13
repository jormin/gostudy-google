package parser

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/jormin/go-study/crawler/engine"
	"github.com/jormin/go-study/crawler/model"
	"github.com/jormin/go-study/helper"
	"github.com/jormin/go-study/modules/log"
	"strconv"
	"strings"
)

// 解析个人信息
func ParseProfile(contents string) (result engine.ParseResult) {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(contents))
	if err != nil {
		log.Error("Parse profile error: %v", err)
		return result
	}
	profile := model.NewProfile()
	// 个人信息
	dom.Find(".m-title").Each(func(i int, s *goquery.Selection) {
		switch i {
		case 0:
			profile.BasicInfo.Description = s.Next().Text()
		case 4:
			if s.Text() == "她的动态" {
				profile.NormalInfo.Gender = 1
				profile.Condition.Gender = 0
			} else {
				profile.NormalInfo.Gender = 0
				profile.Condition.Gender = 1
			}
		default:
			return
		}
	})
	s := dom.Find(".m-userInfo").Find(".top")
	avatarAttr, _ := s.Find(".logo").Attr("style")
	info := s.Find(".right")
	profile.BasicInfo.ID, _ = strconv.Atoi(info.Find(".id").Text()[5:])
	profile.BasicInfo.Avatar = avatarAttr[21 : len(avatarAttr)-2]
	profile.BasicInfo.Nickname = info.Find(".nickName").Text()
	des := strings.Split(dom.Find(".m-userInfoFixed").Find(".des").Text(), " | ")
	profile.BasicInfo.City = des[0]
	profile.NormalInfo.Age = des[1]
	profile.NormalInfo.Education = des[2]
	profile.NormalInfo.Marital = des[3]
	profile.NormalInfo.Height = des[4]
	profile.NormalInfo.Salary = des[5]
	classes := [2]string{"purple-btns", "pink-btns"}
	for _, class := range classes {
		ParseTag(dom, class, &profile.NormalInfo)
	}
	// 择偶条件
	ParseTag(dom, "gray-btns", &profile.Condition)
	log.Info("%+v", profile)
	result.Items = append(result.Items, profile.BasicInfo)
	return result
}

func ParseTag(dom *goquery.Document, class string, normalInfo *model.NormalInfo) {
	var yes = 1
	var no = 0
	s := dom.Find("." + class).First().Children()
	s.Each(func(i int, sub *goquery.Selection) {
		tag := sub.Text()
		if class == "purple-btns" && i == len(s.Nodes)-2 {
			normalInfo.Job = tag
			return
		}
		if helper.StringMultiIndex(tag, []string{"未婚", "已婚"}) {
			normalInfo.Marital = tag
			return
		}
		if strings.Index(tag, "岁") != -1 {
			normalInfo.Age = tag
			return
		}
		if strings.Index(tag, "cm") != -1 {
			normalInfo.Height = tag
			return
		}
		if strings.Index(tag, "月薪") == 0 {
			normalInfo.Salary = tag
			return
		}
		if strings.Index(tag, "工作地") == 0 {
			normalInfo.WorkPlace = tag[10:]
			return
		}
		if strings.Index(tag, "籍贯") == 0 {
			normalInfo.NativePlace = tag[7:]
			return
		}
		if strings.Index(tag, "体型") == 0 {
			normalInfo.Stature = tag[7:]
			return
		}
		if strings.Index(tag, "何时结婚") == 0 {
			normalInfo.WhenToMarry = tag[13:]
			return
		}
		if strings.Index(tag, "是否想要孩子") == 0 {
			tag = tag[19:]
			switch tag {
			case "想要孩子":
				normalInfo.IsWantBaby = yes
			case "不想要孩子":
				normalInfo.IsWantBaby = no
			}
			return
		}
		if strings.Index(tag, "喝酒") != -1 {
			if tag == "不喝酒" {
				normalInfo.IsDrink = no
			} else {
				normalInfo.IsDrink = yes
			}
			return
		}
		if strings.Index(tag, "吸烟") != -1 {
			if tag == "不吸烟" {
				normalInfo.IsSmoke = no
			} else {
				normalInfo.IsSmoke = yes
			}
			return
		}
		if strings.Index(tag, "座") != -1 {
			normalInfo.Constellation = tag
			return
		}
		if strings.Index(tag, "族") != -1 {
			normalInfo.Nation = tag
			return
		}
		if helper.StringMultiIndex(tag, []string{"小学", "初中", "高中", "中专", "大专", "本科", "研究生", "硕士", "博士"}) {
			normalInfo.Education = tag
			return
		}
		if strings.Index(tag, "kg") != -1 {
			normalInfo.Weight = tag
			return
		}
		if strings.Index(tag, "买车") != -1 {
			if tag == "未买车" {
				normalInfo.HasCar = no
			} else {
				normalInfo.HasCar = yes
			}
			return
		}
		if strings.Index(tag, "小孩") != -1 {
			if tag == "没有小孩" {
				normalInfo.HasBaby = no
			} else {
				normalInfo.HasBaby = yes
			}
			return
		}
		if strings.Index(tag, "购房") != -1 || strings.Index(tag, "住") != -1 {
			if tag == "已购房" {
				normalInfo.HasHouse = yes
			} else {
				normalInfo.HasCar = yes
			}
			return
		}
	})
}
