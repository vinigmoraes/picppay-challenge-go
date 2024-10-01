package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser(t *testing.T) {

	t.Run("Should return true when user is not able to transfer money", func(t *testing.T) {
		user := User{Type: COMMON}

		assert.True(t, user.IsAbleToTransferMoney())
	})

	t.Run("Should return false when user is not able to transfer money", func(t *testing.T) {
		user := User{Type: SELLER}

		assert.False(t, user.IsAbleToTransferMoney())
	})
}
