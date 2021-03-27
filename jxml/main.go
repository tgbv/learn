package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type Carrots struct {
	Sweet bool `json:"id,omitempty"`
	Qty   uint `json:"-"`
	Color string
	Magic struct {
		A int
	}
}

func main() {

	carr := Carrots{
		Sweet: true,
		Qty:   34,
		Color: "orange",
		Magic: struct{ A int }{23243},
	}

	//json

	// encode
	conv, err := json.MarshalIndent(carr, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(conv))

	// decode
	whatever := map[string]interface{}{}
	err = json.Unmarshal(conv, &whatever)
	if err != nil {
		panic(err)
	}
	fmt.Println(whatever)

	// xml
	conv, err = xml.Marshal(carr)
	if err != nil {
		//panic(err)
	}

	fmt.Println(string(conv))

}
