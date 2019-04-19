// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datadog

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Datadog user resource. This can be used to create and manage Datadog users.
type DatadogUser struct {
	s *pulumi.ResourceState
}

// NewDatadogUser registers a new resource with the given unique name, arguments, and options.
func NewDatadogUser(ctx *pulumi.Context,
	name string, args *DatadogUserArgs, opts ...pulumi.ResourceOpt) (*DatadogUser, error) {
	if args == nil || args.Email == nil {
		return nil, errors.New("missing required argument 'Email'")
	}
	if args == nil || args.Handle == nil {
		return nil, errors.New("missing required argument 'Handle'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["accessRole"] = nil
		inputs["disabled"] = nil
		inputs["email"] = nil
		inputs["handle"] = nil
		inputs["isAdmin"] = nil
		inputs["name"] = nil
		inputs["role"] = nil
	} else {
		inputs["accessRole"] = args.AccessRole
		inputs["disabled"] = args.Disabled
		inputs["email"] = args.Email
		inputs["handle"] = args.Handle
		inputs["isAdmin"] = args.IsAdmin
		inputs["name"] = args.Name
		inputs["role"] = args.Role
	}
	inputs["verified"] = nil
	s, err := ctx.RegisterResource("datadog:index/datadogUser:DatadogUser", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DatadogUser{s: s}, nil
}

// GetDatadogUser gets an existing DatadogUser resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatadogUser(ctx *pulumi.Context,
	name string, id pulumi.ID, state *DatadogUserState, opts ...pulumi.ResourceOpt) (*DatadogUser, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["accessRole"] = state.AccessRole
		inputs["disabled"] = state.Disabled
		inputs["email"] = state.Email
		inputs["handle"] = state.Handle
		inputs["isAdmin"] = state.IsAdmin
		inputs["name"] = state.Name
		inputs["role"] = state.Role
		inputs["verified"] = state.Verified
	}
	s, err := ctx.ReadResource("datadog:index/datadogUser:DatadogUser", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DatadogUser{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *DatadogUser) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *DatadogUser) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Role description for user. Can be `st` (standard user), `adm` (admin user) or `ro` (read-only user).  Default is `st`.
func (r *DatadogUser) AccessRole() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["accessRole"])
}

// Whether the user is disabled
func (r *DatadogUser) Disabled() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["disabled"])
}

// Email address for user
func (r *DatadogUser) Email() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["email"])
}

// The user handle, must be a valid email.
func (r *DatadogUser) Handle() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["handle"])
}

// (Optional) Whether the user is an administrator. **Warning**: the corresponding query parameter is ignored by the Datadog API, thus the argument would always trigger an execution plan.
func (r *DatadogUser) IsAdmin() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["isAdmin"])
}

// Name for user
func (r *DatadogUser) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// Role description for user. **Warning**: the corresponding query parameter is ignored by the Datadog API, thus the argument would always trigger an execution plan.
func (r *DatadogUser) Role() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["role"])
}

// Returns true if Datadog user is verified
func (r *DatadogUser) Verified() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["verified"])
}

// Input properties used for looking up and filtering DatadogUser resources.
type DatadogUserState struct {
	// Role description for user. Can be `st` (standard user), `adm` (admin user) or `ro` (read-only user).  Default is `st`.
	AccessRole interface{}
	// Whether the user is disabled
	Disabled interface{}
	// Email address for user
	Email interface{}
	// The user handle, must be a valid email.
	Handle interface{}
	// (Optional) Whether the user is an administrator. **Warning**: the corresponding query parameter is ignored by the Datadog API, thus the argument would always trigger an execution plan.
	IsAdmin interface{}
	// Name for user
	Name interface{}
	// Role description for user. **Warning**: the corresponding query parameter is ignored by the Datadog API, thus the argument would always trigger an execution plan.
	Role interface{}
	// Returns true if Datadog user is verified
	Verified interface{}
}

// The set of arguments for constructing a DatadogUser resource.
type DatadogUserArgs struct {
	// Role description for user. Can be `st` (standard user), `adm` (admin user) or `ro` (read-only user).  Default is `st`.
	AccessRole interface{}
	// Whether the user is disabled
	Disabled interface{}
	// Email address for user
	Email interface{}
	// The user handle, must be a valid email.
	Handle interface{}
	// (Optional) Whether the user is an administrator. **Warning**: the corresponding query parameter is ignored by the Datadog API, thus the argument would always trigger an execution plan.
	IsAdmin interface{}
	// Name for user
	Name interface{}
	// Role description for user. **Warning**: the corresponding query parameter is ignored by the Datadog API, thus the argument would always trigger an execution plan.
	Role interface{}
}