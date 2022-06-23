// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT

package linux

import (
	"testing"

	"github.com/aws/private-amazon-cloudwatch-agent-staging/tool/runtime"

	"github.com/stretchr/testify/assert"
)

func TestDisk_ToMap(t *testing.T) {
	expectedKey := "disk"
	expectedValue := map[string]interface{}{"resources": []string{"*"}, "measurement": []string{"used_percent", "inodes_free"}}
	ctx := &runtime.Context{}
	conf := new(Disk)
	conf.Enable()
	key, value := conf.ToMap(ctx)
	assert.Equal(t, expectedKey, key)
	assert.Equal(t, expectedValue, value)
}
