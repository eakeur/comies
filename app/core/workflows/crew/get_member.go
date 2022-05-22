package crew

import (
	"context"
	"gomies/app/core/entities/member"
	"gomies/app/sdk/fault"
)

func (w workflow) GetByKey(ctx context.Context, key member.Key) (member.Member, error) {

	ct, err := w.crew.GetByKey(ctx, key)
	if err != nil {
		return member.Member{}, fault.Wrap(err)
	}

	return ct, nil
}
