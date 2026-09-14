package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: ascii-art <image-path>")
	}

	generateASCII(os.Args[1])
}

func generateASCII(imgPath string) {
	const ASCII_IMG_SIZE = 256

	img, err := getImg(imgPath)
	if err != nil {
		log.Fatal(err)
	}

	resizedImg := getResizedImg(img, ASCII_IMG_SIZE)

	writeASCII(resizedImg)
}

func getImg(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open image: %w", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("Failed to decode image: %w", err)
	}

	return img, nil
}

func getResizedImg(src image.Image, width int) image.Image {
	const HEIGHT_SCALING = 2
	bounds := src.Bounds()

	srcW := bounds.Dx()
	srcH := bounds.Dy()

	height := srcH * width / srcW / HEIGHT_SCALING

	dst := image.NewRGBA(image.Rect(0, 0, width, height))

	for y := range height {
		for x := range width {
			srcX := x * srcW / width
			srcY := y * srcH / height

			dst.Set(x, y, src.At(srcX+bounds.Min.X, srcY+bounds.Min.Y))
		}
	}

	return dst
}

func writeASCII(img image.Image) {
	chars := []rune(" .:-=+*#%@")
	bounds := img.Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()

			r8 := uint8(r >> 8)
			g8 := uint8(g >> 8)
			b8 := uint8(b >> 8)

			gray := (r8 + g8 + b8) / 3
			idx := int(gray) * (len(chars) - 1) / 255

			fmt.Printf("%c", chars[idx])
		}

		fmt.Printf("\n")
	}
}
