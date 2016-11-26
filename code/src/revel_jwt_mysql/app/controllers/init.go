package controllers

import "github.com/revel/revel"

func init() {
	revel.OnAppStart(InitDB)
	// 下記でDBの初期設定を行なっている。
	revel.InterceptMethod((*GorpController).Begin, revel.BEFORE)
	revel.InterceptMethod((*GorpController).Commit, revel.AFTER)
	revel.InterceptMethod((*GorpController).Rollback, revel.FINALLY)
}