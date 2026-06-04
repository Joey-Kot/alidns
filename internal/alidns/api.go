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
)

type DNSAPI interface {
	AddDomainRecord(ctx context.Context, req *alidns20150109.AddDomainRecordRequest) (*alidns20150109.AddDomainRecordResponseBody, error)
	DeleteSubDomainRecords(ctx context.Context, req *alidns20150109.DeleteSubDomainRecordsRequest) (*alidns20150109.DeleteSubDomainRecordsResponseBody, error)
	DescribeDomainRecords(ctx context.Context, req *alidns20150109.DescribeDomainRecordsRequest) ([]*alidns20150109.DescribeDomainRecordsResponseBodyDomainRecordsRecord, error)
	UpdateDomainRecord(ctx context.Context, req *alidns20150109.UpdateDomainRecordRequest) (*alidns20150109.UpdateDomainRecordResponseBody, error)
}
