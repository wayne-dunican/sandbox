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

// Uses OPA GO SDK library

import (
	"bytes"
	"context"
	"fmt"
	"os"

	"github.com/open-policy-agent/opa/sdk"
	sdktest "github.com/open-policy-agent/opa/sdk/test"
)

// ^^^ IMPORT sdktest "github.com/open-policy-agent/opa/sdk/test"

/*

   The SDK package contains high-level APIs for embedding OPA inside of Go programs and obtaining the output of query evaluation.

   Use the low-level github.com/open-policy-agent/opa/rego package to embed OPA as a library inside services written in Go,
   when only policy evaluation — and no other capabilities of OPA, like the management features — are desired.
   If you’re unsure which one to use, the SDK is probably the better option.

*/

func main() {

	/*
	 Context is a built-in package in the Go standard library that provides a powerful toolset
	 for managing concurrent operations. It enables the propagation of cancellation signals,
	 deadlines, and values across goroutines, ensuring that related operations can gracefully terminate when necessary.

	*/

	/*

	 A Goroutine is a lightweight thread managed by Go runtime

	*/
	ctx := context.Background()

	// Read policy from a file
	policy, err := os.ReadFile("access_policy.rego")
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(string(policy))

	// create a mock bundle server with validate_policy policy rule
	// https://www.openpolicyagent.org/docs/latest/management-bundles/
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

	fmt.Println("server.URL(): ", server.URL())

	// create an instance of the OPA object
	opa, err := sdk.New(ctx, sdk.Options{
		ID:     "opa-test-1",
		Config: bytes.NewReader(config),
	})
	if err != nil {
		// handle error.
		fmt.Println("This is an error")
	}

	defer opa.Stop(ctx)

	// get the policy decision for the specified input
	if result, err := opa.Decision(ctx, sdk.DecisionOptions{
		Path: "/example/validate_policy",
		Input: map[string]interface{}{
			"policyType":    "onap.test.opa",
			"policyVersion": 1.1,
			"taskLogic":     "some logic here",
		},
	}); err != nil {
		fmt.Println("URL invalid")
	} else if decision, ok := result.Result.(bool); !ok || !decision {
		fmt.Printf("Valid policy? %v\n", decision)
	} else {
		fmt.Printf("Valid policy? %v\n", decision)
	}
}
