package main

import "fmt"

type PasswordError struct {
	Len         int
	RequiredLen int
}

func (err PasswordError) Error() string {
	return "암호길이가 짧습니다."
}

type Account struct {
	Id       string
	Password string
}

func RegisterAccount(id, password string) (*Account, error) {
	const minimumLength = 8
	if len(password) < minimumLength {
		return nil, PasswordError{len(password), minimumLength}
	}
	return &Account{id, password}, nil
}

func main() {
	if account, err := RegisterAccount("jini", "jini123"); err != nil {
		if errInfo, ok := err.(PasswordError); ok {
			fmt.Printf("최소 길이: %d, 입력 길이: %d\n", errInfo.RequiredLen, errInfo.Len)
		}
	} else {
		fmt.Println("회원가입 성공! ", account)
	}
}
