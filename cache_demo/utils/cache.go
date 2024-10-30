package utils

import (
	"encoding/json"
	"log"
	"regexp"
	"strconv"
	"strings"
)

const (
	B = 1 << (10 * iota)
	KB
	MB
	GB
	TB
)

func ParseSize(size string) (string, int64) {
	re, err := regexp.Compile("[0-9]+")
	if err != nil {
		log.Fatal(err.Error())
	}
	unit := string(re.ReplaceAll([]byte(size), []byte("")))
	num, _ := strconv.ParseInt(strings.Replace(size, unit, "", 1), 10, 64)
	unit = strings.ToUpper(unit)
	var numTemp int64 = 0
	switch unit {
	case "B":
		numTemp = num * B
	case "KB":
		numTemp = num * KB
	case "MB":
		numTemp = num * MB
	case "GB":
		numTemp = num * GB
	case "TB":
		numTemp = num * TB
	default:
		numTemp = 0
	}
	if numTemp == 0 {
		log.Println("bad input, set size to default")
		numTemp = 100 * MB
		num = 100
		unit = "MB"
	}
	size = strconv.FormatInt(num, 10) + unit
	return size, numTemp

}

func GetValSize(val interface{}) int64 {
	bytes, _ := json.Marshal(val)
	size := int64(len(bytes))
	return size
}
