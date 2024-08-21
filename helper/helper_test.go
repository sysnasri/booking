package helper

import "testing"



var tests = []struct {

	name string
	dividend float32
	divisor float32
	expected float32
	isErr bool

}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 0.0, true},
	{"expect-5", 100.0, 20.0, 5.0, false},

}

func TestDevide(t *testing.T){
	for _,tt := range tests{
		g,err := Devide(tt.dividend,tt.divisor)
		if tt.isErr {
			if err == nil {
				t.Error("Expected error but there is no",tt.name)


		}
	
	
	} else {
		if err != nil {
			t.Error("did not exepect an error but got one!",err.Error())
		}
	}
	if g != tt.expected {
		t.Errorf("expected %f but got %f",tt.expected,g)

	}


}
}
// func TestDevid(t *testing.T){

// 	_,err := Devid(10.0, 1.0)
// 	if err != nil{
// 		t.Error("Problem in Deviding!")
// 	}


// }

// func TestDevidZero (t *testing.T){


// }