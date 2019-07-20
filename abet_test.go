package adet

import "testing"

func TestFixPunctuation(t *testing.T) {
	var cases = []struct {
		Text string
		Exp  string
	}{
		{Text: "привет:андрей", Exp: "привет: андрей"},
		{Text: "a.ek", Exp: "a. ek"},
		{Text: "a,f", Exp: "a, f"},
		{Text: "sdfd!sfd", Exp: "sdfd! sfd"},
		//TODO	{Text:"::",Exp:":"},
		{Text: "a!", Exp: "a!"},
	}
	for _, tt := range cases {
		actual := fixPunctuation(tt.Text)
		if actual != tt.Exp {
			t.Errorf("actual: %s != %s", actual, tt.Exp)
		}
	}

}
