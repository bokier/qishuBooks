package common

import "regexp"

/*
	common.AnalysisRegular() 用来封装匹配和解析函数
	1. 传入 common.analysisHttpRrl 中的 bytes
	2. 输出 经过正则匹配后的内容， 类型是[][]string
 */

func AnalysisRegular(bodyBytes []byte,urlRules string) [][]string {
	regexp := regexp.MustCompile(urlRules)
	return regexp.FindAllStringSubmatch(string(bodyBytes),-1)
}