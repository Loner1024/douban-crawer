package parser

import (
	"douban-book-crawler/crawler/engine"
	"douban-book-crawler/crawler/model"
	"io/ioutil"
	"testing"
)

func TestParseBook(t *testing.T) {
	content, err := ioutil.ReadFile("book_test_data.html")
	if err != nil {
		panic(err)
	}
	var bookTestData = engine.Item{
		Url: "https://book.douban.com/subject/4913064/",
		Id:  "4913064",
		Payload: model.Book{
			Name:          "活着",
			Author:        "余华",
			Publisher:     "作家出版社",
			PublishYear:   "2012-8-1",
			Pages:         191,
			Price:         20.00,
			ISBN:          "9787506365437",
			Summary:       `<p>《活着(新版)》讲述了农村人福贵悲惨的人生遭遇。福贵本是个阔少爷，可他嗜赌如命，终于赌光了家业，一贫如洗。他的父亲被他活活气死，母亲则在穷困中患了重病，福贵前去求药，却在途中被国民党抓去当壮丁。经过几番波折回到家里，才知道母亲早已去世，妻子家珍含辛茹苦地养大两个儿女。此后更加悲惨的命运一次又一次降临到福贵身上，他的妻子、儿女和孙子相继死去，最后只剩福贵和一头老牛相依为命，但老人依旧活着，仿佛比往日更加洒脱与坚强。</p>    <p>《活着(新版)》荣获意大利格林扎纳•卡佛文学奖最高奖项（1998年）、台湾《中国时报》10本好书奖（1994年）、香港“博益”15本好书奖（1994年）、第三届世界华文“冰心文学奖”（2002年），入选香港《亚洲周刊》评选的“20世纪中文小说百年百强”、中国百位批评家和文学编辑评选的“20世纪90年代最有影响的10部作品”。</p>`,
			AuthorSummary: `<p>余华，1960年出生，1983年开始写作。至今已经出版长篇小说4部，中短篇小说集6部，随笔集4部。主要作品有《兄弟》《活着》《许三观卖血记》《在细雨中呼喊》等。其作品已被翻译成20多种语言在美国、英国、法国、德国、意大利、西班牙、荷兰、瑞典、挪威、希腊、俄罗斯、保加利亚、匈牙利、捷克、塞尔维亚、斯洛伐克、波兰、巴西、以色列、日本、韩国、越南、泰国和印度等国出版。曾获意大利格林扎纳·卡佛文学奖（1998年）、法国文学和艺术骑士勋章（2004年）、中华图书特殊贡献奖（2005年）、法国国际信使外国小说奖（2008年）等。</p>`,
		},
	}

	result := ParseBook(content, "活着", "https://book.douban.com/subject/4913064/")
	if result.Items[0] != bookTestData {
		t.Errorf(" excepted: %+v\nbut got: %+v", bookTestData, result.Items[0])
	}

}
