package url

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParse(t *testing.T) {
	const rawurl = "https://foo.com/go"

	want := &URL{
		Scheme: "https",
		Host:   "foo.com",
		Path:   "go",
	}

	got, err := Parse(rawurl)
	if err != nil {
		t.Fatalf("Parse(%q) err = %q, want nil", rawurl, err)
	}
	if *got != *want {
		t.Errorf("Parse(%q):\n\tgot: %q\n\twant: %q", rawurl, got, want)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Parse(%q) mismatch (-want +got):\n%s", rawurl, diff)
	}
}

func TestURLPort(t *testing.T) {
	tests := []struct {
		in   string
		port string
	}{
		{in: "foo.com:80", port: "80"},
		{in: "foo.com:", port: ""},
		{in: "foo.com", port: ""},
		{in: "1.2.3.4:90", port: "90"},
		{in: "1.2.3.4", port: ""},
	}

	for _, tt := range tests {
		u := &URL{Host: tt.in}
		if got, want := u.Port(), tt.port; got != want {
			t.Errorf("for host %q; got %q; want %q", tt.in, got, want)
		}
	}
}

func TestParseInvalidURLs(t *testing.T) {
	tests := map[string]string{
		"missing scheme": "foo.com",
	}
	for name, in := range tests {
		t.Run(name, func(t *testing.T) {
			if _, err := Parse(in); err == nil {
				t.Errorf("Parse(%q)=nil; want an error", in)
			}
		})
	}
}

func BenchmarkURLString(b *testing.B) {
	b.ReportAllocs()
	b.Logf("Loop %d times\n", b.N)
	u := &URL{Scheme: "https", Host: "foo.com", Path: "go"}
	for i := 0; i < b.N; i++ {
		u.String()
	}
}
