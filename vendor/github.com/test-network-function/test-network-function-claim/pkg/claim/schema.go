// Copyright (C) 2020 Red Hat, Inc.
//
// This program is free software; you can redistribute it and/or modify it under the terms of the GNU General Public
// License as published by the Free Software Foundation; either version 2 of the License, or (at your option) any later
// version.
//
// This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied
// warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License along with this program; if not, write to the Free
// Software Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301, USA.

//
// Code generated by `test-network-function-claim/cmd/generate/generate.go` on: 2021-10-20 10:45:48.026387017 -0500 CDT m=+0.000652323
//
// `https://github.com/a-h/generate` provides a generic set of interfaces to convert JSON schema into
// workable GoLang struct implementations.  However, the code generator is limited and does not allow
// type remapping.  By default, JSON Schema "object" types are remapped to custom struct definitions.
// This becomes a problem in our case, as we do not define certain facets such as "Hosts" or
// "LshwOutput".  This CLI driven generator augments the stock generator to allow overrides to generic
// "map[string]interface{}", which is capable of representing arbitrary JSON data.
//
// Warning:  Do not edit this file by hand.  Instead, use Makefile targets.
//

// Code generated by schema-generate. DO NOT EDIT.

package claim

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// Claim
type Claim struct {

	// Tests within test-network-function often require configuration.  For example, the generic test suite requires listing all CNF containers.  This information is used to derive per-container IP address information, which is then used as input to the connectivity test suite.  Test suites within test-network-function may use multiple configurations, but each with a unique name.
	Configurations map[string]interface{} `json:"configurations"`
	Metadata       *Metadata              `json:"metadata"`

	// An OpenShift cluster is composed of an arbitrary number of Nodes used for platform and application services.  Since a claim must be reproducible, a variety of per-Node information must be collected and stored in the claim.  Node names are unique within a given OpenShift cluster.
	Nodes map[string]interface{} `json:"nodes"`

	// The test-network-function test results.  Results are a JSON representation of the JUnit output.
	RawResults map[string]interface{} `json:"rawResults"`

	// The results for each unique test case.
	Results  map[string]interface{} `json:"results,omitempty"`
	Versions *Versions              `json:"versions"`
}

// Identifier identifier is a per testcase unique identifier.
type Identifier struct {

	// url stores the unique url for a test.
	Url string `json:"url"`

	// version stores the semantic version of the test.
	Version string `json:"version"`
}

// Metadata
type Metadata struct {

	// The UTC end time of a claim evaluation.  This is recorded when the test-network-function test suite completes.
	EndTime string `json:"endTime"`

	// The UTC start time of a claim evaluation.  This is recorded when the test-network-function test suite is invoked.
	StartTime string `json:"startTime"`
}

// Result result is the result of running a testcase.
type Result struct {

	// Ginkgo writer output during the test run.
	CapturedTestOutput string `json:"CapturedTestOutput"`

	// The duration of the test in nanoseconds.
	Duration int `json:"duration"`

	// The end time of the test.
	EndTime string `json:"endTime,omitempty"`

	// The content of the line where the failure happened
	FailureLineContent string `json:"failureLineContent"`

	// The Filename and line number where the failure happened
	FailureLocation string `json:"failureLocation"`

	// Describes the test failure in detail.
	FailureReason string `json:"failureReason"`

	// The start time of the test.
	StartTime string `json:"startTime"`

	// The test result state: INVALID SPEC STATE, pending,skipped,passed,failed,aborted,panicked,interrupted
	State string `json:"state"`

	// The test identifier
	TestID *Identifier `json:"testID,omitempty"`

	// The JUnit test text.
	TestText string `json:"testText"`
}

// Root A test-network-function claim is an attestation of the tests performed, the results and the various configurations.  Since a claim must be reproducible, it also includes an overview of the systems under test and their physical configurations.
type Root struct {
	Claim *Claim `json:"claim"`
}

// Versions
type Versions struct {

	// The Kubernetes release version.
	K8s string `json:"k8s,omitempty"`

	// The oc client release version.
	OcClient string `json:"ocClient,omitempty"`

	// OCP cluster release version.
	Ocp string `json:"ocp,omitempty"`

	// The test-network-function (tnf) release version.
	Tnf string `json:"tnf"`

	// The test-network-function (tnf) Git Commit.
	TnfGitCommit string `json:"tnfGitCommit,omitempty"`
}

func (strct *Claim) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Configurations" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "configurations" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"configurations\": ")
	if tmp, err := json.Marshal(strct.Configurations); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Metadata" field is required
	if strct.Metadata == nil {
		return nil, errors.New("metadata is a required field")
	}
	// Marshal the "metadata" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"metadata\": ")
	if tmp, err := json.Marshal(strct.Metadata); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Nodes" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "nodes" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"nodes\": ")
	if tmp, err := json.Marshal(strct.Nodes); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "RawResults" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "rawResults" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"rawResults\": ")
	if tmp, err := json.Marshal(strct.RawResults); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "results" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"results\": ")
	if tmp, err := json.Marshal(strct.Results); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Versions" field is required
	if strct.Versions == nil {
		return nil, errors.New("versions is a required field")
	}
	// Marshal the "versions" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"versions\": ")
	if tmp, err := json.Marshal(strct.Versions); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Claim) UnmarshalJSON(b []byte) error {
	configurationsReceived := false
	metadataReceived := false
	nodesReceived := false
	rawResultsReceived := false
	versionsReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "configurations":
			if err := json.Unmarshal([]byte(v), &strct.Configurations); err != nil {
				return err
			}
			configurationsReceived = true
		case "metadata":
			if err := json.Unmarshal([]byte(v), &strct.Metadata); err != nil {
				return err
			}
			metadataReceived = true
		case "nodes":
			if err := json.Unmarshal([]byte(v), &strct.Nodes); err != nil {
				return err
			}
			nodesReceived = true
		case "rawResults":
			if err := json.Unmarshal([]byte(v), &strct.RawResults); err != nil {
				return err
			}
			rawResultsReceived = true
		case "results":
			if err := json.Unmarshal([]byte(v), &strct.Results); err != nil {
				return err
			}
		case "versions":
			if err := json.Unmarshal([]byte(v), &strct.Versions); err != nil {
				return err
			}
			versionsReceived = true
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if configurations (a required property) was received
	if !configurationsReceived {
		return errors.New("\"configurations\" is required but was not present")
	}
	// check if metadata (a required property) was received
	if !metadataReceived {
		return errors.New("\"metadata\" is required but was not present")
	}
	// check if nodes (a required property) was received
	if !nodesReceived {
		return errors.New("\"nodes\" is required but was not present")
	}
	// check if rawResults (a required property) was received
	if !rawResultsReceived {
		return errors.New("\"rawResults\" is required but was not present")
	}
	// check if versions (a required property) was received
	if !versionsReceived {
		return errors.New("\"versions\" is required but was not present")
	}
	return nil
}

func (strct *Identifier) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Url" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "url" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"url\": ")
	if tmp, err := json.Marshal(strct.Url); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Version" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "version" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"version\": ")
	if tmp, err := json.Marshal(strct.Version); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Identifier) UnmarshalJSON(b []byte) error {
	urlReceived := false
	versionReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "url":
			if err := json.Unmarshal([]byte(v), &strct.Url); err != nil {
				return err
			}
			urlReceived = true
		case "version":
			if err := json.Unmarshal([]byte(v), &strct.Version); err != nil {
				return err
			}
			versionReceived = true
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if url (a required property) was received
	if !urlReceived {
		return errors.New("\"url\" is required but was not present")
	}
	// check if version (a required property) was received
	if !versionReceived {
		return errors.New("\"version\" is required but was not present")
	}
	return nil
}

func (strct *Metadata) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "EndTime" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "endTime" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"endTime\": ")
	if tmp, err := json.Marshal(strct.EndTime); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "StartTime" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "startTime" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"startTime\": ")
	if tmp, err := json.Marshal(strct.StartTime); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Metadata) UnmarshalJSON(b []byte) error {
	endTimeReceived := false
	startTimeReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "endTime":
			if err := json.Unmarshal([]byte(v), &strct.EndTime); err != nil {
				return err
			}
			endTimeReceived = true
		case "startTime":
			if err := json.Unmarshal([]byte(v), &strct.StartTime); err != nil {
				return err
			}
			startTimeReceived = true
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if endTime (a required property) was received
	if !endTimeReceived {
		return errors.New("\"endTime\" is required but was not present")
	}
	// check if startTime (a required property) was received
	if !startTimeReceived {
		return errors.New("\"startTime\" is required but was not present")
	}
	return nil
}

func (strct *Result) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "CapturedTestOutput" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "CapturedTestOutput" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"CapturedTestOutput\": ")
	if tmp, err := json.Marshal(strct.CapturedTestOutput); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Duration" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "duration" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"duration\": ")
	if tmp, err := json.Marshal(strct.Duration); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "endTime" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"endTime\": ")
	if tmp, err := json.Marshal(strct.EndTime); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "FailureLineContent" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "failureLineContent" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"failureLineContent\": ")
	if tmp, err := json.Marshal(strct.FailureLineContent); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "FailureLocation" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "failureLocation" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"failureLocation\": ")
	if tmp, err := json.Marshal(strct.FailureLocation); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "FailureReason" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "failureReason" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"failureReason\": ")
	if tmp, err := json.Marshal(strct.FailureReason); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "StartTime" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "startTime" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"startTime\": ")
	if tmp, err := json.Marshal(strct.StartTime); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "State" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "state" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"state\": ")
	if tmp, err := json.Marshal(strct.State); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "testID" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"testID\": ")
	if tmp, err := json.Marshal(strct.TestID); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "TestText" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "testText" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"testText\": ")
	if tmp, err := json.Marshal(strct.TestText); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Result) UnmarshalJSON(b []byte) error {
	CapturedTestOutputReceived := false
	durationReceived := false
	failureLineContentReceived := false
	failureLocationReceived := false
	failureReasonReceived := false
	startTimeReceived := false
	stateReceived := false
	testTextReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "CapturedTestOutput":
			if err := json.Unmarshal([]byte(v), &strct.CapturedTestOutput); err != nil {
				return err
			}
			CapturedTestOutputReceived = true
		case "duration":
			if err := json.Unmarshal([]byte(v), &strct.Duration); err != nil {
				return err
			}
			durationReceived = true
		case "endTime":
			if err := json.Unmarshal([]byte(v), &strct.EndTime); err != nil {
				return err
			}
		case "failureLineContent":
			if err := json.Unmarshal([]byte(v), &strct.FailureLineContent); err != nil {
				return err
			}
			failureLineContentReceived = true
		case "failureLocation":
			if err := json.Unmarshal([]byte(v), &strct.FailureLocation); err != nil {
				return err
			}
			failureLocationReceived = true
		case "failureReason":
			if err := json.Unmarshal([]byte(v), &strct.FailureReason); err != nil {
				return err
			}
			failureReasonReceived = true
		case "startTime":
			if err := json.Unmarshal([]byte(v), &strct.StartTime); err != nil {
				return err
			}
			startTimeReceived = true
		case "state":
			if err := json.Unmarshal([]byte(v), &strct.State); err != nil {
				return err
			}
			stateReceived = true
		case "testID":
			if err := json.Unmarshal([]byte(v), &strct.TestID); err != nil {
				return err
			}
		case "testText":
			if err := json.Unmarshal([]byte(v), &strct.TestText); err != nil {
				return err
			}
			testTextReceived = true
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if CapturedTestOutput (a required property) was received
	if !CapturedTestOutputReceived {
		return errors.New("\"CapturedTestOutput\" is required but was not present")
	}
	// check if duration (a required property) was received
	if !durationReceived {
		return errors.New("\"duration\" is required but was not present")
	}
	// check if failureLineContent (a required property) was received
	if !failureLineContentReceived {
		return errors.New("\"failureLineContent\" is required but was not present")
	}
	// check if failureLocation (a required property) was received
	if !failureLocationReceived {
		return errors.New("\"failureLocation\" is required but was not present")
	}
	// check if failureReason (a required property) was received
	if !failureReasonReceived {
		return errors.New("\"failureReason\" is required but was not present")
	}
	// check if startTime (a required property) was received
	if !startTimeReceived {
		return errors.New("\"startTime\" is required but was not present")
	}
	// check if state (a required property) was received
	if !stateReceived {
		return errors.New("\"state\" is required but was not present")
	}
	// check if testText (a required property) was received
	if !testTextReceived {
		return errors.New("\"testText\" is required but was not present")
	}
	return nil
}

func (strct *Root) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "Claim" field is required
	if strct.Claim == nil {
		return nil, errors.New("claim is a required field")
	}
	// Marshal the "claim" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"claim\": ")
	if tmp, err := json.Marshal(strct.Claim); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Root) UnmarshalJSON(b []byte) error {
	claimReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "claim":
			if err := json.Unmarshal([]byte(v), &strct.Claim); err != nil {
				return err
			}
			claimReceived = true
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if claim (a required property) was received
	if !claimReceived {
		return errors.New("\"claim\" is required but was not present")
	}
	return nil
}

func (strct *Versions) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "k8s" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"k8s\": ")
	if tmp, err := json.Marshal(strct.K8s); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "ocClient" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"ocClient\": ")
	if tmp, err := json.Marshal(strct.OcClient); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "ocp" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"ocp\": ")
	if tmp, err := json.Marshal(strct.Ocp); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Tnf" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "tnf" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"tnf\": ")
	if tmp, err := json.Marshal(strct.Tnf); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "tnfGitCommit" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"tnfGitCommit\": ")
	if tmp, err := json.Marshal(strct.TnfGitCommit); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Versions) UnmarshalJSON(b []byte) error {
	tnfReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "k8s":
			if err := json.Unmarshal([]byte(v), &strct.K8s); err != nil {
				return err
			}
		case "ocClient":
			if err := json.Unmarshal([]byte(v), &strct.OcClient); err != nil {
				return err
			}
		case "ocp":
			if err := json.Unmarshal([]byte(v), &strct.Ocp); err != nil {
				return err
			}
		case "tnf":
			if err := json.Unmarshal([]byte(v), &strct.Tnf); err != nil {
				return err
			}
			tnfReceived = true
		case "tnfGitCommit":
			if err := json.Unmarshal([]byte(v), &strct.TnfGitCommit); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if tnf (a required property) was received
	if !tnfReceived {
		return errors.New("\"tnf\" is required but was not present")
	}
	return nil
}