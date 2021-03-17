// WARNING: This application will write a JPEG file to your hard drive if
// you click on the "Convert to grayscale" button.

package main

import (
	"fmt"
	"image"
	"path/filepath"

	"image/color"
	"image/jpeg"
	"log"
	"math"
	"net"
	"net/http"
	"os"

	"github.com/zserge/lorca"
)

func main() {
	ui, err := lorca.New("", "", 800, 600)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	ui.Bind("changePhotoGo", func(name, caption string) {
		width, height := getImageDimension(name)
		evalStr := fmt.
			Sprintf("changePhotoJS(\"%s\",\"%s\",\"%d\",\"%d\");", name, caption, width, height)
		ui.Eval(evalStr)
	})

	ui.Bind("grayScaleGo", func(path string) {
		filename := filepath.Base(path)
		grayScale(filename)
		filename = "GS" + "_" + filename
		evalStr := fmt.Sprintf("loadGrayScale(\"%s\");", filename)
		ui.Eval(evalStr)
	})

	curDir, err := os.Getwd()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	go http.Serve(ln, http.FileServer(http.Dir(curDir)))

	ui.Load(fmt.Sprintf("http://%s", ln.Addr()))

	<-ui.Done()
}

func getImageDimension(imagePath string) (int, int) {
	file, err := os.Open(imagePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	image, _, err := image.DecodeConfig(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", imagePath, err)
	}
	return image.Width, image.Height
}

func grayScale(filename string) {
	infile, err := os.Open(filename)

	if err != nil {
		log.Printf("failed opening %s: %s", filename, err)
		panic(err.Error())
	}
	defer infile.Close()

	imgSrc, _, err := image.Decode(infile)
	if err != nil {
		panic(err.Error())
	}

	// Create a new grayscale image
	bounds := imgSrc.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	grayScale := image.NewGray(image.Rectangle{image.Point{0, 0}, image.Point{w, h}})
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			imageColor := imgSrc.At(x, y)
			rr, gg, bb, _ := imageColor.RGBA()
			r := math.Pow(float64(rr), 2.2)
			g := math.Pow(float64(gg), 2.2)
			b := math.Pow(float64(bb), 2.2)
			m := math.Pow(0.2125*r+0.7154*g+0.0721*b, 1/2.2)
			Y := uint16(m + 0.5)
			grayColor := color.Gray{uint8(Y >> 8)}
			grayScale.Set(x, y, grayColor)
		}
	}

	// Encode the grayscale image to the new file
	newFilename := "GS" + "_" + filename
	newfile, err := os.Create(newFilename)
	if err != nil {
		log.Printf("failed creating %s: %s", newFilename, err)
		panic(err.Error())
	}
	defer newfile.Close()
	jpeg.Encode(newfile, grayScale, nil)
}
