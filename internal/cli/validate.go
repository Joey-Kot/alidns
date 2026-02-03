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
