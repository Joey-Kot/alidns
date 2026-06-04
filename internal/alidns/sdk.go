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
