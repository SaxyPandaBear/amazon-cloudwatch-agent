// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT

package windows_events

import (
	"github.com/aws/private-amazon-cloudwatch-agent-staging/translator"
	"github.com/aws/private-amazon-cloudwatch-agent-staging/translator/jsonconfig/mergeJsonRule"
	"github.com/aws/private-amazon-cloudwatch-agent-staging/translator/jsonconfig/mergeJsonUtil"
	parent "github.com/aws/private-amazon-cloudwatch-agent-staging/translator/translate/logs/logs_collected"
)

var ChildRule = map[string]translator.Rule{}

type WindowsEvent struct {
}

const SectionKey = "windows_events"

func GetCurPath() string {
	return parent.GetCurPath() + SectionKey + "/"
}

func RegisterRule(ruleName string, r translator.Rule) {
	ChildRule[ruleName] = r
}

func (w *WindowsEvent) ApplyRule(input interface{}) (returnKey string, returnVal interface{}) {
	im := input.(map[string]interface{})
	windowsEventConfig := map[string]interface{}{
		"destination": "cloudwatchlogs",
	}

	if _, ok := im[SectionKey]; ok {
		for _, rule := range ChildRule {
			key, val := rule.ApplyRule(im[SectionKey])
			if key != "" {
				windowsEventConfig[key] = val
			}
		}

		return "inputs", map[string]interface{}{
			"windows_event_log": []interface{}{windowsEventConfig},
		}
	} else {
		translator.AddInfoMessages("", "No windows event log configuration found.")
		return "", '"'
	}
}

var MergeRuleMap = map[string]mergeJsonRule.MergeRule{}

func (w *WindowsEvent) Merge(source map[string]interface{}, result map[string]interface{}) {
	mergeJsonUtil.MergeMap(source, result, SectionKey, MergeRuleMap, GetCurPath())
}

func init() {
	obj := new(WindowsEvent)
	parent.RegisterWindowsRule(SectionKey, obj)
	parent.MergeRuleMap[SectionKey] = obj
}
