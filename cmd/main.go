package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/m-machida1108/prime-factorization/pkg/primenum"
)

const (
	maxIn int64 = 1000000000
)

type Err struct {
	Msg  string
	Code int
}

func (e Err) Error() string {
	return e.Msg
}

func main() {
	flag.Parse()
	if err := exec(flag.Args()); err != nil {
		log.Printf("error: %s \n", err)
	}
}

func exec(args []string) error {
	if err := validationParam(args); err != nil {
		return err
	}
	target, _ := strconv.ParseInt(args[0], 10, 64)
	if primenum.Is(target) {
		log.Printf("%d は素数です。\n", target)
		return nil
	}
	res := primenum.Factorization(target)
	msgArray := make([]string, len(res))
	for i, r := range res {
		msgArray[i] = strconv.FormatInt(r, 10)
	}
	log.Printf("%d を素因数分解すると %s になります。\n", target, strings.Join(msgArray, " × "))

	return nil
}

func validationParam(args []string) error {
	if args == nil {
		return Err{"数値を１つ指定してください", 400}
	}
	if len(args) != 1 {
		return Err{"引数の数が不正です", 400}
	}
	converted, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		return Err{"数値を指定してください", 400}
	}
	if converted > maxIn {
		return Err{fmt.Sprintf("%d以下の数値を指定してください", maxIn), 400}
	}
	if converted <= 0 {
		return Err{"1以上の数値を指定してください", 400}
	}
	return nil
}
