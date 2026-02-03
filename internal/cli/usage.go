package cli

import (
	"errors"
	"flag"
	"fmt"
	"io"
)

type addFlags struct {
	ak       string
	sk       string
	domain   string
	name     string
	rType    string
	value    string
	ttl      int64
	priority int64
	line     string
	output   string
}

type delFlags struct {
	ak     string
	sk     string
	domain string
	name   string
	rType  string
	output string
}

type queryFlags struct {
	ak     string
	sk     string
	domain string
	output string
}

type updateFlags struct {
	ak       string
	sk       string
	recordID string
	name     string
	rType    string
	value    string
	ttl      int64
	priority int64
	line     string
	output   string
}

func parseFlagSet(fs *flag.FlagSet, args []string) (bool, error) {
	if err := fs.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func newAddFlagSet(stderr io.Writer, globalOutput OutputFormat) (*flag.FlagSet, *addFlags) {
	f := &addFlags{}
	fs := flag.NewFlagSet("add", flag.ContinueOnError)
	fs.SetOutput(stderr)

	fs.StringVar(&f.ak, "ak", "", "Alibaba Cloud Access Key ID (必需)")
	fs.StringVar(&f.sk, "sk", "", "Alibaba Cloud Access Key Secret (必需)")
	fs.StringVar(&f.domain, "domain", "", "要添加记录的主域名 (必需)")
	fs.StringVar(&f.name, "name", "", "主机记录 (必需)")
	fs.StringVar(&f.rType, "type", "", "记录类型 (必需)")
	fs.StringVar(&f.value, "value", "", "记录值 (必需)")
	fs.Int64Var(&f.ttl, "ttl", 600, "TTL")
	fs.Int64Var(&f.priority, "priority", 1, "优先级")
	fs.StringVar(&f.line, "line", "default", "线路")
	fs.StringVar(&f.output, "output", string(globalOutput), "output format: json|pretty")
	fs.Usage = func() {
		printAddUsage(stderr, globalOutput)
	}

	return fs, f
}

func newDelFlagSet(stderr io.Writer, globalOutput OutputFormat) (*flag.FlagSet, *delFlags) {
	f := &delFlags{}
	fs := flag.NewFlagSet("del", flag.ContinueOnError)
	fs.SetOutput(stderr)

	fs.StringVar(&f.ak, "ak", "", "Alibaba Cloud Access Key ID (必需)")
	fs.StringVar(&f.sk, "sk", "", "Alibaba Cloud Access Key Secret (必需)")
	fs.StringVar(&f.domain, "domain", "", "要删除记录的主域名 (必需)")
	fs.StringVar(&f.name, "name", "", "主机记录 (必需)")
	fs.StringVar(&f.rType, "type", "", "记录类型 (必需)")
	fs.StringVar(&f.output, "output", string(globalOutput), "output format: json|pretty")
	fs.Usage = func() {
		printDelUsage(stderr, globalOutput)
	}

	return fs, f
}

func newQueryFlagSet(stderr io.Writer, globalOutput OutputFormat) (*flag.FlagSet, *queryFlags) {
	f := &queryFlags{}
	fs := flag.NewFlagSet("query", flag.ContinueOnError)
	fs.SetOutput(stderr)

	fs.StringVar(&f.ak, "ak", "", "Alibaba Cloud Access Key ID (必需)")
	fs.StringVar(&f.sk, "sk", "", "Alibaba Cloud Access Key Secret (必需)")
	fs.StringVar(&f.domain, "domain", "", "要查询的主域名 (必需)")
	fs.StringVar(&f.output, "output", string(globalOutput), "output format: json|pretty")
	fs.Usage = func() {
		printQueryUsage(stderr, globalOutput)
	}

	return fs, f
}

func newUpdateFlagSet(stderr io.Writer, globalOutput OutputFormat) (*flag.FlagSet, *updateFlags) {
	f := &updateFlags{}
	fs := flag.NewFlagSet("update", flag.ContinueOnError)
	fs.SetOutput(stderr)

	fs.StringVar(&f.ak, "ak", "", "Alibaba Cloud Access Key ID (必需)")
	fs.StringVar(&f.sk, "sk", "", "Alibaba Cloud Access Key Secret (必需)")
	fs.StringVar(&f.recordID, "id", "", "解析记录ID (必需)")
	fs.StringVar(&f.name, "name", "", "主机记录 (必需)")
	fs.StringVar(&f.rType, "type", "", "记录类型 (必需)")
	fs.StringVar(&f.value, "value", "", "记录值 (必需)")
	fs.Int64Var(&f.ttl, "ttl", 600, "TTL")
	fs.Int64Var(&f.priority, "priority", 1, "优先级")
	fs.StringVar(&f.line, "line", "default", "线路")
	fs.StringVar(&f.output, "output", string(globalOutput), "output format: json|pretty")
	fs.Usage = func() {
		printUpdateUsage(stderr, globalOutput)
	}

	return fs, f
}

func printRootUsage(w io.Writer) {
	_, _ = fmt.Fprintln(w, `用法:
  alidns [--output json|pretty] <command> [flags]
  alidns help [command]

全局参数:
  --output string
    	output format: json|pretty (default "pretty")
  -h, --help
    	显示帮助

命令:
  add      添加 DNS 记录
  del      删除 DNS 记录
  query    查询 DNS 记录
  update   修改 DNS 记录
  help     显示帮助

示例:
  alidns add -ak AK -sk SK -domain example.com -name www -type A -value 1.2.3.4
  alidns query -ak AK -sk SK -domain example.com --output json
`)

	_, _ = fmt.Fprintln(w, "子命令参数说明:")
	printAddUsage(w, OutputPretty)
	printDelUsage(w, OutputPretty)
	printQueryUsage(w, OutputPretty)
	printUpdateUsage(w, OutputPretty)
}

func printAddUsage(w io.Writer, globalOutput OutputFormat) {
	_, _ = fmt.Fprintln(w, `
用法:
  alidns add [flags]

说明:
  添加 DNS 记录。

参数:`)
	fs, _ := newAddFlagSet(w, globalOutput)
	fs.PrintDefaults()
	_, _ = fmt.Fprintln(w, `
示例:
  alidns add -ak AK -sk SK -domain example.com -name www -type A -value 1.2.3.4
  alidns add -ak AK -sk SK -domain example.com -name @ -type TXT -value hello --output json`)
}

func printDelUsage(w io.Writer, globalOutput OutputFormat) {
	_, _ = fmt.Fprintln(w, `
用法:
  alidns del [flags]

说明:
  删除 DNS 记录。

参数:`)
	fs, _ := newDelFlagSet(w, globalOutput)
	fs.PrintDefaults()
	_, _ = fmt.Fprintln(w, `
示例:
  alidns del -ak AK -sk SK -domain example.com -name www -type A`)
}

func printQueryUsage(w io.Writer, globalOutput OutputFormat) {
	_, _ = fmt.Fprintln(w, `
用法:
  alidns query [flags]

说明:
  查询 DNS 记录列表。

参数:`)
	fs, _ := newQueryFlagSet(w, globalOutput)
	fs.PrintDefaults()
	_, _ = fmt.Fprintln(w, `
示例:
  alidns query -ak AK -sk SK -domain example.com
  alidns query -ak AK -sk SK -domain example.com --output json`)
}

func printUpdateUsage(w io.Writer, globalOutput OutputFormat) {
	_, _ = fmt.Fprintln(w, `
用法:
  alidns update [flags]

说明:
  修改 DNS 记录。

参数:`)
	fs, _ := newUpdateFlagSet(w, globalOutput)
	fs.PrintDefaults()
	_, _ = fmt.Fprintln(w, `
示例:
  alidns update -ak AK -sk SK -id RECORD_ID -name www -type A -value 1.2.3.4`)
}

func printCommandUsage(command string, w io.Writer, globalOutput OutputFormat) error {
	switch command {
	case "add":
		printAddUsage(w, globalOutput)
	case "del":
		printDelUsage(w, globalOutput)
	case "query":
		printQueryUsage(w, globalOutput)
	case "update":
		printUpdateUsage(w, globalOutput)
	default:
		return fmt.Errorf("unknown help command %q", command)
	}
	return nil
}
