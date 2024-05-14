package parser

import (
	"fmt"
	"newcrawler/engine"
	"newcrawler/model"
	"regexp"
	"strconv"
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
	Salary           int
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
var heightRe = regexp.MustCompile(`<td width="180"><span class="grayL">身[^高]+高：</span>([^<]+)</td>`)

var educationRe = regexp.MustCompile(`<td><span class="grayL">学[^历]+历：</span>([^<]+)</td>`)
var salaryRe = regexp.MustCompile(`<td><span class="grayL">月[^薪]+薪：</span>([^<]+)</td>`)
var introduceContentRe = regexp.MustCompile(`<td><span class="grayL">内心独白：</span>([^<]+)</td>`)
var linkURLRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[\d]+)" target="_blank">[^<]+</a>`)

func ParseCity(contents []byte) engine.ParseResult {
	// log.Printf("ParseCityList: got content: %s", contents)
	// re := regexp.MustCompile(cityRe)

	matches := cityRe.FindAll(contents, -1)
	// for _, m := range matches {
	// 	fmt.Printf("%v\n", m)

	// }
	// fmt.Printf("ParseCityList: got %d items", len(matches))
	fmt.Printf("ParseCity: got %d items\n", len(matches))
	// matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	item := model.Profile{}
	for _, m := range matches {
		item.Nickname = extractString(m, nicknameRe)
		item.Age = extractInt(m, ageRe)
		item.Gender = extractString(m, genderRe)
		item.Marriage = extractString(m, marriageRe)
		item.Localtion = extractString(m, locationRe)
		item.Height = extractInt(m, heightRe)
		item.Education = extractString(m, educationRe)
		item.Salary = extractString(m, salaryRe)
		item.IntroduceContent = extractString(m, introduceContentRe)
		item.LinkURL = extractString(m, linkURLRe)

		result.Items = append(result.Items, item)

	}
	return result
	// return engine.ParseResult{}
}

func extractString(contents []byte, re *regexp.Regexp) string {
	// fmt.Printf("extraceString: %s\n", string(contents))
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return " "
	}
}

func extractInt(contents []byte, re *regexp.Regexp) int {
	// fmt.Printf("extraceInt: %s\n", string(contents))
	if i, err := strconv.Atoi(extractString(contents, re)); err == nil {
		return i
	} else {
		// log.Println("extract int error: ", err)
		return 0
	}

}
