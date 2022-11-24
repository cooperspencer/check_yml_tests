package main

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

func Ordinalize(num int) string {

	var ordinalDictionary = map[int]string{
		0: "th",
		1: "st",
		2: "nd",
		3: "rd",
		4: "th",
		5: "th",
		6: "th",
		7: "th",
		8: "th",
		9: "th",
	}

	// math.Abs() is to convert negative number to positive
	floatNum := math.Abs(float64(num))
	positiveNum := int(floatNum)

	if ((positiveNum % 100) >= 11) && ((positiveNum % 100) <= 13) {
		return strconv.Itoa(num) + "th"
	}

	return strconv.Itoa(num) + ordinalDictionary[positiveNum]

}

func readyml(path string) {
	yamlFile, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("yamlFile.Get err   #%v\n", err.Error())
	}

	dec := yaml.NewDecoder(bytes.NewReader(yamlFile))
	i := 1
	for {
		c := make(map[string]interface{})
		err = dec.Decode(&c)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("an error occured in the %s part of the yaml\n", Ordinalize(i))
			fmt.Println(err.Error())
		}

		if !reflect.ValueOf(c).IsZero() {
			i++
		}
		//	err = yaml.Unmarshal(yamlFile, c)
		//	if err != nil {
		//		fmt.Printf("Unmarshal: %v\n", err.Error())
		//	}
	}
}

func main() {
	err := filepath.Walk(".",
		func(path string, i os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if strings.HasSuffix(path, ".yml") || strings.HasSuffix(path, ".yaml") {
				fmt.Println(path)
				readyml(path)
			}

			return nil
		})
	if err != nil {
		fmt.Println(err.Error())
	}
}
