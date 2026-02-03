package alidns

import (
	"context"

	alidns20150109 "github.com/alibabacloud-go/alidns-20150109/v4/client"
	"github.com/alibabacloud-go/tea/tea"
)

const (
	defaultTTL      int64 = 600
	defaultPriority int64 = 1
	defaultLine           = "default"
)

type Service struct {
	api DNSAPI
}

func NewService(api DNSAPI) *Service {
	return &Service{api: api}
}

type AddInput struct {
	DomainName string
	Name       string
	Type       string
	Value      string
	TTL        int64
	Priority   int64
	Line       string
}

type DelInput struct {
	DomainName string
	Name       string
	Type       string
}

type QueryInput struct {
	DomainName string
}

type UpdateInput struct {
	RecordID string
	Name     string
	Type     string
	Value    string
	TTL      int64
	Priority int64
	Line     string
}

func (s *Service) Add(ctx context.Context, in AddInput) (*alidns20150109.AddDomainRecordResponseBody, error) {
	req := &alidns20150109.AddDomainRecordRequest{
		Lang:       tea.String("en"),
		DomainName: tea.String(in.DomainName),
		RR:         tea.String(in.Name),
		Type:       tea.String(in.Type),
		Value:      tea.String(in.Value),
		TTL:        tea.Int64(defaultInt64(in.TTL, defaultTTL)),
		Priority:   tea.Int64(defaultInt64(in.Priority, defaultPriority)),
		Line:       tea.String(defaultString(in.Line, defaultLine)),
	}
	return s.api.AddDomainRecord(ctx, req)
}

func (s *Service) Del(ctx context.Context, in DelInput) (*alidns20150109.DeleteSubDomainRecordsResponseBody, error) {
	req := &alidns20150109.DeleteSubDomainRecordsRequest{
		DomainName: tea.String(in.DomainName),
		RR:         tea.String(in.Name),
		Type:       tea.String(in.Type),
		Lang:       tea.String("en"),
	}
	return s.api.DeleteSubDomainRecords(ctx, req)
}

func (s *Service) Query(ctx context.Context, in QueryInput) ([]*alidns20150109.DescribeDomainRecordsResponseBodyDomainRecordsRecord, error) {
	req := &alidns20150109.DescribeDomainRecordsRequest{
		DomainName: tea.String(in.DomainName),
		Lang:       tea.String("en"),
		Direction:  tea.String("ASC"),
		Status:     tea.String("Enable"),
		PageNumber: tea.Int64(1),
		PageSize:   tea.Int64(500),
		SearchMode: tea.String("LIKE"),
	}
	return s.api.DescribeDomainRecords(ctx, req)
}

func (s *Service) Update(ctx context.Context, in UpdateInput) (*alidns20150109.UpdateDomainRecordResponseBody, error) {
	req := &alidns20150109.UpdateDomainRecordRequest{
		Lang:     tea.String("en"),
		RR:       tea.String(in.Name),
		RecordId: tea.String(in.RecordID),
		Type:     tea.String(in.Type),
		Value:    tea.String(in.Value),
		TTL:      tea.Int64(defaultInt64(in.TTL, defaultTTL)),
		Priority: tea.Int64(defaultInt64(in.Priority, defaultPriority)),
		Line:     tea.String(defaultString(in.Line, defaultLine)),
	}
	return s.api.UpdateDomainRecord(ctx, req)
}

func defaultInt64(v, fallback int64) int64 {
	if v == 0 {
		return fallback
	}
	return v
}

func defaultString(v, fallback string) string {
	if v == "" {
		return fallback
	}
	return v
}
