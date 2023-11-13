package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	Indent     Indent `json:"indent"`
	Quotes     string `json:"quotes"`
	ExportType int    `json:"exportType"`
	Style      Style  `json:"style"`
}

type Indent struct {
	Type  string `json:"type"`
	Count int    `json:"count"`
}

type Style struct {
	Enable bool   `json:"enable"`
	Ext    string `json:"ext"`
}

type FuncTypeMap struct {
	Name string
	ID   string
}

const configFile = "rc.json"

var config Config

var funcTypeMap = []FuncTypeMap{
	{"Default Functional Export", "1"},
	{"Named Functional Export", "2"},
	{"Exported Arrow Function", "3"},
	{"Default Exported Function", "4"},
	{"Named Exported Function", "5"},
}

func loadConfig() error {
	file, err := os.Open(configFile)
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		return err
	}

	return nil
}

func saveConfig() error {
	file, err := os.Create(configFile)
	if err != nil {
		return err
	}
	defer file.Close()

	formattedConfig, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		return err
	}

	_, err = file.Write(formattedConfig)
	if err != nil {
		return err
	}

	return nil
}
