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
	"fmt"
	"strings"
)

type requiredArg struct {
	name  string
	value string
}

func requireAll(args ...requiredArg) error {
	missing := make([]string, 0)
	for _, arg := range args {
		if strings.TrimSpace(arg.value) == "" {
			missing = append(missing, arg.name)
		}
	}
	if len(missing) == 0 {
		return nil
	}
	return fmt.Errorf("错误: 缺少必需的参数:%s", " "+strings.Join(missing, " "))
}
