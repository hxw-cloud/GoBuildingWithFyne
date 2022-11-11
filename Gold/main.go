package main

import (
	"Gold/repository"
	"database/sql"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	_ "github.com/glebarez/go-sqlite"
	"log"
	"net/http"
	"os"
)

type Config struct {
	App                            fyne.App
	InfoLog                        *log.Logger
	ErrorLog                       *log.Logger
	DB                             repository.Repository
	MainWindow                     fyne.Window
	PriceContainer                 *fyne.Container
	ToolBar                        *widget.Toolbar
	PriceChartContainer            *fyne.Container
	Holdings                       [][]interface{}
	HoldingsTable                  *widget.Table
	HTTPClient                     *http.Client
	AddHoldingsPurchaseAmountEntry *widget.Entry
	AddHoldingsPurchaseDataEntry   *widget.Entry
	AddHoldingsPurchasePriceEntry  *widget.Entry
}

var myApp Config

func main() {
	//创建一个fyne应用
	fyneApp := app.NewWithID("ca.gocode.goldwatcher.preferences")
	myApp.App = fyneApp
	myApp.HTTPClient = &http.Client{}
	//创建日志
	myApp.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	myApp.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime)
	//连接数据库
	sqlDb, err := myApp.connectSql()
	if err != nil {
		log.Panicln(err)
	}
	//创建一个数据库
	myApp.setupDB(sqlDb)
	//创建一个fyne窗口
	myApp.MainWindow = fyneApp.NewWindow("GoldWatcher")
	myApp.MainWindow.Resize(fyne.NewSize(770, 410))
	myApp.MainWindow.SetFixedSize(true)
	myApp.MainWindow.SetMaster()
	myApp.makeUI()
	//显示和运行程序
	myApp.MainWindow.ShowAndRun()
}

func (app *Config) connectSql() (*sql.DB, error) {
	path := ""

	if os.Getenv("DB_PATH") != "" {
		path = os.Getenv("DB_PATH")
	} else {
		path = app.App.Storage().RootURI().Path() + "/sql.db"
		app.InfoLog.Println("数据库位置在: ", path)
	}

	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (app *Config) setupDB(sqlDB *sql.DB) {
	app.DB = repository.NewSqLiteRepository(sqlDB)
	err := app.DB.Migrate()
	if err != nil {
		app.ErrorLog.Println(err)
	}

}
