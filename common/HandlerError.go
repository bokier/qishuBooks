package common

import "fmt"

/*
	common.HandlerError() 简化全局错误重复判断流程
	1. 传入 错误和原因
    2. 返回 错误和原因
*/

func HandlerError(err error,why string) {
	if err != nil {
		fmt.Println(why,err)
	}
}