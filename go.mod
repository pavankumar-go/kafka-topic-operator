module github.com/btrace-baader/kafka-topic-operator

go 1.16

require (
	github.com/Shopify/sarama v1.29.1
	github.com/go-logr/logr v0.4.0
	github.com/go-logr/stdr v0.4.0
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.13.0
	github.com/sirupsen/logrus v1.8.1
	github.com/smartystreets/goconvey v1.6.4
	k8s.io/api v0.21.2
	k8s.io/apimachinery v0.21.2
	k8s.io/client-go v0.21.2
	sigs.k8s.io/controller-runtime v0.9.2
)
