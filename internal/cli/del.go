package cli

import (
	"context"
	"fmt"

	"alidns/internal/alidns"
)

func runDel(ctx context.Context, args []string, globalOutput OutputFormat, deps Deps) error {
	fs, f := newDelFlagSet(deps.Stderr, globalOutput)
	helpShown, err := parseFlagSet(fs, args)
	if err != nil {
		return err
	}
	if helpShown {
		return nil
	}

	if err := requireAll(
		requiredArg{name: "-ak", value: f.ak},
		requiredArg{name: "-sk", value: f.sk},
		requiredArg{name: "-domain", value: f.domain},
		requiredArg{name: "-name", value: f.name},
		requiredArg{name: "-type", value: f.rType},
	); err != nil {
		return err
	}

	output, err := ParseOutputFormat(f.output)
	if err != nil {
		return err
	}

	api, err := deps.NewAPI(f.ak, f.sk)
	if err != nil {
		return fmt.Errorf("创建 Alidns Client 失败: %w", err)
	}
	svc := alidns.NewService(api)

	resp, err := svc.Del(ctx, alidns.DelInput{
		DomainName: f.domain,
		Name:       f.name,
		Type:       f.rType,
	})
	if err != nil {
		return err
	}

	return Print(deps.Stdout, resp, output)
}
