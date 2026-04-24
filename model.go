package main

import (
	"crypto/sha256"
	"strings"
	"errors"
)

// data container representing a page or page segment,
// uses sha256 as a checksum to detect if a pages content as different
type Page struct {
	URL 	   string
	Checksum [32]byte
	//Keyword [5]string
}

// Creates a page, returns a page ptr
// if url is empty, or content is nilptr or empty it will return an error
func NewPage(url string, content *string, /*keyword [5]string*/) (*Page, error) {	
	//guard, content must be init or 1> char
	if strings.TrimSpace(url) == "" {
		return nil, errors.New("url can't be empty")
	}
	if content == nil || strings.TrimSpace(*content) == "" {
		return nil, errors.New("content can't be nil or empty")
	}

	//create page
	var p Page 
	p.URL      = url
	p.Checksum = sha256.Sum256([]byte(*content))
	return &p, nil
}


/// embedding created by an ai to represent pages semantic 
/// float32 is used to assure native compatibility with sqlite-vec 
type Embedding [2048]float32

/*
func NewEmbedding(ai *ai.Api, content *string) (*Embedding, error) {
	if content == nil || *content == "" {
		return nil, errors.New("content can't be nil or empty")	
	}

	var e Embedding

	return e, nil 
}
*/
