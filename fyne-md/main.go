package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"io/ioutil"
	"strings"
)

type config struct {
	EditWidget *widget.Entry
	PreviewWidget *widget.RichText
	CurrentFile fyne.URI
	SaveMenuItem *fyne.MenuItem
}

var cfg config

func main() {
	//创建一个app
	apps := app.New()
	apps.Settings().SetTheme(&myTheme{})
	//创建一个app窗口
	windows := apps.NewWindow("MarkDown")
	//获取用户信息
	edit,preview := cfg.makeUI()
	cfg.createMenuItems(windows)
	//设置窗口内容
	windows.SetContent(container.NewHSplit(edit,preview))
	//显示窗口，运行程序
	windows.Resize(fyne.Size{Width : 800,Height: 500})
	windows.CenterOnScreen()
	windows.ShowAndRun()
}

func (app *config) makeUI() (*widget.Entry,*widget.RichText) {
	edit := widget.NewMultiLineEntry()
	preview := widget.NewRichTextFromMarkdown("Please Input:")
	app.EditWidget =edit
	app.PreviewWidget=preview
	edit.OnChanged =preview.ParseMarkdown
	return edit,preview
}

func (app *config ) createMenuItems(windows fyne.Window)  {
	openMenuItem := fyne.NewMenuItem("Open...", app.openFunc(windows))
	saveMenuItem := fyne.NewMenuItem("Save", app.saveFunc(windows))
	app.SaveMenuItem = saveMenuItem
	app.SaveMenuItem.Disabled = true
	saveAsMenuItem := fyne.NewMenuItem("Save as...", app.saveAsFunc(windows))
	fileMenu := fyne.NewMenu("File",openMenuItem,saveMenuItem,saveAsMenuItem)

	menu := fyne.NewMainMenu(fileMenu)
	windows.SetMainMenu(menu)
}

//为文件增加过滤
var filter = storage.NewExtensionFileFilter([]string{".hxw",".HXW"})

//打开文件

func (app *config) openFunc(windows fyne.Window)  func(){
	return func() {
		openDialog := dialog.NewFileOpen(func(closer fyne.URIReadCloser, err error) {
			if err!=nil{
				dialog.ShowError(err,windows)
				return
			}
			if closer == nil {
				return
			}
			defer  closer.Close()

			data ,err := ioutil.ReadAll(closer)
			if err!=nil{
				dialog.ShowError(err,windows)
				return
			}
			app.EditWidget.SetText(string(data))
			app.CurrentFile = closer.URI()
			windows.SetTitle(windows.Title() + " - " + closer.URI().Name())
			app.SaveMenuItem.Disabled = false
		},windows)
		//只能打开filter后缀的文件
		openDialog.SetFilter(filter)
		openDialog.Show()
	}
}

func (app *config ) saveFunc(windows fyne.Window) func() {
	return func() {
		if app.CurrentFile!= nil{
			write,err := storage.Writer(app.CurrentFile)
			if err != nil {
				dialog.ShowError(err,windows)
				return
			}

			_,_=write.Write([]byte(app.EditWidget.Text))
			defer write.Close()
		}
	}
}

//文件保存
func (app *config) saveAsFunc(windows fyne.Window) func() {
	return func(){
		saveDialog := dialog.NewFileSave(func(closer fyne.URIWriteCloser, err error) {
			if err!=nil{
				dialog.ShowError(err,windows)
				return
			}
			if closer == nil{
				//用户取消
				return
			}
			//判断文件后缀名是否是filter结尾的
			if !strings.HasSuffix(strings.ToLower(closer.URI().String()),".hxw"){
				dialog.ShowInformation("错误！","请输入正确的文件名称！",windows)
				return
			}

			//save file
			_,_=closer.Write([]byte(app.EditWidget.Text))
			app.CurrentFile=closer.URI()
			defer closer.Close()

			windows.SetTitle(windows.Title() + " - " + closer.URI().Name())
			app.SaveMenuItem.Disabled =false
		}, windows)

		saveDialog.SetFileName("Untitled.hxw")
		saveDialog.SetFilter(filter)
		saveDialog.Show()
	}
}