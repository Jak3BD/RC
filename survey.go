package main

import "github.com/AlecAivazis/survey/v2"

func selectOption(label string, opts []string) string {
	var res string
	prompt := &survey.Select{
		Message: label,
		Options: opts,
	}
	survey.AskOne(prompt, &res)

	return res
}

func inputOption(label, defaultAnswer string) string {
	var res string
	prompt := &survey.Input{
		Message: label,
		Default: defaultAnswer,
	}
	survey.AskOne(prompt, &res)

	return res
}
