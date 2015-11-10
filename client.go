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

// RunJob runs a Chronos job by the given job name
func (cl Client) RunJob(name, args string) (bool, error) {

	// Check job
	if name == "" {
		return false, errors.New("invalid job name")
	}

	query := name
	if args != "" {
		query += fmt.Sprintf("?arguments=%s", args)
	}

	// Run job
	req, err := http.NewRequest("PUT", cl.URL+"/scheduler/job/"+query, nil)
	res, err := cl.doRequest(req)
	if bytes.Index(res, []byte("not found")) != -1 {
		return true, errors.New(name + " job couldn't be found")
	} else if err != nil {
		return false, errors.New("failed to run job due to " + err.Error())
	}

	return true, nil
}

// DeleteJob deletes a Chronos job by the given job name
func (cl Client) DeleteJob(name string) (bool, error) {

	// Check job
	if name == "" {
		return false, errors.New("invalid job name")
	}

	// Delete job
	req, err := http.NewRequest("DELETE", cl.URL+"/scheduler/job/"+name, nil)
	res, err := cl.doRequest(req)
	if err != nil {
		return false, errors.New("failed to delete job due to " + err.Error())
	} else if bytes.Index(res, []byte("not found")) != -1 {
		//if strings.Index(string(res), "not found") != -1 {
		return true, errors.New(name + " job couldn't be found")
	}

	return true, nil
}

// KillTasks kills Chronos job tasks by the given job name
func (cl Client) KillTasks(name string) (bool, error) {

	// Check job
	if name == "" {
		return false, errors.New("invalid job name")
	}

	// Kill job tasks
	req, err := http.NewRequest("DELETE", cl.URL+"/scheduler/task/kill/"+name, nil)
	_, err = cl.doRequest(req)
	if err != nil && strings.Index(err.Error(), "bad response") != -1 {
		return true, errors.New(name + " job couldn't be found")
	} else if err != nil {
		return false, errors.New("failed to kill tasks due to " + err.Error())
	}

	return true, nil
}

// doRequest makes a request to Chronos by the given request
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
