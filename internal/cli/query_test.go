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
	"bytes"
	"strings"
	"testing"

	"alidns/internal/alidns"
	alidns20150109 "github.com/alibabacloud-go/alidns-20150109/v4/client"
)

func TestQueryEmptyOutputPretty(t *testing.T) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	api := &fakeDNSAPI{queryResp: nil}

	err := Run([]string{"query", "-ak", "ak", "-sk", "sk", "-domain", "example.com", "--output", "pretty"}, Deps{
		Stdout: stdout,
		Stderr: stderr,
		NewAPI: func(_, _ string) (alidns.DNSAPI, error) { return api, nil },
	})
	if err != nil {
		t.Fatalf("Run returned error: %v", err)
	}

	got := strings.TrimSpace(stdout.String())
	if got != "[]" {
		t.Fatalf("unexpected pretty output: %q", got)
	}
}

func TestQueryJSONOutputWithRecord(t *testing.T) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	api := &fakeDNSAPI{queryResp: []*alidns20150109.DescribeDomainRecordsResponseBodyDomainRecordsRecord{{RecordId: nil}}}

	err := Run([]string{"query", "-ak", "ak", "-sk", "sk", "-domain", "example.com", "--output", "json"}, Deps{
		Stdout: stdout,
		Stderr: stderr,
		NewAPI: func(_, _ string) (alidns.DNSAPI, error) { return api, nil },
	})
	if err != nil {
		t.Fatalf("Run returned error: %v", err)
	}

	got := strings.TrimSpace(stdout.String())
	if !strings.HasPrefix(got, "[") || !strings.HasSuffix(got, "]") {
		t.Fatalf("unexpected json output: %q", got)
	}
}
