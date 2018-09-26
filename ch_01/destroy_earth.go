package main

import (
	"fmt"

	"./fleet"
)

func main() {
	message, err := fleet.DestroyAstronomicalObject("地球", "二向箔")
	if err != nil {
		fmt.Printf("error: 摧毁地球失败: %+v\n", err)
		return
	}

	fmt.Println(message)
}
