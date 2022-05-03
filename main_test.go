package main

import (
	"testing"
)

func Test_buildWords(t *testing.T) {
	type args struct {
		words      []string
		capitalize bool
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "return correct words if capitalize flag is false",
			args: args{
				[]string{"Abc", "dEf", "ghI"},
				false,
			},
			want: []string{"abc", "d", "ef", "gh", "i"},
		},
		{
			name: "return correct words if capitalize flag is true",
			args: args{
				[]string{"Abc", "dEf", "ghI"},
				true,
			},
			want: []string{"abc", "D", "Ef", "Gh", "I"},
		},
		{
			name: "return snake if words include non alphabets",
			args: args{
				[]string{"ab=c", "de_f", "gh-i"},
				false,
			},
			want: []string{"ab", "c", "de", "f", "gh", "i"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildWords(tt.args.words, tt.args.capitalize)

			if len(got) != len(tt.want) {
				t.Errorf("buildWords() = %v, want = %v", got, tt.want)
				return
			}

			for i := 0; i < len(got); i++ {
				if got[i] != tt.want[i] {
					t.Errorf("buildWords() = %v, want = %v", got, tt.want)
					return
				}
			}
		})
	}
}
