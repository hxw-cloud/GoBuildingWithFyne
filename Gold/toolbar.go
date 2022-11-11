package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func (app *Config) getToolBar( /*windows fyne.Window*/ ) *widget.Toolbar {
	toolBar := widget.NewToolbar(
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			app.addHoldingsDialog()
		}),
		widget.NewToolbarAction(theme.ViewRefreshIcon(), func() {
			app.refreshPriceContent()
		}),
		widget.NewToolbarAction(theme.SettingsIcon(), func() {

		}),
	)
	return toolBar
}
func (app *Config) addHoldingsDialog() dialog.Dialog {
	addAmountEntry := widget.NewEntry()
	purchaseDataEntry := widget.NewEntry()
	purchasePriceEntry := widget.NewEntry()
	app.AddHoldingsPurchaseAmountEntry = addAmountEntry
	app.AddHoldingsPurchaseDataEntry = purchaseDataEntry
	app.AddHoldingsPurchasePriceEntry = purchasePriceEntry

	purchaseDataEntry.PlaceHolder = "YYYY-MM-DD"
	addForm := dialog.NewForm("Add Gold", "Add", "Cancel", []*widget.FormItem{
		{Text: "Amount in toz", Widget: addAmountEntry}, {Text: "Purchase Price", Widget: purchasePriceEntry}, {Text: "Purchase Date", Widget: purchaseDataEntry},
	}, func(valid bool) {
		if valid {

		}
	}, app.MainWindow)

	addForm.Resize(fyne.Size{Width: 400})
	addForm.Show()

	return addForm
}
