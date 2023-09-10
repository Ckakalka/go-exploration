package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type ErrorsArr []error

func (eArr ErrorsArr) Error() string {
	var builder strings.Builder
	for _, err := range eArr {
		builder.WriteString(err.Error())
		builder.WriteString(" ")
	}
	return builder.String()
}

func MyCheck(s string) error {
	eArr := make(ErrorsArr, 0, 3)
	if len([]rune(s)) >= 20 {
		eArr = append(eArr, errors.New("line is too long"))
	}
	isContainsDigit := false
	countSpaces := 0
	for _, symb := range s {
		if unicode.IsDigit(symb) {
			isContainsDigit = true
		}
		if symb == ' ' {
			countSpaces++
		}
	}
	if isContainsDigit {
		eArr = append(eArr, errors.New("found numbers"))
	}
	if countSpaces != 2 {
		eArr = append(eArr, errors.New("no two spaces"))
	}
	if len(eArr) == 0 {
		return nil
	}
	return eArr
}

func main() {
	for {
		fmt.Printf("Укажите строку (q для выхода): ")
		reader := bufio.NewReader(os.Stdin)
		ret, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}
		ret = strings.TrimRight(ret, "\n")
		if ret == `q` {
			break
		}
		if err = MyCheck(ret); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(`Строка прошла проверку`)
		}
	}
}
