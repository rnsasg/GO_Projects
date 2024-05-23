package k8s_test

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	common "github.com/rnsasg/GO_Projects/go-kubernetes-clientset"
	k8s "github.com/rnsasg/GO_Projects/go-kubernetes-clientset/k8s"
	"k8s.io/client-go/kubernetes/fake"
)

var _ = Describe("Configmap", func() {

	When("Configmap is not avilable", func() {
		It("Should retrun false", func() {
			// Create a fake.Clientset
			checker := k8s.NewConfigMapChecker(fake.NewSimpleClientset())
			ok := checker.IsConfigMapAvailable(common.ObservabilityConfigMap, common.TanzuSystemNamespace, "Unit Test", 20*time.Second, common.MaxRetries)
			Expect(ok).To(Equal(false))
		})
	})
	When("Configmap is avilable", func() {
		It("Should retrun true", func() {
			checker := k8s.NewConfigMapChecker(fake.NewSimpleClientset())
			_, err := checker.CreateObservabilityConfigmap(1 * time.Second)
			Expect(err).ToNot(HaveOccurred())
			ok := checker.IsConfigMapAvailable(common.ObservabilityConfigMap, common.TanzuSystemNamespace, "Unit Test", 20*time.Second, common.MaxRetries)
			Expect(ok).To(Equal(true))
		})
	})

	When("Configmap is become avialable after some time", func() {
		It("Should retrun true", func() {
			checker := k8s.NewConfigMapChecker(fake.NewSimpleClientset())
			go checker.CreateObservabilityConfigmap(10 * time.Second)
			ok := checker.IsConfigMapAvailable(common.ObservabilityConfigMap, common.TanzuSystemNamespace, "Unit Test", 30*time.Second, common.MaxRetries)
			Expect(ok).To(Equal(true))
		})
	})
})
