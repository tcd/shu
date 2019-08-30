package shu

import "strings"

// Credit to VK
// https://www.asciiart.eu/clothing-and-accessories/footwear
func logo() string {
	var lines = []string{
		"        ________",
		"     __(_____  <|",
		"    (____ / <| <|",
		"    (___ /  <| L`-------.",
		"    (__ /   L`--------.  \\",
		"    /  `.    ^^^^^ |   \\  |",
		"   |     \\---------'    |/",
		"   |______|____________/]",
		"   [_____|`-.__________]",
	}
	return strings.Join(lines, "\n")
}
