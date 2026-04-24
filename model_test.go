package main

import (
	"testing"
)


func TestPageCreation(t *testing.T){
	// test to validate
	validContent := "hello world"
	WSContent    := "    "
	EmptyContent := ""	
	var NilContent *string = nil

	test := []struct {
		name 	   string
		url  	   string
		content *string
		wantErr      bool
	}{
		{"Valid Page"   , "test.com", &validContent, false},
		{"Empty URL"    , ""        , &validContent, true },
		{"WS    URL"    , "        ", &validContent, true },
		{"Nil Content"  , "test.com", NilContent   , true },
		{"Empty Content", "test.com", &EmptyContent, true },
		{"WS Content   ", "test.com", &WSContent   , true },
	}

	for _,tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			p, err := NewPage(tt.url, tt.content)

			if (err != nil) != tt.wantErr {
				t.Errorf("NewPage(), error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if p.URL != tt.url {
					t.Errorf("expected URL = %s, got %s", p.URL, tt.url)
				}
				if p.Checksum ==[32]byte{} {
					t.Errorf("expected non-empty hash")
				}
			}
		})
	}
}


