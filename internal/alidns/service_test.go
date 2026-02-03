package alidns

import (
	"context"
	"testing"

	alidns20150109 "github.com/alibabacloud-go/alidns-20150109/v4/client"
	"github.com/alibabacloud-go/tea/tea"
)

type fakeAPI struct {
	addReq    *alidns20150109.AddDomainRecordRequest
	delReq    *alidns20150109.DeleteSubDomainRecordsRequest
	queryReq  *alidns20150109.DescribeDomainRecordsRequest
	updateReq *alidns20150109.UpdateDomainRecordRequest

	addResp    *alidns20150109.AddDomainRecordResponseBody
	delResp    *alidns20150109.DeleteSubDomainRecordsResponseBody
	queryResp  []*alidns20150109.DescribeDomainRecordsResponseBodyDomainRecordsRecord
	updateResp *alidns20150109.UpdateDomainRecordResponseBody
}

func (f *fakeAPI) AddDomainRecord(_ context.Context, req *alidns20150109.AddDomainRecordRequest) (*alidns20150109.AddDomainRecordResponseBody, error) {
	f.addReq = req
	return f.addResp, nil
}

func (f *fakeAPI) DeleteSubDomainRecords(_ context.Context, req *alidns20150109.DeleteSubDomainRecordsRequest) (*alidns20150109.DeleteSubDomainRecordsResponseBody, error) {
	f.delReq = req
	return f.delResp, nil
}

func (f *fakeAPI) DescribeDomainRecords(_ context.Context, req *alidns20150109.DescribeDomainRecordsRequest) ([]*alidns20150109.DescribeDomainRecordsResponseBodyDomainRecordsRecord, error) {
	f.queryReq = req
	return f.queryResp, nil
}

func (f *fakeAPI) UpdateDomainRecord(_ context.Context, req *alidns20150109.UpdateDomainRecordRequest) (*alidns20150109.UpdateDomainRecordResponseBody, error) {
	f.updateReq = req
	return f.updateResp, nil
}

func TestServiceAddBuildsRequest(t *testing.T) {
	api := &fakeAPI{addResp: &alidns20150109.AddDomainRecordResponseBody{RecordId: tea.String("r-1")}}
	svc := NewService(api)

	resp, err := svc.Add(context.Background(), AddInput{
		DomainName: "example.com",
		Name:       "www",
		Type:       "A",
		Value:      "1.2.3.4",
	})
	if err != nil {
		t.Fatalf("Add returned error: %v", err)
	}
	if tea.StringValue(resp.RecordId) != "r-1" {
		t.Fatalf("unexpected record id: %s", tea.StringValue(resp.RecordId))
	}

	req := api.addReq
	if tea.StringValue(req.DomainName) != "example.com" || tea.StringValue(req.RR) != "www" || tea.StringValue(req.Type) != "A" || tea.StringValue(req.Value) != "1.2.3.4" {
		t.Fatalf("unexpected add request core fields: %+v", req)
	}
	if tea.Int64Value(req.TTL) != 600 || tea.Int64Value(req.Priority) != 1 || tea.StringValue(req.Line) != "default" || tea.StringValue(req.Lang) != "en" {
		t.Fatalf("unexpected add request defaults: %+v", req)
	}
}

func TestServiceDelBuildsRequest(t *testing.T) {
	api := &fakeAPI{}
	svc := NewService(api)

	_, err := svc.Del(context.Background(), DelInput{DomainName: "example.com", Name: "www", Type: "A"})
	if err != nil {
		t.Fatalf("Del returned error: %v", err)
	}

	req := api.delReq
	if tea.StringValue(req.DomainName) != "example.com" || tea.StringValue(req.RR) != "www" || tea.StringValue(req.Type) != "A" || tea.StringValue(req.Lang) != "en" {
		t.Fatalf("unexpected del request: %+v", req)
	}
}

func TestServiceQueryBuildsRequestAndReturnsRecords(t *testing.T) {
	api := &fakeAPI{queryResp: []*alidns20150109.DescribeDomainRecordsResponseBodyDomainRecordsRecord{{RecordId: tea.String("r-2")}}}
	svc := NewService(api)

	records, err := svc.Query(context.Background(), QueryInput{DomainName: "example.com"})
	if err != nil {
		t.Fatalf("Query returned error: %v", err)
	}
	if len(records) != 1 || tea.StringValue(records[0].RecordId) != "r-2" {
		t.Fatalf("unexpected query records: %+v", records)
	}

	req := api.queryReq
	if tea.StringValue(req.DomainName) != "example.com" || tea.StringValue(req.Lang) != "en" || tea.StringValue(req.Direction) != "ASC" || tea.StringValue(req.Status) != "Enable" || tea.Int64Value(req.PageNumber) != 1 || tea.Int64Value(req.PageSize) != 500 || tea.StringValue(req.SearchMode) != "LIKE" {
		t.Fatalf("unexpected query request: %+v", req)
	}
}

func TestServiceUpdateBuildsRequest(t *testing.T) {
	api := &fakeAPI{}
	svc := NewService(api)

	_, err := svc.Update(context.Background(), UpdateInput{
		RecordID: "r-3",
		Name:     "www",
		Type:     "A",
		Value:    "1.2.3.5",
	})
	if err != nil {
		t.Fatalf("Update returned error: %v", err)
	}

	req := api.updateReq
	if tea.StringValue(req.RecordId) != "r-3" || tea.StringValue(req.RR) != "www" || tea.StringValue(req.Type) != "A" || tea.StringValue(req.Value) != "1.2.3.5" {
		t.Fatalf("unexpected update request core fields: %+v", req)
	}
	if tea.Int64Value(req.TTL) != 600 || tea.Int64Value(req.Priority) != 1 || tea.StringValue(req.Line) != "default" || tea.StringValue(req.Lang) != "en" {
		t.Fatalf("unexpected update request defaults: %+v", req)
	}
}
