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
