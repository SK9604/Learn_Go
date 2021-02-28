package main

import (
	"errors"
	"fmt"
	"net/http"
)

/* 기본적인 function의 형식
func multiply(a, b int) int { //(a int, b int)와 같다
	return a * b
}
*/

/* go function의 특징 1
func lenAndUpper(name string) (int, string) { //go는 한번에 2개 이상의 타입을 return 할 수 있다.
	return len(name), strings.ToUpper(name) //go의 package는 아주 방대하다
}
*/

/* go function의 특징 2
func lenAndUpper(name string) (length int, uppercase string) { //naked return (미리 리턴할 값을 function을 정의 할 때 설정해준다.)
	defer fmt.Println("I'm done") //func을 실행시킨 후 끝나면 defer를 수행한다.
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}
*/

/* go argument 특징
func repeatMe(words ...string) { //같은 타입의 인자를 여러 개 받을 때 .을 이용하여 한 변수로 모두 받을 수 있다.
	fmt.Println(words)
}
*/

/* go의 for문
func superAdd(numbers ...int) int {
	for index, number := range numbers { //range : loop를 만드는 방법
		fmt.Println(index, number)
	}
	//index를 무시하고 받은 값들의 합을 보내기
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total

	//일반적인 for문으로 loop 작성하기
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	return 1
}
*/
/* switch
func canIDrink(age int) bool {
		if age < 18 {
			return false
		}
		return true
		if koreanAge := age + 2; koreanAge < 20 {
			return false
		}
		return true
		switch {
		case age < 18:
			return false
		case age == 18:
			return true
		}
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}
*/
/*struct
type person struct {
	name    string
	age     int
	favFood []string
}
*/
var errRequestFailed = errors.New("Request failed")

func hitURL(url string) error {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(err, resp.StatusCode)
		return errRequestFailed
	}
	return nil
}

func main() {
	var results = make(map[string]string)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}
	for _, url := range urls {
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}
	for url, result := range results {
		fmt.Println(url, result)
	}

	/*
		dictionary := mydict.Dictionary{}
		baseword := "hello"
		dictionary.Add(baseword, "First")
		dictionary.Search(baseword)
		dictionary.Delete(baseword)
		word, _ := dictionary.Search(baseword)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(word)
		}
	*/
	/*
		definition := "Greeting"
		err := dictionary.Add(word, definition)
		if err != nil {
			fmt.Println(err)
		}
		hello, _ := dictionary.Search(word)
		fmt.Println("found", word, "definition:", hello)
		err2 := dictionary.Add(word, definition)
		if err2 != nil {
			fmt.Println(err2)
		}
	*/
	/*
		definition, err := dictionary.Search("first")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(definition)
		}
	*/
	//account := banking.Account{Owner: "sk", Balance: 10000}
	//fmt.Println(account)
	/*
		account := accounts.NewAccount("sk")
		account.Deposit(10)

			fmt.Println(account.Balance())
			err := account.Withdraw(20)
			if err != nil {
				fmt.Println(err)
			}

		fmt.Println(account)
	*/
	/*struct
	favFood := []string{"ricenoodle", "strawberry"}
	sk := person{"sk", 26, favFood}
	fmt.Println(sk.name, sk.age, sk.favFood)
	*/
	/*
		sk := map[string]string{"name": "sk", "age": "12"} //key:value 형태
		fmt.Println(sk)
		for key, value := range sk {
			fmt.Println(key, value)
		}
	*/
	/*
		names1 := [5]string{"sk", "kim", "kyu"} //Array
		names2 := []string{"sk", "kim", "kyu"}  // Slice = 크기 제한 없는 Array
		names1[3] = "heon"
		names1[4] = "jin"
		//names1[5] = "hee" //array값을 벗어남
		//names2[3] = "heon" //Slice에 새로운 값을 추가할 경우에는 append함수 사용
		names2 = append(names2, "jin")
		fmt.Println(names1)
		fmt.Println(names2)
	*/
	/*a := 2
	b := a
	c := &a //& = 주소
	a = 10
	fmt.Println(&a, a, &b, b, c, *c)
	*c = 20 //* = 주소에 있는 값
	fmt.Println(*c, a)
	*/
	/*
		fmt.Println(canIDrink(18))
		superAdd(1, 2, 3, 4, 5, 6)
		name := "sk" // var name string = "sk"와 같다, func안에서만 작동한다.
		name = "sk kim"
		totalLength, upperName := lenAndUpper("sk")
		length, _ := lenAndUpper("nico") // value값을 무시한다
		fmt.Println(name)
		fmt.Println(multiply(2, 2))
		fmt.Println(totalLength, upperName)
		fmt.Println(length)
		repeatMe("seon", "kyeong", "kim")
	*/
}
