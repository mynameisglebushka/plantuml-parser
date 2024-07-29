package sequence

import (
	"errors"
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
	errNoParticipant = errors.New("undefind participant type")
)

type Participant struct {
	Alias string
	Name  string
	Type  int
}

func parseParticipantFromLine(line string) (*Participant, error) {
	participant := new(Participant)

	words := strings.Split(line, " ")

	_type := checkParticipantType(string(words[0]))
	if _type == -1 {
		return nil, errNoParticipant
	}
	participant.Type = _type

	name, alias := parseNameAndAlias(words[1:])

	participant.Alias = alias
	participant.Name = name

	return participant, nil
}

func parseParticipantFromMessage(participantDefine string) *Participant {
	participant := new(Participant)

	name, alias := parseNameAndAlias(strings.Split(participantDefine, " "))

	participant.Type = ParticipantDefaultType
	participant.Name = name
	participant.Alias = alias

	return participant
}

func parseNameAndAlias(words []string) (name string, alias string) {
	var (
		nameStarted           bool
		squareBracketsStarted bool
	)
	for _, v := range words {
		switch {
		case v == `"`:
			switch {
			case !nameStarted && squareBracketsStarted:
				name = name + " " + v
			case !nameStarted && !squareBracketsStarted:
				nameStarted = true
			case nameStarted && !squareBracketsStarted:
				nameStarted = false
			}
		case strings.HasPrefix(v, `"`) && !nameStarted && !squareBracketsStarted:
			if strings.HasSuffix(v, `"`) {
				name = strings.Trim(v, `"`)
			} else {
				name = strings.TrimPrefix(v, `"`)
				nameStarted = true
			}
		case v == "[":
			squareBracketsStarted = true
		case strings.HasPrefix(v, "[") && !squareBracketsStarted && !nameStarted:
			if strings.HasSuffix(v, "]") {
				name = strings.Trim(v, "[]")
			} else {
				name = strings.TrimPrefix(v, "[")
				squareBracketsStarted = true
			}
		case strings.HasSuffix(v, `"`) && nameStarted && !squareBracketsStarted:
			name = name + " " + strings.TrimSuffix(v, `"`)
			nameStarted = false
		case v == "]":
			squareBracketsStarted = false
		case strings.HasSuffix(v, "]") && squareBracketsStarted && !nameStarted:
			name = name + " " + strings.TrimSuffix(v, "]")
			squareBracketsStarted = false
		case v == "as":
			if nameStarted {
				name = name + " " + v
			} else {
				continue
			}
		default:
			if nameStarted || squareBracketsStarted {
				if name == "" {
					name = v
				} else {
					name = name + " " + v
				}
			} else {
				alias = v
			}
		}
	}

	return name, alias
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
