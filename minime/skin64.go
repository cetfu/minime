package minime

import (
	"image"
	"image/color"
)

// Skin64 returns an avatar generated from a skin that is 64x64 pixels. The skin must also be in the Minecraft format
// otherwise the generated result may not look correct.
func Skin64(src image.Image, scale int) image.Image {
	dst := image.NewRGBA(image.Rect(0, 0, 10*scale, 16*scale))

	// Function to set scaled pixels
	setScaledPixel := func(x, y int, c color.Color) {
		for dx := 0; dx < scale; dx++ {
			for dy := 0; dy < scale; dy++ {
				dst.Set(x*scale+dx, y*scale+dy, c)
			}
		}
	}

	// Head
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			if x == 0 || x == 9 || y == 0 || y == 9 {
				setScaledPixel(x, y, color.Black)
				continue
			}
			setScaledPixel(x, y, findSuitablePixel(src, x+39, y+7, x+7, y+7))
		}
	}

	// Arms
	for x := 0; x < 8; x++ {
		for y := 0; y < 4; y++ {
			if x == 0 || x == 7 || y == 0 || y == 3 {
				setScaledPixel(x+1, y+9, color.Black)
				continue
			}
			yOff := (y - 1) * 7
			if x == 1 {
				setScaledPixel(x+1, y+9, findSuitablePixel(src, 45, 37+yOff, 45, 21+yOff))
			} else if x == 6 {
				setScaledPixel(x+1, y+9, findSuitablePixel(src, 54, 53+yOff, 38, 53+yOff))
			}
		}
	}

	// Body
	for x := 0; x < 4; x++ {
		for y := 0; y < 7; y++ {
			if x == 0 || x == 3 || y == 0 || y == 4 || y == 6 {
				setScaledPixel(x+3, y+9, color.Black)
				continue
			} else if y == 5 {
				if x == 1 {
					setScaledPixel(x+3, y+9, findSuitablePixel(src, 5, 41, 5, 25))
				} else {
					setScaledPixel(x+3, y+9, findSuitablePixel(src, 6, 57, 22, 57))
				}
				continue
			}
			xOff, yOff := (x-1)*4, (y-1)*4
			setScaledPixel(x+3, y+9, findSuitablePixel(src, 21+xOff, 37+yOff, 21+xOff, 21+yOff))
		}
	}
	return dst
}
