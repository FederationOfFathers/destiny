package destiny

import (
	"fmt"
	"strings"
)

// Member structs are good for serializing; into a database for example
type Member struct {
	Name     string
	ID       string
	Icon     string
	Platform int
}

// Turn a member record into a membership
func (m Member) Membership(c *Client) *Membership {
	return &Membership{
		client:   c,
		Name:     m.Name,
		Platform: m.Platform,
		ID:       m.ID,
		Icon:     m.Icon,
	}
}

type Membership struct {
	client   *Client
	Name     string `json:"displayName"`
	Icon     string `json:"iconPath"`
	ID       string `json:"membershipId"`
	Platform int    `json:"membershipType"`
}

func (m *Membership) Member() *Member {
	return &Member{
		Name:     m.Name,
		ID:       m.ID,
		Platform: m.Platform,
		Icon:     m.Icon,
	}
}

func (p *Platform) Memberships(displayName string) (*Membership, error) {
	var rval []*Membership
	success, err := p.getAndUnwrap(fmt.Sprintf("%s/SearchDestinyPlayer/%d/%s/", p.Base, p.ID, displayName), &rval)
	if err != nil {
		return nil, err
	}
	if !success {
		return nil, fmt.Errorf("Failed response from API")
	}
	for _, m := range rval {
		if m.Platform != p.ID {
			continue
		}
		if strings.ToLower(m.Name) != strings.ToLower(displayName) {
			continue
		}
		m.client = p.Client
		return m, nil
	}
	return nil, nil
}
