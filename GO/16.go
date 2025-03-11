package main

import "fmt"

type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

type student struct {
	name   string
	age    int
	marks  []int
	grades map[string]int
}

func addToMarks(s *[]int) {
	*s = append(*s, len(*s))
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
	///////////////////////////////////////
	var s student
	fmt.Printf("%+v", s)
	st := new(student) // Allocated on heap
	// st := student{}    Usually on stack
	st.grades = make(map[string]int)
	st.grades["s"] = 1
	fmt.Printf("%+v\n", st)
	fmt.Printf("%#v\n", st)
	addToMarks(&(st.marks))
	fmt.Println(st.marks)
}
