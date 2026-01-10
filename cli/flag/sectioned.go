package flag

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/spf13/pflag"
)

// NamedFlagSets 按名称组织的 FlagSet 集合。
type NamedFlagSets struct {
	Order    []string
	FlagSets map[string]*pflag.FlagSet
}

// FlagSet 获取或创建一个命名的 FlagSet。
func (nfs *NamedFlagSets) FlagSet(name string) *pflag.FlagSet {
	if nfs.FlagSets == nil {
		nfs.FlagSets = make(map[string]*pflag.FlagSet)
	}
	if _, ok := nfs.FlagSets[name]; !ok {
		nfs.FlagSets[name] = pflag.NewFlagSet(name, pflag.ExitOnError)
		nfs.Order = append(nfs.Order, name)
	}
	return nfs.FlagSets[name]
}

// PrintSections 以分节格式将命名的 FlagSet 集合打印到指定的 io.Writer。
// 每节的宽度由 cols 参数指定。if cols 大于 24，则会添加一个占位符标志以确保正确的换行。
func PrintSections(w io.Writer, fss NamedFlagSets, cols int) {
	for _, name := range fss.Order {
		fs := fss.FlagSets[name]
		if !fs.HasFlags() {
			continue
		}

		wideFS := pflag.NewFlagSet("", pflag.ExitOnError)
		wideFS.AddFlagSet(fs)

		var zzz string
		if cols > 24 {
			zzz = strings.Repeat("z", cols-24)
			wideFS.Int(zzz, 0, strings.Repeat("z", cols-24))
		}

		var buf bytes.Buffer
		_, _ = fmt.Fprintf(&buf, "\n%s flags:\n\n%s", strings.ToUpper(name[:1])+name[1:], wideFS.FlagUsagesWrapped(cols))

		if cols > 24 {
			i := strings.Index(buf.String(), zzz)
			lines := strings.Split(buf.String()[:i], "\n")
			_, _ = fmt.Fprint(w, strings.Join(lines[:len(lines)-1], "\n"))
			_, _ = fmt.Fprintln(w)
		} else {
			_, _ = fmt.Fprint(w, buf.String())
		}
	}
}
