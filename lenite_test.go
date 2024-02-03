package gomaith_test

import (
	"github.com/johnstcn/gomaith"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLenite(t *testing.T) {
	t.Parallel()

	t.Run("Empty", func(t *testing.T) {
		actual := gomaith.Lenite("")
		assert.Equal(t, "", actual)
	})
	t.Run("TooShort", func(t *testing.T) {
		actual := gomaith.Lenite("b")
		assert.Equal(t, "b", actual)
	})
	t.Run("NonLenitable", func(t *testing.T) {
		actual := gomaith.Lenite("éan")
		assert.Equal(t, "éan", actual)
	})
	t.Run("Lenitable", func(t *testing.T) {
		actual := gomaith.Lenite("ceann")
		assert.Equal(t, "cheann", actual)
	})
}
