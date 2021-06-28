package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ResourceOfferSpec defines the desired state of ResourceOffer.
type ResourceOfferSpec struct {
	// ClusterId is the identifier of the cluster that is sending this ResourceOffer.
	// It is the uid of the first master node in you cluster.
	ClusterId string `json:"clusterId"`
	// Images is the list of the images already stored in the cluster.
	Images []corev1.ContainerImage `json:"images,omitempty"`
	// ResourceQuota contains the quantity of resources made available by the cluster.
	ResourceQuota corev1.ResourceQuotaSpec `json:"resourceQuota,omitempty"`
	// Labels contains the label to be added to the virtual node.
	Labels map[string]string `json:"labels,omitempty"`
	// Prices contains the possible prices for every kind of resource (cpu, memory, image).
	Prices corev1.ResourceList `json:"prices,omitempty"`
	// Timestamp is the time instant when this ResourceOffer was created.
	Timestamp metav1.Time `json:"timestamp"`
	// TimeToLive is the time instant until this ResourceOffer will be valid.
	// If not refreshed, an ResourceOffer will expire after 30 minutes.
	TimeToLive metav1.Time `json:"timeToLive"`
	// WithdrawalTimestamp is set when a graceful deletion is requested by the user.
	WithdrawalTimestamp *metav1.Time `json:"withdrawalTimestamp,omitempty"`
}

// OfferPhase describes the phase of the ResourceOffer.
type OfferPhase string

const (
	// ResourceOfferPending indicates a pending phase, an action is required.
	ResourceOfferPending OfferPhase = "Pending"
	// ResourceOfferManualActionRequired indicates that a manual action is required.
	ResourceOfferManualActionRequired OfferPhase = "ManualActionRequired"
	// ResourceOfferAccepted indicates an accepted offer.
	ResourceOfferAccepted OfferPhase = "Accepted"
	// ResourceOfferRefused indicates a refused offer.
	ResourceOfferRefused OfferPhase = "Refused"
)

// VirtualKubeletStatus indicates the observed status of the VirtualKubelet Deployment.
type VirtualKubeletStatus string

const (
	// VirtualKubeletStatusNone indicates that there is no VirtualKubelet Deployment.
	VirtualKubeletStatusNone VirtualKubeletStatus = "None"
	// VirtualKubeletStatusCreated indicates that the VirtualKubelet Deployment has been created.
	VirtualKubeletStatusCreated VirtualKubeletStatus = "Created"
	// VirtualKubeletStatusDeleting indicates that the VirtualKubelet Deployment is deleting.
	VirtualKubeletStatusDeleting VirtualKubeletStatus = "Deleting"
)

// ResourceOfferStatus defines the observed state of ResourceOffer.
type ResourceOfferStatus struct {
	// Phase is the status of this ResourceOffer.
	// When the offer is created it is checked by the operator, which sets this field to "Accepted" or "Refused" on tha base of cluster configuration.
	// If the ResourceOffer is accepted a virtual-kubelet for the foreign cluster will be created.
	// +kubebuilder:validation:Enum="Pending";"ManualActionRequired";"Accepted";"Refused"
	// +kubebuilder:default="Pending"
	Phase OfferPhase `json:"phase"`
	// VirtualKubeletStatus indicates if the virtual-kubelet for this ResourceOffer has been created or not.
	// +kubebuilder:validation:Enum="None";"Created";"Deleting"
	// +kubebuilder:default="None"
	VirtualKubeletStatus VirtualKubeletStatus `json:"virtualKubeletStatus,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName="offer"

// ResourceOffer is the Schema for the resourceOffers API.
// +kubebuilder:printcolumn:name="Status",type=string,JSONPath=`.status.phase`
// +kubebuilder:printcolumn:name="Expiration",type=string,JSONPath=`.spec.timeToLive`
// +kubebuilder:printcolumn:name="VirtualKubeletStatus",type=string,JSONPath=`.status.virtualKubeletStatus`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
type ResourceOffer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ResourceOfferSpec   `json:"spec,omitempty"`
	Status ResourceOfferStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceOfferList contains a list of ResourceOffer.
type ResourceOfferList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourceOffer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ResourceOffer{}, &ResourceOfferList{})
}
