package main

type PdpResponseStatus string

const (
	Success PdpResponseStatus = "SUCCESS"
	Failure PdpResponseStatus = "FAILURE"
)

type PdpResponseDetails struct {
	ResponseTo      string
	ResponseStatus  PdpResponseStatus
	ResponseMessage string
}

func NewPdpResponseDetails(responseTo string, responseStatus PdpResponseStatus, responseMessage string) *PdpResponseDetails {
	return &PdpResponseDetails{
		ResponseTo:      responseTo,
		ResponseStatus:  responseStatus,
		ResponseMessage: responseMessage,
	}
}

func CopyPdpResponseDetails(source *PdpResponseDetails) *PdpResponseDetails {
	return &PdpResponseDetails{
		ResponseTo:      source.ResponseTo,
		ResponseStatus:  source.ResponseStatus,
		ResponseMessage: source.ResponseMessage,
	}
}
