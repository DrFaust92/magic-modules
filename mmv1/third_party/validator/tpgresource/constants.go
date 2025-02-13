package tpgresource

import (
	"errors"

	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

// ErrNoConversion can be returned if a conversion is unable to be returned.

// because of the current state of the system.
// Example: The conversion requires that the resource has already been created
// and is now being updated).
var ErrNoConversion = errors.New("no conversion")

// ErrEmptyIdentityField can be returned when fetching a resource is not possible
// due to the identity field of that resource returning empty.
var ErrEmptyIdentityField = errors.New("empty identity field")

// ErrResourceInaccessible can be returned when fetching an IAM resource
// on a project that has not yet been created or if the service account
// lacks sufficient permissions
var ErrResourceInaccessible = errors.New("resource does not exist or service account is lacking sufficient permissions")

// Global MutexKV
//
// Deprecated: For backward compatibility mutexKV is still working,
// but all new code should use MutexStore in the transport_tpg package instead.
var mutexKV = transport_tpg.MutexStore
