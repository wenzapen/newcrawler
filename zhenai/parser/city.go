package parser

import (
	"fmt"
	"log"
	"newcrawler/engine"
	"regexp"
	// "strings"
)

/*
const cityRe = `<table><tbody><tr><th><a href="http://album.zhenai.com/u/1415391185" target="_blank">思念的感觉</a></th></tr> <tr><td width="180"><span class="grayL">性别：</span>女士</td> <td><span class="grayL">居住地：</span>陕西安康</td></tr> <tr><td width="180"><span class="grayL">年龄：</span>37</td> <td><span class="grayL">学   历：</span>大专</td> <!----></tr> <tr><td width="180"><span class="grayL">婚况：</span>离异</td> <td width="180"><span class="grayL">身   高：</span>160</td></tr></tbody></table> `
*/
// const cityRe = `<table><tbody><tr><th><a href="http://album.zhenai.com/u/[\d]+" target="_blank">[^<]+</a></th></tr> <tr><td width="180"><span class="grayL">性别：</span>[^<]+</td> <td><span class="grayL">居住地：</span>[^<]+</td></tr> <tr><td width="180"><span class="grayL">年龄：</span>[\d]+</td> <td><span class="grayL">学   历：</span>[^<]+</td> <!----></tr> <tr><td width="180"><span class="grayL">婚况：</span>[^<]+</td> <td width="180"><span class="grayL">身   高：</span>[\d]+</td></tr></tbody></table>`
const cityRe = `<table>.*?</table>`

func ParseCity(contents []byte) engine.ParseResult {
	// log.Printf("ParseCityList: got content: %s", contents)
	re := regexp.MustCompile(cityRe)

	matches := re.FindAll(contents, -1)
	for _, m := range matches {
		fmt.Printf("%s\n", m)

	}
	// fmt.Printf("ParseCityList: got %d items", len(matches))
	log.Printf("ParseCity: got %d items", len(matches))
	// matches := re.FindAllSubmatch(contents, -1)

	// result := engine.ParseResult{}
	// for _, m := range matches {
	// 	result.Items = append(result.Items, string(m[1]))
	// 	url := strings.ReplaceAll(string(m[2]), `\u002F`, `/`)
	// 	// log.Println("new url:", url)
	// 	result.Requests = append(result.Requests, engine.Request{
	// 		Url:       url,
	// 		ParseFunc: engine.NilParser,
	// 	})
	// }
	// return result
	return engine.ParseResult{}
}
