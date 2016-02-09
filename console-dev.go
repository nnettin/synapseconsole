// console-dev.go
package main

import (
	"fmt"
	//"io/ioutil"
	"log"
	"os"

	"github.com/jroimartin/gocui"
)

func init() {
	// Change the device for logging to stdout
	log.SetOutput(os.Stdout)
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func layout(g *gocui.Gui) error {

	maxX, maxY := g.Size()
	v, err := g.SetView("mainmenu", 2, 1, maxX-2, maxY-2)

	v.FgColor = gocui.ColorGreen

	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		//		b, err := ioutil.ReadFile("mainmenu")
		//if err != nil {
		//	log.Panicln(err)
		//	panic(err)
		//}
		v.Title = "Main Menu"
		v.Frame = true
		//fmt.Fprintf(v, "%s", b)

		v.Wrap = true
		v.SetCursor(10, 10)
		fmt.Println(v.Cursor())
	}
	return nil

}

func keybindings(g *gocui.Gui) error {

	if err := g.SetKeybinding("mainmenu", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	return nil
}

func main() {

	log.Println("This is a test")

	g := gocui.NewGui()
	if err := g.Init(); err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetLayout(layout)

	if err := keybindings(g); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	//g.SelBgColor = gocui.ColorGreen
	//g.SelFgColor = gocui.ColorBlack
	g.Cursor = true

	g.FgColor = gocui.ColorGreen

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
