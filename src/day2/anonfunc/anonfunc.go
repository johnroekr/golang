package main

import "fmt"

func main() {
	func(temp string) {
		fmt.Println("func", temp)
	}("arg")

	func() {
		fmt.Println("func2")
	}()

	var myFunc = func(temp string) string {
		fmt.Println("func3", temp)
		return "func3" + temp
	}("arg3")

	var myFunc2 = func(temp string) string {
		fmt.Println("func4", temp)
		return "func4" + temp
	}

	// _ = myFunc
	fmt.Println(myFunc)
	fmt.Println(myFunc2("test"))

	/* closure */
	fmt.Println("=============")

	param := "myParam"
	param2 := "myParam2"
	var myFunc3 = func(temp string) string {
		fmt.Println("closure func", temp)
		temp = temp + "-modified"
		param2 = temp
		return temp
	}

	fmt.Println(param)
	fmt.Println("result : ", myFunc3(param))
	fmt.Println(param)
	fmt.Println(param2)

	fmt.Println("=============")
	myClosureFunc := test()
	funResult := myClosureFunc("param")
	fmt.Println(funResult)

	funResult2 := myClosureFunc("param")
	fmt.Println(funResult2)

}

func test() func(temp string) string {
	result := "closerTest"
	return func(temp string) string {
		return temp + result
	}
}
