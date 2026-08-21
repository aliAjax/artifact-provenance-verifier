package application

import "github.com/example/artifact-provenance-verifier/internal/webhook/domain"

func cloneEvent(e domain.Event) domain.Event { return e }
