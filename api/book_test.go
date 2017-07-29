package api

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestBookToJSON(t *testing.T) {
	
	book := Book{Title: "Cloud Native Go", Author: "M.-L. Abe", ISBN: "wwrw9094"}
	json := book.ToJSON()
	
	assert.Equal(t, `{title: "Cloud Native Go", author: "M.-L. Abe", isbn: "wwrw9094"}`, 
		string(json), "Something went wrong")
}

func TestBookFromJSON(t *testing.T) {
	assert.True(t, true, "Implement Me.")
}