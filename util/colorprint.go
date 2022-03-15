package util

import (
	"fmt"
	colorprint "github.com/fatih/color"
)

// 输出有颜色的字体
func ColorPrint(res int, describe string) {
	var color colorprint.Attribute
	switch res {
	case 0:
		color = colorprint.FgGreen
		break
	case 1:
		color = colorprint.FgRed
		break
	case 2:
		color = colorprint.FgYellow
		break
	default:
		break
	}
	colorprint.Set(color)
	fmt.Println(describe)
	colorprint.Unset()
}
