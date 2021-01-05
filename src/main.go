package main

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int { //(a int, b int)와 같다
	return a * b
}

func lenAndUpper(name string) (int, string) { //go는 한번에 2개 이상의 타입을 return 할 수 있다.
	return len(name), strings.ToUpper(name) //go의 package는 아주 방대하다
}

func repeatMe(words ...string) { //같은 타입의 인자를 여러 개 받을 때 .을 이용하여 한 변수로 모두 받을 수 있다.
	fmt.Println(words)
}

func main() {
	name := "sk" // var name string = "sk"와 같다, func안에서만 작동한다.
	name = "sk kim"
	totalLength, upperName := lenAndUpper("sk")
	length, _ := lenAndUpper("nico") // value값을 무시한다
	fmt.Println(name)
	fmt.Println(multiply(2, 2))
	fmt.Println(totalLength, upperName)
	fmt.Println(length)
	repeatMe("seon", "kyeong", "kim")
}
