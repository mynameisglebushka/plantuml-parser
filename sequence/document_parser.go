package sequence

type documentParser struct {
	participantParser *participantParser
}

func newDocumentParser() *documentParser {
	return &documentParser{
		participantParser: newParticipantParser(),
	}
}