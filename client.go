/*
 * chronos-client
 * Copyright (c) 2015 Yieldbot, Inc. (http://github.com/yieldbot/chronos-client)
 * For the full copyright and license information, please view the LICENSE.txt file.
 */

// client provides Chronos operations
package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
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
	return nil, nil
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

// DeleteJob deletes a Chronos job
func (cl Client) DeleteJob(jobName string) (bool, error) {

	// Check job
	if jobName == "" {
		return false, errors.New("invalid job name")
	}

	// Delete job
	res, err := cl.request("DELETE", "/scheduler/job/"+jobName)
	if err != nil {
		return false, errors.New("failed to delete job due to " + err.Error())
	}

	//if strings.Index(string(res), "not found") != -1 {
	if bytes.Index(res, []byte("not found")) != -1 {
		return true, errors.New(jobName + " job couldn't be found")
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

	return data, nil
}
