package json

import (
	"encoding/json"
	"github.com/loopholelabs/scale-go/scalefunc"
)

type output struct {
	Name       string `json:"name"`
	Tag        string `json:"tag"`
	Language   string `json:"language"`
	Middleware bool   `json:"middleware"`
}

func NewList(entries []*scalefunc.ScaleFunc, middleware bool) error {
	rows := make([]output, 0, len(entries))
	for _, scaleFunc := range entries {
		if middleware && !scaleFunc.ScaleFile.Middleware {
			continue
		}
		row := output{scaleFunc.ScaleFile.Name, "latest", scaleFunc.ScaleFile.Build.Language, false}
		if scaleFunc.ScaleFile.Middleware {
			row.Middleware = true
		}
		if scaleFunc.Tag != "" {
			row.Tag = scaleFunc.Tag
		}
		rows = append(rows, row)
	}

	data, err := json.MarshalIndent(rows, "", "  ")
	if err != nil {
		return err
	}

	println(string(data))

	return nil
}
