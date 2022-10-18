package types

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestID_MarshalJSON(t *testing.T) {
	t.Parallel()

	const want = `{"name":"test","id":"44343437426487"}`

	type field struct {
		Name string `json:"name,omitempty"`
		ID   ID     `json:"id,omitempty"`
	}

	arg := field{
		Name: "test",
		ID:   ID(44343437426487),
	}

	got, _ := json.Marshal(arg)

	t.Logf("marshaled id: %s", got)
	assert.Equal(t, want, string(got))

	var res field
	_ = json.Unmarshal(got, &res)

	t.Logf("unmarshaled struct: %v", res)
	assert.Equal(t, arg, res)

}
