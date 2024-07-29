package sequence

import (
	"encoding/json"
	"testing"
)

func TestOpen_Participant(t *testing.T) {
	diagram, err := Open("../example/participant_example.puml")
	if err != nil {
		t.Fatal(err)
	}
	b, err := json.MarshalIndent(diagram, "", "\t")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(b))
}