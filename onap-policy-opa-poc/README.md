
This repository contains the Go code for the OPA PDP Policy Framework Integration PoC

This is experimental at present. 

- opa-policy-api.go
Integrating OPA using the rego opa library. In progress. Not working yet.

- opa-policy-kafka.go


- opa-policy-rest.go
Integrating OPA using the OPA rest library. Working example.

- opa-policy-sdk.go
Integrating OPA using the OPA sdk library. Working example.

- opa-service.go
Spinning up a OPA PDP service. This service receives REST POST requests and queries OPA, returning a result to the user over REST. Working example.

- kafka-comm/
Directory for Kafka communication classes. Work In Progress.

- config/
Directory to store all docker-compose config for components.

- docker-compose.yml
docker-compose file to spin up OPA alongside PF components for docker testing.

- access_policy.rego
Example OPA policy.

- input.json
Example OPA policy input for the above example policy.

