package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"
	"unicode/utf8"
)

func foo() string {
	return "this is a function"
}
func definingVariables() {
	var intNum int = 32767
	intNum = intNum + 1
	fmt.Println(intNum)

	var floatNum float32 = 12345678.9
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	var intNum1 int = 3
	var intNumb2 int = 2
	fmt.Println(intNum1 / intNumb2)

	var myString string = "Hello" + " " + "World"
	fmt.Println(myString)

	fmt.Println(utf8.RuneCountInString("YY"))

	var myRune rune = 'a'
	fmt.Println(myRune)

	var myBoolean bool = false
	fmt.Println(myBoolean)

	var intNum3 rune
	fmt.Println(intNum3)

	myVar := "text"
	fmt.Println(myVar)

	var myOtherVar = "text"
	fmt.Println(myOtherVar)

	var1, var2 := 1, 2
	fmt.Println(var1, var2)

	var myfunvar string = foo()
	fmt.Println(myfunvar)

	const myConst string = "const value"
	fmt.Println(myConst)
}

type gasEngine struct {
	mpg        uint8
	gallons    uint8
	owenerInfo owner
}

type owner struct {
	name string
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

type engine interface {
	milesLeft() uint8
}

var m = sync.RWMutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

var MAX_CHICKEN_PRICE float32 = 5
var MAX_TOFU_PRICE float32 = 3

func main() {
	// definingVariables()
	// usingClauses()
	// workingWithArrays()
	// workingWithIterations()
	// memoryAllocationPerformance()
	// workingWithString()
	//workingWithInterfacesAndStructs()
	//workingWithPointers()
	//workingWithArraysAndPointers()
	//workingWithConcurrency()
	//workingWithConcurrency2()
	//whatNotToDoWithChannels()
	//basicChannelProgram()
	//moreAdvancedChannelProgram()
	workingWithGenericsStructAndFunctions()

}
func workingWithGenericsStructAndFunctions() {
	var instSlice = []int{1, 2, 3}
	fmt.Println(sumIntSlice(instSlice))
	fmt.Println(sumSlice(instSlice))

	var float32Slice = []float32{1, 2, 3}
	fmt.Println(sumFloat32Slice(float32Slice))
	fmt.Println(sumSlice(float32Slice))

	var float64Slice = []float64{1, 2, 3}
	fmt.Println(sumFloat64Slice(float64Slice))
	fmt.Println(sumSlice(float64Slice))

	fmt.Println(isEmpty[int](instSlice))
	fmt.Println(isEmpty(instSlice))

	var contacts []contactInfo = loadJSON[contactInfo]("./contactInfo.json")
	fmt.Printf("\n%+v", contacts)

	var purchases []purchaseInfo = loadJSON[purchaseInfo]("./purchaseInfo.json")
	fmt.Printf("\n%+v", purchases)

	var gasCar = car[gasEngine]{
		carMake:  "Honda",
		carModel: "Civic",
		engine: gasEngine{
			gallons: 12,
			mpg:     40,
		},
	}

	var electricCar = car[electricEngine]{
		carMake:  "Honda",
		carModel: "Civic",
		engine: electricEngine{
			kwh:   57,
			mpkwh: 40,
		},
	}

	fmt.Println(gasCar)
	fmt.Println(electricCar)
}

type contactInfo struct {
	Name  string
	Email string
}

type purchaseInfo struct {
	Name   string
	Price  float32
	Amount int
}
type car[T gasEngine | electricEngine] struct {
	carMake  string
	carModel string
	engine   T
}

func loadJSON[T contactInfo | purchaseInfo](filepath string) []T {
	var data, _ = os.ReadFile(filepath)

	var loaded = []T{}
	json.Unmarshal(data, &loaded)

	return loaded
}

func isEmpty[T any](slice []T) bool {
	return len(slice) == 0
}

func sumIntSlice(slice []int) int {
	var sum int
	for _, v := range slice {
		sum += v
	}
	return sum
}
func sumFloat32Slice(slice []float32) float32 {
	var sum float32
	for _, v := range slice {
		sum += v
	}
	return sum
}

func sumFloat64Slice(slice []float64) float64 {
	var sum float64
	for _, v := range slice {
		sum += v
	}
	return sum
}

func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

func moreAdvancedChannelProgram() {
	var chickenChannel = make(chan string)
	var tofuChannel = make(chan string)

	var websites = []string{"walmart.com", "costco.com", "wholefoods.com"}
	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
		go checkTofuPrices(websites[i], tofuChannel)
	}
	sendMessage(chickenChannel, tofuChannel)
}

func checkTofuPrices(website string, tofuChannel chan string) {
	for {
		time.Sleep(time.Second + 1)
		var tofuPrice = rand.Float32() * 20
		if tofuPrice < MAX_TOFU_PRICE {
			tofuChannel <- website
			break
		}
	}
}
func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second + 1)
		var chickenPrice = rand.Float32() * 20
		if chickenPrice < MAX_CHICKEN_PRICE {
			chickenChannel <- website
			break
		}
	}
}

func sendMessage(chickenChannel chan string, tofuChannel chan string) {
	select {
	case website := <-chickenChannel:
		fmt.Printf("\nFound a deal on chicken at %s", website)
	case website := <-tofuChannel:
		fmt.Printf("\nFound a deal on tofu at %s", website)
	}
}

func basicChannelProgram() {
	var c = make(chan int, 5)
	go process(c)
	for i := range c {
		fmt.Println(i)
		time.Sleep(time.Second * 1)
	}
}
func process(c chan int) {
	defer close(c)
	for i := 0; i < 5; i++ {
		c <- i
	}
	fmt.Println("Exiting process")
}

func whatNotToDoWithChannels() {
	var c = make(chan int)
	c <- 1 //will wait forever until something else reads from the channel
	var i = <-c
	fmt.Println(i)
}
func workingWithConcurrency2() {
	t0 := time.Now()
	for i := 0; i < 2000000; i++ {
		wg.Add(1)
		//go dbCall2(i)
		go count()
	}
	wg.Wait()
	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
}
func dbCall2(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	wg.Done()
}
func count() {
	var res int
	for i := 0; i < 10000; i++ {
		res += 1
	}
	wg.Done()
}
func workingWithConcurrency() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
	fmt.Printf("\nThe results are %v", results)
}
func dbCall(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the database is: ", dbData[i])
	save(dbData[i])
	log()
	results = append(results, dbData[i])
	log()
	wg.Done()
}

func save(result string) {
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

func log() {
	m.RLock()
	fmt.Printf("\nThe current results are: %v", results)
	m.RUnlock()
}

func workingWithArraysAndPointers() {
	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("\nThe memory location of the thing1 array is: %p", &thing1)
	var result [5]float64 = square(&thing1)
	fmt.Printf("\nThe result is %v", result)
	fmt.Printf("\nThe value of thing1 is%v", thing1)
}
func square(thing2 *[5]float64) [5]float64 {
	fmt.Printf("\nThe memory location of the thing2 array is: %p", &thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return *thing2
}

func workingWithPointers() {
	var p *int32 = new(int32)
	var i int32
	fmt.Printf("The value of p points to : %v", *p)
	fmt.Printf("\nThe value of i is: %v", i)
	*p = 10
	var samei *int32 = new(int32)
	samei = &i
	fmt.Printf("\nThe value of samei is: %v", *samei)
	*samei = 100

	fmt.Printf("\nThe value of samei is: %v", *samei)
	fmt.Printf("\nThe value of i is: %v\n", i)

	var slice = []int32{1, 2, 3}
	var sliceCopy = slice
	sliceCopy[2] = 4
	fmt.Println(slice)
	fmt.Println(sliceCopy)
}
func workingWithInterfacesAndStructs() {
	var myEngine gasEngine = gasEngine{25, 1, owner{"Olonyl"}}
	var myElectricEngine electricEngine = electricEngine{25, 15}
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.owenerInfo.name)
	var myEngine2 = struct {
		mpg     uint8
		gallons uint8
	}{25, 25}
	fmt.Println(myEngine2.mpg, myEngine2.gallons)
	fmt.Printf("Total miles left in tank %v\n", myEngine.milesLeft())
	canMakeIt(myEngine, 50)
	canMakeIt(myElectricEngine, 50)
}
func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}
func (e electricEngine) milesLeft() uint8 {
	return e.mpkwh * e.kwh
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("Need to fuel up first!")
	}
}

func workingWithString() {
	var myString = []rune("résumé")
	var indexed = myString[1]
	fmt.Printf("%v", indexed)
	fmt.Printf("\n%v %T\n", indexed, indexed)
	for i, v := range myString {
		fmt.Printf("\n%v %v", i, v)
	}

	var myRune = 'a'
	fmt.Printf("\nmyRune = %v\n", myRune)

	var strSlice = []string{"s", "u", "b", "s", "c", "r", "i", "b", "e"}
	var strBuilder strings.Builder

	var catStr = ""
	for i := range strSlice {
		catStr += strSlice[i]
		strBuilder.WriteString((strSlice[i]))
	}
	fmt.Printf("%v\n", catStr)
	fmt.Printf("%v\n", strBuilder.String())
}
func memoryAllocationPerformance() {
	var n int = 1000000
	var testSlice = []int{}
	var testSlice2 = make([]int, 0, n)

	fmt.Printf("Total time without preallocation: %v\n", timeLoop(testSlice, n))
	fmt.Printf("Total time with preallocation: %v\n", timeLoop(testSlice2, n))
}
func timeLoop(slice []int, n int) time.Duration {
	var t0 = time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(t0)
}
func usingClauses() {
	var printValue string = "Hello world using functions"
	printMe(printValue)

	var numerator int = 11
	var denominator int = 11
	var result, remainder, err = intDivision(numerator, denominator)

	if err != nil {
		fmt.Printf("%s", err.Error())
	} else if remainder == 0 {
		fmt.Printf("The result of the integer division is %v", result)
	} else {
		fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
	}

	switch {
	case err != nil:
		fmt.Printf("\n%s", err.Error())
	case remainder == 0:
		fmt.Printf("\nThe result of the integer division is %v", result)
	default:
		fmt.Printf("\nThe result of the integer division is %v with remainder %v", result, remainder)
	}
	switch remainder {
	case 0:
		fmt.Printf("\nThe division was exact")
	case 1, 2:
		fmt.Printf("\nThe division was close")
	default:
		fmt.Printf("\nThe division was not close")
	}
}
func workingWithArrays() {
	var intArr [3]int32
	intArr[1] = 123

	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	var intArr2 [3]int32 = [3]int32{1, 2, 3}
	fmt.Println(intArr2)

	intArr3 := [...]int32{1, 2, 3, 4}
	fmt.Println(intArr3)

	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Printf("\nThe lenght is %v with capacity %v", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("\nThe lenght is %v with capacity %v\n", len(intSlice), cap(intSlice))

	var intSlice2 []int32 = []int32{8, 9}
	intSlice2 = append(intSlice, intSlice2...)
	fmt.Println(intSlice2)

	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Println(intSlice3)

}
func workingWithIterations() {
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45}
	fmt.Println(myMap2)
	fmt.Println(myMap2["Adam"])
	var age, ok = myMap2["NotExist"]
	if ok {
		fmt.Printf("The age is %v", age)
	} else {
		fmt.Println("Invalid Name")
	}

	for name, age := range myMap2 {
		fmt.Printf("Name: %v, Age: %v \n", name, age)
	}

	var intArr [3]int32 = [3]int32{1, 2, 3}
	for i, v := range intArr {
		fmt.Printf("Index: %v, Value: %v \n", i, v)
	}

	var i int = 0
	for i < 10 {
		fmt.Println(i)
		i = i + 1
	}
	i = 0
	for {
		if i > 10 {
			break
		}
		fmt.Println(i)
		i = i + 1
	}
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("cannot divide by 0")
		return 0, 0, err
	}
	result := numerator / denominator
	remainder := numerator % denominator
	return result, remainder, err
}
