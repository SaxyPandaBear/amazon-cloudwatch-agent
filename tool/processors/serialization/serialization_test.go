// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT

package serialization

import (
	"testing"

	"github.com/aws/private-amazon-cloudwatch-agent-staging/tool/data"
	"github.com/aws/private-amazon-cloudwatch-agent-staging/tool/processors/ssm"
	"github.com/aws/private-amazon-cloudwatch-agent-staging/tool/runtime"

	"github.com/stretchr/testify/assert"
)

func TestProcessor_Process(t *testing.T) {
	ctx := new(runtime.Context)
	conf := new(data.Config)

	Processor.Process(ctx, conf)
	assert.Equal(t, new(runtime.Context), ctx)
	assert.Equal(t, new(data.Config), conf)
}

func TestProcessor_NextProcessor(t *testing.T) {
	nextProcessor := Processor.NextProcessor(nil, nil)
	assert.Equal(t, ssm.Processor, nextProcessor)
}
