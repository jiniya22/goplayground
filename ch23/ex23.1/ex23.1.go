package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()
	rd := bufio.NewReader(file)
	line, _ := rd.ReadString('\n')
	return line, nil
}

func WriteFile(filename string, line string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = fmt.Fprintf(file, line)
	return err
}

func main() {
	filename := "data.txt"
	line, err := ReadFile(filename)
	if err != nil {
		err = WriteFile(filename, "파일이 존재하지 않아서 WriteFile 함수로 새로 생성했습니다.")
		if err != nil {
			fmt.Println("파일 생성중 에러 발생", err)
			return
		}
		line, err = ReadFile(filename)
		if err != nil {
			fmt.Println("파일 읽기 2번째 실패", err)
			return
		}
	}
	fmt.Println("파일에 쓰여진 값은", line)

}
