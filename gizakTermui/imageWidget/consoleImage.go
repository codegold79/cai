package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func main() {
	var imgs []image.Image
	for _, arg := range os.Args[1:] {
		file, err := os.Open(arg)
		if err != nil {
			log.Fatalf("failed to fetch image file: %v", err)
		}
		defer file.Close()

		img, _, err := image.Decode(file)
		if err != nil {
			log.Fatalf("failed to decode fetched image: %v", err)
		}
		imgs = append(imgs, img)
	}
	if len(imgs) == 0 {
		ottoFile, err := os.Open("veba_otto.png")
		if err != nil {
			log.Printf("Unable to retrieve veba_otto.png image file: %v", err)
		}
		defer ottoFile.Close()

		ottoImg, _, err := image.Decode(ottoFile)
		if err != nil {
			log.Printf("Unable to load veba_otto.png image: %v", err)
		}

		imgs = append(imgs, ottoImg)
	}

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	uiImg := widgets.NewImage(nil)
	x := 100
	y := 50
	uiImg.SetRect(0, 0, x, y)
	index := 0

	render := func() {
		uiImg.Image = imgs[index]
		uiImg.SetRect(0, 0, x, y)
		if !uiImg.Monochrome {
			uiImg.Title = fmt.Sprintf("Color %d/%d", index+1, len(imgs))
		} else if !uiImg.MonochromeInvert {
			uiImg.Title = fmt.Sprintf("Monochrome(%d) %d/%d", uiImg.MonochromeThreshold, index+1, len(imgs))
		} else {
			uiImg.Title = fmt.Sprintf("InverseMonochrome(%d) %d/%d", uiImg.MonochromeThreshold, index+1, len(imgs))
		}
		ui.Render(uiImg)
	}
	render()

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "<Left>", "h":
			index = (index + len(imgs) - 1) % len(imgs)
		case "<Right>", "l":
			index = (index + 1) % len(imgs)
		case "<Up>", "k":
			uiImg.MonochromeThreshold++
		case "<Down>", "j":
			uiImg.MonochromeThreshold--
		case "<Enter>":
			uiImg.Monochrome = !uiImg.Monochrome
		case "<Tab>":
			uiImg.MonochromeInvert = !uiImg.MonochromeInvert
		case "x":
			x += 10
		case "z":
			x -= 10
		case "y":
			y += 10
		case "t":
			y -= 10
		}
		render()
	}
}
