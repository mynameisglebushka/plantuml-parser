package sequence

import (
	"errors"
	"regexp"
	"strings"
)

// Participant Types
const (
	ParticipantDefaultType = iota
	ParticipantActorType
	ParticipantBoundaryType
	ParticipantControlType
	ParticipantEntityType
	ParticipantDatabaseType
	ParticipantCollectionsType
	ParticipantDueueType
)

var (
	errNoParticipant = errors.New("undefind")
)

type Participant struct {
	Alias string
	Name  string
	Type  int
}

type participantParser struct {
	re *regexp.Regexp
}

func newParticipantParser() *participantParser {
	re := regexp.MustCompile(`^([^ ]+)[^\S]+([^ ]+)[^\S]+as[^\S]+([^ ]+)$`)
	return &participantParser{
		re: re,
	}
}

func (pp *participantParser) parseParticipant(line []byte) (*Participant, error) {
	p := new(Participant)

	part := pp.re.FindStringSubmatch(string(line))

	// Throw match
	part = part[1:]

	typ := checkParticipantType(part[0])
	if typ == -1 {
		return nil, errNoParticipant
	}
	p.Type = typ

	if len(part) < 3 {
		p.Alias = part[1]
		return p, nil
	}

	var (
		name string
		alias string
	)
	if checkDoubleQuotes(part[1]) {
		name = part[1]
		alias = part[2]
	} else if checkDoubleQuotes(part[2]) {
		name = part[2]
		alias = part[1]
	}

	if alias == "" || name == "" {
		return nil, errNoParticipant
	}

	p.Alias = alias
	p.Name = trimDoubleQuotes(name)

	return p, nil
}

func checkDoubleQuotes(s string) bool {
	return strings.HasPrefix(s, "\"") && strings.HasSuffix(s, "\"")
}

func trimDoubleQuotes(s string) string {
	s = strings.TrimLeft(s, "\"")
	s = strings.TrimRight(s, "\"")
	return s
}

func checkParticipantType(s string) int {
	switch s {
	case "participant":
		return ParticipantDefaultType
	case "actor":
		return ParticipantActorType
	case "boundary":
		return ParticipantBoundaryType
	case "control":
		return ParticipantControlType
	case "entity":
		return ParticipantEntityType
	case "database":
		return ParticipantDatabaseType
	case "collections":
		return ParticipantCollectionsType
	case "queue":
		return ParticipantDueueType
	default:
		return -1
	}
}
