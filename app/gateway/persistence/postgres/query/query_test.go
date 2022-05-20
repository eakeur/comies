package query

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuery(t *testing.T) {
	t.Parallel()

	const script = `
	select
		id, uid, auth_id, name, created_at, updated_at
	from editors
	where
		%s`

	type (
		test struct {
			name       string
			query      Query
			wantScript string
			wantArgs   []interface{}
		}
	)

	for _, c := range []test{
		{
			name: "should return script with no condition",
			query: NewQuery(script).
				Append(false, "id = $%v", "val").And().
				Append(false, "uid = $%v", "val1").Or().
				Append(false, "auth_id = $%v", "val2"),
			wantScript: fmt.Sprintf(script, ""),
			wantArgs:   []interface{}{},
		},
		{
			name: "should return script with one condition",
			query: NewQuery(script).
				Append(false, "id = $%v", "val").And().
				Append(true, "uid = $%v", "val1").Or().
				Append(false, "auth_id = $%v", "val2"),
			wantScript: fmt.Sprintf(script, "uid = $1"),
			wantArgs:   []interface{}{"val1"},
		},
		{
			name: "should return script with two conditions and operator",
			query: NewQuery(script).
				Append(true, "id = $%v", "val").And().
				Append(true, "uid = $%v", "val1").Or().
				Append(false, "auth_id = $%v", "val2"),
			wantScript: fmt.Sprintf(script, "id = $1 and uid = $2"),
			wantArgs:   []interface{}{"val", "val1"},
		},
		{
			name: "should return script with all conditions and operators",
			query: NewQuery(script).
				Append(true, "id = $%v", "val").And().
				Append(true, "uid = $%v", "val1").Or().
				Append(true, "(auth_id = $%v or auth_id = $%v)", "val2", "val3"),
			wantScript: fmt.Sprintf(script, "id = $1 and uid = $2 or (auth_id = $3 or auth_id = $4)"),
			wantArgs:   []interface{}{"val", "val1", "val2", "val3"},
		},
	} {
		c := c

		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, c.wantScript, c.query.Script())
			assert.Equal(t, c.wantArgs, c.query.Args)
		})
	}
}
