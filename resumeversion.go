// Copyright 2016 Jesse Allen.

// Update applies the data from a json file to templates
package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type AlternateDelims struct {
	Left, Right string
}

var DELIMS_FOR_TEX = AlternateDelims{Left: "<<", Right: ">>"}

type resumeVersion struct {
	bytes.Buffer
	template       *template.Template
	outputFileName string
	delims         AlternateDelims
}

func newResumeVersion(templateFile, outputDir string) (*resumeVersion, error) {
	rv := new(resumeVersion)
	if err := rv.init(templateFile, outputDir); err != nil {
		return nil, err
	}
	return rv, nil
}

func (rv *resumeVersion) init(templateFile, outputDir string) error {
	_, n := filepath.Split(templateFile)

	// assign output file name
	rv.outputFileName = filepath.Join(filepath.Clean(outputDir), strings.TrimSuffix(n, filepath.Ext(n)))

	// new template
	rv.template = template.New(n)

	// check for alternate delims
	switch strings.ToLower(filepath.Ext(rv.outputFileName)) {
	case ".tex":
		_ = rv.template.Delims(DELIMS_FOR_TEX.Left, DELIMS_FOR_TEX.Right)
	}

	// add some utility functions to the templates
	rv.template.Funcs(template.FuncMap{
		"mod": func(a, b int) int { return a % b },
		"add": func(a, b int) int { return a + b },
		"quickfilter": func(s, find string) string {
			replacements := map[string]string{
				"LaTeX": "\\LaTeX{}",
			}
			if rep := replacements[find]; len(rep) > 0 {
				s = strings.Replace(s, find, rep, -1)
			}
			return s
		},
		"title": strings.Title,
	})

	// load template
	if _, err := rv.template.ParseFiles(templateFile); err != nil {
		return err
	}

	return nil
}

// Execute the template
func (rv *resumeVersion) Execute(r Resume) error {
	return rv.template.Execute(rv, r)
}

// Save the file
func (rv *resumeVersion) Save(forceSave bool) error {
	flag := os.O_RDWR | os.O_CREATE | os.O_EXCL
	if forceSave {
		flag = os.O_RDWR | os.O_CREATE | os.O_TRUNC
	}
	out, err := os.OpenFile(rv.outputFileName, flag, 0666)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = rv.WriteTo(out)
	return err
}

func (rv resumeVersion) Name() string {
	return rv.template.Name() + " / " + rv.outputFileName
}
