package main

import (
	"errors"
	"log"
	"regexp"
	"strconv"
)

func generateConfig() {
	err := indent()
	if err != nil {
		log.Println(err)
		generateConfig()
		return
	}

	quotes()
	exportType()
	style()

	err = saveConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func indent() error {
	t := selectOption("Indent type", []string{"tabs", "spaces"})
	c := inputOption("Indent count", "1")

	matched, err := regexp.MatchString("^[0-9]+$", c)
	if err != nil {
		return err
	}

	if !matched {
		return errors.New("indent count must be a number")
	}

	count, err := strconv.Atoi(c)
	if err != nil {
		return err
	}

	config.Indent.Type = t
	config.Indent.Count = count

	return nil
}

func quotes() {
	t := selectOption("Quotes", []string{"double", "single"})
	config.Quotes = t
}

func exportType() {
	var opts []string

	for _, v := range funcTypeMap {
		opts = append(opts, v.Name)
	}

	t := selectOption("Export type", opts)

	for _, v := range funcTypeMap {
		if v.Name == t {
			id, err := strconv.Atoi(v.ID)
			if err != nil {
				log.Fatal(err)
			}

			config.ExportType = id
		}
	}
}

func style() {
	e := selectOption("Enable style modules", []string{"yes", "no"})
	if e == "no" {
		config.Style.Enable = false
		return
	}

	t := selectOption("Style file extension", []string{"css", "scss", "custom"})
	if t == "custom" {
		t = inputOption("Enter custom extension", "")
	}

	config.Style.Enable = true
	config.Style.Ext = t
}
