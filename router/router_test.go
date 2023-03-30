package router

import (
	"testing"
)

func TestMatch(t *testing.T) {
	// 测试路由匹配
	r := RInfo{
		Path: "/api/1.0/*",
	}
	if !r.Match("/api/1.0/123/456") {
		t.Error("匹配错误")
	}
	if !r.Match("/api/1.0/?") {
		t.Error("匹配错误")
	}
	if !r.Match("/api/1.0/") {
		t.Error("匹配错误")
	}
}
