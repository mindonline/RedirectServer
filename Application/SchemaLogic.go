package Application

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func NewSchema() *Schema {
	var r = new(Schema)
	jsonFile, err := os.Open("schema.json")

	if err != nil {
		panic(err)
	} else {
		byteData, _ := ioutil.ReadAll(jsonFile)

		err := jsonFile.Close()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Trying to Unmarshal JSON")
		err = json.Unmarshal(byteData, &r)
		if err != nil {
			fmt.Println(err)
		}
	}
	return r
}
