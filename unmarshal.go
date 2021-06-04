package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json: "First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"First":"james","Last":"bond","Age":44},{"First":"vijendra","Last":"chandel","Age":24}]`
	bs := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)
	var people []person
	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("all the data", people)
	for i, v := range people {
		fmt.Println("person ", i)
		fmt.Println(v.First, v.Last, v.Age)
	}
}
