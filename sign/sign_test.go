package sign

import "testing"

func TestGenerateSign(t *testing.T) {
	timestamp := int64(1751372757946)
	t.Logf("ts=%d", timestamp)
	secret := "your_secret"
	got := GenerateSign(timestamp, secret)

	want := "Aklg052w+AenXom7sNKGk11gwoxlQSjfmzaMjLq2VC4="
	if got != want {
		t.Fatalf("signature mismatch: want %s but got %s", want, got)
	}
}
