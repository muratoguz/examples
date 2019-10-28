package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	/*
		userName := os.Getenv("USERNAME")
		uDomain := os.Getenv("USERDOMAIN")
		goRoot := os.Getenv("GOROOT")
		goPath := os.Getenv("GOPATH")
		homePath := os.Getenv("HOMEPATH")

		fmt.Println("User Name: " + userName)
		fmt.Println("User Domain: " + uDomain)
		fmt.Println("Go Root: " + goRoot)
		fmt.Println("Go Path: " + goPath)
		fmt.Println("Home Path: " + homePath)

		t := time.Date(2019, time.December, 10, 22, 33, 0, 0, time.UTC)
		n := time.Now()
		fmt.Printf("Time: %s\n", t)
		fmt.Printf("Now : %s\n", n)
		fmt.Printf("Month: %s\n", t.Month())
		fmt.Printf("Day: %s\n", t.Day())
		fmt.Printf("Weekday: %s\n", t.Weekday())
	*/

	for i, v := range pow {
		fmt.Printf("2**%d = %d \n", i, v)
	}

	arr := [...]string{"a", "b", "c", "d"}

	for i := range arr {
		fmt.Println("Array item ", i, " = ", arr[i])
	}

	//map example

	var baskentler = map[string]string{"Turkey": "Ankara", "Italy": "Roma", "Spain": "Madrid"}

	for key, val := range baskentler {
		fmt.Println(key, ":", val)
	}

}
