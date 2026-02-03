package alidns

import (
	"context"

	alidns20150109 "github.com/alibabacloud-go/alidns-20150109/v4/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
)

type sdkClient struct {
	client *alidns20150109.Client
}

func NewSDKClient(client *alidns20150109.Client) DNSAPI {
	return &sdkClient{client: client}
}

func (s *sdkClient) AddDomainRecord(_ context.Context, req *alidns20150109.AddDomainRecordRequest) (*alidns20150109.AddDomainRecordResponseBody, error) {
	resp, err := s.client.AddDomainRecordWithOptions(req, &util.RuntimeOptions{})
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, nil
	}
	return resp.Body, nil
}

func (s *sdkClient) DeleteSubDomainRecords(_ context.Context, req *alidns20150109.DeleteSubDomainRecordsRequest) (*alidns20150109.DeleteSubDomainRecordsResponseBody, error) {
	resp, err := s.client.DeleteSubDomainRecordsWithOptions(req, &util.RuntimeOptions{})
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, nil
	}
	return resp.Body, nil
}

func (s *sdkClient) DescribeDomainRecords(_ context.Context, req *alidns20150109.DescribeDomainRecordsRequest) ([]*alidns20150109.DescribeDomainRecordsResponseBodyDomainRecordsRecord, error) {
	resp, err := s.client.DescribeDomainRecordsWithOptions(req, &util.RuntimeOptions{})
	if err != nil {
		return nil, err
	}
	if resp == nil || resp.Body == nil || resp.Body.DomainRecords == nil || resp.Body.DomainRecords.Record == nil {
		return []*alidns20150109.DescribeDomainRecordsResponseBodyDomainRecordsRecord{}, nil
	}
	return resp.Body.DomainRecords.Record, nil
}

func (s *sdkClient) UpdateDomainRecord(_ context.Context, req *alidns20150109.UpdateDomainRecordRequest) (*alidns20150109.UpdateDomainRecordResponseBody, error) {
	resp, err := s.client.UpdateDomainRecordWithOptions(req, &util.RuntimeOptions{})
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, nil
	}
	return resp.Body, nil
}
