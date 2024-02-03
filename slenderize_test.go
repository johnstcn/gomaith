package gomaith_test

import (
	"fmt"
	"testing"

	"github.com/johnstcn/gomaith"
	"github.com/stretchr/testify/assert"
)

func ExampleSlenderize() {
	fmt.Println(gomaith.Slenderize("bonn"))
	// output: boinn
}

func TestSlenderize(t *testing.T) {
	t.Parallel()

	t.Run("Empty", func(t *testing.T) {
		assert.Empty(t, gomaith.Slenderize(""))
	})

	t.Run("NoEndingConsonant", func(t *testing.T) {
		actual := gomaith.Slenderize("dara")
		assert.Equal(t, "dara", actual)
	})

	t.Run("EndingBroadConsonant", func(t *testing.T) {
		actual := gomaith.Slenderize("bonn")
		assert.Equal(t, "boinn", actual)
	})
}
