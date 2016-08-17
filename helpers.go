package validate

import (
	"bytes"
	"errors"
	"text/template"
)

func FormatErrorMessage(tmpl *template.Template, params map[string]interface{}, value interface{}) error {
	data := make(map[string]interface{})
	for key, value := range params {
		data[key] = value
	}
	data["Value"] = value
	buf := bytes.NewBuffer(make([]byte, 0))
	if err := tmpl.Execute(buf, data); err != nil {
		return errors.New("an error occurred formatting the error message")
	}
	return &Error{Message: buf.String(), Params: params}
}

type Error struct {
	Message string
	Params  map[string]interface{}
}

func (e Error) Error() string { return e.Message }
