// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT

//go:build !windows

package keu

import (
	"time"

	"github.com/aws/amazon-cloudwatch-agent-test/environment"
	. "github.com/aws/amazon-cloudwatch-agent-test/test/awsneuron/resources"
	"github.com/aws/amazon-cloudwatch-agent-test/test/metric"
	"github.com/aws/amazon-cloudwatch-agent-test/test/status"
	"github.com/aws/amazon-cloudwatch-agent-test/test/test_runner"
)

const (
	awsNeuronMetricIndicator = "_neuron"
)

var expectedDimsToMetrics = map[string][]string{
    "ClusterName-ClusterQueue-Status": {
        KueuePendingWorkflows
    },
    "ClusterName-ClusterQueue": {
        KueuePendingWorkflows,
        KueueEvictedWorkflowsTotal,
        KueueAdmittedActiveWorkloads,
    },
    "ClusterName-Status": {
        KueuePendingWorkflows,
    },
    "ClusterName": {
        KueuePendingWorkflows,
        KueueEvictedWorkflowsTotal,
        KueueAdmittedActiveWorkloads,
        KueueClusterQueueResourceUsage,
        KueueClusterQueueNominalUsage,
        KueueClusterQueueBorrowingLimit
    },
    "ClusterName-Reason": {
        KueueEvictedWorkflowsTotal,
    },
    "ClusterName-ClusterQueue-Reason": {
        KueueEvictedWorkflowsTotal,
    },
    "ClusterName-ClusterQueue-Reason": {
        KueueEvictedWorkflowsTotal,
    },
    "ClusterName-ClusterQueue-Resource-Flavor": {
        KueueClusterQueueResourceUsage,
        KueueClusterQueueNominalUsage,
        KueueClusterQueueBorrowingLimit
    },
    "ClusterName-ClusterQueue-Resource": {
        KueueClusterQueueResourceUsage,
        KueueClusterQueueNominalUsage,
        KueueClusterQueueBorrowingLimit
    },
    "ClusterName-ClusterQueue-Flavor": {
        KueueClusterQueueResourceUsage,
        KueueClusterQueueNominalUsage,
        KueueClusterQueueBorrowingLimit
    },
}

type KueueTestRunner struct {
	test_runner.BaseTestRunner
	testName string
	env      *environment.MetaData
}

var _ test_runner.ITestRunner = (*KueueTestRunner)(nil)

func (t *KueueTestRunner) Validate() status.TestGroupResult {
	var testResults []status.TestResult
	testResults = append(testResults, metric.ValidateMetrics(t.env, awsNeuronMetricIndicator, expectedDimsToMetrics)...)
	testResults = append(testResults, metric.ValidateLogs(t.env))
	testResults = append(testResults, metric.ValidateLogsFrequency(t.env))
	return status.TestGroupResult{
		Name:        t.GetTestName(),
		TestResults: testResults,
	}
}

func (t *AwsNeuronTestRunner) GetTestName() string {
	return t.testName
}

func (t *AwsNeuronTestRunner) GetAgentConfigFileName() string {
	return ""
}

func (t *AwsNeuronTestRunner) GetAgentRunDuration() time.Duration {
	return 25 * time.Minute
}

func (t *AwsNeuronTestRunner) GetMeasuredMetrics() []string {
	return nil
}
