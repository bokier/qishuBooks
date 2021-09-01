package engine

import (
	"fmt"
	"qishuBook/common"
	"qishuBook/controller"
	"sync"
)

// 这里是项目总控
var (
	// 用于存放区全部分页信息的数据管道
	chanBookInfoUrls    chan string
	wg 					sync.WaitGroup
)

func RunEngineV1(url string) {
	// 初始化管道
	chanBookInfoUrls = make(chan string,100000)
	urlStr := controller.GetBookTypesV1(url)
	go func() {
		for _,v := range urlStr {
			GetBookPages(v)
		}
	}()

	for i := 0;i<10; i++ {
		wg.Add(1)
		go ChanToInfos()
	}
	wg.Wait()
}

// 这里定义的是获取全部类型的全部分页
func GetBookPages(url string)  {
	bytes := common.AnalysisHttpUrl(url)
	matches := common.AnalysisRegular(bytes,common.TypeInfoUrl)
	for _,m := range matches {
		chanBookInfoUrls <- common.SplicingUrl+m[1]
	}
}

// 这里定义的是单一页面的作品信息,其中包括 1，书籍作者，2，书籍大小，3，书籍等级，4，网站的更新时间。
func GetBookInfos(url string) {
	//bytes := common.AnalysisHttpUrl(url)
	//matches := common.AnalysisRegular(bytes,common.BookInfoRule)
	//for _,m := range matches {
	//	fmt.Printf("Author:%v, Size:%vMB, Level:%v, Date:%v\n",m[1],m[2],m[3],m[7])
	//}
	bytes := common.AnalysisHttpUrl(url)
	matches := common.AnalysisRegular(bytes,common.TBookInfoRuleAll)
	for _,m := range matches {
		fmt.Printf("Author:%v, BookName:%v, Size:%vMB, Level:%v, Date:%v\n",m[1],m[9],m[2],m[3],m[4])
	}
}

// 这里是用于测试匹配规则
func GetBookInfosTest(url string) {
	bytes := common.AnalysisHttpUrl(url)
	fmt.Printf("%v\n",string(bytes))
	matches := common.AnalysisRegular(bytes,common.TBookInfoRuleAll)
	fmt.Printf("length:%v\n",len(matches))
	for _,m := range matches {
		fmt.Printf("Author:%v, BookName:%v, Size:%vMB, Level:%v, Date:%v\n",m[1],m[9],m[2],m[3],m[4])
	}
}

func ChanToInfos() {
	for url := range chanBookInfoUrls {
		GetBookInfos(url)
    }
  wg.Done()
}
