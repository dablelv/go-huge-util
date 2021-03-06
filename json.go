package util

import (
	"bytes"
	"encoding/json"
)

// ToIndentJSON converts the golang value to indent JSON string, such as a struct, map, slice, array and so on
func ToIndentJSON(obj interface{}) (string, error) {
	bs, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	err = json.Indent(buf, bs, "", "\t")
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
