package destiny

import (
	"encoding/json"
	"fmt"
)

func (m *Membership) RawAccountStats() ([]byte, error) {
	var buf json.RawMessage
	success, err := m.client.getAndUnwrap(
		fmt.Sprintf(
			"https://www.bungie.net/platform/destiny/Stats/Account/%d/%s/",
			m.Platform,
			m.ID,
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
