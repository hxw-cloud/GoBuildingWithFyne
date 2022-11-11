package main

import "testing"

func TestApp_getToolBar(t *testing.T) {
	toolbar := testApp.getToolBar()

	if len(toolbar.Items) != 4 {
		t.Error("wrong number of items in toolbar:", len(toolbar.Items))
	}
}
