package sequence

import (
	"reflect"
	"testing"
)

func Test_parseParticipantLine(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name    string
		args    args
		want    *Participant
		wantErr bool
	}{
		{
			name: "alias first",
			args: args{
				`participant Foo as "Bar"`,
			},
			want: &Participant{
				Name:  "Bar",
				Alias: "Foo",
				Type:  ParticipantDefaultType,
			},
			wantErr: false,
		},
		{
			name: "alias second",
			args: args{
				`participant "Foo" as Bar`,
			},
			want: &Participant{
				Name:  "Foo",
				Alias: "Bar",
				Type:  ParticipantDefaultType,
			},
			wantErr: false,
		},
		{
			name: "actor type",
			args: args{
				`actor Foo as "Bar"`,
			},
			want: &Participant{
				Name:  "Bar",
				Alias: "Foo",
				Type:  ParticipantActorType,
			},
			wantErr: false,
		},
		{
			name: "database type and multi word name",
			args: args{
				`database " Data Base " as DB`,
			},
			want: &Participant{
				Name:  "Data Base",
				Alias: "DB",
				Type:  ParticipantDatabaseType,
			},
			wantErr: false,
		},
		{
			name: "wrong type",
			args: args{
				"abra-cadabra component",
			},
			wantErr: true,
		},
		{
			name: "only alias",
			args: args{
				"participant Foo",
			},
			want: &Participant{
				Alias: "Foo",
				Type:  ParticipantDefaultType,
			},
			wantErr: false,
		},
		{
			name: "name in bracet",
			args: args{
				"participant alias [ first second third ]",
			},
			want: &Participant{
				Alias: "alias",
				Name:  "first second third",
				Type:  ParticipantDefaultType,
			},
			wantErr: false,
		},
		{
			name: "name in bracet with quotes",
			args: args{
				`participant alias [ first "second" third ]`,
			},
			want: &Participant{
				Alias: "alias",
				Name:  `first "second" third`,
				Type:  ParticipantDefaultType,
			},
			wantErr: false,
		},
		// TODO: Return error on this case
		// {
		// 	name: "quotes inside name",
		// 	args: args{
		// 		`participant "first "second" third" as digits`,
		// 	},
		// 	want: &Participant{
		// 		Alias: "digits",
		// 		Name: `first "second" third`,
		// 		Type: ParticipantDefaultType,
		// 	},
		// 	wantErr: false,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseParticipantFromLine(tt.args.line)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseParticipantLine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseParticipantLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
