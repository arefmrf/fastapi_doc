// Code generated by "stringer -type=Age"; DO NOT EDIT.

package main

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[CHILDERN-0]
	_ = x[ADOLESCENTS-1]
	_ = x[ADULTS-2]
}

const _Age_name = "CHILDERNADOLESCENTSADULTS"

var _Age_index = [...]uint8{0, 8, 19, 25}

func (i Age) String() string {
	if i < 0 || i >= Age(len(_Age_index)-1) {
		return "Age(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Age_name[_Age_index[i]:_Age_index[i+1]]
}
