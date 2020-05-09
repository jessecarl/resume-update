// Copyright 2016 Jesse Allen.

// Update applies the data from a json file to templates
package main

import (
	"encoding/json"
	"errors"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type stringSlice []string

func (ss *stringSlice) String() string {
	s := ""
	for _, si := range *ss {
		s += si
	}
	return s
}

func (ss *stringSlice) Set(value string) error {
	for _, s := range strings.Split(value, ",") {
		*ss = append(*ss, s)
	}
	return nil
}

func main() {
	var (
		inputFile     string
		outputDir     string
		templateFiles stringSlice
		forceSave     bool
	)
	flag.StringVar(&inputFile, "i", "", "[required] json input file containing résumé data")
	flag.StringVar(&outputDir, "o", ".", "output directory to save résumé files to")
	flag.Var(&templateFiles, "t", "[required] individual template files to generate résumés. For each template file 'resume.md.tmpl', an output file 'resume.md' will be created")
	flag.BoolVar(&forceSave, "f", false, "force save over existing files")
	flag.Parse()

	if len(inputFile) == 0 || len(templateFiles) == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if err := createOutputDirIfNotExist(outputDir); err != nil {
		log.Fatal(err)
	}

	// grab the data
	resume, err := getResumeFromFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Loaded Résumé data from %q", inputFile)

	// grab template files
	var rvs []*resumeVersion
	for _, t := range templateFiles {
		if rv, err := newResumeVersion(t, outputDir); err != nil {
			log.Printf("Failed to load %q: %v", t, err)
		} else {
			rvs = append(rvs, rv)
			log.Printf("Loaded template: %q", t)
		}
	}

	// execute templates
	var executed []*resumeVersion
	for _, rv := range rvs {
		if err := rv.Execute(*resume); err != nil {
			log.Printf("Failed to execute %q: %v", rv.Name(), err)
		} else {
			executed = append(executed, rv)
			log.Printf("Executed template: %q", rv.Name())
		}
	}

	// save résumé files
	for _, rv := range executed {
		if err := rv.Save(forceSave); err != nil {
			log.Printf("Failed to save %q: %v", rv.Name(), err)
		} else {
			log.Printf("Saved: %q", rv.Name())
		}
	}
}

func getResumeFromFile(in string) (resume *Resume, err error) {
	var resumeJson []byte
	resume = new(Resume)
	resumeJson, err = ioutil.ReadFile(in)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resumeJson, resume)
	if err != nil {
		return nil, err
	}
	return resume, nil
}

func createOutputDirIfNotExist(outputDir string) error {
	info, err := os.Stat(outputDir)
	// determine if the directory exists and is a directory
	// [x] error if it exists and is not a directory
	// [x] do nothing if it exists and is a directory
	// [ ] create the directory if it does not exist
	if err == nil {
		if info.IsDir() {
			return nil
		} else {
			return errors.New("Cannot replace file with directory")
		}
	} else if !os.IsNotExist(err) {
		return err
	}
	return os.MkdirAll(outputDir, os.ModePerm|os.ModeDir)
}
