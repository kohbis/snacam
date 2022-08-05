package main

import (
	"testing"
)

func Test_buildWords(t *testing.T) {
	type args struct {
		words  []string
		ccType CamelCaseType
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "return correct words if ccType is None",
			args: args{
				[]string{"Abc", "dEf", "ghI"},
				NONE,
			},
			want: []string{"abc", "d", "ef", "gh", "i"},
		},
		{
			name: "return correct words if ccType is Upper",
			args: args{
				[]string{"Abc", "dEf", "ghI"},
				UCC,
			},
			want: []string{"Abc", "D", "Ef", "Gh", "I"},
		},
		{
			name: "return correct words if ccType is Lower",
			args: args{
				[]string{"Abc", "dEf", "ghI"},
				LCC,
			},
			want: []string{"abc", "D", "Ef", "Gh", "I"},
		},
		{
			name: "return correct words if ccType is Lower and includes non alphabets",
			args: args{
				[]string{"Abc=def_gh-I"},
				LCC,
			},
			want: []string{"abc", "Def", "Gh", "I"},
		},
		{
			name: "return correct words if words include non alphabets",
			args: args{
				[]string{"ab=c", "de_f", "gh-i"},
				NONE,
			},
			want: []string{"ab", "c", "de", "f", "gh", "i"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildWords(tt.args.words, tt.args.ccType)

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
