package parser

import (
	"fmt"
	"log"
	"newcrawler/engine"
	"regexp"
	// "strings"
)

/*
	Nickname         string
	Gender           int
	Age              int
	Marriage         string
	Localtion        string
	Height           int
	Education        string
	Salary           string
	introduceContent string
	LinkURL          string
const cityRe = `<table><tbody><tr><th><a href="http://album.zhenai.com/u/1415391185" target="_blank">思念的感觉</a></th></tr> <tr><td width="180"><span class="grayL">性别：</span>女士</td> <td><span class="grayL">居住地：</span>陕西安康</td></tr> <tr><td width="180"><span class="grayL">年龄：</span>37</td> <td><span class="grayL">学   历：</span>大专</td> <!----></tr> <tr><td width="180"><span class="grayL">婚况：</span>离异</td> <td width="180"><span class="grayL">身   高：</span>160</td></tr></tbody></table> `
*/
// const cityRe = `<table><tbody><tr><th><a href="http://album.zhenai.com/u/[\d]+" target="_blank">[^<]+</a></th></tr> <tr><td width="180"><span class="grayL">性别：</span>[^<]+</td> <td><span class="grayL">居住地：</span>[^<]+</td></tr> <tr><td width="180"><span class="grayL">年龄：</span>[\d]+</td> <td><span class="grayL">学   历：</span>[^<]+</td> <!----></tr> <tr><td width="180"><span class="grayL">婚况：</span>[^<]+</td> <td width="180"><span class="grayL">身   高：</span>[\d]+</td></tr></tbody></table>`
var cityRe = regexp.MustCompile(`<table>.*?</table>`)
var ageRe = regexp.MustCompile(`<td width="180"><span class="grayL">年龄：</span>([\d]+)</td>`)
var nicknameRe = regexp.MustCompile(`<th><a href="http://album.zhenai.com/u/[\d]+" target="_blank">([^<]+)</a></th>`)
var genderRe = regexp.MustCompile(`<td width="180"><span class="grayL">性别：</span>([^<]+)</td>`)
var marriageRe = regexp.MustCompile(`<td width="180"><span class="grayL">婚况：</span>([^<]+)</td>`)
var locationRe = regexp.MustCompile(`<td><span class="grayL">居住地：</span>([^<]+)</td>`)
var heightRe = regexp.MustCompile(`<td width="180"><span class="grayL">身   高：</span>([\d]+)</td>`)
var EducationRe = regexp.MustCompile(`<td><span class="grayL">学   历：</span>([^<]+)</td>`)
var SalaryRe = regexp.MustCompile(`<td><span class="grayL">月薪：</span>([^<]+)</td>`)
var introduceContentRe = regexp.MustCompile(`<td><span class="grayL">内心独白：</span>([^<]+)</td>`)
var linkURLRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[\d]+)" target="_blank">[^<]+</a>`)

func ParseCity(contents []byte) engine.ParseResult {
	// log.Printf("ParseCityList: got content: %s", contents)
	// re := regexp.MustCompile(cityRe)

	matches := cityRe.FindAll(contents, -1)
	for _, m := range matches {
		fmt.Printf("%s\n", m)

	}
	// fmt.Printf("ParseCityList: got %d items", len(matches))
	log.Printf("ParseCity: got %d items", len(matches))
	// matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	item := model.Profile{}
	for _, m := range matches {
		item.Nickname = extractString(m, nicknameRe)
		if age, err := strconv.Atoi(extractString(m, ageRe)); err == nil {
			item.Age = age
		}
		item.Gender = extractString(m, genderRe)
		item.Marriage = extractString(m, marriageRe)
		item.Localtion = extractString(m, locationRe)
		if height, err := strconv.Atoi(extractString(m, heightRe)); err == nil {
			item.Height = height
		}
		item.Education = extractString(m, EducationRe)
		if salary, err := strconv.Atoi(extractString(m, SalaryRe)); err == nil {
			item.Salary = salary
		}

		item.introduceContent = extractString(m, introduceContentRe)
		item.LinkURL = extractString(m, linkURLRe)

		result.Items = append(result.Items, item)

	}
	return result
	// return engine.ParseResult{}
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
