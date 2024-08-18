package specifications

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/quii/go-specs-greet/acceptance-tests/dsl"
)

func SpecsForGreet(t *testing.T, greeter dsl.Greeter) {
	t.Run("when greet with empty name, then say hello world", func(t *testing.T) {
		// Act
		actual, err := greeter.Greet("")

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, "Hello, world", actual)
	})

	t.Run("when greet with name, then say hello with the name", func(t *testing.T) {
		// Act
		actual, err := greeter.Greet("Mike")

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, "Hello, Mike", actual)
	})
}

func SpecsForCurse(t *testing.T, meany dsl.Meany) {
	t.Run("when greet with name from meany, then curse with the name", func(t *testing.T) {
		// Act
		actual, err := meany.Curse("Chris")

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, "Go to hell, Chris!", actual)
	})
}
