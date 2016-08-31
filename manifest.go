package destiny

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ManifestAssetDatabase struct {
	Version int    `json:"version"`
	Path    string `json:"path"`
}

type Manifest struct {
	Version            string                  `json:"version"`
	AssetContentPath   string                  `json:"mobileAssetContentPath"`
	WorldContentPaths  map[string]string       `json:"mobileWorldContentPaths"`
	GearCDN            map[string]string       `json:"mobileGearCDN"`
	GearAssetDataBases []ManifestAssetDatabase `json:"mobileGearAssetDataBases"`
}

func (m *Manifest) Get(path string) (*http.Response, error) {
	return http.Get(fmt.Sprintf("http://www.bungie.net%s", path))
}

func (c *Client) Manifest() (*Manifest, error) {
	var rval *Manifest
	_, err := c.getAndUnwrap("https://www.bungie.net/Platform/Destiny/Manifest/", &rval)
	return rval, err
}

func (c *Client) RawManifestData() ([]byte, error) {
	var raw json.RawMessage
	_, err := c.getAndUnwrap("https://www.bungie.net/Platform/Destiny/Manifest/", &raw)
	return raw, err
}
