package main

import (
	"fmt"
)

type ToscaConceptIdentifier struct {
	Name    string
	Version string
}

func NewToscaConceptIdentifier(name, version string) *ToscaConceptIdentifier {
	return &ToscaConceptIdentifier{
		Name:    name,
		Version: version,
	}
}

func NewToscaConceptIdentifierFromKey(key PfKey) *ToscaConceptIdentifier {
	return &ToscaConceptIdentifier{
		Name:    key.Name,
		Version: key.Version,
	}
}

func (id *ToscaConceptIdentifier) ValidatePapRest() error {
	if id.Name == "" || id.Version == "" {
		return fmt.Errorf("name and version must be non-empty")
	}
	return nil
}

type PfKey struct {
	Name    string
	Version string
}
