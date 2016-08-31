package destiny

import (
	"encoding/json"
	"fmt"
)

func (m *Membership) RawAggregateActivityStats(characterID string) ([]byte, error) {
	var buf json.RawMessage
	success, err := m.client.getAndUnwrapData(
		fmt.Sprintf(
			"https://www.bungie.net/platform/destiny/Stats/AggregateActivityStats/%d/%s/%s/",
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
