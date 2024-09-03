package types

type Modality string

const (
	Text  Modality = "text"
	Image Modality = "image"
)

type EmbeddingModel interface {
	GetVendor() string
	GetModelName() string
	GetModelDescription() string
	GetMaxInputSequenceLength() int
	GetOutputDimensions() int //TODO some models may have multiple output dimensions this should be a list
	GetLicense() string
	GetDocumentationURL() string
	GetMaxBatchSize() int
	GetModalities() []Modality
	GetURN() string //e.g. urn:embedding:1/openai/text-embedding-ada-002:1?max_batch_size=100
}
