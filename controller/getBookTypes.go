package controller

import (
	"qishuBook/common"
)

/*
	controller.GetBookTypes() 获取书籍类型页面
	1. 传入 common.analysisHttpUrl 转化的 bytes
    2.1 返回类型名称
 	2.2 返回类型名称对应的url地址
 */


// 这里获取 类型和url 并返回
func GetBookTypes(url string) map[string]string {
	bytes := common.AnalysisHttpUrl(url)
	matches := common.AnalysisRegular(bytes,common.SYUrl)

	BooksInfoMap := make(map[string]string,len(matches)+1)
	for _,m := range matches {
		BooksInfoMap[m[3]] = common.SplicingUrl+m[1]
	}
	return BooksInfoMap
}

func GetBookTypesV1(url string) (urls []string) {
	bytes := common.AnalysisHttpUrl(url)
	matches := common.AnalysisRegular(bytes,common.SYUrl)
	// fmt.Printf("[info] 共找到 %d 个类型\n",len(matches))
	for _,m := range matches {
		url := common.SplicingUrl+m[1]
		urls = append(urls,url)
	}
	return
}
