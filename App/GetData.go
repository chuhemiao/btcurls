package main

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/bitly/go-simplejson"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"qihuodata/Common"
	"reflect"
	"strings"
	"sync"
	"time"
)

type HotData struct {
	Code    int
	Message string
	Data    interface{}
}

type Spider struct {
	DataType string
}

// 保存数据

func SaveDataToJson(data interface{}) string {
	Message := HotData{}
	Message.Code = 0
	Message.Message = "获取成功"
	Message.Data = data
	jsonStr, err := json.Marshal(Message)
	if err != nil {
		log.Fatal("序列化json错误", err)
	}
	return string(jsonStr)

}

// chainnews
func (spider Spider) GetChainnews() []map[string]interface{} {
	url := "https://www.chainnews.com/"
	timeout := time.Duration(5 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}
	request.Header.Add("User-Agent", `Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Mobile Safari/537.36`)
	res, err := client.Do(request)
	if err != nil {
		fmt.Println("抓取" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}
	defer res.Body.Close()
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("抓取" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}
	var allData []map[string]interface{}
	document.Find(".feed-post-title").Each(func(i int, selection *goquery.Selection) {
		text := selection.Find("a").Text()
		url, boolUrl := selection.Find("a").Attr("href")
		if boolUrl {
			allData = append(allData, map[string]interface{}{"title": text, "url": "https://www.chainnews.com" + url})
		}
	})
	return allData
}

// huoxing24
func (spider Spider) GetHuoxing() []map[string]interface{} {
	url := "https://news.huoxing24.com/list/0"
	timeout := time.Duration(5 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取火星财经" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}
	request.Header.Add("User-Agent", `Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Mobile Safari/537.36`)
	res, err := client.Do(request)
	if err != nil {
		fmt.Println("抓取火星财经" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}
	defer res.Body.Close()
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("抓取火星财经" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}
	var allData []map[string]interface{}
	document.Find(".item-right").Each(func(i int, selection *goquery.Selection) {
		text := selection.Find("h5").Text()
		url, boolUrl := selection.Find("a").Attr("href")
		if boolUrl {
			allData = append(allData, map[string]interface{}{"title": text, "url": "https://news.huoxing24.com" + url})
		}
	})
	fmt.Printf("抓取dsdsdsdsdq：", allData)
	return allData
}

// yangtuobtc
func (spider Spider) GetYangtuo() []map[string]interface{} {
	url := "https://www.yangtuobtc.com/"
	timeout := time.Duration(5 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取羊驼财经" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}
	request.Header.Add("User-Agent", `Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Mobile Safari/537.36`)
	res, err := client.Do(request)
	if err != nil {
		fmt.Println("抓取羊驼财经" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}
	defer res.Body.Close()
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("抓取羊驼财经" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}
	var allData []map[string]interface{}
	document.Find(".list ul").Each(func(i int, selection *goquery.Selection) {
		text := selection.Find("a").Text()
		url, boolUrl := selection.Find("a").Attr("href")
		if boolUrl {
			allData = append(allData, map[string]interface{}{"title": text, "url": "https://www.yangtuobtc.com/" + url})
		}
	})
	fmt.Printf("抓取dsdsdsdsdq：", allData)
	return allData
}

// chaindd
func (spider Spider) GetChaindd() []map[string]interface{} {
	url := "https://www.chaindd.com/"
	timeout := time.Duration(5 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取链得得" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}
	request.Header.Add("User-Agent", `Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Mobile Safari/537.36`)
	res, err := client.Do(request)
	if err != nil {
		fmt.Println("抓取链得得" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}
	defer res.Body.Close()
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("抓取链得得" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}
	var allData []map[string]interface{}
	document.Find(".post_part").Each(func(i int, selection *goquery.Selection) {
		text := selection.Find("a").Text()
		text = strings.Replace(text, "\n", "", -1)
		text = strings.Replace(text, " ", "", -1)
		url, boolUrl := selection.Find("a").Attr("href")
		if boolUrl {
			allData = append(allData, map[string]interface{}{"title": text, "url": "https://www.chaindd.com/" + url})
		}
	})
	fmt.Printf("抓取dsdsdsdsdq：", allData)
	return allData
}

// bitcoin86
func (spider Spider) GetBitcoin86() []map[string]interface{} {
	url := "http://www.bitcoin86.com/"
	timeout := time.Duration(5 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取bitcoin86" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}
	request.Header.Add("User-Agent", `Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Mobile Safari/537.36`)
	res, err := client.Do(request)
	if err != nil {
		fmt.Println("抓取bitcoin86" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}
	defer res.Body.Close()
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("抓取bitcoin86" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}
	var allData []map[string]interface{}
	document.Find(".excerpt").Each(func(i int, selection *goquery.Selection) {
		text := selection.Find("a").Text()
		url, boolUrl := selection.Find("a").Attr("href")
		if boolUrl {
			allData = append(allData, map[string]interface{}{"title": text, "url": "http://www.bitcoin86.com/" + url})
		}
	})
	fmt.Printf("抓取dsdsdsdsdq：", allData)
	return allData
}

// people
func (spider Spider) GetPeople() []map[string]interface{} {
	url := "http://capital.people.com.cn/"
	timeout := time.Duration(5 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取人民创投" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}
	request.Header.Add("User-Agent", `Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Mobile Safari/537.36`)
	res, err := client.Do(request)
	if err != nil {
		fmt.Println("抓取人民创投" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}
	defer res.Body.Close()
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("抓取人民创投" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}
	var allData []map[string]interface{}
	document.Find(".hdNews strong").Each(func(i int, selection *goquery.Selection) {
		text := selection.Find("a").Text()
		text, _ = decodeToGBK(text)
		url, boolUrl := selection.Find("a").Attr("href")
		if boolUrl {
			allData = append(allData, map[string]interface{}{"title": text, "url": "http://capital.people.com.cn/" + url})
		}
	})
	fmt.Printf("抓取dsdsdsdsdq：", allData)
	return allData
}

// feixiaohao
func (spider Spider) GetFeixiaohao() []map[string]interface{} {
	url := "https://www.feixiaohao.com/news/"
	timeout := time.Duration(5 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取非小号" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}
	request.Header.Add("User-Agent", `Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Mobile Safari/537.36`)
	res, err := client.Do(request)
	if err != nil {
		fmt.Println("抓取非小号" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}
	defer res.Body.Close()
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("抓取非小号" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}
	var allData []map[string]interface{}
	document.Find(".info").Each(func(i int, selection *goquery.Selection) {
		text := selection.Find("a").Text()
		url, boolUrl := selection.Find("a").Attr("href")
		if boolUrl {
			allData = append(allData, map[string]interface{}{"title": text, "url": "https://www.feixiaohao.com/news/" + url})
		}
	})
	fmt.Printf("抓取dsdsdsdsdq：", allData)
	return allData
}

// orange
func (spider Spider) GetOrange() []map[string]interface{} {
	url := "https://orange.xyz/"
	timeout := time.Duration(5 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取橙皮书" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}
	request.Header.Add("User-Agent", `Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Mobile Safari/537.36`)
	res, err := client.Do(request)
	if err != nil {
		fmt.Println("抓取橙皮书" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}
	defer res.Body.Close()
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("抓取橙皮书" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}
	var allData []map[string]interface{}
	document.Find(".article-wrapper").Each(func(i int, selection *goquery.Selection) {
		text := selection.Find(".b-title").Text()
		text = strings.Replace(text, " ", "", -1)
		url, boolUrl := selection.Find("a").Attr("href")
		if boolUrl {
			allData = append(allData, map[string]interface{}{"title": text, "url": "https://orange.xyz/" + url})
		}
	})
	fmt.Printf("抓取dsdsdsdsdq：", allData)
	return allData
}

// ccvalue
func (spider Spider) GetCcvalue() []map[string]interface{} {
	url := "https://www.ccvalue.cn/"
	timeout := time.Duration(10 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取碳链价值" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}
	request.Header.Add("User-Agent", `Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Mobile Safari/537.36`)
	res, err := client.Do(request)
	if err != nil {
		fmt.Println("抓取碳链价值" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}
	defer res.Body.Close()
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("抓取碳链价值" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}
	var allData []map[string]interface{}
	document.Find(".news-item .title").Each(func(i int, selection *goquery.Selection) {
		text := selection.Find("a").Text()
		//text = strings.Replace(text, " ", "", -1)
		url, boolUrl := selection.Find("a").Attr("href")
		if boolUrl {
			allData = append(allData, map[string]interface{}{"title": text, "url": "https://www.ccvalue.cn/" + url})
		}
	})
	fmt.Printf("抓取dsdsdsdsdq：", allData)
	return allData
}

// ethfans
func (spider Spider) GetEthfans() []map[string]interface{} {
	url := "https://ethfans.org"
	timeout := time.Duration(10 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取ethfans" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}
	request.Header.Add("User-Agent", `Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Mobile Safari/537.36`)
	res, err := client.Do(request)
	if err != nil {
		fmt.Println("抓取ethfans" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}
	defer res.Body.Close()
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("抓取ethfans" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}
	var allData []map[string]interface{}
	document.Find(".post-info").Each(func(i int, selection *goquery.Selection) {
		text := selection.Find("a").Text()
		//text = strings.Replace(text, " ", "", -1)
		url, boolUrl := selection.Find("a").Attr("href")
		if boolUrl {
			allData = append(allData, map[string]interface{}{"title": text, "url": "https://ethfans.org" + url})
		}
	})
	fmt.Printf("抓取dsdsdsdsdq：", allData)
	return allData
}

// qukuaiwang
func (spider Spider) GetQukuaiwang() []map[string]interface{} {
	url := "http://www.qukuaiwang.com.cn/"
	timeout := time.Duration(10 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取ethfans" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}
	request.Header.Add("User-Agent", `Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Mobile Safari/537.36`)
	res, err := client.Do(request)
	if err != nil {
		fmt.Println("抓取ethfans" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}
	defer res.Body.Close()
	document, err := goquery.NewDocumentFromReader(res.Body)
	//fmt.Printf("抓取d999999：",document.Text())
	if err != nil {
		fmt.Println("抓取ethfans" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}
	var allData []map[string]interface{}
	document.Find(".cooperation").Each(func(i int, selection *goquery.Selection) {
		text := selection.Find("a").Text()
		//text = strings.Replace(text, " ", "", -1)
		url, boolUrl := selection.Find("a").Attr("href")
		if boolUrl {
			allData = append(allData, map[string]interface{}{"title": text, "url": "http://www.qukuaiwang.com.cn/" + url})
		}
	})
	fmt.Printf("抓取dsdsdsdsdq：", allData)
	return allData
}

// fengli
func (spider Spider) GetFengli() []map[string]interface{} {
	url := "https://www.fengli.com"
	timeout := time.Duration(5 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取fengli" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}
	request.Header.Add("User-Agent", `Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Mobile Safari/537.36`)
	res, err := client.Do(request)
	if err != nil {
		fmt.Println("抓取fengli" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}
	defer res.Body.Close()
	document, err := goquery.NewDocumentFromReader(res.Body)
	//fmt.Printf("抓取d999999：",document.Text())
	if err != nil {
		fmt.Println("抓取fengli" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}
	var allData []map[string]interface{}
	document.Find(".news-list-ul li").Each(func(i int, selection *goquery.Selection) {
		text := selection.Find("a").Text()
		url, boolUrl := selection.Find("a").Attr("href")
		if boolUrl {
			allData = append(allData, map[string]interface{}{"title": text, "url": "https://www.fengli.com" + url})
		}
	})
	fmt.Printf("抓取dsdsdsdsdq：", allData)
	return allData
}

// theblockbeats
func (spider Spider) GetTheblockbeats() []map[string]interface{} {
	url := "https://www.theblockbeats.com"
	timeout := time.Duration(5 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取theblockbeats" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}
	request.Header.Add("User-Agent", `Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Mobile Safari/537.36`)
	res, err := client.Do(request)
	if err != nil {
		fmt.Println("抓取theblockbeats" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}
	defer res.Body.Close()
	document, err := goquery.NewDocumentFromReader(res.Body)
	//fmt.Printf("抓取d999999：",document.Text())
	if err != nil {
		fmt.Println("抓取theblockbeats" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}
	var allData []map[string]interface{}
	document.Find(".boxxxx").Each(func(i int, selection *goquery.Selection) {
		text := selection.Find(".im_title").Text()
		url, boolUrl := selection.Find("a").Attr("href")
		if boolUrl {
			allData = append(allData, map[string]interface{}{"title": text, "url": "https://www.theblockbeats.com" + url})
		}
	})
	fmt.Printf("抓取dsdsdsdsdq：", allData)
	return allData
}

// cailu
func (spider Spider) GetCailu() []map[string]interface{} {
	url := "https://apisns.cailuw.com/v1/posts/hot?type=hot&limit=20&page=1&now_time=&now_id="
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("抓取" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}

	str, _ := ioutil.ReadAll(res.Body)
	js, err2 := simplejson.NewJson(str)
	//fmt.Printf("抓取btc：",js)
	if err2 != nil {
		fmt.Println("抓取" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}
	var allData []map[string]interface{}
	i := 1
	for i < 25 {
		// 如果json是多级  则依次get获取
		test := js.Get("data").Get("data").GetIndex(i).MustMap()

		if test["title"] != nil {

			// /long-post/12910789382554864?h=xG4GXD7Z2dgm

			allData = append(allData, map[string]interface{}{"title": test["title"], "url": "https://www.cailuw.com/long-post/" + test["id"].(string) + "?h=" + test["user_id"].(string)})
		}
		i++
	}

	fmt.Printf("抓取dsdsdsdsdq：", allData)
	return allData
}

// 8btc
func (spider Spider) GetBtc() []map[string]interface{} {
	url := "https://webapi.8btc.com/bbt_api/news/list?num=15"
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("抓取" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}

	str, _ := ioutil.ReadAll(res.Body)
	js, err2 := simplejson.NewJson(str)
	//fmt.Printf("抓取btc：",js)
	if err2 != nil {
		fmt.Println("抓取" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}
	var allData []map[string]interface{}
	i := 1
	for i < 25 {
		// 如果json是多级  则依次get获取
		test := js.Get("data").Get("list").GetIndex(i).MustMap()

		if test["title"] != nil {
			value, ok := test["id"].(json.Number)
			if !ok {
				fmt.Println("It's not ok for type string")
			}

			allData = append(allData, map[string]interface{}{"title": test["title"], "url": "https://www.8btc.com/article/" + value.String()})
		}
		i++
	}

	fmt.Printf("抓取dsdsdsdsdq：", allData)

	return allData
}

// jinse
func (spider Spider) GetJinse() []map[string]interface{} {
	url := "https://api.jinse.com/v6/www/information/list?catelogue_key=www&limit=23&information_id=11555&flag=down"
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("抓取" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}

	str, _ := ioutil.ReadAll(res.Body)
	js, err2 := simplejson.NewJson(str)
	//fmt.Printf("抓取btc：",js)
	if err2 != nil {
		fmt.Println("抓取" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}
	var allData []map[string]interface{}
	i := 1
	for i < 25 {
		// 如果json是多级  则依次get获取
		test := js.Get("list").GetIndex(i).Get("extra").MustMap()

		if test["title"] != nil {
			allData = append(allData, map[string]interface{}{"title": test["title"], "url": test["topic_url"]})
		}
		i++
	}

	fmt.Printf("抓取dsdsdsdsdq：", allData)

	return allData
}

// xcong
func (spider Spider) GetXcong() []map[string]interface{} {
	url := "https://cong-api.xcong.com/apiv1/dashboard/chosen_page?limit=20&cursor="
	timeout := time.Duration(5 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取theblockbeats" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}
	request.Header.Add("User-Agent", `Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1`)
	request.Header.Add("x-client-type", `mweb`)
	request.Header.Add("x-device-id", `mcong-16eda5f5-962f-a669-4e0b-13a7c62f7534`)
	request.Header.Add("x-ivanka-app", `cong|web|1.0.4|11.1|0`)
	request.Header.Add("x-ivanka-platform", `cong-platform`)
	request.Header.Add("Referer", `https://m.xcong.com/news/1005678`)
	res, err := client.Do(request)
	if err != nil {
		fmt.Println("抓取theblockbeats" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}
	defer res.Body.Close()
	str, _ := ioutil.ReadAll(res.Body)
	js, err2 := simplejson.NewJson(str)
	//fmt.Printf("抓取btc：",js)
	if err2 != nil {
		fmt.Println("抓取" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}
	var allData []map[string]interface{}
	i := 1
	for i < 25 {
		// 如果json是多级  则依次get获取
		test := js.Get("data").Get("items").GetIndex(i).Get("resource").MustMap()
		if test["title"] != nil {
			value, ok := test["id"].(json.Number)
			if !ok {
				fmt.Println("It's not ok for type string")
			}

			allData = append(allData, map[string]interface{}{"title": test["title"], "url": "https://xcong.com/articles/" + value.String()})
		}
		i++
	}

	fmt.Printf("抓取dsdsdsdsdq：", allData)
	return allData
}

// bishijie
func (spider Spider) GetBishijie() []map[string]interface{} {
	//currentTime:=time.Now().Unix()
	// signature 问题

	//fmt.Printf("抓取dsdsdsdsdq：",currentTime)
	url := "https://iapi.bishijie.com/news/list?page=1&size=20&signature=f6a09bb905b6e8b139a42710e82c1890&ts=1576029537"
	timeout := time.Duration(5 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取bishijie" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}
	request.Header.Add("User-Agent", `Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1`)
	request.Header.Add("packetBit", `h5`)
	request.Header.Add("lang", `zh-cn`)
	request.Header.Add("Origin", `https://m.bishijie.com`)
	request.Header.Add("timeZone", `+08:00`)
	request.Header.Add("version", `1.0`)
	request.Header.Add("Sec-Fetch-Mode", `cors`)
	request.Header.Add("Referer", `https://m.bishijie.com/`)
	request.Header.Add("uuid", `469E5B85667247BE181356E5FBA16594`)
	res, err := client.Do(request)
	if err != nil {
		fmt.Println("抓取theblockbeats" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}
	defer res.Body.Close()
	str, _ := ioutil.ReadAll(res.Body)
	js, err2 := simplejson.NewJson(str)
	fmt.Printf("抓取btc：", js)
	if err2 != nil {
		fmt.Println("抓取" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}
	var allData []map[string]interface{}
	i := 1
	for i < 25 {
		// 如果json是多级  则依次get获取
		test := js.Get("data").Get("list").GetIndex(i).MustMap()
		if test["title"] != nil {
			value, ok := test["news_id"].(json.Number)
			if !ok {
				fmt.Println("It's not ok for type string")
			}

			allData = append(allData, map[string]interface{}{"title": test["title"], "url": "https://www.bishijie.com/shendu_" + value.String()})
		}
		i++
	}

	fmt.Printf("抓取dsdsdsdsdq：", allData)
	return allData
}

type Pfatures struct {
	Ocode   int    `json:"o_code"`
	Omsg    int    `json:"o_msg"`
	Ocursor []byte `json:"o_cursor"`
}

/**
出现乱码转换为gbk格式
simplifiedchinese包
*/
func decodeToGBK(text string) (string, error) {

	dst := make([]byte, len(text)*2)
	tr := simplifiedchinese.GB18030.NewDecoder()
	nDst, _, err := tr.Transform(dst, []byte(text), true)
	if err != nil {
		return text, err
	}

	return string(dst[:nDst]), nil
}

/**
执行每个分类数据
*/
func ExecGetData(spider Spider) {
	// 抓取数据开始  传入当前开始抓取的类别
	reflectValue := reflect.ValueOf(spider)
	dataType := reflectValue.MethodByName("Get" + spider.DataType)
	data := dataType.Call(nil)
	//fmt.Printf("抓取数据qqqqq：",data)
	//fmt.Println()
	originData := data[0].Interface().([]map[string]interface{})
	fmt.Printf("抓取数据999999：", spider.DataType)
	start := time.Now()
	Common.MySql{}.GetConn().Where(map[string]string{"dataType": spider.DataType}).Update("hotData2", map[string]string{"str": SaveDataToJson(originData)})
	group.Done()
	seconds := time.Since(start).Seconds()
	fmt.Printf("抓取数据3333333333：", originData)
	//fmt.Println()
	fmt.Printf("耗费 %.2fs 秒完成抓取%s", seconds, spider.DataType)
	//fmt.Println()

}

func reverse7(x int) int {
	tmp := 0
	for x != 0 {
		tmp = tmp*10 + x%10
		x = x / 10
	}
	if tmp > 1<<31-1 || tmp < -(1<<31) {
		return 0
	}
	return tmp
}

var group sync.WaitGroup

func main() {
	allData := []string{
		//"Chainnews",
		//"Huoxing",
		//"Yangtuo",
		//"Chaindd",
		//"Bitcoin86",
		//"People",
		//"Feixiaohao",
		//"Orange",
		//"Ccvalue",
		//"Ethfans",
		//"Qukuaiwang",
		//"Fengli",
		//"Theblockbeats",
		//"Btc",
		//"Xcong",
		//"Jinse",
		//"Bishijie",
		//"Cailu",
		//"Test",
	}
	res := reverse7(784566)
	fmt.Println("sdsdsdsdsd", res)
	//fmt.Println("开始抓取" + strconv.Itoa(len(allData)) + "种数据类型")
	group.Add(len(allData))
	var spider Spider
	for _, value := range allData {
		//fmt.Println("开始抓取1111" + value)
		spider = Spider{DataType: value}
		//fmt.Print(spider)
		go ExecGetData(spider)
	}
	group.Wait()
	//fmt.Print("完成抓取22222")

}
