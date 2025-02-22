package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Username string `json:"username,omitempty"`
	Password string `json:"-"` // Will be ignored in JSON
}

func main() {
	p := &person{Name: "John Doe", Age: 30, Email: "john.doe@example.com", Username: "", Password: "1234"}

	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(string(b))
	// Unmarshal
	data := `{"name":"John Doe","age":30,"email":"john.doe@example.com"}`

	var r person
	err = json.Unmarshal([]byte(data), &r)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(r.Name, r.Age, r.Email, r.Username, r.Password)

	//

	numbers := []int{1, 2, 3, 4, 5}

	res, err := json.Marshal(numbers)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(string(res))
	matrix := [2][3]int{{1, 2, 3}, {4, 5, 6}}

	res, err = json.Marshal(matrix)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(string(res))

	//

	dict := map[string]interface{}{
		"name":  "John Doe",
		"age":   30,
		"email": "john.doe@example.com",
	}

	res, err = json.Marshal(dict)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(string(res))

	// https://github.com/tidwall/gjson/tree/master
	// https://github.com/tidwall/sjson
}
