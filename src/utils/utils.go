package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func MakeGetRequest(url string) (string, error) {
	var myClient = &http.Client{Timeout: 3 * time.Second}
	resp, err := myClient.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", errors.New("non-200 status code")
	}
	// fmt.Println(resp)
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("error in reading response body")
	}
	return string(resBody), nil
}

func AppendToMarkdown(readFileName, writeFileName, data string) bool {
	markdownData, err := os.ReadFile(readFileName)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	newString := string(markdownData) + data + "\n"
	// fmt.Print(newString)
	err = os.WriteFile(writeFileName, []byte(newString), 0644)
	return err == nil
}
