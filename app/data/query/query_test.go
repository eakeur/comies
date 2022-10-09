package query

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestQuery(t *testing.T) {
	t.Parallel()

	const script = `
	select
		id, uid, auth_id, name, created_at, updated_at
	from editors
	where
		%query%`

	const scriptNoWhere = `
	select
		id, uid, auth_id, name, created_at, updated_at
	from editors
		%where_query%`

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
				Where(false, "id = $%v", "val").And().
				Where(false, "uid = $%v", "val1").Or().
				Where(false, "auth_id = $%v", "val2"),
			wantScript: strings.Replace(script, "%query%", "", 1),
			wantArgs:   []interface{}{},
		},
		{
			name: "should return script with one condition",
			query: NewQuery(script).
				Where(false, "id = $%v", "val").And().
				Where(true, "uid = $%v", "val1").Or().
				Where(false, "auth_id = $%v", "val2"),
			wantScript: strings.Replace(script, "%query%", "uid = $1", 1),
			wantArgs:   []interface{}{"val1"},
		},
		{
			name: "should return script with two conditions and operator",
			query: NewQuery(script).
				Where(true, "id = $%v", "val").And().
				Where(true, "uid = $%v", "val1").Or().
				Where(false, "auth_id = $%v", "val2"),
			wantScript: strings.Replace(script, "%query%", "id = $1 and uid = $2", 1),
			wantArgs:   []interface{}{"val", "val1"},
		},
		{
			name: "should return script with all conditions and operators",
			query: NewQuery(script).
				Where(true, "id = $%v", "val").And().
				Where(true, "uid = $%v", "val1").Or().
				Where(true, "(auth_id = $%v or auth_id = $%v)", "val2", "val3"),
			wantScript: strings.Replace(script, "%query%", "id = $1 and uid = $2 or (auth_id = $3 or auth_id = $4)", 1),
			wantArgs:   []interface{}{"val", "val1", "val2", "val3"},
		},
		{
			name: "should return script with all conditions and operators with where",
			query: NewQuery(scriptNoWhere).
				Where(true, "id = $%v", "val").And().
				Where(true, "uid = $%v", "val1").Or().
				Where(true, "(auth_id = $%v or auth_id = $%v)", "val2", "val3"),
			wantScript: strings.Replace(scriptNoWhere, "%where_query%", "where id = $1 and uid = $2 or (auth_id = $3 or auth_id = $4)", 1),
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
