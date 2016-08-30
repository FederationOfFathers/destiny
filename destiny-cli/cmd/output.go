package cmd

import (
	"encoding/json"
	"os"
)

func jsonOut(i interface{}) error {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("  ", "  ")
	return enc.Encode(i)
}
