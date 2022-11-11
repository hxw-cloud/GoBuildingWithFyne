package main

import (
	"Gold/repository"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

func (app *Config) holdingsTab() *fyne.Container {
	app.Holdings = app.getHoldingsSlice()
	app.HoldingsTable = app.getHoldingsTable()

	holdingsContainer := container.NewBorder(
		nil,
		nil,
		nil,
		nil,
		container.NewAdaptiveGrid(1, app.HoldingsTable),
	)
	return holdingsContainer
}

func (app *Config) getHoldingsTable() *widget.Table {

	table := widget.NewTable(
		func() (int, int) {
			return len(app.Holdings), len(app.Holdings[0])
		},
		func() fyne.CanvasObject {
			ctr := container.NewVBox(widget.NewLabel(""))
			return ctr
		},
		func(id widget.TableCellID, object fyne.CanvasObject) {
			if id.Col == (len(app.Holdings[0])-1) && id.Row != 0 {
				//按钮
				w := widget.NewButtonWithIcon(" Delete ", theme.DeleteIcon(), func() {
					dialog.ShowConfirm("Delete?", "", func(deleted bool) {
						if deleted {
							ids, _ := strconv.Atoi(app.Holdings[id.Row][0].(string))
							err := app.DB.DeleteHolding(int64(ids))
							if err != nil {
								app.ErrorLog.Println(err)
							}
						}
						app.refreshHoldingsTable()
					}, app.MainWindow)
				})
				w.Importance = widget.HighImportance
				object.(*fyne.Container).Objects = []fyne.CanvasObject{
					w,
				}
			} else {
				//处理文本使文本符合要求
				object.(*fyne.Container).Objects = []fyne.CanvasObject{
					widget.NewLabel(app.Holdings[id.Row][id.Col].(string)),
				}
			}
		})
	colWidths := []float32{50, 200, 200, 110}
	for i := 0; i < len(colWidths); i++ {
		table.SetColumnWidth(i, colWidths[i])
	}
	return table
}

func (app *Config) getHoldingsSlice() [][]interface{} {
	var slices [][]interface{}
	holdings, err := app.currentHoldings()
	if err != nil {
		app.ErrorLog.Println(err)
	}
	slices = append(slices, []interface{}{"ID", "Amount", "Price", "Date", "Delete?"})
	for _, x := range holdings {
		var currentRow []interface{}
		currentRow = append(currentRow, strconv.FormatInt(x.ID, 10))
		currentRow = append(currentRow, fmt.Sprintf("%d toz", x.Amount))
		currentRow = append(currentRow, fmt.Sprintf("$%.2f", float32(x.PurchasePrice)))
		currentRow = append(currentRow, x.PurchaseDate.Format("2006-01-02"))
		currentRow = append(currentRow, widget.NewButton(" Delete ", func() {}))
		slices = append(slices, currentRow)
	}
	return slices
}

func (app *Config) currentHoldings() ([]repository.Holdings, error) {
	holdings, err := app.DB.AllHoldings()
	if err != nil {
		app.ErrorLog.Println(err)
		return nil, err
	}
	return holdings, nil
}
