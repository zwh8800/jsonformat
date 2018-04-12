package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	indent := flag.Int("indent", 2, "specify indent count")
	flag.Parse()

	obj := parseJSON(os.Stdin)
	fmt.Println(jsonStringify(obj, *indent))
}

func parseJSON(input io.Reader) interface{} {
	data, err := ioutil.ReadAll(input)
	if err != nil {
		return nil
	}
	var obj map[string]interface{}
	if err := json.Unmarshal(data, &obj); err != nil {
		return nil
	}

	return obj
}

func jsonStringify(obj interface{}, indent int) string {
	data, err := json.MarshalIndent(obj, "", strings.Repeat(" ", indent))
	if err != nil {
		return ""
	}
	return string(data)
}
