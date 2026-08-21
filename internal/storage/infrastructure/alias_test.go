package infrastructure

import (
	"sync"
	"testing"

	"github.com/example/artifact-provenance-verifier/internal/platform"
)

func exerciseConcurrentReaders(action func()) {
	start := make(chan struct{})
	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() { defer wg.Done(); <-start; action() }()
	}
	close(start)
	wg.Wait()
}

func TestArtifactTagsConcurrentIsolation(t *testing.T) {
	m := NewMemory()
	_, err := m.CreateArtifact(nil, platform.Artifact{ID: "art-1", Organization: "acme", Repository: "payments", Name: "api", Digest: "sha256:0123456789abcdef", Tags: []string{"stable"}})
	if err != nil {
		t.Fatal(err)
	}
	first, err := m.GetArtifact(nil, "art-1")
	if err != nil {
		t.Fatal(err)
	}
	first.Tags[0] = "caller-mutated"
	second, err := m.GetArtifact(nil, "art-1")
	if err != nil {
		t.Fatal(err)
	}
	if second.Tags[0] != "stable" {
		t.Fatalf("repository state changed: %#v", second.Tags)
	}
	start := make(chan struct{})
	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-start
			a, e := m.GetArtifact(nil, "art-1")
			if e != nil {
				t.Errorf("get artifact: %v", e)
				return
			}
			a.Tags[0] = "worker-mutated"
		}()
	}
	close(start)
	wg.Wait()
	last, _ := m.GetArtifact(nil, "art-1")
	if last.Tags[0] != "stable" {
		t.Fatalf("concurrent caller changed state: %#v", last.Tags)
	}
}

func TestArtifactCreateReturnsIsolatedCopy(t *testing.T) {
	m := NewMemory()
	a, err := m.CreateArtifact(nil, platform.Artifact{ID: "art-copy", Name: "api", Digest: "sha256:abcdef0123456789", Tags: []string{"one"}})
	if err != nil {
		t.Fatal(err)
	}
	a.Tags[0] = "changed"
	exerciseConcurrentReaders(func() { _, _ = m.GetArtifact(nil, "art-copy") })
	got, _ := m.GetArtifact(nil, "art-copy")
	if got.Tags[0] != "one" {
		t.Fatalf("create result shares storage: %#v", got.Tags)
	}
}

func TestArtifactCreateCopiesInput(t *testing.T) {
	m := NewMemory()
	input := platform.Artifact{ID: "art-input", Name: "worker", Digest: "sha256:1111222233334444", Tags: []string{"trusted"}}
	_, _ = m.CreateArtifact(nil, input)
	input.Tags[0] = "caller"
	exerciseConcurrentReaders(func() { _, _ = m.GetArtifact(nil, "art-input") })
	got, _ := m.GetArtifact(nil, "art-input")
	if got.Tags[0] != "trusted" {
		t.Fatalf("create retained input slice: %#v", got.Tags)
	}
}

func TestArtifactListReturnsIsolatedCopy(t *testing.T) {
	m := NewMemory()
	_, _ = m.CreateArtifact(nil, platform.Artifact{ID: "art-list", Name: "api", Digest: "sha256:fedcba9876543210", Tags: []string{"one"}})
	items, err := m.ListArtifacts(nil, "api")
	if err != nil || len(items) != 1 {
		t.Fatalf("list: %v %#v", err, items)
	}
	items[0].Tags[0] = "changed"
	exerciseConcurrentReaders(func() { _, _ = m.ListArtifacts(nil, "api") })
	got, _ := m.GetArtifact(nil, "art-list")
	if got.Tags[0] != "one" {
		t.Fatalf("list result shares storage: %#v", got.Tags)
	}
}

func TestAttestationPutCopiesPredicate(t *testing.T) {
	m := NewMemory()
	input := platform.Attestation{ArtifactID: "art-put", BodyDigest: "sha256:put", Predicate: map[string]any{"builder": "ci"}}
	_, _ = m.PutAttestation(nil, input)
	input.Predicate["builder"] = "caller"
	exerciseConcurrentReaders(func() { _, _ = m.ListAttestations(nil, "art-put") })
	got, _ := m.ListAttestations(nil, "art-put")
	if got[0].Predicate["builder"] != "ci" {
		t.Fatalf("put retained predicate map: %#v", got[0].Predicate)
	}
}

func TestAttestationListReturnsDeepCopy(t *testing.T) {
	m := NewMemory()
	_, _ = m.PutAttestation(nil, platform.Attestation{ArtifactID: "art-1", BodyDigest: "sha256:body", Predicate: map[string]any{"builder": "ci"}})
	items, err := m.ListAttestations(nil, "art-1")
	if err != nil || len(items) != 1 {
		t.Fatalf("list attestations: %v %#v", err, items)
	}
	items[0].Predicate["builder"] = "caller"
	exerciseConcurrentReaders(func() { _, _ = m.ListAttestations(nil, "art-1") })
	got, _ := m.ListAttestations(nil, "art-1")
	if got[0].Predicate["builder"] != "ci" {
		t.Fatalf("predicate map shares storage: %#v", got[0].Predicate)
	}
}

func TestSnapshotRawSavesIndependentCopy(t *testing.T) {
	s := NewSnapshotStore()
	data := []byte("snapshot")
	s.SaveRaw("k", data)
	data[0] = 'X'
	got, err := s.LoadRaw("k")
	exerciseConcurrentReaders(func() { _, _ = s.LoadRaw("k") })
	if err != nil || string(got) != "snapshot" {
		t.Fatalf("saved bytes changed: %q %v", got, err)
	}

}

func TestSnapshotRawLoadsIndependentCopy(t *testing.T) {
	s := NewSnapshotStore()
	s.SaveRaw("k", []byte("snapshot"))
	got, err := s.LoadRaw("k")
	if err != nil {
		t.Fatal(err)
	}
	got[0] = 'Y'
	exerciseConcurrentReaders(func() { _, _ = s.LoadRaw("k") })
	again, _ := s.LoadRaw("k")
	if string(again) != "snapshot" {
		t.Fatalf("loaded bytes changed storage: %q", again)
	}
}
