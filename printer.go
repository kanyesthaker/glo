package printer

import (
	. "types"
	"strings"
	"fmt"
)

func pr_str(m types.MalType) string {
	switch t := m.(type) {
	case types.List:
		arr := make([]string, 0, len(t.Val))
		for _, e := range(t.Val) {
			arr = append(arr, pr_str(t.Val))
		}
		return strings.Join(arr, "")
	case types.Symbol:
		return t.Val
	default:
		return fmt.Sprintf("%v", m)
	}
}