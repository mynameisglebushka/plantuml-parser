package sequence

import (
	"reflect"
	"testing"
)

func Test_participantParser_parseParticipant(t *testing.T) {
	type args struct {
		line []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *Participant
		wantErr bool
	}{
		{
			name: "participant",
			args: args{
				[]byte(`participant Foo as "Bar"`),
			},
			want: &Participant{
				Name: "Bar",
				Alias: "Foo",
				Type: ParticipantDefaultType,
			},
			wantErr: false,
		},
		{
			name: "actor",
			args: args{
				[]byte(`actor Foo as "Bar"`),
			},
			want: &Participant{
				Name: "Bar",
				Alias: "Foo",
				Type: ParticipantActorType,
			},
			wantErr: false,
		},
		{
			name: "name is first",
			args: args{
				[]byte(`participant "Foo" as Bar`),
			},
			want: &Participant{
				Name: "Foo",
				Alias: "Bar",
				Type: ParticipantDefaultType,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pp := newParticipantParser()
			got, err := pp.parseParticipant(tt.args.line)
			if (err != nil) != tt.wantErr {
				t.Errorf("participantParser.parseParticipant() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("participantParser.parseParticipant() = %v, want %v", got, tt.want)
			}
		})
	}
}
