/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this ManagedPolicyAttachment.
func (mg *ManagedPolicyAttachment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PermissionSetArn),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.ForProvider.PermissionSetArnRef,
		Selector:     mg.Spec.ForProvider.PermissionSetArnSelector,
		To: reference.To{
			List:    &PermissionSetList{},
			Managed: &PermissionSet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PermissionSetArn")
	}
	mg.Spec.ForProvider.PermissionSetArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PermissionSetArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this PermissionSetInlinePolicy.
func (mg *PermissionSetInlinePolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceArn),
		Extract:      resource.ExtractParamPath("instance_arn", false),
		Reference:    mg.Spec.ForProvider.InstanceArnRef,
		Selector:     mg.Spec.ForProvider.InstanceArnSelector,
		To: reference.To{
			List:    &PermissionSetList{},
			Managed: &PermissionSet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceArn")
	}
	mg.Spec.ForProvider.InstanceArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceArnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PermissionSetArn),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.ForProvider.PermissionSetArnRef,
		Selector:     mg.Spec.ForProvider.PermissionSetArnSelector,
		To: reference.To{
			List:    &PermissionSetList{},
			Managed: &PermissionSet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PermissionSetArn")
	}
	mg.Spec.ForProvider.PermissionSetArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PermissionSetArnRef = rsp.ResolvedReference

	return nil
}
