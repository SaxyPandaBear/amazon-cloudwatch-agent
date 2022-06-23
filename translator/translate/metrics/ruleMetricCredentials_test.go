// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT

package metrics

import (
	"encoding/json"
	"testing"

	"github.com/aws/private-amazon-cloudwatch-agent-staging/translator/context"
	"github.com/aws/private-amazon-cloudwatch-agent-staging/translator/translate/agent"
	"github.com/stretchr/testify/assert"
)

func TestWithAgentConfig(t *testing.T) {
	agent.Global_Config.Credentials = map[string]interface{}{}
	ctx := context.CurrentContext()
	ctx.SetCredentials(map[string]string{})
	c := new(MetricsCreds)
	var input interface{}
	err := json.Unmarshal([]byte(`{ "credentials" : {"access_key":"metric_ak", "secret_key":"metric_sk", "token": "dummy_token", "profile": "dummy_profile", "role_arn": "role_value"}}`), &input)
	if err == nil {
		_, returnVal := c.ApplyRule(input)
		assert.Equal(t, "role_value", returnVal.(map[string]interface{})["role_arn"], "Expected to be equal")
	} else {
		panic(err)
	}

	agent.Global_Config.Role_arn = "global_role_arn_test"
	err = json.Unmarshal([]byte(`{ "credentials" : {"access_key":"metric_ak", "secret_key":"metric_sk", "token": "dummy_token", "profile": "dummy_profile", "role_arn": "role_value"}}`), &input)
	if err == nil {
		_, returnVal := c.ApplyRule(input)
		assert.Equal(t, "role_value", returnVal.(map[string]interface{})["role_arn"], "Expected to be equal")
	} else {
		panic(err)
	}

	agent.Global_Config.Role_arn = "global_role_arn_test"
	err = json.Unmarshal([]byte(`{ "credentials" : {"access_key":"metric_ak", "secret_key":"metric_sk", "token": "dummy_token", "profile": "dummy_profile"}}`), &input)
	if err == nil {
		_, returnVal := c.ApplyRule(input)
		assert.Equal(t, "global_role_arn_test", returnVal.(map[string]interface{})["role_arn"], "Expected to be equal")
	} else {
		panic(err)
	}
}
