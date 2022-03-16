package preferences

import "fmt"

type Preferences map[string]string

func (d Preferences) GetPreference(module, group string, keys ...string) (results map[string]string) {
	parent := fmt.Sprintf("%s/%s/", module, group)

	for detailKey, detail := range d {
		for _, key := range keys {
			if detailKey == parent+key {
				results[key] = detail
			}
		}
	}

	return results
}
