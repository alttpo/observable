package observable

import "testing"

func Test_observerImpl_Equals(t *testing.T) {
	type fields struct {
		key      string
		observer ObserverFunc
	}
	type args struct {
		other Observer
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Keys compare",
			fields: fields{
				key:      "a",
				observer: nil,
			},
			args: args{
				other: NewObserver("a", nil),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &observerImpl{
				key:      tt.fields.key,
				observer: tt.fields.observer,
			}
			if got := o.Equals(tt.args.other); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}
