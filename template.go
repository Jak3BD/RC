package main

import (
	"fmt"
	"strings"
)

func generateIndentation() string {
	if config.Indent.Type == "spaces" {
		return strings.Repeat(" ", config.Indent.Count)
	}
	return strings.Repeat("\t", config.Indent.Count)
}

func functionExport(name string) string {
	indentation := generateIndentation()
	quote := `"`
	if config.Quotes == "single" {
		quote = `'`
	}

	var styleImport, divClass string
	if config.Style.Enable {
		styleImport = fmt.Sprintf("import style from %s./%s.%s%s;\n\n", quote, name, config.Style.Ext, quote)
		divClass = "className={style.div}"
	}

	return styleImport + fmt.Sprintf(
		"export default function %s() {\n%sreturn(\n%s<div %s>\n%s<p>%s works!</p>\n%s</div>\n%s)\n}",
		name,
		indentation,
		indentation+indentation,
		divClass,
		indentation+indentation+indentation,
		name,
		indentation+indentation,
		indentation,
	)
}
