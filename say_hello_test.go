package learn_go_one

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTableSayHello(t *testing.T) {
	tests := []struct {
		Name, Param, Expectation string
	}{
		{
			Name:        "SayHelloEko",
			Param:       "Eko",
			Expectation: "hello world from Eko",
		},
		{
			Name:        "SayHelloFebri",
			Param:       "Febri",
			Expectation: "hello world from Febri",
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			result := SayHello(test.Param)
			require.Equal(t, test.Expectation, result, "ada yang salah")
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Eko", func(t *testing.T) {
		result := SayHello("eko")
		assert.Equal(t, "hello world from eko", result, "ada yg salah")
	})

	t.Run("Febri", func(t *testing.T) {
		result := SayHello("febri")
		assert.Equal(t, "hello world from febri", result, "ada yg salah")
	})
}

func TestMain(m *testing.M) {
	fmt.Println("BEFORE")
	m.Run()
	fmt.Println("AFTER")
}

func TestSayHello(t *testing.T) {
	result := SayHello("eko")
	assert.Equal(t, "hello world from eko", result, "ada yg salah")
	// if result != "hello world from eko" {
	// 	// t.Fail()
	// 	t.Error("ada yang salah")
	// }
}

func TestSayHelloFebri(t *testing.T) {
	result := SayHello("febri")
	assert.Equal(t, "hello world from febri", result, "ada yg salah")
	// if result != "hello world from febri" {
	// 	// t.FailNow()
	// 	t.Fatal("ada yang salah")
	// }
}

func TestSayHelloSkip(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("ups linux not allowed")
	}
	result := SayHello("febri")
	assert.Equal(t, "hello world from febri", result, "ada yg salah")
}
