// Code generated by "stringer -type=ExperienceType"; DO NOT EDIT

package main

import "fmt"

const _ExperienceType_name = "workvolunteer"

var _ExperienceType_index = [...]uint8{0, 4, 13}

func (i ExperienceType) String() string {
	if i < 0 || i >= ExperienceType(len(_ExperienceType_index)-1) {
		return fmt.Sprintf("ExperienceType(%d)", i)
	}
	return _ExperienceType_name[_ExperienceType_index[i]:_ExperienceType_index[i+1]]
}
