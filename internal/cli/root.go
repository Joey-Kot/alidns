package cli

import (
	"context"
	"flag"
	"fmt"
	"io"

	"alidns/internal/alidns"
)

type APIFactory func(accessKeyID, accessKeySecret string) (alidns.DNSAPI, error)

type Deps struct {
	Stdout io.Writer
	Stderr io.Writer
	NewAPI APIFactory
}

func NewDefaultDeps(stdout, stderr io.Writer) Deps {
	return Deps{
		Stdout: stdout,
		Stderr: stderr,
		NewAPI: func(accessKeyID, accessKeySecret string) (alidns.DNSAPI, error) {
			client, err := alidns.CreateClient(accessKeyID, accessKeySecret)
			if err != nil {
				return nil, err
			}
			return alidns.NewSDKClient(client), nil
		},
	}
}

func Run(args []string, deps Deps) error {
	if deps.Stdout == nil || deps.Stderr == nil {
		return fmt.Errorf("invalid deps: stdout/stderr is nil")
	}
	if deps.NewAPI == nil {
		return fmt.Errorf("invalid deps: NewAPI is nil")
	}

	rootFlags := flag.NewFlagSet("alidns", flag.ContinueOnError)
	rootFlags.SetOutput(deps.Stderr)
	outputRaw := rootFlags.String("output", string(OutputPretty), "output format: json|pretty")
	help := rootFlags.Bool("help", false, "show help")
	rootFlags.BoolVar(help, "h", false, "show help")
	rootFlags.Usage = func() {
		printRootUsage(deps.Stderr)
	}

	if err := rootFlags.Parse(args); err != nil {
		return err
	}
	if *help {
		rootFlags.Usage()
		return nil
	}

	globalOutput, err := ParseOutputFormat(*outputRaw)
	if err != nil {
		return err
	}

	rest := rootFlags.Args()
	if len(rest) == 0 {
		rootFlags.Usage()
		return fmt.Errorf("missing command")
	}

	cmd, cmdArgs := rest[0], rest[1:]
	ctx := context.Background()
	switch cmd {
	case "add":
		return runAdd(ctx, cmdArgs, globalOutput, deps)
	case "del":
		return runDel(ctx, cmdArgs, globalOutput, deps)
	case "query":
		return runQuery(ctx, cmdArgs, globalOutput, deps)
	case "update":
		return runUpdate(ctx, cmdArgs, globalOutput, deps)
	case "help":
		if len(cmdArgs) == 0 {
			rootFlags.Usage()
			return nil
		}
		if len(cmdArgs) > 1 {
			return fmt.Errorf("help 仅支持一个子命令")
		}
		if err := printCommandUsage(cmdArgs[0], deps.Stderr, globalOutput); err != nil {
			rootFlags.Usage()
			return err
		}
		return nil
	default:
		rootFlags.Usage()
		return fmt.Errorf("unknown command %q", cmd)
	}
}
