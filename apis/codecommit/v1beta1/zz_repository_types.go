/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type RepositoryObservation struct {

	// The ARN of the repository
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The URL to use for cloning the repository over HTTPS.
	CloneURLHTTP *string `json:"cloneUrlHttp,omitempty" tf:"clone_url_http,omitempty"`

	// The URL to use for cloning the repository over SSH.
	CloneURLSSH *string `json:"cloneUrlSsh,omitempty" tf:"clone_url_ssh,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the repository
	RepositoryID *string `json:"repositoryId,omitempty" tf:"repository_id,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type RepositoryParameters struct {

	// The default branch of the repository. The branch specified here needs to exist.
	// +kubebuilder:validation:Optional
	DefaultBranch *string `json:"defaultBranch,omitempty" tf:"default_branch,omitempty"`

	// The description of the repository. This needs to be less than 1000 characters
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// RepositorySpec defines the desired state of Repository
type RepositorySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RepositoryParameters `json:"forProvider"`
}

// RepositoryStatus defines the observed state of Repository.
type RepositoryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RepositoryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Repository is the Schema for the Repositorys API. Provides a CodeCommit Repository Resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Repository struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RepositorySpec   `json:"spec"`
	Status            RepositoryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RepositoryList contains a list of Repositorys
type RepositoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Repository `json:"items"`
}

// Repository type metadata.
var (
	Repository_Kind             = "Repository"
	Repository_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Repository_Kind}.String()
	Repository_KindAPIVersion   = Repository_Kind + "." + CRDGroupVersion.String()
	Repository_GroupVersionKind = CRDGroupVersion.WithKind(Repository_Kind)
)

func init() {
	SchemeBuilder.Register(&Repository{}, &RepositoryList{})
}
