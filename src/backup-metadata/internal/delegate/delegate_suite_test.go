package delegate_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDelegate(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Delegate Suite")
}
