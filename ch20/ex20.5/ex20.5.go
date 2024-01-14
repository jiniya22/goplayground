package main

import "fmt"

type Member interface {
	Introduce() string
}

type BasicMember struct {
	Name string
	Age  int
}

func (v *BasicMember) Introduce() string {
	return fmt.Sprintf("%s는 기본회원이며 나이는 %d 입니다.", v.Name, v.Age)
}

type VipMember struct {
	Name string
	Age  int
}

func (v *VipMember) Introduce() string {
	return fmt.Sprintf("%s는 Vip회원이며 나이는 %d 입니다.", v.Name, v.Age)
}

func main() {
	var member Member = &VipMember{"sol", 26}
	fmt.Println(member.Introduce())
	fmt.Printf("member의 타입은 %T\n\n", member)

	vipMember := member.(*VipMember)
	fmt.Println(vipMember.Introduce())
	fmt.Printf("vipMember의 타입은 %T\n\n", vipMember)

	basicMember := member.(*BasicMember) // Runtime Exception 발생!
	fmt.Println(basicMember.Introduce())
	fmt.Printf("normalMember의 타입은 %T", basicMember)
}
