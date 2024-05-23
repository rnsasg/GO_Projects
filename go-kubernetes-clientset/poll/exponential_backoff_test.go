package poll_test

import (
	"log/slog"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rnsasg/GO_Projects/go-kubernetes-clientset/poll"
)

var _ = Describe("ExponentialBackoff", func() {
	Context("RetryMode Timeout", func() {
		When("Within timeout", func() {
			It("Should be successfull if callback funtion return true", func() {
				config := poll.RetryConfig{
					InitialBackoff: 1 * time.Second,
					MaxBackoff:     5 * time.Second,
					Timeout:        1 * time.Minute,
					MaxRetries:     10,
					Mode:           poll.RetryTimeout,
				}
				fn := func() bool {
					slog.Info("Callback function processing ...")
					<-time.After(10 * time.Second)
					return true
				}
				ok := poll.Retry(fn, config, "Unified-Observability Commonn", "Unit Test")
				Expect(ok).To(Equal(true))
			})
		})

		When("timeout expires", func() {
			It("Should be fail if callback funtion return false", func() {
				config := poll.RetryConfig{
					InitialBackoff: 1 * time.Second,
					MaxBackoff:     5 * time.Second,
					Timeout:        1 * time.Minute,
					MaxRetries:     10,
					Mode:           poll.RetryTimeout,
				}
				fn := func() bool {
					slog.Info("Callback function processing ...")
					<-time.After(5 * time.Second)
					return false
				}
				ok := poll.Retry(fn, config, "Unified-Observability Commonn", "Unit Test")
				Expect(ok).To(Equal(false))
			})
		})
	})

	Context("RetryMode MaxRetries", func() {
		When("Within retry limit", func() {
			It("Should be successfull if callback funtion return true", func() {
				config := poll.RetryConfig{
					InitialBackoff: 1 * time.Second,
					MaxBackoff:     5 * time.Second,
					Timeout:        10 * time.Minute,
					MaxRetries:     10,
					Mode:           poll.RetryMaxRetries,
				}
				fn := func() bool {
					slog.Info("Callback function processing ...")
					<-time.After(5 * time.Second)
					return true
				}
				ok := poll.Retry(fn, config, "Unified-Observability Commonn", "Unit Test")
				Expect(ok).To(Equal(true))
			})
		})

		When("Retry limit reached", func() {
			It("Should be fail if callback funtion return false", func() {
				config := poll.RetryConfig{
					InitialBackoff: 1 * time.Second,
					MaxBackoff:     5 * time.Second,
					Timeout:        1 * time.Minute,
					MaxRetries:     10,
					Mode:           poll.RetryMaxRetries,
				}
				fn := func() bool {
					slog.Info("Callback function processing ...")
					<-time.After(5 * time.Second)
					return false
				}
				ok := poll.Retry(fn, config, "Unified-Observability Commonn", "Unit Test")
				Expect(ok).To(Equal(false))
			})
		})
	})

	It("Should fail if retry mode is invalid ", func() {
		config := poll.RetryConfig{
			InitialBackoff: 1 * time.Second,
			MaxBackoff:     5 * time.Second,
			Timeout:        1 * time.Minute,
			MaxRetries:     10,
			Mode:           3,
		}
		fn := func() bool {
			slog.Info("Callback function processing ...")
			<-time.After(5 * time.Second)
			return false
		}
		ok := poll.Retry(fn, config, "Unified-Observability Commonn", "Unit Test")
		Expect(ok).To(Equal(false))
	})

})
