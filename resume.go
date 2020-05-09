// Copyright 2016 Jesse Allen.

package main

import (
	"time"
)

// Container for all information used to present a CV or Résumé
type Resume struct {
	Contact     Contact      `json:"contact"`
	Summary     string       `json:"summary"`
	Experiences []Experience `json:"experiences"`
	Skills      []Skill      `json:"skills"`
	Education   []Education  `json:"education"`
	Awards      []Award      `json:"awards"`
}

type Experience struct {
	Type         ExperienceType `json:"type"`
	CVOnly       bool           `json:"cvOnly"`
	Organization string         `json:"organization"`
	Location     string         `json:"location"`
	Role         string         `json:"role"`
	StartDate    time.Time      `json:"startDate"`
	EndDate      time.Time      `json:"endDate"`
	Tasks        []string       `json:"tasks"`
}

type ExperienceType int

//go:generate stringer -type=ExperienceType
//go:generate jsonenums -type=ExperienceType

const (
	work ExperienceType = iota
	volunteer
)

type Skill struct {
	Type SkillType `json:"type"`
	Name string    `json:"name"`
}

type SkillType int

//go:generate stringer -type=SkillType
//go:generate jsonenums -type=SkillType

const (
	fluent SkillType = iota
	proficient
	competent
	familiar
)

type Education struct {
	Institution string `json:"institution"`
	Location    string `json:"location"`
	Degree      []struct {
		Type         string   `json:"type"`
		FieldOfStudy []string `json:"fieldOfStudy"`
		Minor        []string `json:"minor"`
	} `json:"degree"`
	Date time.Time `json:"date"`
}

type Award struct {
	Organization string    `json:"organization"`
	Award        string    `json:"award"`
	Description  string    `json:"description"`
	Date         time.Time `json:"date"`
}
