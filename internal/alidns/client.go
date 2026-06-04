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
