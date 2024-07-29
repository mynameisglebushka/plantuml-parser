package sequence

import "bytes"

const (
	kindEmptyLine = iota
	kindTextLine
	kindParticipantLine
	kindCreateParticipantLine
	kindParticipantMultiLine
	kindParticipantMultiLineEnd
)

func lineKindOf(line []byte, prevKind int) (kind int) {
	switch {
	case len(line) == 0:
		return kindEmptyLine
	case bytes.HasPrefix(line, []byte("participant")) ||
		bytes.HasPrefix(line, []byte("actor")) ||
		bytes.HasPrefix(line, []byte("boundary")) ||
		bytes.HasPrefix(line, []byte("control")) ||
		bytes.HasPrefix(line, []byte("entity")) ||
		bytes.HasPrefix(line, []byte("database")) ||
		bytes.HasPrefix(line, []byte("collections")) ||
		bytes.HasPrefix(line, []byte("queue")):
		if bytes.HasSuffix(line, []byte("[")) {
			return kindParticipantMultiLine
		}
		return kindParticipantLine
	case bytes.HasPrefix(line, []byte("create")):
		return kindCreateParticipantLine
	case prevKind == kindParticipantMultiLine:
		if bytes.HasSuffix(line, []byte("]")) {
			return kindParticipantMultiLineEnd
		}
		return kindParticipantMultiLine
	default:
		return kindTextLine
	}
}
