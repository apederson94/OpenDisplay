package main

import (
	"fmt"
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

func testView() *widgets.QWidget {
	displayArea = widgets.NewQWidget(nil, 0)
	var view = widgets.NewQTableView(nil)
	var tabBar = widgets.NewQTabBar(nil)

	var row = widgets.NewQRow

	var button = widgets.NewQPushButton2("TEST", nil)
	button.ConnectClicked(func(flag bool) {
		widgets.QApplication_Beep()

	})

	var layout = widgets.NewQVBoxLayout()

	layout.AddWidget(view, 0, core.Qt__AlignCenter)
	layout.AddWidget(button, 0, core.Qt__AlignCenter)
	layout.AddWidget(tabBar, 0, core.Qt__AlignCenter)

	displayArea.SetLayout(layout)

	return displayArea
}

func main() {

	fmt.Println("Loading image : ", imageFileName)

	mainApp = widgets.NewQApplication(len(os.Args), os.Args)

	testView().Show()

	widgets.QApplication_Exec()
}
