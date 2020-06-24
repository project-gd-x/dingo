package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImports_Add(t *testing.T) {
	imports := Imports{}

	i1 := imports.Add("github.com/core/application/view/signup", "signup")
	i2 := imports.Add("github.com/core/application/signup", "signup")
	i3 := imports.Add("github.com/core/signup", "signup")
	i4 := imports.Add("github.com/core/application/view/signup", "signup")
	i5 := imports.Add("github.com/core/signup", "signup")

	assert.Equal(t, &Import{Name: "github.com/core/application/view/signup", LocalName: "signup", Renamed: false, OriginalLocalName: ""}, i1)
	assert.Equal(t, &Import{Name: "github.com/core/application/signup", LocalName: "signup_1", Renamed: true, OriginalLocalName: "signup"}, i2)
	assert.Equal(t, &Import{Name: "github.com/core/signup", LocalName: "signup_2", Renamed: true, OriginalLocalName: "signup"}, i3)
	assert.Equal(t, &Import{Name: "github.com/core/application/view/signup", LocalName: "signup", Renamed: false, OriginalLocalName: ""}, i4)
	assert.Equal(t, &Import{Name: "github.com/core/signup", LocalName: "signup_2", Renamed: true, OriginalLocalName: "signup"}, i5)

	expected := Imports{
		Items: []*Import{
			{Name: "github.com/core/application/view/signup", LocalName: "signup", Renamed: false, OriginalLocalName: ""},
			{Name: "github.com/core/application/signup", LocalName: "signup_1", Renamed: true, OriginalLocalName: "signup"},
			{Name: "github.com/core/signup", LocalName: "signup_2", Renamed: true, OriginalLocalName: "signup"},
		},
	}

	assert.Equal(t, expected, imports)
}
