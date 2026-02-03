package cli

import (
	"bytes"
	"context"
	"errors"
	"strings"
	"testing"

	"alidns/internal/alidns"
	alidns20150109 "github.com/alibabacloud-go/alidns-20150109/v4/client"
	"github.com/alibabacloud-go/tea/tea"
)

type fakeDNSAPI struct {
	addCalled    bool
	delCalled    bool
	queryCalled  bool
	updateCalled bool

	addResp    *alidns20150109.AddDomainRecordResponseBody
	delResp    *alidns20150109.DeleteSubDomainRecordsResponseBody
	queryResp  []*alidns20150109.DescribeDomainRecordsResponseBodyDomainRecordsRecord
	updateResp *alidns20150109.UpdateDomainRecordResponseBody

	err error
}

func (f *fakeDNSAPI) AddDomainRecord(_ context.Context, _ *alidns20150109.AddDomainRecordRequest) (*alidns20150109.AddDomainRecordResponseBody, error) {
	f.addCalled = true
	return f.addResp, f.err
}

func (f *fakeDNSAPI) DeleteSubDomainRecords(_ context.Context, _ *alidns20150109.DeleteSubDomainRecordsRequest) (*alidns20150109.DeleteSubDomainRecordsResponseBody, error) {
	f.delCalled = true
	return f.delResp, f.err
}

func (f *fakeDNSAPI) DescribeDomainRecords(_ context.Context, _ *alidns20150109.DescribeDomainRecordsRequest) ([]*alidns20150109.DescribeDomainRecordsResponseBodyDomainRecordsRecord, error) {
	f.queryCalled = true
	return f.queryResp, f.err
}

func (f *fakeDNSAPI) UpdateDomainRecord(_ context.Context, _ *alidns20150109.UpdateDomainRecordRequest) (*alidns20150109.UpdateDomainRecordResponseBody, error) {
	f.updateCalled = true
	return f.updateResp, f.err
}

func TestRunDispatchAdd(t *testing.T) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	api := &fakeDNSAPI{addResp: &alidns20150109.AddDomainRecordResponseBody{RequestId: tea.String("req-1")}}

	err := Run([]string{"add", "-ak", "ak", "-sk", "sk", "-domain", "example.com", "-name", "www", "-type", "A", "-value", "1.2.3.4", "--output", "json"}, Deps{
		Stdout: stdout,
		Stderr: stderr,
		NewAPI: func(_, _ string) (alidns.DNSAPI, error) { return api, nil },
	})
	if err != nil {
		t.Fatalf("Run returned error: %v", err)
	}
	if !api.addCalled {
		t.Fatal("add command was not dispatched")
	}
	if got := strings.TrimSpace(stdout.String()); !strings.Contains(got, `"RequestId":"req-1"`) {
		t.Fatalf("unexpected output: %s", got)
	}
}

func TestRunMissingRequiredFlags(t *testing.T) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}

	err := Run([]string{"add", "-ak", "ak", "-domain", "example.com", "-name", "www", "-type", "A", "-value", "1.2.3.4"}, Deps{
		Stdout: stdout,
		Stderr: stderr,
		NewAPI: func(_, _ string) (alidns.DNSAPI, error) { return &fakeDNSAPI{}, nil },
	})
	if err == nil {
		t.Fatal("expected missing required flags error")
	}
	if !strings.Contains(err.Error(), "-sk") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestRunOutputOverrideBySubcommand(t *testing.T) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	api := &fakeDNSAPI{queryResp: []*alidns20150109.DescribeDomainRecordsResponseBodyDomainRecordsRecord{}}

	err := Run([]string{"--output", "pretty", "query", "-ak", "ak", "-sk", "sk", "-domain", "example.com", "--output", "json"}, Deps{
		Stdout: stdout,
		Stderr: stderr,
		NewAPI: func(_, _ string) (alidns.DNSAPI, error) { return api, nil },
	})
	if err != nil {
		t.Fatalf("Run returned error: %v", err)
	}
	if got := stdout.String(); got != "[]\n" {
		t.Fatalf("expected json one-line empty array, got %q", got)
	}
}

func TestRunServiceErrorPath(t *testing.T) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	api := &fakeDNSAPI{err: errors.New("boom")}

	err := Run([]string{"del", "-ak", "ak", "-sk", "sk", "-domain", "example.com", "-name", "www", "-type", "A"}, Deps{
		Stdout: stdout,
		Stderr: stderr,
		NewAPI: func(_, _ string) (alidns.DNSAPI, error) { return api, nil },
	})
	if err == nil || !strings.Contains(err.Error(), "boom") {
		t.Fatalf("expected boom error, got: %v", err)
	}
}

func TestRunRootHelpIncludesSubcommandFlags(t *testing.T) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}

	err := Run([]string{"-h"}, Deps{
		Stdout: stdout,
		Stderr: stderr,
		NewAPI: func(_, _ string) (alidns.DNSAPI, error) { return &fakeDNSAPI{}, nil },
	})
	if err != nil {
		t.Fatalf("Run returned error: %v", err)
	}

	helpText := stderr.String()
	if !strings.Contains(helpText, "alidns add [flags]") {
		t.Fatalf("help should include add usage, got: %s", helpText)
	}
	if !strings.Contains(helpText, "-ak string") || !strings.Contains(helpText, "-domain string") {
		t.Fatalf("help should include subcommand flags, got: %s", helpText)
	}
	if !strings.Contains(helpText, "alidns update [flags]") || !strings.Contains(helpText, "-id string") {
		t.Fatalf("help should include update flags, got: %s", helpText)
	}
}

func TestRunHelpAddShowsAddUsage(t *testing.T) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}

	err := Run([]string{"help", "add"}, Deps{
		Stdout: stdout,
		Stderr: stderr,
		NewAPI: func(_, _ string) (alidns.DNSAPI, error) { return &fakeDNSAPI{}, nil },
	})
	if err != nil {
		t.Fatalf("Run returned error: %v", err)
	}

	helpText := stderr.String()
	if !strings.Contains(helpText, "alidns add [flags]") {
		t.Fatalf("help add should include add usage, got: %s", helpText)
	}
	if !strings.Contains(helpText, "-value string") {
		t.Fatalf("help add should include add flags, got: %s", helpText)
	}
}

func TestRunSubcommandHelpFlagReturnsNil(t *testing.T) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	called := false

	err := Run([]string{"add", "-h"}, Deps{
		Stdout: stdout,
		Stderr: stderr,
		NewAPI: func(_, _ string) (alidns.DNSAPI, error) {
			called = true
			return &fakeDNSAPI{}, nil
		},
	})
	if err != nil {
		t.Fatalf("Run returned error: %v", err)
	}
	if called {
		t.Fatal("NewAPI should not be called when printing help")
	}
	if !strings.Contains(stderr.String(), "alidns add [flags]") {
		t.Fatalf("subcommand help should include add usage, got: %s", stderr.String())
	}
}
