package sequence

import (
	"fmt"
	"os"
	"path/filepath"
)

type Diagram struct {
	Participants []*Participant

	fPath string
	path string
}

func newDiagram() *Diagram {
	return &Diagram{

	}
}

func Open(path string) (*Diagram, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Open.ReadFile %s: %v", path, err)
	}

	wd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("Open.Getwd: %v", err)
	}

	diagram := newDiagram()
	diagram.fPath = filepath.Join(wd, path)
	diagram.path = path

	parse(diagram, content)

	return diagram, nil
}

func ParseDiagram(content []byte) (*Diagram, error) {
	diagram := newDiagram()

	parse(diagram, content)

	return diagram, nil
}