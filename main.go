package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/twistedasylummc/minime/minime"
	"image"
	"image/png"
	"log"
	"strings"
	"syscall/js"
)

func generateMiniMe(base64String string, scale int, slim bool) string {
	base64String = strings.TrimPrefix(base64String, "data:image/png;base64,")

	decoded, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		log.Fatalf("base64 decode error: %v", err)
	}

	img, err := png.Decode(bytes.NewReader(decoded))
	if err != nil {
		log.Fatalf("png decode error: %v", err)
	}

	bounds := img.Bounds()
	if !(bounds.Dx() == 64 && bounds.Dy() == 64) && !(bounds.Dx() == 128 && bounds.Dy() == 128) {
		fmt.Printf("Input file must be 64x64 or 128x128 pixels, got %dx%d\n", bounds.Dx(), bounds.Dy())
		return ""
	}
	res := bounds.Dx() / 64

	var dst image.Image
	if res == 1 {
		dst = minime.Skin64(img, scale)
	} else {
		dst = minime.Skin128(img, scale, slim)
	}
	var buf bytes.Buffer
	err = png.Encode(&buf, dst)
	if err != nil {
		log.Fatalf("png encode error: %v", err)
	}

	encoded := base64.StdEncoding.EncodeToString(buf.Bytes())

	encodedWithPrefix := "data:image/png;base64," + encoded

	return encodedWithPrefix
}

func main() {
	fmt.Println("Go Web Assembly")
	js.Global().Set("generateMiniMe", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		base64String := args[0].String()
		scale := args[1].Int()
		slim := args[2].Bool()

		return generateMiniMe(base64String, scale, slim)
	}))
	select {}
}
