package poll_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestPoll(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Poll Suite")
}
