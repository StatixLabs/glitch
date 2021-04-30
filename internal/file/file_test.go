package file_test

import (
	"os"

	"github.com/jakealves/glitch/internal/file"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Files", func() {

	var (
		testQuery string
		testData  []byte
		err       error
	)

	Context("when asked to write a file", func() {
		BeforeEach(func() {
			testData = []byte("SELECT * FROM NOTHING;")
			err = file.WriteContentsToFile("testWrite.sql", testData)
		})

		AfterEach(func() {
			err = os.Remove("testWrite.sql")
			Expect(err).To(Not(HaveOccurred()))
		})

		It("should not have errored", func() {
			Expect(err).To(Not(HaveOccurred()))
		})
	})

	Context("when asked to read a file", func() {
		BeforeEach(func() {
			testQuery, err = file.ReadContentsFromFile("read_test_file.sql")
		})

		It("it should return the string from the file", func() {
			Expect(testQuery).To(Equal("SELECT * FROM SOMETHING;"))
		})

		It("should not have errored", func() {
			Expect(err).To(Not(HaveOccurred()))
		})
	})
})
