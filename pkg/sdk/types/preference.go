package types

import (
	"encoding/json"
)

type (
	Preferences []Preference

	Preference struct {
		Module string
		Group  string
		Key    string
		Value  string
	}
)

func (p Preferences) ToJSON() (string, error) {
	marshal, err := json.Marshal(p)
	if err != nil {
		return "", err
	}

	return string(marshal), nil
}

func FromJSON(input string) (p Preferences, err error) {
	err = json.Unmarshal([]byte(input), &p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (p Preferences) GetPreference(module, group, key, defaultValue string) string {
	for _, pref := range p {
		if pref.Module == module && pref.Group == group && pref.Key == key {
			return pref.Value
		}
	}

	return defaultValue
}
