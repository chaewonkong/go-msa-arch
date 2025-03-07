package app_test

import (
	. "github.com/chaewonkong/go-msa-arch/internal/app"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadConfig(t *testing.T) {
	t.Run("Not panics if config is nil", func(t *testing.T) {
		assert.NotPanics(t, func() {
			err := ReadConfig(nil)
			assert.Error(t, err)
		})
	})
}
