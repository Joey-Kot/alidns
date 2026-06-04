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
	"context"
	"fmt"

	"alidns/internal/alidns"
	alidns20150109 "github.com/alibabacloud-go/alidns-20150109/v4/client"
)

func runQuery(ctx context.Context, args []string, globalOutput OutputFormat, deps Deps) error {
	fs, f := newQueryFlagSet(deps.Stderr, globalOutput)
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

	records, err := svc.Query(ctx, alidns.QueryInput{DomainName: f.domain})
	if err != nil {
		return err
	}
	if records == nil {
		records = []*alidns20150109.DescribeDomainRecordsResponseBodyDomainRecordsRecord{}
	}

	return Print(deps.Stdout, records, output)
}
