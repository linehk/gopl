package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/linehk/gopl/ch7/exercise7.15/eval"
)

func main() {
	fmt.Print("Expr: ")

	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	expr, err := eval.Parse(string(b))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n")

	env := Env(expr)
	fmt.Printf("Expr = %g\n", expr.Eval(env))
}

func Env(expr eval.Expr) eval.Env {
	env := make(eval.Env)
	scanner := bufio.NewScanner(os.Stdin)
	for _, v := range expr.Vars() {

		fmt.Printf("%s: ", v)

		if !scanner.Scan() {
			log.Fatalf("not enough var!")
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		val, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			log.Fatal(err)
		}

		// 填充表达式所需的变量
		env[v] = val
	}
	return env
}
