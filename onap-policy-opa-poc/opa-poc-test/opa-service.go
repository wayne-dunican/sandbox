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

/*
	Example curl commands:

	POST valid policy data	 -> curl -H 'Content-Type: application/json' -H 'Accept: application/json' -d '{"policyType":"onap.test.opa","policyVersion":1.1,"taskLogic":"Some logic"}'  -X POST http://localhost:8282/opa/policies/data
	POST invalid policy data -> curl -H 'Content-Type: application/json' -H 'Accept: application/json' -d '{"policyType":"onap.test.policy.framework.apex","policyVersion":0.1,"taskLogic":"Some logic"}'  -X POST http://localhost:8282/opa/policies/data
*/

package main

import (
	"bytes"
	"context"
	"os"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/open-policy-agent/opa/sdk"
	sdktest "github.com/open-policy-agent/opa/sdk/test"
)

type inputData struct {
	PolicyType    string  `json:"policyType"`
	PolicyVersion float64 `json:"policyVersion"`
	TaskLogic     string  `json:"taskLogic"`
}

func main() {

	// Spin up OPA Instance
	ctx := context.Background()

	// Read policy from a file
	policy, err := os.ReadFile("access_policy.rego")
	if err != nil {
		fmt.Println(err)
	}

	// create a mock bundle server with validate_policy policy running
	server, err := sdktest.NewServer(sdktest.MockBundle("/bundles/bundle.tar.gz", map[string]string{
		"example.rego": string(policy),
	}))
	if err != nil {
		fmt.Println("Cannot spin-up opa...")
	}

	defer server.Stop()

	// provide the OPA configuration which specifies
	// fetching policy bundles from the mock server
	// and logging decisions locally to the console
	config := []byte(fmt.Sprintf(`{
		"services": {
			"test": {
				"url": %q
			}
		},
		"bundles": {
			"test": {
				"resource": "/bundles/bundle.tar.gz"
			}
		},
		"decision_logs": {
			"console": true
		}
	}`, server.URL()))

	// create an instance of the OPA object
	opa, err := sdk.New(ctx, sdk.Options{
		ID:     "opa-test-1",
		Config: bytes.NewReader(config),
	})
	if err != nil {
		fmt.Println("Could not spin up OPA instance...")
	}

	defer opa.Stop(ctx)
	// OPA should be now running

	// Set up routes...
	router := gin.Default()
	router.POST("/opa/policies/data", func(c *gin.Context) {

		var json inputData
		c.BindJSON(&json)

		if result, err := opa.Decision(ctx, sdk.DecisionOptions{
			Path:  "/example/validate_policy",
			Input: json,
		}); err != nil {
			c.String(400, "Policy and/or input invalid")
		} else if decision, ok := result.Result.(bool); !ok || !decision {
			c.String(200, fmt.Sprintln("Decision: Policy is invalid"))
		} else {
			c.String(202, fmt.Sprintln("Decision: Policy is valid"))
		}
	})

	// Start service
	router.Run("localhost:8282")
}
