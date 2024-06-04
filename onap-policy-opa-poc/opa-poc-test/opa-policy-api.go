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

// IN PROGRESS: Not working

import (
	"context"
	"fmt"

	"github.com/open-policy-agent/opa/rego"
)

func main() {
	module := `
	package example


	default validate_policy = false


	validate_policy {
		input.policyType == "onap.test.opa"
		input.policyVersion >= 1
		input.taskLogic != null
	}
	`

	ctx := context.TODO()

	query, err := rego.New(
		rego.Query("x = data.example.validate_policy"),
		rego.Module("example.validate_policy", module),
	).PrepareForEval(ctx)

	if err != nil {
		fmt.Println("Query error.")
	}

	input := map[string]interface{}{
		"method": "GET",
		"subject": map[string]interface{}{
			"policyType":    "onap.test.opa",
			"policyVersion": 1.1,
			"taskLogic":     "some logic here",
		},
	}

	results, err := query.Eval(ctx, rego.EvalInput(input))

	if err != nil {
		fmt.Println("Error caught. Results could not be generated.")
	}
	if !results.Allowed() {
		fmt.Println("False")
	} else {
		fmt.Println("True")
	}

}
