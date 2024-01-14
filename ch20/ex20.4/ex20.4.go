package main

type Member interface {
	Introduce()
}

func main() {
	var member Member
	member.Introduce() // 컴파일 에러 발생!
}
