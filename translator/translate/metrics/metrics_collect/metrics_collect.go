// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT

package metrics_collect

import (
	"sort"

	"github.com/aws/private-amazon-cloudwatch-agent-staging/translator"
	"github.com/aws/private-amazon-cloudwatch-agent-staging/translator/config"
	"github.com/aws/private-amazon-cloudwatch-agent-staging/translator/jsonconfig/mergeJsonRule"
	"github.com/aws/private-amazon-cloudwatch-agent-staging/translator/jsonconfig/mergeJsonUtil"
	parent "github.com/aws/private-amazon-cloudwatch-agent-staging/translator/translate/metrics"
)

type Rule translator.Rule

var ChildRule = map[string]Rule{}

var (
	windowsMetricCollectRule = map[string]Rule{}
	linuxMetricCollectRule   = map[string]Rule{}
	darwinMetricCollectRule  = map[string]Rule{}
)

const SectionKey = "metrics_collected"
const ObjectKey = "object"

func GetCurPath() string {
	curPath := parent.GetCurPath() + SectionKey + "/"
	return curPath
}

func RegisterLinuxRule(ruleName string, r Rule) {
	linuxMetricCollectRule[ruleName] = r
}

func RegisterDarwinRule(ruleName string, r Rule) {
	darwinMetricCollectRule[ruleName] = r
}

func RegisterWindowsRule(ruleName string, r Rule) {
	windowsMetricCollectRule[ruleName] = r
}

type CollectMetrics struct {
}

func (c *CollectMetrics) ApplyRule(input interface{}) (returnKey string, returnVal interface{}) {
	im := input.(map[string]interface{})
	result := map[string]interface{}{}

	var targetRuleMap map[string]Rule
	switch translator.GetTargetPlatform() {
	case config.OS_TYPE_LINUX:
		targetRuleMap = linuxMetricCollectRule
	case config.OS_TYPE_DARWIN:
		targetRuleMap = darwinMetricCollectRule
	case config.OS_TYPE_WINDOWS:
		targetRuleMap = windowsMetricCollectRule
	default:
		// NOTE: we should panic but now there are many unit tests not setting the global var
		targetRuleMap = linuxMetricCollectRule
		//panic("unknown target platform " + translator.GetTargetPlatform())
	}

	//Check if this plugin exist in the input instance
	//If not, not process
	if _, ok := im[SectionKey]; !ok {
		returnKey = ""
		returnVal = ""
	} else {
		//If yes, process it
		for _, rule := range getOrderedRules(targetRuleMap) {
			key, val := rule.ApplyRule(im[SectionKey])

			//If key == "", then no instance of this class in input
			if key != "" {
				result[key] = val
			}
		}
	}
	returnKey = "inputs"
	returnVal = result
	return
}

// Adding alphabet order to the Rules
func getOrderedRules(ruleMap map[string]Rule) []Rule {
	var orderedRules []Rule
	var orderedRuleNames []string
	for ruleName := range ruleMap {
		orderedRuleNames = append(orderedRuleNames, ruleName)
	}
	sort.Strings(orderedRuleNames)
	for _, orderedRuleName := range orderedRuleNames {
		orderedRules = append(orderedRules, ruleMap[orderedRuleName])
	}
	return orderedRules
}

var MergeRuleMap = map[string]mergeJsonRule.MergeRule{}

func (c *CollectMetrics) Merge(source map[string]interface{}, result map[string]interface{}) {
	mergeJsonUtil.MergeMap(source, result, SectionKey, MergeRuleMap, GetCurPath())
}

func init() {
	c := new(CollectMetrics)
	parent.RegisterRule(SectionKey, c)
	parent.MergeRuleMap[SectionKey] = c
}
