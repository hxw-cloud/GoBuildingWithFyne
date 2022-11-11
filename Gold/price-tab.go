package main

import (
	"bytes"
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	image2 "image"
	"image/png"
	"io"
	"os"
	"strings"
)

func (app *Config) pricesTab() *fyne.Container {
	chart := app.getChart()
	chartContainer := container.NewVBox(chart)
	app.PriceChartContainer = chartContainer
	return chartContainer
}

func (app *Config) getChart() *canvas.Image {
	apiURL := fmt.Sprintf("https://goldprice.org/charts/gold_3d_b_o_%s_x.png", strings.ToLower(currency))
	var image *canvas.Image
	err := app.downloadFile(apiURL, "gold.png")
	if err != nil {
		image = canvas.NewImageFromResource(resourceUnreachablePng)
	} else {
		image = canvas.NewImageFromFile("download/gold.png")
	}
	image.SetMinSize(fyne.Size{
		Width:  770,
		Height: 410,
	})

	image.FillMode = canvas.ImageFillOriginal
	return image
}

func (app Config) downloadFile(URL, fileName string) error {
	//从url中获取返回信息
	resp, err := app.HTTPClient.Get(URL)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New("连接失败，未能下载images")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	image, _, err := image2.Decode(bytes.NewReader(body))
	if err != nil {
		return err
	}
	out, err := os.Create(fmt.Sprintf("./download/%s", fileName))
	if err != nil {
		return err
	}

	err = png.Encode(out, image)
	if err != nil {
		return err
	}

	return nil
}
