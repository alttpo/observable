package observable

import (
	"reflect"
	"testing"
)

func TestList_Append(t *testing.T) {
	type args struct {
		newElement interface{}
	}
	type want struct {
		event Event
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Append",
			args: args{
				newElement: 1,
			},
			want: want{
				event: Event{
					Operation: ListAppend,
					Value:     1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := NewList()

			first := true
			gotEvent := false
			o.Subscribe(NewObserver("a", func(event Event) {
				if first {
					first = false
					return
				}
				gotEvent = true
				if expected, actual := tt.want.event, event; !reflect.DeepEqual(actual, expected) {
					t.Fatalf("Event = %v, want %v", actual, expected)
				}
			}))
			if first {
				t.Fatalf("No event received; expected lset []")
			}

			o.Append(tt.args.newElement)
			if !gotEvent {
				t.Fatalf("No event received; wanted event")
			}
		})
	}
}

//func TestList_Concat(t *testing.T) {
//	type fields struct {
//		lock      sync.Mutex
//		observers []Observer
//		list      []interface{}
//	}
//	type args struct {
//		newElements []interface{}
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			o := &List{
//				lock:      tt.fields.lock,
//				observers: tt.fields.observers,
//				list:      tt.fields.list,
//			}
//		})
//	}
//}
//
//func TestList_List(t *testing.T) {
//	type fields struct {
//		lock      sync.Mutex
//		observers []Observer
//		list      []interface{}
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		want   []interface{}
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			o := &List{
//				lock:      tt.fields.lock,
//				observers: tt.fields.observers,
//				list:      tt.fields.list,
//			}
//			if got := o.List(); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("List() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestList_Set(t *testing.T) {
//	type fields struct {
//		lock      sync.Mutex
//		observers []Observer
//		list      []interface{}
//	}
//	type args struct {
//		newList []interface{}
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			o := &List{
//				lock:      tt.fields.lock,
//				observers: tt.fields.observers,
//				list:      tt.fields.list,
//			}
//		})
//	}
//}
//
//func TestList_Unsubscribe(t *testing.T) {
//	type fields struct {
//		lock      sync.Mutex
//		observers []Observer
//		list      []interface{}
//	}
//	type args struct {
//		observer Observer
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			o := &List{
//				lock:      tt.fields.lock,
//				observers: tt.fields.observers,
//				list:      tt.fields.list,
//			}
//		})
//	}
//}
