package sequence

import (
	"bytes"
	"log"
)

type diagramParser struct {
	diagram *Diagram
	lines   [][]byte // array of content

	lineNum  int // number of next line in diagram plain text
	kind     int // type of current line
	prevKind int // type of previos line
}

func newDiagramParser(diagram *Diagram, content []byte) *diagramParser {
	diap := &diagramParser{
		diagram: diagram,
	}

	content = bytes.ReplaceAll(content, []byte("\r\n"), []byte("\n"))
	diap.lines = bytes.Split(content, []byte("\n"))

	var (
		wspaces = "\t\n\v\f\r \x85\xA0"

		line []byte
		x    int
	)
	for x, line = range diap.lines {
		diap.lines[x] = bytes.TrimRight(line, wspaces)
	}

	return diap
}

func (parser *diagramParser) next() (line []byte, eof bool) {
	parser.prevKind = parser.kind

	if parser.lineNum >= len(parser.lines) {
		return nil, true
	}

	line = parser.lines[parser.lineNum]
	parser.lineNum++

	parser.kind = lineKindOf(line, parser.prevKind)

	return line, false
}

func parse(diagram *Diagram, content []byte) {
	parser := newDiagramParser(diagram, content)

	var (
		collectedMultiline string
	)
	for {
		line, ok := parser.next()
		if ok {
			return
		}
		switch parser.kind {
		case kindEmptyLine:
			continue
		case kindParticipantLine:
			parser.parseParticipant(string(line))
		case kindCreateParticipantLine:
			line = bytes.TrimPrefix(line, []byte("create "))
			if !(bytes.HasPrefix(line, []byte("participant")) ||
				bytes.HasPrefix(line, []byte("actor")) ||
				bytes.HasPrefix(line, []byte("boundary")) ||
				bytes.HasPrefix(line, []byte("control")) ||
				bytes.HasPrefix(line, []byte("entity")) ||
				bytes.HasPrefix(line, []byte("database")) ||
				bytes.HasPrefix(line, []byte("collections")) ||
				bytes.HasPrefix(line, []byte("queue"))) {
				line = append([]byte("participant "), line...)
			}
			parser.parseParticipant(string(line))
		case kindParticipantMultiLine:
			collectedMultiline = collectedMultiline + string(line)
		case kindParticipantMultiLineEnd:
			parser.parseParticipant(collectedMultiline + string(line))
			collectedMultiline = ""
		}

	}

}

func (parser *diagramParser) parseParticipant(line string) {
	participant, err := parseParticipantFromLine(line)
	if err != nil {
		log.Printf("parseParticipant: %v\n", err)
		return
	}

	parser.diagram.Participants = append(parser.diagram.Participants, participant)
}
