package material_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMaterial(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Material Suite")
}
