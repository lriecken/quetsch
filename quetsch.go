package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Quetsch - Image size optimizer")
	fmt.Println("Copyright (C) 2021 - Lennart Riecken <l.riecken@posteo.de>")
	fmt.Println(`
    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU Affero General Public License as
    published by the Free Software Foundation, either version 3 of the
    License, or (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU Affero General Public License for more details.

    You should have received a copy of the GNU Affero General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.`)
	input := flag.String("i", "", "Input file")
	output := flag.String("o", "", "Output file (e.g. target.jpg or target.webp or target.avif")
	size := flag.String("s", "", "Minimal size (e.g. 1920x1080)")
	quality := flag.Int("q", 75, "Quality 1-100")
	targetsize := flag.Int("t", 128*1024, "Targetsize in bytes")

	flag.Parse()
	jpeg := false
	webp := false
	avif := false

	minDims := strings.Split(*size,"x")
	if len(minDims) != 2 {
		fmt.Println("Error parsing size. Should be in format 1920x1080")
		os.Exit(1)
	}
	minWidth, err := strconv.ParseInt(minDims[0], 10, 32)
	if err != nil {
		fmt.Println("Error parsing size. Should be in format 1920x1080")
		os.Exit(1)
	}
	minHeight, err := strconv.ParseInt(minDims[1], 10, 32)
	if err != nil {
		fmt.Println("Error parsing size. Should be in format 1920x1080")
		os.Exit(1)
	}

	if *input == "" {
		fmt.Println("Missing input file")
		os.Exit(1)
	}
	if *output == "" {
		fmt.Println("Missing output file")
		os.Exit(1)
	}
	if strings.Contains(*output, "jpg") {
		jpeg = true
	}
	if strings.Contains(*output, "webp") {
		webp = true
	}
	if strings.Contains(*output, "avif") {
		avif = true
	}
	if *size == "" {
		fmt.Println("Missing minimal size")
		os.Exit(1)
	}
	if jpeg {
		optimizeSize(*input, *output, int(minWidth), int(minHeight), *quality, *targetsize, encodeJpeg)
	} else if webp {
		optimizeSize(*input, *output, int(minWidth), int(minHeight), *quality, *targetsize, encodeWebP)
	} else if avif {
		optimizeSize(*input, *output, int(minWidth), int(minHeight), *quality, *targetsize, encodeAvif)
	} else {
		fmt.Println("Cannot identify output format")
		os.Exit(1)
	}
}





