// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT

package linux

import (
	"testing"

	"github.com/aws/private-amazon-cloudwatch-agent-staging/tool/runtime"

	"github.com/stretchr/testify/assert"
)

func TestDiskIO_ToMap(t *testing.T) {
	expectedKey := "diskio"
	expectedValue := map[string]interface{}{"resources": []string{"*"}, "measurement": []string{"io_time", "write_bytes", "read_bytes", "writes", "reads"}}
	ctx := &runtime.Context{}
	conf := new(DiskIO)
	conf.Enable()
	key, value := conf.ToMap(ctx)
	assert.Equal(t, expectedKey, key)
	assert.Equal(t, expectedValue, value)
}
