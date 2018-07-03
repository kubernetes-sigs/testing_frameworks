/*
Copyright 2016 The Kubernetes Authors.

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

package base

// TODO 	fuzz "github.com/google/gofuzz"

// TODO 	v1 "k8s.io/api/core/v1"
// TODO 	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
// TODO 	kubeletconfigv1beta1 "k8s.io/kubernetes/pkg/kubelet/apis/kubeletconfig/v1beta1"
// TODO 	kubeproxyconfigv1alpha1 "k8s.io/kubernetes/pkg/proxy/apis/kubeproxyconfig/v1alpha1"

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MasterConfiguration contains a list of elements which make up master's
// configuration object.
type MasterConfiguration struct {
	// TODO 	metav1.TypeMeta

	// `kubeadm init`-only information. These fields are solely used the first time `kubeadm init` runs.
	// After that, the information in the fields ARE NOT uploaded to the `kubeadm-config` ConfigMap
	// that is used by `kubeadm upgrade` for instance.

	// BootstrapTokens is respected at `kubeadm init` time and describes a set of Bootstrap Tokens to create.
	// This information IS NOT uploaded to the kubeadm cluster configmap, partly because of its sensitive nature
	BootstrapTokens []BootstrapToken

	// NodeRegistration holds fields that relate to registering the new master node to the cluster
	NodeRegistration NodeRegistrationOptions

	// Cluster-wide configuration
	// TODO: Move these fields under some kind of ClusterConfiguration or similar struct that describes
	// one cluster. Eventually we want this kind of spec to align well with the Cluster API spec.

	// API holds configuration for the k8s apiserver.
	API API
	// KubeProxy holds configuration for the k8s service proxy.
	KubeProxy KubeProxy
	// Etcd holds configuration for etcd.
	Etcd Etcd
	// KubeletConfiguration holds configuration for the kubelet.
	KubeletConfiguration KubeletConfiguration
	// Networking holds configuration for the networking topology of the cluster.
	Networking Networking
	// KubernetesVersion is the target version of the control plane.
	KubernetesVersion string

	// APIServerExtraArgs is a set of extra flags to pass to the API Server or override
	// default ones in form of <flagname>=<value>.
	// TODO: This is temporary and ideally we would like to switch all components to
	// use ComponentConfig + ConfigMaps.
	APIServerExtraArgs map[string]string
	// ControllerManagerExtraArgs is a set of extra flags to pass to the Controller Manager
	// or override default ones in form of <flagname>=<value>
	// TODO: This is temporary and ideally we would like to switch all components to
	// use ComponentConfig + ConfigMaps.
	ControllerManagerExtraArgs map[string]string
	// SchedulerExtraArgs is a set of extra flags to pass to the Scheduler or override
	// default ones in form of <flagname>=<value>
	// TODO: This is temporary and ideally we would like to switch all components to
	// use ComponentConfig + ConfigMaps.
	SchedulerExtraArgs map[string]string

	// APIServerExtraVolumes is an extra set of host volumes mounted to the API server.
	APIServerExtraVolumes []HostPathMount
	// ControllerManagerExtraVolumes is an extra set of host volumes mounted to the
	// Controller Manager.
	ControllerManagerExtraVolumes []HostPathMount
	// SchedulerExtraVolumes is an extra set of host volumes mounted to the scheduler.
	SchedulerExtraVolumes []HostPathMount

	// APIServerCertSANs sets extra Subject Alternative Names for the API Server
	// signing cert.
	APIServerCertSANs []string
	// CertificatesDir specifies where to store or look for all required certificates.
	CertificatesDir string

	// ImageRepository is the container registry to pull control plane images from.
	ImageRepository string

	// CIImageRepository is the container registry for core images generated by CI.
	// Useful for running kubeadm with images from CI builds.
	// +k8s:conversion-gen=false
	CIImageRepository string

	// UnifiedControlPlaneImage specifies if a specific container image should be
	// used for all control plane components.
	UnifiedControlPlaneImage string

	// AuditPolicyConfiguration defines the options for the api server audit system.
	AuditPolicyConfiguration AuditPolicyConfiguration

	// FeatureGates enabled by the user.
	FeatureGates map[string]bool

	// The cluster name
	ClusterName string
}

// API struct contains elements of API server address.
type API struct {
	// AdvertiseAddress sets the IP address for the API server to advertise.
	AdvertiseAddress string
	// ControlPlaneEndpoint sets a stable IP address or DNS name for the control plane; it
	// can be a valid IP address or a RFC-1123 DNS subdomain, both with optional TCP port.
	// In case the ControlPlaneEndpoint is not specified, the AdvertiseAddress + BindPort
	// are used; in case the ControlPlaneEndpoint is specified but without a TCP port,
	// the BindPort is used.
	// Possible usages are:
	// e.g. In an cluster with more than one control plane instances, this field should be
	// assigned the address of the external load balancer in front of the
	// control plane instances.
	// e.g.  in environments with enforced node recycling, the ControlPlaneEndpoint
	// could be used for assigning a stable DNS to the control plane.
	ControlPlaneEndpoint string
	// BindPort sets the secure port for the API Server to bind to.
	// Defaults to 6443.
	BindPort int32
}

// NodeRegistrationOptions holds fields that relate to registering a new master or node to the cluster, either via "kubeadm init" or "kubeadm join"
type NodeRegistrationOptions struct {

	// Name is the `.Metadata.Name` field of the Node API object that will be created in this `kubeadm init` or `kubeadm joiń` operation.
	// This field is also used in the CommonName field of the kubelet's client certificate to the API server.
	// Defaults to the hostname of the node if not provided.
	Name string

	// CRISocket is used to retrieve container runtime info. This information will be annotated to the Node API object, for later re-use
	CRISocket string

	// Taints specifies the taints the Node API object should be registered with. If this field is unset, i.e. nil, in the `kubeadm init` process
	// TODO 	// it will be defaulted to []v1.Taint{'node-role.kubernetes.io/master=""'}. If you don't want to taint your master node, set this field to an
	// empty slice, i.e. `taints: {}` in the YAML file. This field is solely used for Node registration.
	// TODO 	Taints []v1.Taint

	// KubeletExtraArgs passes through extra arguments to the kubelet. The arguments here are passed to the kubelet command line via the environment file
	// kubeadm writes at runtime for the kubelet to source. This overrides the generic base-level configuration in the kubelet-config-1.X ConfigMap
	// Flags have higher higher priority when parsing. These values are local and specific to the node kubeadm is executing on.
	KubeletExtraArgs map[string]string
}

// Networking contains elements describing cluster's networking configuration.
type Networking struct {
	// ServiceSubnet is the subnet used by k8s services. Defaults to "10.96.0.0/12".
	ServiceSubnet string
	// PodSubnet is the subnet used by pods.
	PodSubnet string
	// DNSDomain is the dns domain used by k8s services. Defaults to "cluster.local".
	DNSDomain string
}

// BootstrapToken describes one bootstrap token, stored as a Secret in the cluster
// TODO: The BootstrapToken object should move out to either k8s.io/client-go or k8s.io/api in the future
// (probably as part of Bootstrap Tokens going GA). It should not be staged under the kubeadm API as it is now.
type BootstrapToken struct {
	// Token is used for establishing bidirectional trust between nodes and masters.
	// Used for joining nodes in the cluster.
	Token *BootstrapTokenString
	// Description sets a human-friendly message why this token exists and what it's used
	// for, so other administrators can know its purpose.
	Description string
	// TTL defines the time to live for this token. Defaults to 24h.
	// Expires and TTL are mutually exclusive.
	// TODO 	TTL *metav1.Duration
	// Expires specifies the timestamp when this token expires. Defaults to being set
	// dynamically at runtime based on the TTL. Expires and TTL are mutually exclusive.
	// TODO 	Expires *metav1.Time
	// Usages describes the ways in which this token can be used. Can by default be used
	// for establishing bidirectional trust, but that can be changed here.
	Usages []string
	// Groups specifies the extra groups that this token will authenticate as when/if
	// used for authentication
	Groups []string
}

// Etcd contains elements describing Etcd configuration.
type Etcd struct {

	// Local provides configuration knobs for configuring the local etcd instance
	// Local and External are mutually exclusive
	Local *LocalEtcd

	// External describes how to connect to an external etcd cluster
	// Local and External are mutually exclusive
	External *ExternalEtcd
}

// Fuzz is a dummy function here to get the roundtrip tests working in cmd/kubeadm/app/apis/kubeadm/fuzzer working.
// As we split the monolith-etcd struct into two smaller pieces with pointers and they are mutually exclusive, roundtrip
// tests that randomize all values in this struct isn't feasible. Instead, we override the fuzzing function for .Etcd with
// this func by letting Etcd implement the fuzz.Interface interface. As this func does nothing, we rely on the values given
// in fuzzer/fuzzer.go for the roundtrip tests, which is exactly what we want.
// TODO: Remove this function when we remove the v1alpha1 API
// TODO func (e Etcd) Fuzz(c fuzz.Continue) {}

// LocalEtcd describes that kubeadm should run an etcd cluster locally
type LocalEtcd struct {

	// Image specifies which container image to use for running etcd.
	// If empty, automatically populated by kubeadm using the image
	// repository and default etcd version.
	Image string

	// DataDir is the directory etcd will place its data.
	// Defaults to "/var/lib/etcd".
	DataDir string

	// ExtraArgs are extra arguments provided to the etcd binary
	// when run inside a static pod.
	ExtraArgs map[string]string

	// ServerCertSANs sets extra Subject Alternative Names for the etcd server signing cert.
	ServerCertSANs []string
	// PeerCertSANs sets extra Subject Alternative Names for the etcd peer signing cert.
	PeerCertSANs []string
}

// ExternalEtcd describes an external etcd cluster
type ExternalEtcd struct {

	// Endpoints of etcd members. Useful for using external etcd.
	// If not provided, kubeadm will run etcd in a static pod.
	Endpoints []string
	// CAFile is an SSL Certificate Authority file used to secure etcd communication.
	CAFile string
	// CertFile is an SSL certification file used to secure etcd communication.
	CertFile string
	// KeyFile is an SSL key file used to secure etcd communication.
	KeyFile string
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NodeConfiguration contains elements describing a particular node.
// TODO: This struct should be replaced by dynamic kubelet configuration.
type NodeConfiguration struct {
	// TODO 	metav1.TypeMeta

	// NodeRegistration holds fields that relate to registering the new master node to the cluster
	NodeRegistration NodeRegistrationOptions

	// CACertPath is the path to the SSL certificate authority used to
	// secure comunications between node and master.
	// Defaults to "/etc/kubernetes/pki/ca.crt".
	CACertPath string
	// DiscoveryFile is a file or url to a kubeconfig file from which to
	// load cluster information.
	DiscoveryFile string
	// DiscoveryToken is a token used to validate cluster information
	// fetched from the master.
	DiscoveryToken string
	// DiscoveryTokenAPIServers is a set of IPs to API servers from which info
	// will be fetched. Currently we only pay attention to one API server but
	// hope to support >1 in the future.
	DiscoveryTokenAPIServers []string
	// DiscoveryTimeout modifies the discovery timeout
	// TODO 	DiscoveryTimeout *metav1.Duration
	// TLSBootstrapToken is a token used for TLS bootstrapping.
	// Defaults to Token.
	TLSBootstrapToken string
	// Token is used for both discovery and TLS bootstrapping.
	Token string
	// The cluster name
	ClusterName string

	// DiscoveryTokenCACertHashes specifies a set of public key pins to verify
	// when token-based discovery is used. The root CA found during discovery
	// must match one of these values. Specifying an empty set disables root CA
	// pinning, which can be unsafe. Each hash is specified as "<type>:<value>",
	// where the only currently supported type is "sha256". This is a hex-encoded
	// SHA-256 hash of the Subject Public Key Info (SPKI) object in DER-encoded
	// ASN.1. These hashes can be calculated using, for example, OpenSSL:
	// openssl x509 -pubkey -in ca.crt openssl rsa -pubin -outform der 2>&/dev/null | openssl dgst -sha256 -hex
	DiscoveryTokenCACertHashes []string

	// DiscoveryTokenUnsafeSkipCAVerification allows token-based discovery
	// without CA verification via DiscoveryTokenCACertHashes. This can weaken
	// the security of kubeadm since other nodes can impersonate the master.
	DiscoveryTokenUnsafeSkipCAVerification bool

	// FeatureGates enabled by the user.
	FeatureGates map[string]bool
}

// KubeletConfiguration contains elements describing initial remote configuration of kubelet.
type KubeletConfiguration struct {
	// TODO 	BaseConfig *kubeletconfigv1beta1.KubeletConfiguration
}

// GetControlPlaneImageRepository returns name of image repository
// for control plane images (API,Controller Manager,Scheduler and Proxy)
// It will override location with CI registry name in case user requests special
// Kubernetes version from CI build area.
// (See: kubeadmconstants.DefaultCIImageRepository)
func (cfg *MasterConfiguration) GetControlPlaneImageRepository() string {
	if cfg.CIImageRepository != "" {
		return cfg.CIImageRepository
	}
	return cfg.ImageRepository
}

// HostPathMount contains elements describing volumes that are mounted from the
// host.
type HostPathMount struct {
	// Name of the volume inside the pod template.
	Name string
	// HostPath is the path in the host that will be mounted inside
	// the pod.
	HostPath string
	// MountPath is the path inside the pod where hostPath will be mounted.
	MountPath string
	// Writable controls write access to the volume
	Writable bool
	// PathType is the type of the HostPath.
	// TODO 	PathType v1.HostPathType
}

// KubeProxy contains elements describing the proxy configuration.
type KubeProxy struct {
	// TODO 	Config *kubeproxyconfigv1alpha1.KubeProxyConfiguration
}

// AuditPolicyConfiguration holds the options for configuring the api server audit policy.
type AuditPolicyConfiguration struct {
	// Path is the local path to an audit policy.
	Path string
	// LogDir is the local path to the directory where logs should be stored.
	LogDir string
	// LogMaxAge is the number of days logs will be stored for. 0 indicates forever.
	LogMaxAge *int32
	//TODO(chuckha) add other options for audit policy.
}

// CommonConfiguration defines the list of common configuration elements and the getter
// methods that must exist for both the MasterConfiguration and NodeConfiguration objects.
// This is used internally to deduplicate the kubeadm preflight checks.
type CommonConfiguration interface {
	GetCRISocket() string
	GetNodeName() string
	GetKubernetesVersion() string
}

// GetCRISocket will return the CRISocket that is defined for the MasterConfiguration.
// This is used internally to deduplicate the kubeadm preflight checks.
func (cfg *MasterConfiguration) GetCRISocket() string {
	return cfg.NodeRegistration.CRISocket
}

// GetNodeName will return the NodeName that is defined for the MasterConfiguration.
// This is used internally to deduplicate the kubeadm preflight checks.
func (cfg *MasterConfiguration) GetNodeName() string {
	return cfg.NodeRegistration.Name
}

// GetKubernetesVersion will return the KubernetesVersion that is defined for the MasterConfiguration.
// This is used internally to deduplicate the kubeadm preflight checks.
func (cfg *MasterConfiguration) GetKubernetesVersion() string {
	return cfg.KubernetesVersion
}

// GetCRISocket will return the CRISocket that is defined for the NodeConfiguration.
// This is used internally to deduplicate the kubeadm preflight checks.
func (cfg *NodeConfiguration) GetCRISocket() string {
	return cfg.NodeRegistration.CRISocket
}

// GetNodeName will return the NodeName that is defined for the NodeConfiguration.
// This is used internally to deduplicate the kubeadm preflight checks.
func (cfg *NodeConfiguration) GetNodeName() string {
	return cfg.NodeRegistration.Name
}

// GetKubernetesVersion will return an empty string since KubernetesVersion is not a
// defined property for NodeConfiguration. This will just cause the regex validation
// of the defined version to be skipped during the preflight checks.
// This is used internally to deduplicate the kubeadm preflight checks.
func (cfg *NodeConfiguration) GetKubernetesVersion() string {
	return ""
}
