package llama_test

import (
	"os"

	"github.com/go-skynet/go-llama.cpp"
	. "github.com/go-skynet/go-llama.cpp"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("LLama binding", func() {
	Context("Declaration", func() {
		It("fails with no model", func() {
			model, err := New("not-existing")
			Expect(err).To(HaveOccurred())
			Expect(model).To(BeNil())
		})
	})
	Context("Test loading model", func() {
		modelPath := os.Getenv("MODELS_PATH")

		model, err := New(modelPath, llama.EnableF16Memory, llama.SetContext(128), llama.EnableEmbeddings)
		Expect(err).ToNot(HaveOccurred())
		Expect(model).ToNot(BeNil())
	})
})
