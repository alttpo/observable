package observable

//func TestObject_Object(t *testing.T) {
//	type fields struct {
//		lock      sync.Mutex
//		observers []Observer
//		object    interface{}
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		want   interface{}
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			o := &Object{
//				lock:      tt.fields.lock,
//				observers: tt.fields.observers,
//				object:    tt.fields.object,
//			}
//			if got := o.Object(); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Object() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestObject_Set(t *testing.T) {
//	type fields struct {
//		lock      sync.Mutex
//		observers []Observer
//		object    interface{}
//	}
//	type args struct {
//		object interface{}
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
//			o := &Object{
//				lock:      tt.fields.lock,
//				observers: tt.fields.observers,
//				object:    tt.fields.object,
//			}
//		})
//	}
//}
//
//func TestObject_Subscribe(t *testing.T) {
//	type fields struct {
//		lock      sync.Mutex
//		observers []Observer
//		object    interface{}
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
//			o := &Object{
//				lock:      tt.fields.lock,
//				observers: tt.fields.observers,
//				object:    tt.fields.object,
//			}
//		})
//	}
//}
//
//func TestObject_Unsubscribe(t *testing.T) {
//	type fields struct {
//		lock      sync.Mutex
//		observers []Observer
//		object    interface{}
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
//			o := &Object{
//				lock:      tt.fields.lock,
//				observers: tt.fields.observers,
//				object:    tt.fields.object,
//			}
//		})
//	}
//}
