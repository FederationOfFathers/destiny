package destiny

import (
	"encoding/json"
	"fmt"
)

func (m *Membership) RawCharacterStats(characterID string) ([]byte, error) {
	var buf json.RawMessage
	success, err := m.client.getAndUnwrap(
		fmt.Sprintf(
			"https://www.bungie.net/platform/destiny/Stats/%d/%s/%s/",
			m.Platform,
			m.ID,
			characterID,
		),
		&buf,
	)
	if err != nil {
		return nil, err
	}
	if !success {
		return nil, fmt.Errorf("Invalid API Response")
	}
	return buf, err
}
