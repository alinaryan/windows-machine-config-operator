// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	operatorv1 "github.com/openshift/api/operator/v1"
)

// DNSSpecApplyConfiguration represents an declarative configuration of the DNSSpec type for use
// with apply.
type DNSSpecApplyConfiguration struct {
	Servers           []ServerApplyConfiguration           `json:"servers,omitempty"`
	UpstreamResolvers *UpstreamResolversApplyConfiguration `json:"upstreamResolvers,omitempty"`
	NodePlacement     *DNSNodePlacementApplyConfiguration  `json:"nodePlacement,omitempty"`
	ManagementState   *operatorv1.ManagementState          `json:"managementState,omitempty"`
	OperatorLogLevel  *operatorv1.DNSLogLevel              `json:"operatorLogLevel,omitempty"`
	LogLevel          *operatorv1.DNSLogLevel              `json:"logLevel,omitempty"`
	Cache             *DNSCacheApplyConfiguration          `json:"cache,omitempty"`
}

// DNSSpecApplyConfiguration constructs an declarative configuration of the DNSSpec type for use with
// apply.
func DNSSpec() *DNSSpecApplyConfiguration {
	return &DNSSpecApplyConfiguration{}
}

// WithServers adds the given value to the Servers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Servers field.
func (b *DNSSpecApplyConfiguration) WithServers(values ...*ServerApplyConfiguration) *DNSSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithServers")
		}
		b.Servers = append(b.Servers, *values[i])
	}
	return b
}

// WithUpstreamResolvers sets the UpstreamResolvers field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UpstreamResolvers field is set to the value of the last call.
func (b *DNSSpecApplyConfiguration) WithUpstreamResolvers(value *UpstreamResolversApplyConfiguration) *DNSSpecApplyConfiguration {
	b.UpstreamResolvers = value
	return b
}

// WithNodePlacement sets the NodePlacement field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NodePlacement field is set to the value of the last call.
func (b *DNSSpecApplyConfiguration) WithNodePlacement(value *DNSNodePlacementApplyConfiguration) *DNSSpecApplyConfiguration {
	b.NodePlacement = value
	return b
}

// WithManagementState sets the ManagementState field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ManagementState field is set to the value of the last call.
func (b *DNSSpecApplyConfiguration) WithManagementState(value operatorv1.ManagementState) *DNSSpecApplyConfiguration {
	b.ManagementState = &value
	return b
}

// WithOperatorLogLevel sets the OperatorLogLevel field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OperatorLogLevel field is set to the value of the last call.
func (b *DNSSpecApplyConfiguration) WithOperatorLogLevel(value operatorv1.DNSLogLevel) *DNSSpecApplyConfiguration {
	b.OperatorLogLevel = &value
	return b
}

// WithLogLevel sets the LogLevel field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LogLevel field is set to the value of the last call.
func (b *DNSSpecApplyConfiguration) WithLogLevel(value operatorv1.DNSLogLevel) *DNSSpecApplyConfiguration {
	b.LogLevel = &value
	return b
}

// WithCache sets the Cache field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Cache field is set to the value of the last call.
func (b *DNSSpecApplyConfiguration) WithCache(value *DNSCacheApplyConfiguration) *DNSSpecApplyConfiguration {
	b.Cache = value
	return b
}