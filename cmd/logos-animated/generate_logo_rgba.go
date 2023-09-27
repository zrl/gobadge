package logos_animated

import (
	"fmt"
	"image/color"
	"image/gif"
	"log"
	"os"
	"strings"
	"text/template"
)

func GenerateLogoRGBA_AnimatedFile(filepath string) {
	colors := generateLogoRGBAs(filepath)
	colorsStrs := convertToStrings(colors)
	generateFile(colorsStrs)
}

func generateLogoRGBAs(filepath string) [][]color.RGBA {
	file, _ := os.Open(filepath)
	img, err := gif.DecodeAll(file)
	if err != nil {
		log.Fatal("failed to decode image file", err)
	}

	log.Printf("frames: %d\n", len(img.Image))

	colorsArr := make([][]color.RGBA, 0)

	//yBuffer := 128 - img.Config.Height
	//xBuffer := 160 - img.Config.Width

	for i, frame := range img.Image {
		if i%5 != 0 {
			continue
		}
		colors := make([]color.RGBA, 0)
		for y := 0; y < frame.Bounds().Max.Y; y++ {
			for x := 0; x < frame.Bounds().Max.X; x++ {
				r, g, b, _ := frame.At(x, y).RGBA()
				colors = append(colors, color.RGBA{R: uint8(r >> 8), G: uint8(g >> 8), B: uint8(b >> 8), A: uint8(255)})
			}
		}
		colorsArr = append(colorsArr, colors)
	}

	return colorsArr
}

func convertToStrings(colors [][]color.RGBA) []string {
	contents := make([]string, 0, len(colors))
	var err error

	for _, arr := range colors {
		var content strings.Builder
		content.WriteRune('{')
		for i, c := range arr {
			str := fmt.Sprintf("{%d, %d, %d, %d}", c.R, c.G, c.B, c.A)
			_, err = content.WriteString(str)
			if err != nil {
				log.Fatal("failed to write string")
			}

			if i < len(arr)-1 {
				_, err = content.WriteString(", ")
				if err != nil {
					log.Fatal("failed to write string")
				}
			}
		}
		content.WriteRune('}')
		contents = append(contents, content.String())
	}

	return contents
}

func generateFile(colorsStrs []string) {
	tmp, err := template.ParseFiles("./cmd/logos-animated/logo-template.txt")
	if err != nil {
		log.Fatal("failed to parse template", err)
	}

	f, err := os.Create("logo_animated.go")
	if err != nil {
		log.Fatal("failed to create output file", err)
	}

	type Colors struct {
		Str string
	}

	c := Colors{Str: strings.Join(colorsStrs, ", ")}

	err = tmp.Execute(f, c)
	if err != nil {
		log.Fatal("failed to execute template", err)
	}
}
