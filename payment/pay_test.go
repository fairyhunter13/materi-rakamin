package payment

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	t.Run("success payment bank transfer", func(t *testing.T) {
		pay := New(TypeBankTransfer, "71331")
		require.NotNil(t, pay)
		require.Nil(t, pay.Validate())
		res := pay.Pay()
		require.NotEmpty(t, res)
		fmt.Println(res)
	})
}
