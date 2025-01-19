package main

import "fmt"

const LoginTocken2 string = "rushimal2dd"

func main() {
	var username string = "rushimani"
	fmt.Println(username)
	fmt.Printf("the variable of type %T \n", username)

	var isLoggedin bool = true
	fmt.Println(isLoggedin)
	fmt.Printf("the varible of type %T \n", isLoggedin)

	var floatNumber float32 = 20.4656645
	fmt.Println(floatNumber)
	fmt.Printf("the varible of type %T \n", floatNumber)

	var floatNumber2 float64 = 9.3842555845884225
	fmt.Println(floatNumber2)
	fmt.Printf("the varible of type %T \n", floatNumber2)

	var anotherValue int
	fmt.Println(anotherValue)
	fmt.Printf("the varible of type of %T \n", anotherValue)

	var anotherValue2 string
	fmt.Println(anotherValue2)
	fmt.Printf("the varible of type of %T \n", anotherValue2)

	numberOfUsers := 20.256
	fmt.Println(numberOfUsers)
	fmt.Printf("the variable of type %T \n", numberOfUsers)

	const LoginTocken string = "rushimal"
	fmt.Println(LoginTocken)

	fmt.Println(LoginTocken2)
}
