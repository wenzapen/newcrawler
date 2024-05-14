package parser

import (
	"newcrawler/engine"
	"regexp"
	"strings"
)

// const cityListRe = `"linkContent":"阿坝","linkURL":"http:\u002F\u002Fwww.zhenai.com\u002Fzhenghun\u002Faba"`
const cityListRe = `"linkContent":"([^"]+)","linkURL":"(http:\\u002F\\u002Fwww.zhenai.com\\u002Fzhenghun\\u002F[^"]+)"`

func ParseCityList(contents []byte) engine.ParseResult {
	// fmt.Printf("ParseCityList: got content: %s", contents)
	re := regexp.MustCompile(cityListRe)

	// matches := re.FindAll(contents, -1)
	// for _, m := range matches {
	// 	log.Printf("%s\n", m)

	// }
	// fmt.Printf("ParseCityList: got %d items", len(matches))
	// log.Printf("ParseCityList: got %d items", len(matches))
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, string(m[1]))
		url := strings.ReplaceAll(string(m[2]), `\u002F`, `/`)
		// log.Println("new url:", url)
		result.Requests = append(result.Requests, engine.Request{
			Url:       url,
			ParseFunc: ParseCity,
		})
	}
	return result
	// return engine.ParseResult{}
}
