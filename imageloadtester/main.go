package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

func loadImage(photoUrl string, card *widget.Card) {
	// fmt.Println("Sleep start")
	// time.Sleep(time.Second * time.Duration(rand.Int64N(10))) // Simulate a really long download
	// fmt.Println("Waking up")
	uri, err := storage.ParseURI(photoUrl)
	if err != nil {
		fyne.LogError("parsing url", err)
		return
	}
	image := canvas.NewImageFromURI(uri)
	image.FillMode = canvas.ImageFillContain
	card.SetContent(image)
}

func main() {
	app := app.New()
	win := app.NewWindow("Image Loading Tester")
	win.Resize(fyne.NewSize(800, 800))

	var cards []fyne.CanvasObject
	for i := 0; i < len(images); i++ {
		card := widget.NewCard("Image Loading Tester", "This is a test to see if the image loading is working.", nil)
		card.SetContent(canvas.NewRectangle(color.Black))
		go loadImage(images[i], card)
		cards = append(cards, card)
	}
	win.SetContent(container.NewScroll(container.NewGridWrap(fyne.NewSize(200, 200), cards...)))

	fmt.Println("************************************************************************************Starting run loop")
	win.ShowAndRun()
}

// Create an array of strings
var images = []string{
	"https://live.staticflickr.com/65535/54200664117_029134f1a1_z.jpg",
	"https://live.staticflickr.com/65535/54201842210_b2950ee76b_z.jpg",
	"https://live.staticflickr.com/65535/54200481547_7004003877_z.jpg",
	"https://live.staticflickr.com/65535/54200563977_e3a0e95eb2_z.jpg",
	"https://live.staticflickr.com/65535/54200283076_29750eee93_z.jpg",
	"https://live.staticflickr.com/65535/54201407369_e87aae88b1_z.jpg",
	"https://live.staticflickr.com/65535/54201629548_7499bb66d8_z.jpg",
	"https://live.staticflickr.com/65535/54200521348_6a5522e2ff_z.jpg",
	"https://live.staticflickr.com/65535/54201688919_19f3a260af_z.jpg",
	"https://live.staticflickr.com/65535/54199390104_cb6be3eb2a_z.jpg",
	"https://live.staticflickr.com/65535/54201319701_1ea74fb83f_z.jpg",
	"https://live.staticflickr.com/65535/54201449438_c894fbd86d_z.jpg",
	"https://live.staticflickr.com/65535/54200503702_7d54d275f5_z.jpg",
	"https://live.staticflickr.com/65535/54199621171_4f039f31e5_z.jpg",
	"https://live.staticflickr.com/65535/54201228505_9296452413_z.jpg",
	"https://live.staticflickr.com/65535/54201002653_1be9919050_z.jpg",
	"https://live.staticflickr.com/65535/54200508730_e8c966b813_z.jpg",
	"https://live.staticflickr.com/65535/54201637349_c6f431230d_z.jpg",
	"https://live.staticflickr.com/65535/54200159197_92b0d1d510_z.jpg",
	"https://live.staticflickr.com/65535/54201352573_e4087605af_z.jpg",
	"https://live.staticflickr.com/65535/54199999333_16d7e4ce3c_z.jpg",
	"https://live.staticflickr.com/65535/54201575919_82bede49a0_z.jpg",
	"https://live.staticflickr.com/65535/54200099558_37d619d030_z.jpg",
	"https://live.staticflickr.com/65535/54200200549_9509d7e977_z.jpg",
	"https://live.staticflickr.com/65535/54198939877_04cd6b4093_z.jpg",
	"https://live.staticflickr.com/65535/54190842963_9790b599e0_z.jpg",
	"https://live.staticflickr.com/65535/54198785806_9a339a94e5_z.jpg",
	"https://live.staticflickr.com/65535/54200771525_89f8d831c9_z.jpg",
	"https://live.staticflickr.com/65535/54201426565_afa8962f73_z.jpg",
	"https://live.staticflickr.com/65535/54199898932_5ae67a6f8e_z.jpg",
	"https://live.staticflickr.com/65535/54200729571_62abc7e56d_z.jpg",
	"https://live.staticflickr.com/65535/54199379612_a4981d9735_z.jpg",
	"https://live.staticflickr.com/65535/54199479790_8ddbacc0bf_z.jpg",
	"https://live.staticflickr.com/65535/54182836587_c1bc6b36a7_z.jpg",
	"https://live.staticflickr.com/65535/54201281249_ae3a7abbf7_z.jpg",
	"https://live.staticflickr.com/65535/54201322133_c2b41f11d9_z.jpg",
	"https://live.staticflickr.com/65535/54199888079_70f8a6fb50_z.jpg",
	"https://live.staticflickr.com/65535/54184160455_0f897d8365_z.jpg",
	"https://live.staticflickr.com/65535/54200801821_7d0f6e2286_z.jpg",
	"https://live.staticflickr.com/65535/54199500968_a24c024857_z.jpg",
	"https://live.staticflickr.com/65535/54200533229_fba780da9f_z.jpg",
	"https://live.staticflickr.com/65535/54199379582_eb83315d7b_z.jpg",
	"https://live.staticflickr.com/65535/54200129726_e040c7b177_z.jpg",
	"https://live.staticflickr.com/65535/54199068937_96eacfe56b_z.jpg",
}
