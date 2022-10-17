package repo

import (
	"testing"
)

func TestGeneration(t *testing.T) {

	var linkTests = []Link{
		{lenShortLink: 10,
			FullLink: "test1.com"},
		{lenShortLink: 10,
			FullLink: "test2.com"},
		{lenShortLink: 10,
			FullLink: "test3.com"},
	}

	for _, test := range linkTests {
		if output := test.Generate(); len(output) != test.lenShortLink {
			t.Errorf("Short link length %v not equal to expected length %v", len(output), test.lenShortLink)
		}
	}

}
