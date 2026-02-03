package alidns

import (
	"fmt"

	alidns20150109 "github.com/alibabacloud-go/alidns-20150109/v4/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/credentials-go/credentials"
)

func CreateClient(accessKeyID, accessKeySecret string) (*alidns20150109.Client, error) {
	if accessKeyID == "" {
		return nil, fmt.Errorf("AccessKeyId is required")
	}
	if accessKeySecret == "" {
		return nil, fmt.Errorf("AccessKeySecret is required")
	}

	config := new(credentials.Config).
		SetType("access_key").
		SetAccessKeyId(accessKeyID).
		SetAccessKeySecret(accessKeySecret)

	akCredential, err := credentials.NewCredential(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create credential using config: %w", err)
	}

	openapiConfig := &openapi.Config{
		Credential: akCredential,
		Endpoint:   tea.String("alidns.cn-hangzhou.aliyuncs.com"),
	}

	client, err := alidns20150109.NewClient(openapiConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create alidns client: %w", err)
	}

	return client, nil
}
