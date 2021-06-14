package utils_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jakealves/glitch/lib/utils"
)

var _ = Describe("Running the ForceKeysToUppercase function", func() {
	Context("when provided a map with mixed case", func() {
		startingStruct := map[string]interface{}{
			"STRING":  "delicious",
			"InT":     64,
			"boolean": true,
		}
		It("all keys should be forced to lowercase", func() {
			endingStruct := map[string]interface{}{
				"string":  "delicious",
				"int":     64,
				"boolean": true,
			}
			Expect(utils.ForceKeysToLowercase(startingStruct)).To(Equal(endingStruct))
		})
	})
})
