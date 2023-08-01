package cmd

import "testing"

var output = `[{" Age":" 37"," Country":" Portugal","Player":"Christiano Ronaldo"},{" Age":" 36"," Country":" Argentina","Player":"Lionel Messi"},{" Age":" 30"," Country":" Brazil","Player":"Neymar Jr"}]`

func TestCSVToJSONCommand(t *testing.T) {
	got := output
	want := output + "!"

	if got != want {
		t.Errorf("Did not return expected json result")
	}
}
