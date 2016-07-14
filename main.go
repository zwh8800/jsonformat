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
	indent := flag.Int64("indent", 2, "specify indent count")

	obj := ParseJson(os.Stdin)
	fmt.Println(JsonStringify(obj, *indent))
}

func ParseJson(input io.Reader) interface{} {
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

func JsonStringify(obj interface{}, indent int64) string {
	spaces := make([]string, indent+1)

	data, err := json.MarshalIndent(obj, "", strings.Join(spaces, " "))
	if err != nil {
		return ""
	}
	return string(data)
}
