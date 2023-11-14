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
		divClass = " className={style.div}"
	}

	component := fmt.Sprintf(
		"function %s() {\n%sreturn(\n%s<div%s>\n%s<p>%s works!</p>\n%s</div>\n%s)\n}",
		name,
		indentation,
		indentation+indentation,
		divClass,
		indentation+indentation+indentation,
		name,
		indentation+indentation,
		indentation,
	)

	switch config.ExportType {
	case 1:
		return styleImport + "export default " + component
	case 2:
		return styleImport + "export " + component
	case 3:
		return styleImport + fmt.Sprintf("export const %s = () => {\n%s};\n", name, component)
	case 4:
		return styleImport + component + fmt.Sprintf("\nexport default %s;\n", name)
	case 5:
		return styleImport + component + fmt.Sprintf("\nexport { %s };\n", name)
	default:
		return styleImport + "export default " + component
	}
}
