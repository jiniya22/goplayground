package publicpkg

import "fmt"

const (
	PI = 3.1415
	pi = 3.14
)

var ScreenSize = 1080
var screenHeight int

func PublicFunc() {
	const MyConst = 100
	fmt.Println("PublicFunc 는 공개함수입니다!", MyConst)
}

func privateFunc() {
	fmt.Println("privateFunc 는 비공개함수입니다!")
}

type MyInt int
type myString string

type MyStruct struct {
	Age  int
	Name string
}

func (m MyStruct) PublicMethod() {
	fmt.Println("MyStruct 에 들어있는 공개 메서드")
}

func (m MyStruct) privateMethod() {
	fmt.Println("MyStruct 에 들어있는 비공개 메서드")
}

type myPrivateStruct struct {
	Age  int
	Name string
}

func (m myPrivateStruct) PrivateMethod() {
	fmt.Println("myPrivateStruct 에 들어있는 비공개 메서드")
}
