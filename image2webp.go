package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/chai2010/webp"
)

var Input_src = flag.String("src", "", "-src='source image(png or jpeg) path'")
var Input_mode = flag.String("mode", "lossless_rgba", "-mode=lossless_rgba|lossless_rgb|lossless_gray|rgba90|rgba80|rgba70|rgb90|rgb80|rgb70|gray90|gray80|gray70")
var Input_dest = flag.String("dest", "", "-dest='destination webp path'")
var Input_help = flag.String("help", "", "-src='source image(png or jpeg) path'' -dest='destination webp path'")
var Input_usage = flag.String("usage", "", "./image2webp -src='source image(png or jpeg) path' -dest='destination webp path'")

func main() {
	flag.Parse()
	if _, find := os.Stat(*Input_src); find != nil {
		fmt.Println("File " + *Input_src + " not exists.")
		return
	}
	img2Webp(*Input_src, *Input_dest)
}

func img2Webp(src string, dest string) {
	if _, find := os.Stat(src); find == nil {
		lowSrc := strings.ToLower(src)
		if strings.HasSuffix(lowSrc, ".png") {
			err := png2Webp(src, dest)
			if err != nil {
				log.Println(err)
			}
		}
		if strings.HasSuffix(lowSrc, ".jpeg") || strings.HasSuffix(lowSrc, ".jpg") {
			err := jpeg2Webp(src, dest)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func png2Webp(src string, dest string) error {
	imgByte, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	pngImg, ok := png.Decode(bytes.NewReader(imgByte))
	if ok != nil {
		return ok
	}

	return convertHelper(pngImg, src, dest)
}

func jpeg2Webp(src string, dest string) error {
	imgByte, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	jpegImg, ok := jpeg.Decode(bytes.NewReader(imgByte))
	if ok != nil {
		return ok
	}

	return convertHelper(jpegImg, src, dest)
}

func convertHelper(img image.Image, src string, dest string) error {
	var webpByte []byte
	switch *Input_mode {
	case "lossless_rgba":
		webpByte, _ = webp.EncodeLosslessRGBA(img)
	case "lossless_rgb":
		webpByte, _ = webp.EncodeLosslessRGB(img)
	case "lossless_gray":
		webpByte, _ = webp.EncodeLosslessGray(img)
	case "rgba90":
		webpByte, _ = webp.EncodeRGBA(img, 90.0)
	case "rgba80":
		webpByte, _ = webp.EncodeRGBA(img, 80.0)
	case "rgba70":
		webpByte, _ = webp.EncodeRGBA(img, 70.0)
	case "rgb90":
		webpByte, _ = webp.EncodeRGB(img, 90.0)
	case "rgb80":
		webpByte, _ = webp.EncodeRGB(img, 80.0)
	case "rgb70":
		webpByte, _ = webp.EncodeRGB(img, 70.0)
	case "gray90":
		webpByte, _ = webp.EncodeGray(img, 90.0)
	case "gray80":
		webpByte, _ = webp.EncodeGray(img, 80.0)
	case "gray70":
		webpByte, _ = webp.EncodeGray(img, 70.0)
	default:
		webpByte, _ = webp.EncodeLosslessRGBA(img)
	}
	fileInfo, _ := os.Stat(src)
	ioutil.WriteFile(dest, webpByte, fileInfo.Mode())
	return nil
}
