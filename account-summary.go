package destiny

import (
	"fmt"
	"io"
	"os"
)

type AccountSummary struct {
}

func (m *Membership) AccountSummary() (*AccountSummary, error) {
	rsp, err := m.client.Get(fmt.Sprintf("https://www.bungie.net/platform/destiny/%d/Account/%s/Summary/", m.Platform, m.ID))
	if err != nil {
		return nil, err
	}
	io.Copy(os.Stdout, rsp.Body)
	return nil, nil
}
