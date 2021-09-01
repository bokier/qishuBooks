package common

import (
	"io"
	"net/http"
)

/*
	common.AnalysisHttpUrl() 将http页面转化成bytes
	1. 传入 url 地址
    2. 返回 bytes
 */

func AnalysisHttpUrl(url string)[]byte {
	resp,err := http.Get(url)
	HandlerError(err,"[error] find a err in http.Get ")

	// 延迟关闭打开的body
	defer resp.Body.Close()

	bytes,err := io.ReadAll(resp.Body)
	HandlerError(err,"[error] find a err in io.ReadAll ")
	return bytes
}