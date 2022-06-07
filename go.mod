module github.com/mvavilov/observability-operator/v3

go 1.16

require (
	github.com/Shopify/logrus-bugsnag v0.0.0-20171204204709-577dee27f20d // indirect
	github.com/blang/semver v3.5.1+incompatible
	github.com/bshuster-repo/logrus-logstash-hook v1.0.2 // indirect
	github.com/bugsnag/bugsnag-go v2.1.2+incompatible // indirect
	github.com/bugsnag/panicwrap v1.3.4 // indirect
	github.com/docker/go-metrics v0.0.1 // indirect
	github.com/docker/libtrust v0.0.0-20160708172513-aabc10ec26b7 // indirect
	github.com/garyburd/redigo v1.6.3 // indirect
	github.com/ghodss/yaml v1.0.0
	github.com/go-logr/logr v0.2.1
	github.com/goccy/go-yaml v1.9.5
	github.com/gorilla/handlers v1.5.1 // indirect
	github.com/integr8ly/grafana-operator/v3 v3.10.3
	github.com/onsi/ginkgo v1.14.0
	github.com/onsi/gomega v1.10.1
	github.com/openshift/api v3.9.1-0.20190924102528-32369d4db2ad+incompatible
	github.com/operator-framework/api v0.3.20
	github.com/operator-framework/operator-registry v1.12.6-0.20200611222234-275301b779f8
	github.com/pkg/errors v0.9.1
	github.com/prometheus-operator/prometheus-operator v0.43.0
	github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring v0.43.0
	github.com/redhat-developer/observability-operator/v3 v3.0.10
	github.com/sirupsen/logrus v1.8.1
	github.com/yvasiyarov/go-metrics v0.0.0-20150112132944-c25f46c4b940 // indirect
	github.com/yvasiyarov/gorelic v0.0.7 // indirect
	github.com/yvasiyarov/newrelic_platform_go v0.0.0-20160601141957-9c099fbc30e9 // indirect
	k8s.io/api v0.19.2
	k8s.io/apimachinery v0.19.2
	k8s.io/client-go v12.0.0+incompatible
	sigs.k8s.io/controller-runtime v0.6.2
)

replace k8s.io/client-go => k8s.io/client-go v0.19.2

replace github.com/opencontainers/runc => github.com/opencontainers/runc v1.0.3

replace github.com/containerd/containerd => github.com/containerd/containerd v1.4.13

replace github.com/gogo/protobuf => github.com/gogo/protobuf v1.3.2

replace github.com/docker/distribution => github.com/docker/distribution v2.8.0+incompatible

replace go.mongodb.org/mongo-driver => go.mongodb.org/mongo-driver v1.5.1

replace github.com/opencontainers/image-spec => github.com/opencontainers/image-spec v1.0.2
