package domain

type Event struct {
	Type       string
	ResourceID string
	Conclusion string
	URL        string
	Payload    []byte
	Headers    map[string]string
}

func (e Event) Clone() Event        { return e }
func (e Event) ClonePayload() Event { return e }
func (e Event) CloneHeaders() Event { return e }
