package infrastructure

import (
	"encoding/json"
	"fmt"
	"github.com/example/artifact-provenance-verifier/internal/platform"
)

type CycloneDocument struct {
	BomFormat   string `json:"bomFormat"`
	SpecVersion string `json:"specVersion"`
	Components  []struct {
		Name     string `json:"name"`
		Version  string `json:"version"`
		PURL     string `json:"purl"`
		Licenses []struct {
			License struct {
				ID string `json:"id"`
			} `json:"license"`
		} `json:"licenses"`
	} `json:"components"`
}

func ParseCyclone(data []byte) ([]platform.Component, error) {
	var d CycloneDocument
	if e := json.Unmarshal(data, &d); e != nil {
		return nil, e
	}
	if d.BomFormat != "CycloneDX" {
		return nil, fmt.Errorf("not cyclonedx")
	}
	out := []platform.Component{}
	for _, c := range d.Components {
		lic := ""
		if len(c.Licenses) > 0 {
			lic = c.Licenses[0].License.ID
		}
		out = append(out, platform.Component{Name: c.Name, Version: c.Version, PURL: c.PURL, License: lic})
	}
	return out, nil
}
