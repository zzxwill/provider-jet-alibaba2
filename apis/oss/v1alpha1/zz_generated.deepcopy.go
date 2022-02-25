//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AbortMultipartUploadObservation) DeepCopyInto(out *AbortMultipartUploadObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AbortMultipartUploadObservation.
func (in *AbortMultipartUploadObservation) DeepCopy() *AbortMultipartUploadObservation {
	if in == nil {
		return nil
	}
	out := new(AbortMultipartUploadObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AbortMultipartUploadParameters) DeepCopyInto(out *AbortMultipartUploadParameters) {
	*out = *in
	if in.CreatedBeforeDate != nil {
		in, out := &in.CreatedBeforeDate, &out.CreatedBeforeDate
		*out = new(string)
		**out = **in
	}
	if in.Days != nil {
		in, out := &in.Days, &out.Days
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AbortMultipartUploadParameters.
func (in *AbortMultipartUploadParameters) DeepCopy() *AbortMultipartUploadParameters {
	if in == nil {
		return nil
	}
	out := new(AbortMultipartUploadParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bucket) DeepCopyInto(out *Bucket) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bucket.
func (in *Bucket) DeepCopy() *Bucket {
	if in == nil {
		return nil
	}
	out := new(Bucket)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Bucket) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketList) DeepCopyInto(out *BucketList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Bucket, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketList.
func (in *BucketList) DeepCopy() *BucketList {
	if in == nil {
		return nil
	}
	out := new(BucketList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BucketList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketObservation) DeepCopyInto(out *BucketObservation) {
	*out = *in
	if in.CreationDate != nil {
		in, out := &in.CreationDate, &out.CreationDate
		*out = new(string)
		**out = **in
	}
	if in.ExtranetEndpoint != nil {
		in, out := &in.ExtranetEndpoint, &out.ExtranetEndpoint
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IntranetEndpoint != nil {
		in, out := &in.IntranetEndpoint, &out.IntranetEndpoint
		*out = new(string)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketObservation.
func (in *BucketObservation) DeepCopy() *BucketObservation {
	if in == nil {
		return nil
	}
	out := new(BucketObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketParameters) DeepCopyInto(out *BucketParameters) {
	*out = *in
	if in.ACL != nil {
		in, out := &in.ACL, &out.ACL
		*out = new(string)
		**out = **in
	}
	if in.Bucket != nil {
		in, out := &in.Bucket, &out.Bucket
		*out = new(string)
		**out = **in
	}
	if in.CorsRule != nil {
		in, out := &in.CorsRule, &out.CorsRule
		*out = make([]CorsRuleParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ForceDestroy != nil {
		in, out := &in.ForceDestroy, &out.ForceDestroy
		*out = new(bool)
		**out = **in
	}
	if in.LifecycleRule != nil {
		in, out := &in.LifecycleRule, &out.LifecycleRule
		*out = make([]LifecycleRuleParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Logging != nil {
		in, out := &in.Logging, &out.Logging
		*out = make([]LoggingParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LoggingIsenable != nil {
		in, out := &in.LoggingIsenable, &out.LoggingIsenable
		*out = new(bool)
		**out = **in
	}
	if in.Policy != nil {
		in, out := &in.Policy, &out.Policy
		*out = new(string)
		**out = **in
	}
	if in.RedundancyType != nil {
		in, out := &in.RedundancyType, &out.RedundancyType
		*out = new(string)
		**out = **in
	}
	if in.RefererConfig != nil {
		in, out := &in.RefererConfig, &out.RefererConfig
		*out = make([]RefererConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ServerSideEncryptionRule != nil {
		in, out := &in.ServerSideEncryptionRule, &out.ServerSideEncryptionRule
		*out = make([]ServerSideEncryptionRuleParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.StorageClass != nil {
		in, out := &in.StorageClass, &out.StorageClass
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.TransferAcceleration != nil {
		in, out := &in.TransferAcceleration, &out.TransferAcceleration
		*out = make([]TransferAccelerationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Versioning != nil {
		in, out := &in.Versioning, &out.Versioning
		*out = make([]VersioningParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Website != nil {
		in, out := &in.Website, &out.Website
		*out = make([]WebsiteParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketParameters.
func (in *BucketParameters) DeepCopy() *BucketParameters {
	if in == nil {
		return nil
	}
	out := new(BucketParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketSpec) DeepCopyInto(out *BucketSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketSpec.
func (in *BucketSpec) DeepCopy() *BucketSpec {
	if in == nil {
		return nil
	}
	out := new(BucketSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketStatus) DeepCopyInto(out *BucketStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketStatus.
func (in *BucketStatus) DeepCopy() *BucketStatus {
	if in == nil {
		return nil
	}
	out := new(BucketStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CorsRuleObservation) DeepCopyInto(out *CorsRuleObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CorsRuleObservation.
func (in *CorsRuleObservation) DeepCopy() *CorsRuleObservation {
	if in == nil {
		return nil
	}
	out := new(CorsRuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CorsRuleParameters) DeepCopyInto(out *CorsRuleParameters) {
	*out = *in
	if in.AllowedHeaders != nil {
		in, out := &in.AllowedHeaders, &out.AllowedHeaders
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AllowedMethods != nil {
		in, out := &in.AllowedMethods, &out.AllowedMethods
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AllowedOrigins != nil {
		in, out := &in.AllowedOrigins, &out.AllowedOrigins
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ExposeHeaders != nil {
		in, out := &in.ExposeHeaders, &out.ExposeHeaders
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.MaxAgeSeconds != nil {
		in, out := &in.MaxAgeSeconds, &out.MaxAgeSeconds
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CorsRuleParameters.
func (in *CorsRuleParameters) DeepCopy() *CorsRuleParameters {
	if in == nil {
		return nil
	}
	out := new(CorsRuleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExpirationObservation) DeepCopyInto(out *ExpirationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExpirationObservation.
func (in *ExpirationObservation) DeepCopy() *ExpirationObservation {
	if in == nil {
		return nil
	}
	out := new(ExpirationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExpirationParameters) DeepCopyInto(out *ExpirationParameters) {
	*out = *in
	if in.CreatedBeforeDate != nil {
		in, out := &in.CreatedBeforeDate, &out.CreatedBeforeDate
		*out = new(string)
		**out = **in
	}
	if in.Date != nil {
		in, out := &in.Date, &out.Date
		*out = new(string)
		**out = **in
	}
	if in.Days != nil {
		in, out := &in.Days, &out.Days
		*out = new(float64)
		**out = **in
	}
	if in.ExpiredObjectDeleteMarker != nil {
		in, out := &in.ExpiredObjectDeleteMarker, &out.ExpiredObjectDeleteMarker
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExpirationParameters.
func (in *ExpirationParameters) DeepCopy() *ExpirationParameters {
	if in == nil {
		return nil
	}
	out := new(ExpirationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LifecycleRuleObservation) DeepCopyInto(out *LifecycleRuleObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LifecycleRuleObservation.
func (in *LifecycleRuleObservation) DeepCopy() *LifecycleRuleObservation {
	if in == nil {
		return nil
	}
	out := new(LifecycleRuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LifecycleRuleParameters) DeepCopyInto(out *LifecycleRuleParameters) {
	*out = *in
	if in.AbortMultipartUpload != nil {
		in, out := &in.AbortMultipartUpload, &out.AbortMultipartUpload
		*out = make([]AbortMultipartUploadParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Expiration != nil {
		in, out := &in.Expiration, &out.Expiration
		*out = make([]ExpirationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.NoncurrentVersionExpiration != nil {
		in, out := &in.NoncurrentVersionExpiration, &out.NoncurrentVersionExpiration
		*out = make([]NoncurrentVersionExpirationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NoncurrentVersionTransition != nil {
		in, out := &in.NoncurrentVersionTransition, &out.NoncurrentVersionTransition
		*out = make([]NoncurrentVersionTransitionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Prefix != nil {
		in, out := &in.Prefix, &out.Prefix
		*out = new(string)
		**out = **in
	}
	if in.Transitions != nil {
		in, out := &in.Transitions, &out.Transitions
		*out = make([]TransitionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LifecycleRuleParameters.
func (in *LifecycleRuleParameters) DeepCopy() *LifecycleRuleParameters {
	if in == nil {
		return nil
	}
	out := new(LifecycleRuleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingObservation) DeepCopyInto(out *LoggingObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingObservation.
func (in *LoggingObservation) DeepCopy() *LoggingObservation {
	if in == nil {
		return nil
	}
	out := new(LoggingObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingParameters) DeepCopyInto(out *LoggingParameters) {
	*out = *in
	if in.TargetBucket != nil {
		in, out := &in.TargetBucket, &out.TargetBucket
		*out = new(string)
		**out = **in
	}
	if in.TargetPrefix != nil {
		in, out := &in.TargetPrefix, &out.TargetPrefix
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingParameters.
func (in *LoggingParameters) DeepCopy() *LoggingParameters {
	if in == nil {
		return nil
	}
	out := new(LoggingParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NoncurrentVersionExpirationObservation) DeepCopyInto(out *NoncurrentVersionExpirationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NoncurrentVersionExpirationObservation.
func (in *NoncurrentVersionExpirationObservation) DeepCopy() *NoncurrentVersionExpirationObservation {
	if in == nil {
		return nil
	}
	out := new(NoncurrentVersionExpirationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NoncurrentVersionExpirationParameters) DeepCopyInto(out *NoncurrentVersionExpirationParameters) {
	*out = *in
	if in.Days != nil {
		in, out := &in.Days, &out.Days
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NoncurrentVersionExpirationParameters.
func (in *NoncurrentVersionExpirationParameters) DeepCopy() *NoncurrentVersionExpirationParameters {
	if in == nil {
		return nil
	}
	out := new(NoncurrentVersionExpirationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NoncurrentVersionTransitionObservation) DeepCopyInto(out *NoncurrentVersionTransitionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NoncurrentVersionTransitionObservation.
func (in *NoncurrentVersionTransitionObservation) DeepCopy() *NoncurrentVersionTransitionObservation {
	if in == nil {
		return nil
	}
	out := new(NoncurrentVersionTransitionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NoncurrentVersionTransitionParameters) DeepCopyInto(out *NoncurrentVersionTransitionParameters) {
	*out = *in
	if in.Days != nil {
		in, out := &in.Days, &out.Days
		*out = new(float64)
		**out = **in
	}
	if in.StorageClass != nil {
		in, out := &in.StorageClass, &out.StorageClass
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NoncurrentVersionTransitionParameters.
func (in *NoncurrentVersionTransitionParameters) DeepCopy() *NoncurrentVersionTransitionParameters {
	if in == nil {
		return nil
	}
	out := new(NoncurrentVersionTransitionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RefererConfigObservation) DeepCopyInto(out *RefererConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RefererConfigObservation.
func (in *RefererConfigObservation) DeepCopy() *RefererConfigObservation {
	if in == nil {
		return nil
	}
	out := new(RefererConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RefererConfigParameters) DeepCopyInto(out *RefererConfigParameters) {
	*out = *in
	if in.AllowEmpty != nil {
		in, out := &in.AllowEmpty, &out.AllowEmpty
		*out = new(bool)
		**out = **in
	}
	if in.Referers != nil {
		in, out := &in.Referers, &out.Referers
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RefererConfigParameters.
func (in *RefererConfigParameters) DeepCopy() *RefererConfigParameters {
	if in == nil {
		return nil
	}
	out := new(RefererConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerSideEncryptionRuleObservation) DeepCopyInto(out *ServerSideEncryptionRuleObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerSideEncryptionRuleObservation.
func (in *ServerSideEncryptionRuleObservation) DeepCopy() *ServerSideEncryptionRuleObservation {
	if in == nil {
		return nil
	}
	out := new(ServerSideEncryptionRuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerSideEncryptionRuleParameters) DeepCopyInto(out *ServerSideEncryptionRuleParameters) {
	*out = *in
	if in.KMSMasterKeyID != nil {
		in, out := &in.KMSMasterKeyID, &out.KMSMasterKeyID
		*out = new(string)
		**out = **in
	}
	if in.SseAlgorithm != nil {
		in, out := &in.SseAlgorithm, &out.SseAlgorithm
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerSideEncryptionRuleParameters.
func (in *ServerSideEncryptionRuleParameters) DeepCopy() *ServerSideEncryptionRuleParameters {
	if in == nil {
		return nil
	}
	out := new(ServerSideEncryptionRuleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransferAccelerationObservation) DeepCopyInto(out *TransferAccelerationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransferAccelerationObservation.
func (in *TransferAccelerationObservation) DeepCopy() *TransferAccelerationObservation {
	if in == nil {
		return nil
	}
	out := new(TransferAccelerationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransferAccelerationParameters) DeepCopyInto(out *TransferAccelerationParameters) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransferAccelerationParameters.
func (in *TransferAccelerationParameters) DeepCopy() *TransferAccelerationParameters {
	if in == nil {
		return nil
	}
	out := new(TransferAccelerationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransitionsObservation) DeepCopyInto(out *TransitionsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransitionsObservation.
func (in *TransitionsObservation) DeepCopy() *TransitionsObservation {
	if in == nil {
		return nil
	}
	out := new(TransitionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransitionsParameters) DeepCopyInto(out *TransitionsParameters) {
	*out = *in
	if in.CreatedBeforeDate != nil {
		in, out := &in.CreatedBeforeDate, &out.CreatedBeforeDate
		*out = new(string)
		**out = **in
	}
	if in.Days != nil {
		in, out := &in.Days, &out.Days
		*out = new(float64)
		**out = **in
	}
	if in.StorageClass != nil {
		in, out := &in.StorageClass, &out.StorageClass
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransitionsParameters.
func (in *TransitionsParameters) DeepCopy() *TransitionsParameters {
	if in == nil {
		return nil
	}
	out := new(TransitionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VersioningObservation) DeepCopyInto(out *VersioningObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VersioningObservation.
func (in *VersioningObservation) DeepCopy() *VersioningObservation {
	if in == nil {
		return nil
	}
	out := new(VersioningObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VersioningParameters) DeepCopyInto(out *VersioningParameters) {
	*out = *in
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VersioningParameters.
func (in *VersioningParameters) DeepCopy() *VersioningParameters {
	if in == nil {
		return nil
	}
	out := new(VersioningParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebsiteObservation) DeepCopyInto(out *WebsiteObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebsiteObservation.
func (in *WebsiteObservation) DeepCopy() *WebsiteObservation {
	if in == nil {
		return nil
	}
	out := new(WebsiteObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebsiteParameters) DeepCopyInto(out *WebsiteParameters) {
	*out = *in
	if in.ErrorDocument != nil {
		in, out := &in.ErrorDocument, &out.ErrorDocument
		*out = new(string)
		**out = **in
	}
	if in.IndexDocument != nil {
		in, out := &in.IndexDocument, &out.IndexDocument
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebsiteParameters.
func (in *WebsiteParameters) DeepCopy() *WebsiteParameters {
	if in == nil {
		return nil
	}
	out := new(WebsiteParameters)
	in.DeepCopyInto(out)
	return out
}
