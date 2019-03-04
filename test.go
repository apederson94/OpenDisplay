package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

var (
	displayArea *widgets.QWidget
	scene       *widgets.QGraphicsScene
	//view          *widgets.QColumnView
	item          *widgets.QGraphicsPixmapItem
	mainApp       *widgets.QApplication
	imageFileName string
)

//TODO: create a backend and create a way to display data in columns for daily forecast

func weatherView() *widgets.QWidget {
	displayArea = widgets.NewQWidget(nil, 0)
	var (
		forecastLabel = widgets.NewQLabel2("Current Forecast", nil, core.Qt__Window)
		layout        = widgets.NewQVBoxLayout()
	)

	forecast.Insert

	layout.AddWidget(forecastLabel, 0, core.Qt__AlignCenter)
	layout.AddWidget(forecast, 0, core.Qt__AlignCenter)

	displayArea.SetLayout(layout)

	return displayArea
}

func main() {

	mainApp = widgets.NewQApplication(len(os.Args), os.Args)

	weatherView().Show()

	widgets.QApplication_Exec()
}
