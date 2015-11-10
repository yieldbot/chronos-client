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

// Jobs returns Chronos jobs
func (cl Client) Jobs() ([]Job, error) {

	// Get jobs
	res, err := cl.request("GET", "/scheduler/jobs")
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

// PrintJobs prints jobs
func (cl Client) PrintJobs(pretty bool) error {

	// Get jobs
	jobs, err := cl.Jobs()
	if err != nil {
		return err
	}

	// Parse jobs
	var jsonb []byte

	// If pretty is true then
	if pretty {
		jsonb, err = json.MarshalIndent(jobs, "", "  ")
	} else {
		// Otherwise just parse it
		jsonb, err = json.Marshal(jobs)
	}

	if err != nil {
		return err
	}

	fmt.Printf("%s", jsonb)

	return nil
}

// DeleteJob deletes a Chronos job by the given job name
func (cl Client) DeleteJob(jobName string) (bool, error) {

	// Check job
	if jobName == "" {
		return false, errors.New("invalid job name")
	}

	// Delete job
	res, err := cl.request("DELETE", "/scheduler/job/"+jobName)
	if err != nil {
		return false, errors.New("failed to delete job due to " + err.Error())
	} else if bytes.Index(res, []byte("not found")) != -1 {
		//if strings.Index(string(res), "not found") != -1 {
		return true, errors.New(jobName + " job couldn't be found")
	}

	return true, nil
}

// KillTasks kills Chronos job tasks by the given job name
func (cl Client) KillTasks(jobName string) (bool, error) {

	// Check job
	if jobName == "" {
		return false, errors.New("invalid job name")
	}

	// Delete job
	_, err := cl.request("DELETE", "/scheduler/task/kill/"+jobName)

	if err != nil && strings.Index(err.Error(), "bad response") != -1 {
		return true, errors.New(jobName + " job couldn't be found")
	} else if err != nil {
		return false, errors.New("failed to kill tasks due to " + err.Error())
	}

	return true, nil
}

// request makes a request to the given endpoint
func (cl Client) request(verb string, endpoint string) ([]byte, error) {

	// Init a client
	client := &http.Client{}

	// Do request
	req, err := http.NewRequest(verb, cl.URL+endpoint, nil)
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
