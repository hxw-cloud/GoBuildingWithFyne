package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"time"
)

func (app *Config) makeUI() {
	openPrice, currentPrice, priceChange := app.getPriceText()

	priceContent := container.NewGridWithColumns(3, openPrice, currentPrice, priceChange)

	app.PriceContainer = priceContent

	//工具栏
	toolBar := app.getToolBar( /*app.MainWindow*/ )
	app.ToolBar = toolBar

	priceTabContent := app.pricesTab()
	holdingsTabContent := app.holdingsTab()
	//应用标签
	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Prices", theme.HomeIcon(), priceTabContent),
		container.NewTabItemWithIcon("Holdings", theme.InfoIcon(), holdingsTabContent),
	)
	tabs.SetTabLocation(container.TabLocationTop)

	finalContent := container.NewVBox(priceContent, toolBar, tabs)
	app.MainWindow.SetContent(finalContent)

	go func() {
		for range time.Tick(30 * time.Second) {
			app.refreshPriceContent()
		}
	}()
}

func (app *Config) refreshPriceContent() {
	app.InfoLog.Print("refreshing prices")
	open, current, change := app.getPriceText()
	app.PriceContainer.Objects = []fyne.CanvasObject{open, current, change}
	app.PriceContainer.Refresh()

	chart := app.getChart()
	//log.Printf("%v\n", chart)
	app.PriceChartContainer.Objects = []fyne.CanvasObject{chart}
	//log.Println(app.PriceChartContainer.Objects)
	app.PriceChartContainer.Refresh()

}

func (app *Config) refreshHoldingsTable() {
	app.Holdings = app.getHoldingsSlice()
	app.HoldingsTable.Refresh()
}
