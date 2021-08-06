package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"os"
	"strconv"
)

func openImage(input string) image.Image{
	if inputImage, err := os.Open(input); err != nil {
		fmt.Println("Error opening image")
		os.Exit(1)
	} else {
		if imageData, imageType, err := image.Decode(inputImage); err != nil {
			fmt.Println("Error decoding image. Probably unsupported format.")
			os.Exit(1)
		} else {
			fmt.Println("Input file is: " + imageType)
			return imageData
		}
	}
	return nil
}

func optimizeSize(input string, output string, minWidth int, minHeight int, quality int, targetsize int, encodeImage encoderfunc) {
	image := openImage(input)
	bounds := image.Bounds()
	width := bounds.Max.X
	height := bounds.Max.Y

	minimalScaledWidth, minimalScaledHeight, minScale := getMinimalScale(width, height, minWidth, minHeight)
	minscaleString := fmt.Sprintf("%f", minScale)
	fmt.Println("Origial dimensions are: " + strconv.Itoa(width) + "x" + strconv.Itoa(height))
	fmt.Println("Minimal dimensions are: " + strconv.Itoa(minimalScaledWidth) + "x" + strconv.Itoa(minimalScaledHeight) + " (scale " + minscaleString+ ")")
	fmt.Println("Maximal filesize is: " + strconv.Itoa(targetsize) + " bytes.")

	scaledWidth, scaledHeight, finalQuality, scaledData := determineBestSizeAndQuality(image, width, height, minScale, quality, targetsize, encodeImage)


	fmt.Println("Final settings are: " + strconv.Itoa(scaledWidth) + "x" + strconv.Itoa(scaledHeight) + " at quality: " + strconv.Itoa(finalQuality))
	if len(scaledData) > targetsize {
		fmt.Println("Something went seariously wrong")
	}
	err := ioutil.WriteFile(output, scaledData, 0644)
	if err != nil {
		fmt.Println("Error writing image")
		os.Exit(1)
	}
}

func determineBestSizeAndQuality(image image.Image, width int, height int, minScale float64, quality int, targetSize int, encodeImage encoderfunc) (scaledWidth int, scaledHeight int, optimalQuality int, data []byte) {
	optimalQuality = quality
	minScaledWidth, minScaledHeight := scaleDims(width, height, minScale)
	scaledWidth, scaledHeight = scaleDims(width, height, minScale)
	scaleData := encodeImage(image, optimalQuality, scaledWidth, scaledHeight)
	if len(scaleData) < targetSize {
		// The minimal size of the picture is below target size
		scaledWidth, scaledHeight := scaleDims(width, height, 1.0)
		scaleData = encodeImage(image, optimalQuality, scaledWidth, scaledHeight)
		if len(scaleData) > targetSize {
			// The maximum sized image is too large, we optimize the scale.
			var scale float64
			scale, scaleData = bisectScale(image, width, height, minScale, 1, optimalQuality, targetSize, encodeImage)
			scaledWidth, scaledHeight = scaleDims(width, height, scale)
			return scaledWidth, scaledHeight, optimalQuality, scaleData
		} else {
			// The maximum image dimenstions and quality are small enough
			return scaledWidth, scaledHeight, optimalQuality, scaleData
		}
	} else if len(scaleData) > targetSize {
		// The minimal sized image is still too large. We optimize quality.
		optimalQuality, scaleData = bisectQuality(image, minScaledWidth, minScaledHeight, 0, optimalQuality, targetSize, encodeImage)
		return minScaledWidth, minScaledHeight, optimalQuality, scaleData
	}
	// The minimal sized image matches the target size exactly
	return scaledWidth, scaledHeight, optimalQuality, scaleData
}

func bisectScale(image image.Image, width int, height int, minScale float64, maxScale float64, quality int, targetSize int, encodeImage encoderfunc) (float64, []byte) {
	curScale := (minScale + maxScale) / float64(2)
	swidth, sheight := scaleDims(width, height, curScale)
	data := encodeImage(image, quality, swidth, sheight)
	if len(data) > targetSize {
		if (maxScale - minScale) < 0.001 {
			return curScale, data
		} else {
			return bisectScale(image, width, height, minScale, curScale, quality, targetSize, encodeImage)
		}

	} else if len(data) == targetSize{
		return curScale, data
	} else {
		if (maxScale - minScale) < 0.001 {
			return curScale, data
		} else {
			return bisectScale(image, width, height, curScale, maxScale, quality, targetSize, encodeImage)
		}
	}
}

func bisectQuality(image image.Image, width int, height int, minQ int, maxQ int, targetSize int, encodeImage encoderfunc) (int, []byte) {
	curQ := int(float64(minQ + maxQ) / float64(2))
	data := encodeImage(image, curQ, width, height)
	if len(data) > targetSize {
		if curQ == int(float64(minQ + curQ) / float64(2)) {
			return curQ, data
		} else {
			return bisectQuality(image, width, height, minQ, curQ, targetSize, encodeImage)
		}

	} else if len(data) == targetSize{
		return curQ, data
	} else {
		if  curQ == int(float64(curQ + maxQ) / float64(2)) {
			return curQ, data
		} else {
			return bisectQuality(image, width, height, curQ, maxQ, targetSize, encodeImage)
		}

	}
}
