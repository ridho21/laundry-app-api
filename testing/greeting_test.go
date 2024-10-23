package testing

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestSayHello_Success(t *testing.T){
	expected := "Hello BudiTabuti"
	actual, err := SayHello("BudiTabuti")
	// ____________________________________ Built in
	// if err != nil {
	// 	t.Fatal("SayHello() Test is Failed",err)
	// }

	// if expected != actual {
	// 	t.Fatalf(`SayHello() Test is Failed, actual %v, expected %v`,actual,expected)
	// }
	// ____________________________________ Testify
	assert.NoError(t,err)
	assert.Equal(t,expected,actual)
}

func TestSayHello_Fail(t *testing.T){
	expected := ""
	actual, err := SayHello("")

	// if err == nil {
	// 	t.Fatal("SayHello() Tes is Fail")
	// }

	// if expected != actual {
	// 	t.Fatalf(`SayHello() Test is Failed, actual %v, expected %v`,actual,expected)
	// }
	assert.Error(t,err)
	assert.Equal(t,expected,actual)
}
