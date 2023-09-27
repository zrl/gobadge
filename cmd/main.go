package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/tinygo-org/gobadge/cmd/logos"
	logos_animated "github.com/tinygo-org/gobadge/cmd/logos-animated"
)

const (
	gopherconEU22Logo = "./cmd/assets/gopherconeu-2022.jpg"
	gopherconUK22Logo = "./cmd/assets/gopherconuk-2022.jpg"
	gopherconUS22Logo = "./cmd/assets/gopherconus-2022.jpg"
	gopherconUS23Logo = "./cmd/assets/gopherconus-2023.jpg"
	fosdem23Logo      = "./cmd/assets/fosdem-2023.jpg"
	kubeconEU23       = "./cmd/assets/kubecon-eu-2023.jpg"
	tinygoLogo        = "./cmd/assets/tinygo.jpg"
	skeletor          = "./cmd/assets/skeletor_100.gif"
)

func main() {
	conf := flag.String("conf", tinygoLogo, "Choose the conference logo you want to (e.g.: tinygo, gceu22, gcuk22, gcus22)")
	flag.Parse()

	c := confs()
	logo, ok := c[*conf]
	if !ok {
		fmt.Println("I do not have yet this conf in my catalog.")
		return
	}

	if strings.HasSuffix(logo, ".gif") {
		logos_animated.GenerateLogoRGBA_AnimatedFile(logo)
		return
	}
	logos.GenerateLogoRGBAFile(logo)
}

func confs() map[string]string {
	return map[string]string{
		"gceu22":      gopherconEU22Logo,
		"gcuk22":      gopherconUK22Logo,
		"gcus22":      gopherconUS22Logo,
		"gcus23":      gopherconUS23Logo,
		"fosdem23":    fosdem23Logo,
		"tinygo":      tinygoLogo,
		"kubeconeu23": kubeconEU23,
		"skeletor":    skeletor,
	}
}
