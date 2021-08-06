package main

import (
	"bytes"
	"fmt"
	"github.com/Kagami/go-avif"
	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"os"
	"strconv"
)

type encoderfunc func(image.Image, int, int,  int) []byte

func encodeJpeg(imagedata image.Image, quality int, width int, height int) []byte {
	fmt.Print("Attempting: " + strconv.Itoa(width) + "x" + strconv.Itoa(height) + " at quality " + strconv.Itoa(quality) + "... ")
	resized := resize.Resize(uint(width), uint(height), imagedata, resize.Lanczos2)
	writer := new(bytes.Buffer)
	options := jpeg.Options{quality}
	err := jpeg.Encode(writer, resized, &options)
	if err != nil {
		fmt.Print("Error encoding jpeg image")
		os.Exit(1)
	}
	b := writer.Bytes()
	fmt.Println(strconv.Itoa(len(b)) + " bytes")
	return b
}

func encodeWebP(imagedata image.Image, quality int, width int, height int) []byte {
	fmt.Print("Attempting: " + strconv.Itoa(width) + "x" + strconv.Itoa(height) + " at quality " + strconv.Itoa(quality) + "... ")
	resized := resize.Resize(uint(width), uint(height), imagedata, resize.Lanczos2)
	writer := new(bytes.Buffer)
	options, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, float32(quality))
	if err != nil {
		fmt.Print("Error encoding webp image")
		os.Exit(1)
	}
	err = webp.Encode(writer, resized, options)
	if err != nil {
		panic(err)
	}
	b := writer.Bytes()
	fmt.Println(strconv.Itoa(len(b)) + " bytes")
	return b
}

func encodeAvif(imagedata image.Image, quality int, width int, height int) []byte {
	avifQuality := int(float64(63) * (1 - (float64(quality) / 100)))
	fmt.Print("Attempting: " + strconv.Itoa(width) + "x" + strconv.Itoa(height) + " at quality " + strconv.Itoa(quality) + " avif quality " + strconv.Itoa(avifQuality) + "... ")
	resized := resize.Resize(uint(width), uint(height), imagedata, resize.Lanczos2)

	writer := new(bytes.Buffer)
	options := avif.Options{
		Threads:        0,
		Speed:          0,
		Quality:        avifQuality,
		SubsampleRatio: nil,
	}
	err := avif.Encode(writer, resized, &options)
	if err != nil {
		fmt.Print("Error encoding avif image")
		os.Exit(1)
	}
	b := writer.Bytes()
	fmt.Println(strconv.Itoa(len(b)) + " bytes")
	return b
}