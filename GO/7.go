package main

import "fmt"

func xMap() {
	var a map[string]string = make(map[string]string)
	fmt.Println(a)
	a["en"] = "English"
	fmt.Println(a)

	var b map[string]string = map[string]string{"en": "English", "fa": "Farsi"}

	fmt.Println(b)
	c := b["en"]
	fmt.Println(c)
	d, found := b["qq"]
	fmt.Println(d, "--", found)
	b["qq"] = "UFO Language"
	d, found = b["qq"]
	fmt.Println(d, "--", found)
	delete(b, "qq")
	fmt.Println(b)

	for key, value := range b {
		fmt.Println(key, "-->", value)
	}
}

func main() {
	xMap()
}
