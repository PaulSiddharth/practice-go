package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"os"
	"sync"
)

var signals = []string{"test"}
var wg sync.WaitGroup //pointer
var mut sync.Mutex

func main() {
	// var name string
	// fmt.Println(name)
	// fmt.Println("Hello from sid")

	// // array
	// var fruitList [4]string
	// fruitList[0] = "Apple"

	// fmt.Println(len(fruitList))

	// // slice
	// var fruits = []string{"apple", "toamto"}

	// fruits = append(fruits, "mango")

	// fmt.Println(fruits)
	// fruits = append(fruits[:1], fruits[2:]...)
	// fmt.Println(fruits)

	// //map
	// languages := make(map[string]string)
	// languages["JS"] = "Javscript"
	// languages["PY"] = "Python"
	// languages["RB"] = "Ruby"

	// delete(languages, "RB")
	// fmt.Println(languages)

	// for key, val := range languages {
	// 	fmt.Printf("Key %v, value %v \n", key, val)
	// }

	// sid := User{"Siddhartha", "sid@go.dev", 30}
	// fmt.Println(sid)
	// fmt.Printf("user details %+v\n", sid)

	// if num := 3; num < 10 {
	// 	fmt.Println("Num less than 10")
	// } else {
	// 	fmt.Println("Num greater than 10")

	// }

	// res := adder(2, 3)
	// fmt.Println(res)

	// res1 := ProAdder(2, 3, 7)
	// fmt.Println(res1)

	// // working with file

	// content := "New content in my file"
	// file, fErr := os.Create("./myText.txt")
	// if fErr != nil {
	// 	panic(fErr)
	// }
	// length, err := io.WriteString(file, content)
	// if err != nil {
	// 	panic(err)
	// }
	// readFile("./myText.txt")
	// fmt.Println("length is,", length)
	// defer file.Close()

	// arr := closestPrimes(19, 31)
	// fmt.Println(arr)
	// EncodeJson()
	// DecodeJson()
	// go greeter("Hello")
	// greeter("world")

	websiteList := []string{
		"https://github.com",
		"https://go.dev",
		"https://google.com",
	}

	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()

	fmt.Println(signals)
}

func getStatusCode(endPoint string) {
	defer wg.Done()

	res, err := http.Get(endPoint)

	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endPoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s \n", res.StatusCode, endPoint)

	}

}

func greeter(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}

type courses struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func EncodeJson() {
	lcoCourse := []courses{
		{"React JS", 299, "mychallen.in", "abcd123", []string{"web", "js"}},
		{"Node JS", 299, "mychallen.in", "abcd123", []string{"backend", "js"}},
		{"Full JS", 299, "mychallen.in", "abcd123", nil},
	}

	// package to json date

	finalJson, err := json.MarshalIndent(lcoCourse, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "React JS",
		"Price": 299,
		"website": "mychallen.in",
		"tags": [
				"web",
				"js"
		]
    }
	`)

	var lcoCourse courses

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("Valid json")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v \n", lcoCourse)
	}

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v \n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Keys is %v and value is %v and type is %T \n", k, v, v)
	}
}

func readFile(filename string) {
	dataByte, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dataByte))
}

func adder(val1 int, val2 int) int {
	return val1 + val2
}

func ProAdder(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

type User struct {
	Name  string
	Email string
	Age   int
}

func closestPrimes(left int, right int) []int {

	primes := make([]int, 0)
	minArr := make([]int, 2)
	if left%2 == 0 {
		left++
	}

	for i := left; i <= right; i += 2 {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}

	// if we not have 2 prime number
	if len(primes) < 2 {
		return []int{-1, -1}
	}

	leftIndex := 0
	rightIndex := 1
	min := math.MaxInt16

	for rightIndex < len(primes) {
		if (primes[rightIndex] - primes[leftIndex]) < min {
			min = primes[rightIndex] - primes[leftIndex]
			minArr[0] = primes[leftIndex]
			minArr[1] = primes[rightIndex]
		}
		leftIndex++
		rightIndex++
	}
	return minArr
}

func isPrime(num int) bool {

	for i := 3; i < num/2; i += 2 {
		if num%i == 0 {
			return false
		}
	}
	return true
}
