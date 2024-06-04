/*
#
# ============LICENSE_START================================================
# ONAP
# =========================================================================
# Copyright (C) 2024 Nordix Foundation.
# =========================================================================
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
# ============LICENSE_END==================================================
#
*/

package main

// Communicates with OPA over REST

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/rs/zerolog/log"
)

type Input struct {
	Input string `json:"input`
}

func main() {
	getPolicies()

	postURL := url.URL{
		Host:   "localhost:8181",
		Path:   "v1/data/example/validate_policy",
		Scheme: "http",
	}

	newMessage := Input{"\"policyType\":\"onap.test.opa\", \"policyVersion\":1.0, \"taskLogic\": \"somelogic\""}

	responseBody, err := triggerPolicy(postURL, newMessage)
	if err != nil {
		log.Fatal()
	}
	fmt.Printf("POST response body %s", responseBody)

	//getMessages()
}

func getPolicies() []byte {
	client := http.Client{}
	respGet, errGet := client.Get("http://localhost:8181/v1/policies")
	if errGet != nil {
		fmt.Printf(("Error %s, errGet"))
	}
	defer respGet.Body.Close()

	// allows a function to postpone the execution of a statement until the surrounding function has completed

	body, err := io.ReadAll(respGet.Body)
	if err != nil {
		fmt.Printf(("Error %s, err"))
	}
	fmt.Printf("http://localhost:8181/v1/policies: %s", body)

	return body
}

func triggerPolicy(url url.URL, body any) (string, error) {
	bodyBytes, err := json.Marshal(&body)
	if err != nil {
		return "", nil
	}

	reader := bytes.NewReader(bodyBytes)

	log.Info().Msg(string(bodyBytes))

	request, err := http.NewRequest(http.MethodPost, url.String(), reader)
	if err != nil {
		return "", nil
	}

	request.Header.Set("Content-Type", "application/json")

	httpClient := &http.Client{}

	resp, err := httpClient.Do(request)
	if err != nil {
		return "", nil
	}

	// Close response body
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Fatal()
		}
	}()

	// Read response body
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", nil
	}

	if resp.StatusCode >= 400 && resp.StatusCode <= 500 {
		return string(responseBody), errors.New("400/500 status code error")
	}

	return string(responseBody), nil
}
