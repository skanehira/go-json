package ast

import "testing"

func TestObject(t *testing.T) {
	obj := Object{
		Elements: []Element{
			{"name", String{"gorilla"}},
			{"types", Array{
				Values: []Value{
					Bool{true},
					Bool{false},
					Null{},
					Int{10},
					Float{55.55},
					Int{-11},
				},
			},
			},
		},
	}

	want := `{"name":"gorilla","types":[true,false,null,10,55.55,-11]}`
	if obj.String() != want {
		t.Fatalf("wrong object. want=%s got=%s", obj.String(), want)
	}
}

func TestArray(t *testing.T) {
	array := Array{
		Values: []Value{
			Bool{true},
			Bool{false},
			Null{},
			Int{10},
			String{"banana"},
			Object{
				Elements: []Element{
					{"name", String{"gorilla"}},
					{"age", Int{27}},
				},
			},
			Array{
				Values: []Value{},
			},
			Object{},
		},
	}

	want := `[true,false,null,10,"banana",{"name":"gorilla","age":27},[],{}]`

	if array.String() != want {
		t.Fatalf("wrong array. want=%s, got=%s", want, array.String())
	}
}
