package main

import (
	"fmt"
	"github.com/bytedance/gopkg/lang/stringx"
	"unicode/utf8"
)

func main() {
	fmt.Printf("Sub-[0:100]=%s\n", stringx.Sub("", 0, 100))
	fmt.Printf("Sub-facgbheidjk[3:9]=%s\n", stringx.Sub("facgbheidjk", 3, 9))
	fmt.Printf("Sub-facgbheidjk[-50:100]=%s\n", stringx.Sub("facgbheidjk", -50, 100))
	fmt.Printf("Sub-facgbheidjk[-3:length]=%s\n", stringx.Sub("facgbheidjk", -3, utf8.RuneCountInString("facgbheidjk")))
	fmt.Printf("Sub-facgbheidjk[-3:-1]=%s\n", stringx.Sub("facgbheidjk", -3, -1))
	fmt.Printf("Sub-zh英文hun排[2:5]=%s\n", stringx.Sub("zh英文hun排", 2, 5))
	fmt.Printf("Sub-zh英文hun排[2:-1]=%s\n", stringx.Sub("zh英文hun排", 2, -1))
	fmt.Printf("Sub-zh英文hun排[-100:-1]=%s\n", stringx.Sub("zh英文hun排", -100, -1))
	fmt.Printf("Sub-zh英文hun排[-100:-90]=%s\n", stringx.Sub("zh英文hun排", -100, -90))
	fmt.Printf("Sub-zh英文hun排[-10:-90]=%s\n", stringx.Sub("zh英文hun排", -10, -90))

}
