package external_svc_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestExternalSvc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ExternalSvc Suite")
}
