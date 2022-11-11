package main

import "testing"

func TestConfig_getHoldings(t *testing.T) {
	all, err := testApp.currentHoldings()
	if err != nil {
		t.Error("从数据库中获取数据失败: ", err)
	}

	if len(all) != 2 {
		t.Error("数据库内容条目错误: ", len(all))
	}
}

func TestConfig_getHoldingsSlice(t *testing.T) {
	slice := testApp.getHoldingsSlice()

	if len(slice) != 3 {
		t.Error("项目条数不为3，项目条数为:", len(slice))
	}
}
