package tpl

import (
	"bytes"
	"text/template"
)

/**
 * 模板
 * @author jensen.chen
 * @date 2022/7/12
 */
func ProcessTemplate(tpl string, vars interface{}) (string, error) {
	var buf bytes.Buffer
	tmpl, err := template.New("tmpl").Parse(tpl)

	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(&buf, vars); err != nil {
		return "", err
	} else {
		return buf.String(), nil
	}
}
