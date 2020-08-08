package services

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	layoutISO = "2006-01-02"
	VARFORMAT = "?date="
)

func LoadData(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return responseData
}

func DateToUnix(date string) string {
	var unixDate int64
	if date == "" {
		unixDate = time.Now().Unix()
	} else {
		formatDate, _ := time.Parse(layoutISO, date)
		unixDate = formatDate.Unix()
	}
	return VARFORMAT + strconv.FormatInt(unixDate, 10)
}
