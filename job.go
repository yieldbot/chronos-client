/*
 * chronos-client
 * Copyright (c) 2015 Yieldbot, Inc. (http://github.com/yieldbot/chronos-client)
 * For the full copyright and license information, please view the LICENSE.txt file.
 */

package client

// Job represents Chronos job configuration
type Job struct {

	// REF: https://mesos.github.io/chronos/docs/api.html#job-configuration

	Name                 string        `json:"name"`
	Description          string        `json:"description"`
	Command              string        `json:"command"`
	Arguments            []string      `json:"arguments"`
	Shell                bool          `json:"shell"`
	Epsilon              string        `json:"epsilon"`
	Executor             string        `json:"executor"`
	ExecutorFlags        string        `json:"executorFlags"`
	Retries              int           `json:"retries"`
	Owner                string        `json:"owner"`
	OwnerName            string        `json:"ownerName"`
	Async                bool          `json:"async"`
	SuccessCount         int           `json:"successCount"`
	ErrorCount           int           `json:"errorCount"`
	LastSuccess          string        `json:"lastSuccess"`
	LastError            string        `json:"lastError"`
	Cpus                 float64       `json:"cpus"`
	Mem                  float64       `json:"mem"`
	Disk                 float64       `json:"disk"`
	Disabled             bool          `json:"disabled"`
	Uris                 []string      `json:"uris"`
	Schedule             string        `json:"schedule"`
	ScheduleTimeZone     string        `json:"scheduleTimeZone"`
	Parents              []string      `json:"parents"`
	RunAsUser            string        `json:"runAsUser"`
	Container            *JobContainer `json:"container"`
	DataJob              bool          `json:"dataJob"`
	EnvironmentVariables []*EnvVars    `json:"environmentVariables"`
	Constraints          [][]string    `json:"constraints"`
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
