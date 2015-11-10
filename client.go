/*
 * chronos-client
 * Copyright (c) 2015 Yieldbot, Inc. (http://github.com/yieldbot/chronos-client)
 * For the full copyright and license information, please view the LICENSE.txt file.
 */

// Package client provides Chronos operations
package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Client represents the Chronos client interface
type Client struct {
	URL string
}

// Jobs returns the Chronos jobs
func (cl Client) Jobs() ([]Job, error) {

	// Get jobs
	req, err := http.NewRequest("GET", cl.URL+"/scheduler/jobs", nil)
	res, err := cl.doRequest(req)
	if err != nil {
		return nil, errors.New("failed to fetch jobs due to " + err.Error())
	}

	// Parse jobs
	var jobs []Job
	if err = json.Unmarshal(res, &jobs); err != nil {
		return nil, errors.New("failed to unmarshal JSON data due to " + err.Error())
	}

	return jobs, nil
}

// PrintJobs prints the Chronos jobs
func (cl Client) PrintJobs(pretty bool) error {

	// Get jobs
	jobs, err := cl.Jobs()
	if err != nil {
		return err
	}

	// Parse jobs
	var buf []byte

	// If pretty is true then
	if pretty {
		buf, err = json.MarshalIndent(jobs, "", "  ")
	} else {
		// Otherwise just parse it
		buf, err = json.Marshal(jobs)
	}

	if err != nil {
		return err
	}

	fmt.Printf("%s", buf)

	return nil
}

// AddJob adds a Chronos job
func (cl Client) AddJob(jsonContent string) (bool, error) {

	// Check job
	buf := []byte(jsonContent)
	var job Job
	if err := json.Unmarshal(buf, &job); err != nil {
		return false, errors.New("failed to unmarshal JSON data due to " + err.Error())
	}

	// Add job
	req, err := http.NewRequest("POST", cl.URL+"/scheduler/iso8601", bytes.NewBuffer(buf))
	req.Header.Set("Content-Type", "application/json")
	_, err = cl.doRequest(req)
	if err != nil {
		return false, errors.New("failed to add job due to " + err.Error())
	}

	return true, nil
}

// AddDepJob adds a Chronos dependent job
func (cl Client) AddDepJob(jsonContent string) (bool, error) {

	// Check job
	buf := []byte(jsonContent)
	var job Job
	if err := json.Unmarshal(buf, &job); err != nil {
		return false, errors.New("failed to unmarshal JSON data due to " + err.Error())
	}

	// Add job
	req, err := http.NewRequest("POST", cl.URL+"/scheduler/dependency", bytes.NewBuffer(buf))
	req.Header.Set("Content-Type", "application/json")
	_, err = cl.doRequest(req)
	if err != nil {
		return false, errors.New("failed to add dependent job due to " + err.Error())
	}

	return true, nil
}

// RunJob runs a Chronos job
func (cl Client) RunJob(jobName, args string) (bool, error) {

	// Check job
	if jobName == "" {
		return false, errors.New("invalid job name")
	}

	query := jobName
	if args != "" {
		query += fmt.Sprintf("?arguments=%s", args)
	}

	// Run job
	req, err := http.NewRequest("PUT", cl.URL+"/scheduler/job/"+query, nil)
	res, err := cl.doRequest(req)
	if bytes.Index(res, []byte("not found")) != -1 {
		return true, errors.New(jobName + " job couldn't be found")
	} else if err != nil {
		return false, errors.New("failed to run job due to " + err.Error())
	}

	return true, nil
}

// DeleteJob deletes a Chronos job
func (cl Client) DeleteJob(jobName string) (bool, error) {

	// Check job
	if jobName == "" {
		return false, errors.New("invalid job name")
	}

	// Delete job
	req, err := http.NewRequest("DELETE", cl.URL+"/scheduler/job/"+jobName, nil)
	res, err := cl.doRequest(req)
	if err != nil {
		return false, errors.New("failed to delete job due to " + err.Error())
	} else if bytes.Index(res, []byte("not found")) != -1 {
		//if strings.Index(string(res), "not found") != -1 {
		return true, errors.New(jobName + " job couldn't be found")
	}

	return true, nil
}

// KillJobTasks kills the Chronos job tasks
func (cl Client) KillJobTasks(jobName string) (bool, error) {

	// Check job
	if jobName == "" {
		return false, errors.New("invalid job name")
	}

	// Kill job tasks
	req, err := http.NewRequest("DELETE", cl.URL+"/scheduler/task/kill/"+jobName, nil)
	_, err = cl.doRequest(req)
	if err != nil && strings.Index(err.Error(), "bad response") != -1 {
		return true, errors.New(jobName + " job couldn't be found")
	} else if err != nil {
		return false, errors.New("failed to kill tasks due to " + err.Error())
	}

	return true, nil
}

// UpdateJobTaskProgress updates a Chronos job task progress
func (cl Client) UpdateJobTaskProgress(jobName, taskID, jsonContent string) (bool, error) {

	// Check job and task
	if jobName == "" {
		return false, errors.New("invalid job name")
	}
	if taskID == "" {
		return false, errors.New("invalid task id")
	}

	// Update job task progress
	req, err := http.NewRequest("POST", cl.URL+"/scheduler/job/"+jobName+"/task/"+taskID+"/progress", bytes.NewBuffer([]byte(jsonContent)))
	req.Header.Set("Content-Type", "application/json")
	_, err = cl.doRequest(req)
	if err != nil {
		return false, errors.New("failed to update job task progress due to " + err.Error())
	}

	return true, nil
}

// DepGraph returns the Chronos dependency Graph in the dotfile format
func (cl Client) DepGraph() (string, error) {

	// Get the graph
	req, err := http.NewRequest("GET", cl.URL+"/scheduler/graph/dot", nil)
	res, err := cl.doRequest(req)
	if err != nil {
		return "", errors.New("failed to fetch graph due to " + err.Error())
	}

	return string(res), nil
}

// doRequest makes a request to Chronos REST API
func (cl Client) doRequest(req *http.Request) ([]byte, error) {

	// Init a client
	client := &http.Client{}

	// Do request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read data
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return data, errors.New("bad response: " + fmt.Sprintf("%d", resp.StatusCode))
	}

	return data, nil
}
