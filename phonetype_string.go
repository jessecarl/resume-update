// Code generated by "stringer -type=PhoneType"; DO NOT EDIT

package main

import "fmt"

const _PhoneType_name = "unknownmobilehomeofficeother"

var _PhoneType_index = [...]uint8{0, 7, 13, 17, 23, 28}

func (i PhoneType) String() string {
	if i < 0 || i >= PhoneType(len(_PhoneType_index)-1) {
		return fmt.Sprintf("PhoneType(%d)", i)
	}
	return _PhoneType_name[_PhoneType_index[i]:_PhoneType_index[i+1]]
}
