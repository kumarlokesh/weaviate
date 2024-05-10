// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RaftStatistics The definition of Raft statistics.
//
// swagger:model RaftStatistics
type RaftStatistics struct {

	// applied index
	AppliedIndex string `json:"appliedIndex,omitempty"`

	// commit index
	CommitIndex string `json:"commitIndex,omitempty"`

	// fsm pending
	FsmPending string `json:"fsmPending,omitempty"`

	// last contact
	LastContact string `json:"lastContact,omitempty"`

	// last log index
	LastLogIndex string `json:"lastLogIndex,omitempty"`

	// last log term
	LastLogTerm string `json:"lastLogTerm,omitempty"`

	// last snapshot index
	LastSnapshotIndex string `json:"lastSnapshotIndex,omitempty"`

	// last snapshot term
	LastSnapshotTerm string `json:"lastSnapshotTerm,omitempty"`

	// Weaviate Raft nodes.
	LatestConfiguration interface{} `json:"latestConfiguration,omitempty"`

	// latest configuration index
	LatestConfigurationIndex string `json:"latestConfigurationIndex,omitempty"`

	// num peers
	NumPeers string `json:"numPeers,omitempty"`

	// protocol version
	ProtocolVersion string `json:"protocolVersion,omitempty"`

	// protocol version max
	ProtocolVersionMax string `json:"protocolVersionMax,omitempty"`

	// protocol version min
	ProtocolVersionMin string `json:"protocolVersionMin,omitempty"`

	// snapshot version max
	SnapshotVersionMax string `json:"snapshotVersionMax,omitempty"`

	// snapshot version min
	SnapshotVersionMin string `json:"snapshotVersionMin,omitempty"`

	// state
	State string `json:"state,omitempty"`

	// term
	Term string `json:"term,omitempty"`
}

// Validate validates this raft statistics
func (m *RaftStatistics) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this raft statistics based on context it is used
func (m *RaftStatistics) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RaftStatistics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RaftStatistics) UnmarshalBinary(b []byte) error {
	var res RaftStatistics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
