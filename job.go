/*
 * chronos-client
 * Copyright (c) 2015 Yieldbot, Inc.
 * For the full copyright and license information, please view the LICENSE.txt file.
 */

package client

// Job represents Chronos job configuration
type Job struct {

	// REF: https://mesos.github.io/chronos/docs/api.html#job-configuration

	Name                 string        `json:"name"`
	Description          string        `json:"description"`
	Command              string        `json:"command,omitempty"`
	Arguments            []string      `json:"arguments,omitempty"`
	Shell                bool          `json:"shell"`
	Epsilon              string        `json:"epsilon,omitempty"`
	Executor             string        `json:"executor,omitempty"`
	ExecutorFlags        string        `json:"executorFlags,omitempty"`
	Retries              int           `json:"retries,omitempty"`
	Owner                string        `json:"owner,omitempty"`
	OwnerName            string        `json:"ownerName,omitempty"`
	Async                bool          `json:"async,omitempty"`
	SuccessCount         int           `json:"successCount,omitempty"`
	ErrorCount           int           `json:"errorCount,omitempty"`
	LastSuccess          string        `json:"lastSuccess,omitempty"`
	LastError            string        `json:"lastError,omitempty"`
	Cpus                 float64       `json:"cpus"`
	Mem                  float64       `json:"mem"`
	Disk                 float64       `json:"disk,omitempty"`
	Disabled             bool          `json:"disabled"`
	Uris                 []string      `json:"uris"`
	Schedule             string        `json:"schedule"`
	ScheduleTimeZone     string        `json:"scheduleTimeZone,omitempty"`
	Parents              []string      `json:"parents,omitempty"`
	RunAsUser            string        `json:"runAsUser,omitempty"`
	Container            *JobContainer `json:"container"`
	DataJob              bool          `json:"dataJob"`
	EnvironmentVariables []*EnvVars    `json:"environmentVariables"`
	Constraints          [][]string    `json:"constraints,omitempty"`
}

// JobContainer represents Chronos job container field
type JobContainer struct {
	Type           string        `json:"type"`
	Image          string        `json:"image"`
	ForcePullImage bool          `json:"forcePullImage"`
	Network        string        `json:"network"`
	Volumes        []interface{} `json:"arguments"`
}

// EnvVars represents Chronos job environmentVariables field
type EnvVars struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
