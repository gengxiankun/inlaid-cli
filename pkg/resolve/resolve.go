package resolve

import (
	"github.com/dlclark/regexp2"
)

func ResolveGitCommitContext(history string) string {
	expr := `(?<=git commit -am ["|']).*(?=["|'])`
	reg, _ := regexp2.Compile(expr, 0)
	m, _ := reg.FindStringMatch(history)
	if m == nil {
		return "" 
	}
	return m.String()
}
