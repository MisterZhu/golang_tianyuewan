package xcxapi

import (
	"encoding/json"
	"fmt"
	"gindiary/model"
	"gindiary/response"
	"gindiary/util"
	"gindiary/util/errmsg"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

// type EeapiData struct {
// 	Data model.XcxAnalyModel `json:"data"`
// }

type EeapiData struct {
	Status int    `json:"status"`
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Type   string `json:"type"`
	Data   struct {
		// OpenId        string
		// UserId        string
		Title         string   `json:"title"`
		Cover         string   `json:"cover"`
		DownloadImage string   `json:"download_image"`
		Video         string   `json:"video"`
		URL           string   `json:"url"`
		Down          string   `json:"down"`
		Images        []string `json:"images"`
		BigFile       bool     `json:"bigFile"`
	} `json:"data"`
}

/*
免费解析URL
*/
func XcxFreeAnalysisURL(c *gin.Context) {
	Url := c.PostForm("url")
	fmt.Printf("\nUrl = %s\n", Url)

	// 使用正则表达式提取链接
	re := regexp.MustCompile(`https?://[^\s]+`)
	links := re.FindAllString(Url, -1)
	if len(links) <= 0 {
		response.Fail(c, errmsg.GetErrMsg(errmsg.ERR_analys_fail), nil)
		return
	}
	fmt.Printf("ShouldBind-url:%s\n", links)
	urlStr := links[0]
	url := "https://dlpanda.com/xiaohongshu?url=%s&token=G7eRpMaa"
	if strings.Contains(urlStr, "v.douyin.com") {
		url = "https://dlpanda.com/en?url=%s&token=G7eRpMaa"
	}
	url = fmt.Sprintf(url, links[0])

	// 发起请求
	res, err := http.Get(url)
	if err != nil {
		response.Fail(c, errmsg.GetErrMsg(errmsg.ERR_analys_fail), nil)
		fmt.Println("HTTP GET 请求失败：", err)
		return
	}
	defer res.Body.Close()

	// 读取网页内容
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	// 解析HTML标签，获取文案内容
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
	if err != nil {
		panic(err)
	}

	// 获取所有的图片链接
	images := []string{}
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		src, _ := s.Attr("src")
		images = append(images, src)
	})
	for i := 0; i < len(images); i++ {
		if strings.Contains(images[i], "https://dlpanda.com") {
			images = append(images[:i], images[i+1:]...)
			i--
		}
	}

	// 获取所有的视频链接
	videos := []string{}
	doc.Find("video source").Each(func(i int, s *goquery.Selection) {
		src, _ := s.Attr("src")
		videos = append(videos, src)
	})
	for i := 0; i < len(videos); i++ {
		if strings.Contains(videos[i], "https://dlpanda.com") {
			videos = append(videos[:i], videos[i+1:]...)
			i--
		}
	}
	fmt.Println("读取响应体数据images：", images)

	fmt.Println("读取响应体数据videos：", videos)
	if len(videos) > 0 {
		analyModel := model.XcxAnalyModel{
			OpenId:        "",
			UserId:        "",
			OriginURL:     Url,
			Title:         "",
			Cover:         "",
			DownloadImage: "",
			Video:         videos[0],
			URL:           "",
			Down:          "",
			Images:        "",
			BigFile:       false,
		}
		// 输出解析后的数据
		code := errmsg.SUCCSE
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"data": analyModel,
			"msg":  errmsg.GetErrMsg(code),
		})
	} else if len(images) > 0 {
		imagesStr := strings.Join(images, ",")
		analyModel := model.XcxAnalyModel{
			OpenId:        "",
			UserId:        "",
			OriginURL:     Url,
			Title:         "",
			Cover:         "",
			DownloadImage: "",
			Video:         "",
			URL:           "",
			Down:          "",
			Images:        imagesStr,
			BigFile:       false,
		}
		// 输出解析后的数据
		code := errmsg.SUCCSE
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"data": analyModel,
			"msg":  errmsg.GetErrMsg(code),
		})
	} else {
		fmt.Println("解析 URL 数据失败：", err)
		response.Fail(c, errmsg.GetErrMsg(errmsg.ERR_analys_fail), nil)
	}
}

/*
自用免费解析URL
*/
func XcxOneSelfFreeAnalysisURL(c *gin.Context) {
	Url := c.PostForm("url")
	fmt.Printf("\nUrl = %s\n", Url)
	xcxUser := c.Value("user").(model.XcxUser)
	if xcxUser.QueryCount > 0 {
		xcxUser.QueryCount -= 1
	} else {
		response.Fail(c, errmsg.GetErrMsg(errmsg.ERR_analys_undercount), nil)
		return
	}
	fmt.Printf("QueryCount-url:%d\n", xcxUser.QueryCount)

	// 使用正则表达式提取链接
	re := regexp.MustCompile(`https?://[^\s]+`)
	links := re.FindAllString(Url, -1)
	if len(links) <= 0 {
		response.Fail(c, errmsg.GetErrMsg(errmsg.ERR_analys_fail), nil)
		return
	}
	fmt.Printf("ShouldBind-url:%s\n", links)
	urlStr := links[0]
	url := "https://dlpanda.com/xiaohongshu?url=%s&token=G7eRpMaa"
	if strings.Contains(urlStr, "v.douyin.com") {
		url = "https://dlpanda.com/en?url=%s&token=G7eRpMaa"
	}
	url = fmt.Sprintf(url, links[0])

	// 发起请求
	res, err := http.Get(url)
	if err != nil {
		response.Fail(c, errmsg.GetErrMsg(errmsg.ERR_analys_fail), nil)
		fmt.Println("HTTP GET 请求失败：", err)
		return
	}
	defer res.Body.Close()

	// 读取网页内容
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	// 解析HTML标签，获取文案内容
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
	if err != nil {
		panic(err)
	}

	// 获取所有的图片链接
	images := []string{}
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		src, _ := s.Attr("src")
		images = append(images, src)
	})
	for i := 0; i < len(images); i++ {
		if strings.Contains(images[i], "https://dlpanda.com") {
			images = append(images[:i], images[i+1:]...)
			i--
		}
	}

	// 获取所有的视频链接
	videos := []string{}
	doc.Find("video source").Each(func(i int, s *goquery.Selection) {
		src, _ := s.Attr("src")
		videos = append(videos, src)
	})
	for i := 0; i < len(videos); i++ {
		if strings.Contains(videos[i], "https://dlpanda.com") {
			videos = append(videos[:i], videos[i+1:]...)
			i--
		}
	}
	fmt.Println("读取响应体数据images：", images)

	fmt.Println("读取响应体数据videos：", videos)
	// 获取上下文中小程序用户信息
	fmt.Printf("xcxUser :%s\n", xcxUser.UserId)

	if len(videos) > 0 {
		model.EditXcxUserQueryCount(&xcxUser)

		analyModel := model.XcxAnalyModel{
			OpenId:        xcxUser.OpenId,
			UserId:        xcxUser.UserId,
			OriginURL:     Url,
			Title:         "",
			Cover:         "",
			DownloadImage: "",
			Video:         videos[0],
			URL:           "",
			Down:          "",
			Images:        "",
			BigFile:       false,
		}
		model.CreateXcxAnaly(&analyModel)

		// 输出解析后的数据
		code := errmsg.SUCCSE
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"data": analyModel,
			"msg":  errmsg.GetErrMsg(code),
		})
	} else if len(images) > 0 {
		model.EditXcxUserQueryCount(&xcxUser)

		imagesStr := strings.Join(images, ",")
		analyModel := model.XcxAnalyModel{
			OpenId:        xcxUser.OpenId,
			UserId:        xcxUser.UserId,
			OriginURL:     Url,
			Title:         "",
			Cover:         "",
			DownloadImage: "",
			Video:         "",
			URL:           "",
			Down:          "",
			Images:        imagesStr,
			BigFile:       false,
		}
		model.CreateXcxAnaly(&analyModel)

		// 输出解析后的数据
		code := errmsg.SUCCSE
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"data": analyModel,
			"msg":  errmsg.GetErrMsg(code),
		})
	} else {
		fmt.Println("解析 URL 数据失败：", err)
		response.Fail(c, errmsg.GetErrMsg(errmsg.ERR_analys_fail), nil)
	}
}

/*
解析URL
*/
func XcxAnalysisURL(c *gin.Context) {

	Url := c.PostForm("url")
	fmt.Printf("\nUrl = %s\n", Url)

	// 使用正则表达式提取链接
	re := regexp.MustCompile(`https?://[^\s]+`)
	links := re.FindAllString(Url, -1)
	if len(links) <= 0 {
		response.Fail(c, errmsg.GetErrMsg(errmsg.ERR_analys_fail), nil)
		return
	}
	fmt.Printf("ShouldBind-url:%s\n", links[0])
	url := "http://eeapi.cn/api/video/%s/%s/?url=%s"
	url = fmt.Sprintf(url, util.EEapiUToken, util.EEapiUId, links)

	// 发起请求
	res, err := http.Get(url)
	if err != nil {
		response.Fail(c, errmsg.GetErrMsg(errmsg.ERR_analys_fail), nil)
		fmt.Println("HTTP GET 请求失败：", err)
		return
	}
	defer res.Body.Close()

	// 读取响应体的数据
	body, err := io.ReadAll(res.Body)
	if err != nil {
		response.Fail(c, errmsg.GetErrMsg(errmsg.ERR_analys_fail), nil)
		fmt.Println("读取响应体数据失败：", err)
		return
	}

	// 解析 JSON 数据
	var eeapiModel EeapiData
	err = json.Unmarshal(body, &eeapiModel)
	if err != nil {
		response.Fail(c, errmsg.GetErrMsg(errmsg.ERR_analys_fail), nil)
		fmt.Println("解析 JSON 数据失败：", err)
		return
	}

	if eeapiModel.Status != 101 {
		fmt.Println("解析 URL 数据失败：", err)
		response.Fail(c, errmsg.GetErrMsg(errmsg.ERR_analys_fail), nil)
		return
	}

	jsonImages := strings.Join(eeapiModel.Data.Images, ",")

	// 获取上下文中小程序用户信息
	xcxUser := c.Value("user").(model.XcxUser)
	if xcxUser.QueryCount > 0 {
		xcxUser.QueryCount -= 1
	} else {
		response.Fail(c, errmsg.GetErrMsg(errmsg.ERR_analys_undercount), nil)
		return
	}
	fmt.Printf("QueryCount-url:%d\n", xcxUser.QueryCount)

	model.EditXcxUserQueryCount(&xcxUser)
	analyModel := model.XcxAnalyModel{
		OpenId:        xcxUser.OpenId,
		UserId:        xcxUser.UserId,
		OriginURL:     Url,
		Title:         eeapiModel.Data.Title,
		Cover:         eeapiModel.Data.Cover,
		DownloadImage: eeapiModel.Data.DownloadImage,
		Video:         eeapiModel.Data.Video,
		URL:           eeapiModel.Data.URL,
		Down:          eeapiModel.Data.Down,
		Images:        jsonImages,
		BigFile:       eeapiModel.Data.BigFile,
	}

	// 输出解析后的数据
	fmt.Println(analyModel.UserId, eeapiModel.Data.Title)
	model.CreateXcxAnaly(&analyModel)
	code := errmsg.SUCCSE
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": analyModel,
		"msg":  errmsg.GetErrMsg(code),
	})
}

// 查询分类
func XcxGetAnalysis(c *gin.Context) {
	size, _ := strconv.Atoi(c.PostForm("size"))
	page, _ := strconv.Atoi(c.PostForm("page"))
	fmt.Printf("\nsize :%d\n", size)
	fmt.Printf("\npage :%d\n", page)

	switch {
	case size > 100:
		size = 100
	case size <= 0:
		size = 10
	}
	if page == 0 {
		page = 1
	}
	xcxUser := c.Value("user").(model.XcxUser)

	data := model.GetAnalys(size, page, xcxUser.UserId)
	code := errmsg.SUCCSE
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
		"msg":  errmsg.GetErrMsg(code),
	})
}
