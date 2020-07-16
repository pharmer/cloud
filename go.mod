module go.bytebuilders.dev/resource-model

go 1.12

require (
	cloud.google.com/go v0.49.0 // indirect
	github.com/Azure/azure-sdk-for-go v36.1.0+incompatible
	github.com/Azure/go-autorest/autorest v0.9.3-0.20191028180845-3492b2aff503
	github.com/Azure/go-autorest/autorest/adal v0.8.1-0.20191028180845-3492b2aff503
	github.com/Azure/go-autorest/autorest/date v0.1.0
	github.com/Azure/go-autorest/autorest/to v0.3.1-0.20191028180845-3492b2aff503 // indirect
	github.com/Azure/go-autorest/autorest/validation v0.2.1-0.20191028180845-3492b2aff503 // indirect
	github.com/JamesClonk/vultr v2.0.1+incompatible
	github.com/appscode/go v0.0.0-20200323182826-54e98e09185a
	github.com/aws/aws-sdk-go v1.28.2
	github.com/davecgh/go-spew v1.1.1
	github.com/digitalocean/godo v1.37.0
	github.com/go-openapi/spec v0.19.3
	github.com/golang/groupcache v0.0.0-20191027212112-611e8accdfc9 // indirect
	github.com/google/gofuzz v1.1.0
	github.com/gophercloud/gophercloud v0.6.0 // indirect
	github.com/gopherjs/gopherjs v0.0.0-20191106031601-ce3c9ade29de // indirect
	github.com/hashicorp/golang-lru v0.5.3 // indirect
	github.com/linode/linodego v0.19.0
	github.com/mattn/go-isatty v0.0.9 // indirect
	github.com/moul/gotty-client v1.8.0 // indirect
	github.com/onsi/gomega v1.8.1
	github.com/packethost/packngo v0.2.0
	github.com/pborman/uuid v1.2.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.2.1 // indirect
	github.com/prometheus/common v0.7.0 // indirect
	github.com/prometheus/procfs v0.0.6 // indirect
	github.com/renstrom/fuzzysearch v0.0.0-00010101000000-000000000000 // indirect
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/scaleway/scaleway-cli v1.10.2-0.20190910153627-9984ffe9b94e
	github.com/skratchdot/open-golang v0.0.0-20200116055534-eef842397966
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.5
	go.opencensus.io v0.22.2 // indirect
	go.uber.org/atomic v1.5.0 // indirect
	golang.org/x/crypto v0.0.0-20200429183012-4b2356b1ed79 // indirect
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	golang.org/x/net v0.0.0-20200324143707-d3edc9973b7e
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	golang.org/x/sys v0.0.0-20200509044756-6aff5f38e54f // indirect
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
	google.golang.org/api v0.14.0
	gopkg.in/ini.v1 v1.51.0
	k8s.io/apimachinery v0.18.3
	k8s.io/client-go v12.0.0+incompatible
	k8s.io/kube-openapi v0.0.0-20200410145947-61e04a5be9a6
	kmodules.xyz/client-go v0.0.0-20200525195850-2fd180961371
	kmodules.xyz/crd-schema-fuzz v0.0.0-20200521005638-2433a187de95
	moul.io/anonuuid v1.3.1 // indirect
	sigs.k8s.io/controller-runtime v0.6.0
	sigs.k8s.io/yaml v1.2.0
)

replace (
	bitbucket.org/ww/goautoneg => gomodules.xyz/goautoneg v0.0.0-20120707110453-a547fc61f48d
	git.apache.org/thrift.git => github.com/apache/thrift v0.13.0
	github.com/Azure/azure-sdk-for-go => github.com/Azure/azure-sdk-for-go v35.0.0+incompatible
	github.com/Azure/go-ansiterm => github.com/Azure/go-ansiterm v0.0.0-20170929234023-d6e3b3328b78
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v13.0.0+incompatible
	github.com/Azure/go-autorest/autorest => github.com/Azure/go-autorest/autorest v0.9.0
	github.com/Azure/go-autorest/autorest/adal => github.com/Azure/go-autorest/autorest/adal v0.5.0
	github.com/Azure/go-autorest/autorest/azure/auth => github.com/Azure/go-autorest/autorest/azure/auth v0.2.0
	github.com/Azure/go-autorest/autorest/date => github.com/Azure/go-autorest/autorest/date v0.1.0
	github.com/Azure/go-autorest/autorest/mocks => github.com/Azure/go-autorest/autorest/mocks v0.2.0
	github.com/Azure/go-autorest/autorest/to => github.com/Azure/go-autorest/autorest/to v0.2.0
	github.com/Azure/go-autorest/autorest/validation => github.com/Azure/go-autorest/autorest/validation v0.1.0
	github.com/Azure/go-autorest/logger => github.com/Azure/go-autorest/logger v0.1.0
	github.com/Azure/go-autorest/tracing => github.com/Azure/go-autorest/tracing v0.5.0
	github.com/imdario/mergo => github.com/imdario/mergo v0.3.5
	github.com/prometheus/client_golang => github.com/prometheus/client_golang v1.0.0
	github.com/renstrom/fuzzysearch => github.com/lithammer/fuzzysearch v1.0.2
	go.etcd.io/etcd => go.etcd.io/etcd v0.0.0-20191023171146-3cf2f69b5738
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
	k8s.io/api => github.com/kmodules/api v0.18.4-0.20200524125823-c8bc107809b9
	k8s.io/apimachinery => github.com/kmodules/apimachinery v0.19.0-alpha.0.0.20200520235721-10b58e57a423
	k8s.io/apiserver => github.com/kmodules/apiserver v0.18.4-0.20200521000930-14c5f6df9625
	k8s.io/client-go => k8s.io/client-go v0.18.3
	k8s.io/kubernetes => github.com/kmodules/kubernetes v1.19.0-alpha.0.0.20200521033432-49d3646051ad
)
