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
	chars := []rune("@%#*+=-:. ")
	fmt.Println("Helloo", chars)

	file, err := os.Open("random-person.jpeg")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	img, format, err := image.Decode(file)
	if err != nil {
		log.Fatalf("Failed to decode image: %v", err)
	}
	fmt.Printf("Image format is: %s\n", format)

	bounds := img.Bounds()
	fmt.Println(bounds)

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
