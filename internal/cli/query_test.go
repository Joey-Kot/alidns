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
