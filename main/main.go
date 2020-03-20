package main

import (
	"fmt"
	"go_code/chapter12/familyaccount/utils"
)

func main() {
	fmt.Println("这个是面向对象完成")
	utils.NewFamilyAccount().MainMenu()
}