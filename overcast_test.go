package overcast_test

import (
	"testing"

	"github.com/hezhizhen/overcast"
)

func TestDuration(t *testing.T) {
	codeNewbie := "https://overcast.fm/itunes919219256/codenewbie"
	duration := overcast.Duration(codeNewbie)
	t.Logf("Total duration: %f minutes (%f hours)", duration.Minutes(), duration.Hours())
}
