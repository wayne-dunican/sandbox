package example

default validate_policy = false

validate_policy {
	input.policyType == "onap.test.opa"
	input.policyVersion >= 1
	input.taskLogic != null
}
