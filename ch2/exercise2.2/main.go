package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	temp := flag.Float64("temp", 0, "tempconv")
	length := flag.Float64("length", 0, "lengthconv")
	weight := flag.Float64("weight", 0, "weightconv")

	flag.Parse()

	if !isAllDefault(*temp, *length, *weight) {
		fmt.Println(tempString(*temp))
		fmt.Println(lengthString(*length))
		fmt.Println(weightString(*weight))
		return
	}

	tempPrefix := "temp: "
	lengthPrefix := "length: "
	weightPrefix := "weight: "

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		switch {
		case strings.HasPrefix(line, tempPrefix):
			fmt.Println(tempString(cToF(line, tempPrefix)))
		case strings.HasPrefix(line, lengthPrefix):
			fmt.Println(lengthString(cToF(line, lengthPrefix)))
		case strings.HasPrefix(line, weightPrefix):
			fmt.Println(weightString(cToF(line, weightPrefix)))
		}
	}
}

func cToF(line, prefix string) float64 {
	line = strings.TrimPrefix(line, prefix)
	f, err := strconv.ParseFloat(line, 64)
	if err != nil {
		log.Fatalf("parse float64 err: %v", err)
	}
	return f
}

func tempString(temp float64) string {
	return fmt.Sprintf("%s = %s = %s", Celsius(temp), CToK(Celsius(temp)), CToF(Celsius(temp)))
}

func lengthString(length float64) string {
	return fmt.Sprintf("%s = %s", Meter(length), MToF(Meter(length)))
}

func weightString(weight float64) string {
	return fmt.Sprintf("%s = %s", Kilogram(weight), KToP(Kilogram(weight)))
}

func isAllDefault(temp, length, weight float64) bool {
	return isDefault(temp) && isDefault(length) && isDefault(weight)
}

func isDefault(flag float64) bool { return flag == 0 }
