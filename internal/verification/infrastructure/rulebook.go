package infrastructure

type Rulebook struct {
	Version          string
	RequireSignature bool
	RequireSBOM      bool
	MaxSeverity      string
}

func DefaultRulebook() Rulebook {
	return Rulebook{Version: "v1", RequireSignature: true, RequireSBOM: true, MaxSeverity: "high"}
}
