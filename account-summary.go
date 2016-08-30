package destiny

import (
	"encoding/json"
	"fmt"
)

func (m *Membership) RawAccountSummary() ([]byte, error) {
	var buf json.RawMessage
	success, err := m.client.getAndUnwrapData(
		fmt.Sprintf("https://www.bungie.net/platform/destiny/%d/Account/%s/Summary/", m.Platform, m.ID),
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
