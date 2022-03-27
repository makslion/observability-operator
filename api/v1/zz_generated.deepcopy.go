// +build !ignore_autogenerated

/*


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

package v1

import (
	monitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertmanagerConfigGlobal) DeepCopyInto(out *AlertmanagerConfigGlobal) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertmanagerConfigGlobal.
func (in *AlertmanagerConfigGlobal) DeepCopy() *AlertmanagerConfigGlobal {
	if in == nil {
		return nil
	}
	out := new(AlertmanagerConfigGlobal)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertmanagerConfigReceiver) DeepCopyInto(out *AlertmanagerConfigReceiver) {
	*out = *in
	if in.PagerDutyConfigs != nil {
		in, out := &in.PagerDutyConfigs, &out.PagerDutyConfigs
		*out = make([]PagerDutyConfig, len(*in))
		copy(*out, *in)
	}
	if in.WebhookConfigs != nil {
		in, out := &in.WebhookConfigs, &out.WebhookConfigs
		*out = make([]WebhookConfig, len(*in))
		copy(*out, *in)
	}
	if in.EmailConfig != nil {
		in, out := &in.EmailConfig, &out.EmailConfig
		*out = make([]EmailConfig, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertmanagerConfigReceiver.
func (in *AlertmanagerConfigReceiver) DeepCopy() *AlertmanagerConfigReceiver {
	if in == nil {
		return nil
	}
	out := new(AlertmanagerConfigReceiver)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertmanagerConfigRoot) DeepCopyInto(out *AlertmanagerConfigRoot) {
	*out = *in
	if in.Global != nil {
		in, out := &in.Global, &out.Global
		*out = new(AlertmanagerConfigGlobal)
		**out = **in
	}
	if in.Route != nil {
		in, out := &in.Route, &out.Route
		*out = new(AlertmanagerConfigRoute)
		(*in).DeepCopyInto(*out)
	}
	if in.Receivers != nil {
		in, out := &in.Receivers, &out.Receivers
		*out = make([]AlertmanagerConfigReceiver, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertmanagerConfigRoot.
func (in *AlertmanagerConfigRoot) DeepCopy() *AlertmanagerConfigRoot {
	if in == nil {
		return nil
	}
	out := new(AlertmanagerConfigRoot)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertmanagerConfigRoute) DeepCopyInto(out *AlertmanagerConfigRoute) {
	*out = *in
	if in.Match != nil {
		in, out := &in.Match, &out.Match
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Routes != nil {
		in, out := &in.Routes, &out.Routes
		*out = make([]AlertmanagerConfigRoute, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertmanagerConfigRoute.
func (in *AlertmanagerConfigRoute) DeepCopy() *AlertmanagerConfigRoute {
	if in == nil {
		return nil
	}
	out := new(AlertmanagerConfigRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertmanagerIndex) DeepCopyInto(out *AlertmanagerIndex) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertmanagerIndex.
func (in *AlertmanagerIndex) DeepCopy() *AlertmanagerIndex {
	if in == nil {
		return nil
	}
	out := new(AlertmanagerIndex)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DexConfig) DeepCopyInto(out *DexConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DexConfig.
func (in *DexConfig) DeepCopy() *DexConfig {
	if in == nil {
		return nil
	}
	out := new(DexConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmailConfig) DeepCopyInto(out *EmailConfig) {
	*out = *in
	out.EmailHeader = in.EmailHeader
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmailConfig.
func (in *EmailConfig) DeepCopy() *EmailConfig {
	if in == nil {
		return nil
	}
	out := new(EmailConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmailSubject) DeepCopyInto(out *EmailSubject) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmailSubject.
func (in *EmailSubject) DeepCopy() *EmailSubject {
	if in == nil {
		return nil
	}
	out := new(EmailSubject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaIndex) DeepCopyInto(out *GrafanaIndex) {
	*out = *in
	if in.Dashboards != nil {
		in, out := &in.Dashboards, &out.Dashboards
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DashboardLabelSelector != nil {
		in, out := &in.DashboardLabelSelector, &out.DashboardLabelSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaIndex.
func (in *GrafanaIndex) DeepCopy() *GrafanaIndex {
	if in == nil {
		return nil
	}
	out := new(GrafanaIndex)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Observability) DeepCopyInto(out *Observability) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Observability.
func (in *Observability) DeepCopy() *Observability {
	if in == nil {
		return nil
	}
	out := new(Observability)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Observability) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObservabilityList) DeepCopyInto(out *ObservabilityList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Observability, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObservabilityList.
func (in *ObservabilityList) DeepCopy() *ObservabilityList {
	if in == nil {
		return nil
	}
	out := new(ObservabilityList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ObservabilityList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObservabilitySpec) DeepCopyInto(out *ObservabilitySpec) {
	*out = *in
	if in.ConfigurationSelector != nil {
		in, out := &in.ConfigurationSelector, &out.ConfigurationSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(Storage)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(corev1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.SelfContained != nil {
		in, out := &in.SelfContained, &out.SelfContained
		*out = new(SelfContained)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObservabilitySpec.
func (in *ObservabilitySpec) DeepCopy() *ObservabilitySpec {
	if in == nil {
		return nil
	}
	out := new(ObservabilitySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObservabilityStatus) DeepCopyInto(out *ObservabilityStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObservabilityStatus.
func (in *ObservabilityStatus) DeepCopy() *ObservabilityStatus {
	if in == nil {
		return nil
	}
	out := new(ObservabilityStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObservatoriumIndex) DeepCopyInto(out *ObservatoriumIndex) {
	*out = *in
	if in.DexConfig != nil {
		in, out := &in.DexConfig, &out.DexConfig
		*out = new(DexConfig)
		**out = **in
	}
	if in.RedhatSsoConfig != nil {
		in, out := &in.RedhatSsoConfig, &out.RedhatSsoConfig
		*out = new(RedhatSsoConfig)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObservatoriumIndex.
func (in *ObservatoriumIndex) DeepCopy() *ObservatoriumIndex {
	if in == nil {
		return nil
	}
	out := new(ObservatoriumIndex)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PagerDutyConfig) DeepCopyInto(out *PagerDutyConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PagerDutyConfig.
func (in *PagerDutyConfig) DeepCopy() *PagerDutyConfig {
	if in == nil {
		return nil
	}
	out := new(PagerDutyConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusIndex) DeepCopyInto(out *PrometheusIndex) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PodMonitors != nil {
		in, out := &in.PodMonitors, &out.PodMonitors
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.PodMonitorLabelSelector != nil {
		in, out := &in.PodMonitorLabelSelector, &out.PodMonitorLabelSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.PodMonitorNamespaceSelector != nil {
		in, out := &in.PodMonitorNamespaceSelector, &out.PodMonitorNamespaceSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.ServiceMonitorLabelSelector != nil {
		in, out := &in.ServiceMonitorLabelSelector, &out.ServiceMonitorLabelSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.ServiceMonitorNamespaceSelector != nil {
		in, out := &in.ServiceMonitorNamespaceSelector, &out.ServiceMonitorNamespaceSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.RuleLabelSelector != nil {
		in, out := &in.RuleLabelSelector, &out.RuleLabelSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.RuleNamespaceSelector != nil {
		in, out := &in.RuleNamespaceSelector, &out.RuleNamespaceSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.ProbeLabelSelector != nil {
		in, out := &in.ProbeLabelSelector, &out.ProbeLabelSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.ProbeNamespaceSelector != nil {
		in, out := &in.ProbeNamespaceSelector, &out.ProbeNamespaceSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusIndex.
func (in *PrometheusIndex) DeepCopy() *PrometheusIndex {
	if in == nil {
		return nil
	}
	out := new(PrometheusIndex)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PromtailIndex) DeepCopyInto(out *PromtailIndex) {
	*out = *in
	if in.NamespaceLabelSelector != nil {
		in, out := &in.NamespaceLabelSelector, &out.NamespaceLabelSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.DaemonSetLabelSelector != nil {
		in, out := &in.DaemonSetLabelSelector, &out.DaemonSetLabelSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PromtailIndex.
func (in *PromtailIndex) DeepCopy() *PromtailIndex {
	if in == nil {
		return nil
	}
	out := new(PromtailIndex)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedhatSsoConfig) DeepCopyInto(out *RedhatSsoConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedhatSsoConfig.
func (in *RedhatSsoConfig) DeepCopy() *RedhatSsoConfig {
	if in == nil {
		return nil
	}
	out := new(RedhatSsoConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RemoteWriteIndex) DeepCopyInto(out *RemoteWriteIndex) {
	*out = *in
	if in.QueueConfig != nil {
		in, out := &in.QueueConfig, &out.QueueConfig
		*out = new(monitoringv1.QueueConfig)
		**out = **in
	}
	if in.WriteRelabelConfigs != nil {
		in, out := &in.WriteRelabelConfigs, &out.WriteRelabelConfigs
		*out = make([]monitoringv1.RelabelConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemoteWriteIndex.
func (in *RemoteWriteIndex) DeepCopy() *RemoteWriteIndex {
	if in == nil {
		return nil
	}
	out := new(RemoteWriteIndex)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryConfig) DeepCopyInto(out *RepositoryConfig) {
	*out = *in
	if in.Grafana != nil {
		in, out := &in.Grafana, &out.Grafana
		*out = new(GrafanaIndex)
		(*in).DeepCopyInto(*out)
	}
	if in.Prometheus != nil {
		in, out := &in.Prometheus, &out.Prometheus
		*out = new(PrometheusIndex)
		(*in).DeepCopyInto(*out)
	}
	if in.Alertmanager != nil {
		in, out := &in.Alertmanager, &out.Alertmanager
		*out = new(AlertmanagerIndex)
		**out = **in
	}
	if in.Promtail != nil {
		in, out := &in.Promtail, &out.Promtail
		*out = new(PromtailIndex)
		(*in).DeepCopyInto(*out)
	}
	if in.Observatoria != nil {
		in, out := &in.Observatoria, &out.Observatoria
		*out = make([]ObservatoriumIndex, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryConfig.
func (in *RepositoryConfig) DeepCopy() *RepositoryConfig {
	if in == nil {
		return nil
	}
	out := new(RepositoryConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryIndex) DeepCopyInto(out *RepositoryIndex) {
	*out = *in
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(corev1.Secret)
		(*in).DeepCopyInto(*out)
	}
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(RepositoryConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryIndex.
func (in *RepositoryIndex) DeepCopy() *RepositoryIndex {
	if in == nil {
		return nil
	}
	out := new(RepositoryIndex)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryInfo) DeepCopyInto(out *RepositoryInfo) {
	*out = *in
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(corev1.Secret)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryInfo.
func (in *RepositoryInfo) DeepCopy() *RepositoryInfo {
	if in == nil {
		return nil
	}
	out := new(RepositoryInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelfContained) DeepCopyInto(out *SelfContained) {
	*out = *in
	if in.DisableRepoSync != nil {
		in, out := &in.DisableRepoSync, &out.DisableRepoSync
		*out = new(bool)
		**out = **in
	}
	if in.DisableObservatorium != nil {
		in, out := &in.DisableObservatorium, &out.DisableObservatorium
		*out = new(bool)
		**out = **in
	}
	if in.DisablePagerDuty != nil {
		in, out := &in.DisablePagerDuty, &out.DisablePagerDuty
		*out = new(bool)
		**out = **in
	}
	if in.DisableDeadmansSnitch != nil {
		in, out := &in.DisableDeadmansSnitch, &out.DisableDeadmansSnitch
		*out = new(bool)
		**out = **in
	}
	if in.DisableSmtp != nil {
		in, out := &in.DisableSmtp, &out.DisableSmtp
		*out = new(bool)
		**out = **in
	}
	if in.DisableBlackboxExporter != nil {
		in, out := &in.DisableBlackboxExporter, &out.DisableBlackboxExporter
		*out = new(bool)
		**out = **in
	}
	if in.SelfSignedCerts != nil {
		in, out := &in.SelfSignedCerts, &out.SelfSignedCerts
		*out = new(bool)
		**out = **in
	}
	if in.OverrideSelectors != nil {
		in, out := &in.OverrideSelectors, &out.OverrideSelectors
		*out = new(bool)
		**out = **in
	}
	if in.FederatedMetrics != nil {
		in, out := &in.FederatedMetrics, &out.FederatedMetrics
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PodMonitorLabelSelector != nil {
		in, out := &in.PodMonitorLabelSelector, &out.PodMonitorLabelSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.PodMonitorNamespaceSelector != nil {
		in, out := &in.PodMonitorNamespaceSelector, &out.PodMonitorNamespaceSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.ServiceMonitorLabelSelector != nil {
		in, out := &in.ServiceMonitorLabelSelector, &out.ServiceMonitorLabelSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.ServiceMonitorNamespaceSelector != nil {
		in, out := &in.ServiceMonitorNamespaceSelector, &out.ServiceMonitorNamespaceSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.RuleLabelSelector != nil {
		in, out := &in.RuleLabelSelector, &out.RuleLabelSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.RuleNamespaceSelector != nil {
		in, out := &in.RuleNamespaceSelector, &out.RuleNamespaceSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.ProbeLabelSelector != nil {
		in, out := &in.ProbeLabelSelector, &out.ProbeLabelSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.ProbeNamespaceSelector != nil {
		in, out := &in.ProbeNamespaceSelector, &out.ProbeNamespaceSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.GrafanaDashboardLabelSelector != nil {
		in, out := &in.GrafanaDashboardLabelSelector, &out.GrafanaDashboardLabelSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	in.AlertManagerResourceRequirement.DeepCopyInto(&out.AlertManagerResourceRequirement)
	in.PrometheusResourceRequirement.DeepCopyInto(&out.PrometheusResourceRequirement)
	in.PrometheusOperatorResourceRequirement.DeepCopyInto(&out.PrometheusOperatorResourceRequirement)
	if in.GrafanaResourceRequirement != nil {
		in, out := &in.GrafanaResourceRequirement, &out.GrafanaResourceRequirement
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	in.GrafanaOperatorResourceRequirement.DeepCopyInto(&out.GrafanaOperatorResourceRequirement)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelfContained.
func (in *SelfContained) DeepCopy() *SelfContained {
	if in == nil {
		return nil
	}
	out := new(SelfContained)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Storage) DeepCopyInto(out *Storage) {
	*out = *in
	if in.PrometheusStorageSpec != nil {
		in, out := &in.PrometheusStorageSpec, &out.PrometheusStorageSpec
		*out = new(monitoringv1.StorageSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Storage.
func (in *Storage) DeepCopy() *Storage {
	if in == nil {
		return nil
	}
	out := new(Storage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookConfig) DeepCopyInto(out *WebhookConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookConfig.
func (in *WebhookConfig) DeepCopy() *WebhookConfig {
	if in == nil {
		return nil
	}
	out := new(WebhookConfig)
	in.DeepCopyInto(out)
	return out
}
