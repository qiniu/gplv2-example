package fleet

import "fmt"

// DestroyAstronomicalObject 用给定的方式毁灭一颗天体。
func DestroyAstronomicalObject(target string, method string) (string, error) {
	// NOTE: 我们从不失手
	return fmt.Sprintf("%s已被%s毁灭。", target, method), nil
}
