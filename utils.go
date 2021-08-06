package main

func scaleDims(origwidth int, origheight int, scale float64) (newwidth int, newheight int ) {
	newwidth = int(float64(origwidth) * scale)
	newheight = int(float64(origheight) * scale)
	return newwidth, newheight
}

func getMinimalScale(width int, height int, minwidth int, minheight int) ( int,int, float64 ){
	origratio := float64(width) / float64(height)
	minratio := float64(minwidth) / float64(minheight)
	if origratio < minratio {
		minheight = int(float64(minwidth) / origratio)
		scale := float64(minwidth) / float64(width)
		return minwidth, minheight, scale
	} else if origratio > minratio {
		minwidth = int(float64(minheight) * origratio)
		scale := float64(minheight) / float64(height)
		return minwidth, minheight, scale
	} else {
		scale := float64(minheight) / float64(height)
		return minwidth, minheight, scale
	}

}