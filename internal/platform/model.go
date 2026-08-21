package platform

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Artifact struct {
	ID           string    `json:"id,omitempty"`
	Organization string    `json:"organization"`
	Repository   string    `json:"repository"`
	Name         string    `json:"name"`
	Digest       string    `json:"digest"`
	Algorithm    string    `json:"algorithm"`
	Architecture string    `json:"architecture"`
	Version      string    `json:"version"`
	Commit       string    `json:"commit"`
	Status       string    `json:"status"`
	BuiltAt      time.Time `json:"built_at,omitempty"`
	Tags         []string  `json:"tags,omitempty"`
}
type Attestation struct {
	ID            string         `json:"id,omitempty"`
	ArtifactID    string         `json:"artifact_id,omitempty"`
	Type          string         `json:"type"`
	Subject       string         `json:"subject"`
	PredicateType string         `json:"predicate_type"`
	BodyDigest    string         `json:"body_digest"`
	Signature     string         `json:"signature,omitempty"`
	Status        string         `json:"status,omitempty"`
	CreatedAt     time.Time      `json:"created_at,omitempty"`
	Predicate     map[string]any `json:"predicate"`
}
type Component struct {
	ID           string   `json:"id,omitempty"`
	Name         string   `json:"name"`
	Version      string   `json:"version"`
	PURL         string   `json:"purl,omitempty"`
	License      string   `json:"license,omitempty"`
	ExternalRefs []string `json:"external_refs,omitempty"`
}
type DependencyEdge struct{ From, To string }
type SBOM struct {
	ID           string           `json:"id,omitempty"`
	ArtifactID   string           `json:"artifact_id,omitempty"`
	Format       string           `json:"format"`
	Serial       string           `json:"serial,omitempty"`
	Digest       string           `json:"digest,omitempty"`
	Components   []Component      `json:"components"`
	Dependencies []DependencyEdge `json:"dependencies"`
	CreatedAt    time.Time        `json:"created_at,omitempty"`
}
type Vulnerability struct {
	ID             string    `json:"id,omitempty"`
	Identifier     string    `json:"identifier"`
	PackagePattern string    `json:"package_pattern"`
	AffectedRange  string    `json:"affected_range,omitempty"`
	FixedVersion   string    `json:"fixed_version,omitempty"`
	Severity       string    `json:"severity"`
	Withdrawn      bool      `json:"withdrawn,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
}
type CheckResult struct {
	Name, Severity, Status, Message string
	Details                         map[string]any
}
type Verification struct {
	ID          string        `json:"id,omitempty"`
	ArtifactID  string        `json:"artifact_id"`
	InputDigest string        `json:"input_digest,omitempty"`
	Conclusion  string        `json:"conclusion"`
	RuleVersion string        `json:"rule_version"`
	Checks      []CheckResult `json:"checks"`
	CreatedAt   time.Time     `json:"created_at,omitempty"`
	Reused      bool          `json:"reused,omitempty"`
}

func NormalizeDigest(v string) (string, error) {
	v = strings.TrimSpace(strings.ToLower(v))
	if !strings.Contains(v, ":") {
		return "", fmt.Errorf("digest must include algorithm")
	}
	p := strings.SplitN(v, ":", 2)
	if len(p[1]) < 16 {
		return "", fmt.Errorf("digest value too short")
	}
	if _, err := hex.DecodeString(p[1]); err != nil {
		return "", fmt.Errorf("invalid digest: %w", err)
	}
	return p[0] + ":" + p[1], nil
}
func HashJSON(v any) (string, error) {
	b, e := json.Marshal(v)
	if e != nil {
		return "", e
	}
	h := sha256.Sum256(b)
	return "sha256:" + hex.EncodeToString(h[:]), nil
}
func ValidSeverity(v string) bool {
	switch strings.ToLower(v) {
	case "critical", "high", "medium", "low", "unknown":
		return true
	}
	return false
}
func SeverityRank(v string) int {
	switch strings.ToLower(v) {
	case "critical":
		return 4
	case "high":
		return 3
	case "medium":
		return 2
	case "low":
		return 1
	}
	return 0
}
func ValidateArtifact(a Artifact) error {
	if a.Organization == "" || a.Repository == "" || a.Name == "" {
		return fmt.Errorf("organization, repository and name are required")
	}
	d, e := NormalizeDigest(a.Digest)
	if e != nil {
		return e
	}
	a.Digest = d
	if a.Algorithm == "" {
		return fmt.Errorf("algorithm is required")
	}
	return nil
}
func ValidateComponent(c Component) error {
	if c.Name == "" || c.Version == "" {
		return fmt.Errorf("component name and version are required")
	}
	if c.PURL != "" && !strings.HasPrefix(c.PURL, "pkg:") {
		return fmt.Errorf("invalid purl")
	}
	return nil
}
