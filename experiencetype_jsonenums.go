// generated by jsonenums -type=ExperienceType; DO NOT EDIT

package main

import (
	"encoding/json"
	"fmt"
)

var (
	_ExperienceTypeNameToValue = map[string]ExperienceType{
		"work":      work,
		"volunteer": volunteer,
	}

	_ExperienceTypeValueToName = map[ExperienceType]string{
		work:      "work",
		volunteer: "volunteer",
	}
)

func init() {
	var v ExperienceType
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_ExperienceTypeNameToValue = map[string]ExperienceType{
			interface{}(work).(fmt.Stringer).String():      work,
			interface{}(volunteer).(fmt.Stringer).String(): volunteer,
		}
	}
}

// MarshalJSON is generated so ExperienceType satisfies json.Marshaler.
func (r ExperienceType) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _ExperienceTypeValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid ExperienceType: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so ExperienceType satisfies json.Unmarshaler.
func (r *ExperienceType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("ExperienceType should be a string, got %s", data)
	}
	v, ok := _ExperienceTypeNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid ExperienceType %q", s)
	}
	*r = v
	return nil
}
