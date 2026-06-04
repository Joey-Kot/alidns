// Copyright (C) 2026 Joey Kot <joey.kot.x@gmail.com>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed WITHOUT ANY WARRANTY; without even the
// implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
// See <https://www.gnu.org/licenses/> for more details.

package cli

import (
	"encoding/json"
	"fmt"
	"io"
)

type OutputFormat string

const (
	OutputJSON   OutputFormat = "json"
	OutputPretty OutputFormat = "pretty"
)

func ParseOutputFormat(v string) (OutputFormat, error) {
	format := OutputFormat(v)
	switch format {
	case OutputJSON, OutputPretty:
		return format, nil
	default:
		return "", fmt.Errorf("invalid --output value %q, expected json|pretty", v)
	}
}

func Print(w io.Writer, v any, format OutputFormat) error {
	var (
		data []byte
		err  error
	)

	switch format {
	case OutputJSON:
		data, err = json.Marshal(v)
	case OutputPretty:
		data, err = json.MarshalIndent(v, "", "  ")
	default:
		return fmt.Errorf("unsupported output format %q", format)
	}
	if err != nil {
		return err
	}
	if _, err := w.Write(data); err != nil {
		return err
	}
	_, err = w.Write([]byte("\n"))
	return err
}
