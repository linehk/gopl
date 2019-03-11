package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

var data = make([]XkcdInfo, 700)

func main() {
	for i := 1; i <= 600; i++ {
		/*
			err := Download(i)
			if err != nil {
				fmt.Printf("%v\n", err)
			}
		*/
		err := GetJsonInfo(i)
		if err != nil {
			fmt.Printf("%v\n", err)
		}
		fmt.Printf("%s\n", data[2].Title)
	}
}

func Search(num int) (*XkcdInfo, error) {
	resp, err := http.Get(XkcdURL + strconv.Itoa(num) + "/info.0.json")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("failed: %s", resp.Status)
	}

	var result XkcdInfo
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

func Download(num int) error {
	resp, err := http.Get(XkcdURL + strconv.Itoa(num) + "/info.0.json")
	if err != nil {
		resp.Body.Close()
		return err
	}
	f, err := os.Create("json/" + strconv.Itoa(num) + ".json")
	if err != nil {
		resp.Body.Close()
		return err
	}
	io.Copy(f, resp.Body)
	resp.Body.Close()
	return nil
}

func GetJsonInfo(num int) error {
	f, err := os.OpenFile("json/"+strconv.Itoa(num)+".json", os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	json.NewDecoder(f).Decode(&data)
	return nil
}
