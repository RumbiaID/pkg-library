package htmlsanitizer

import "github.com/microcosm-cc/bluemonday"

func NewSanitizer() *bluemonday.Policy {
	policy := bluemonday.UGCPolicy()

	return policy
}
