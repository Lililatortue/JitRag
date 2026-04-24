package main

import (
	"context"
	"testing"
)

func TestAddPage(t *testing.T){
	//create database
	store, err := NewDatabase(":memory:")
	if err != nil {
		t.Errorf("couldn't create the new database %s", err)
		return
	}

	content  := "hello world"
	page1, err := NewPage("https://github/lililatortue", &content)
	if err != nil {
		t.Errorf("couldn't create the page1 %s", err)
		return
	}
	page2, err := NewPage("https://lol/usuck.com", &content)
	if err != nil {
		t.Errorf("couldn't create the page2 %s", err)
		return
	}


	test := []struct {
		name    string
		page    *Page
	}{
		{"Page 1", page1},
		{"Page 2", page2},
	}
	for i, tt := range test {
		t.Run(tt.name ,func(t *testing.T){
			id ,err := store.InsertPage(context.TODO(),tt.page)	
			
			if err != nil{
				t.Errorf("couldn't add page, %v", err)
				return
			}

			if id != int64(i + 1) {	
				t.Errorf("id was supposed to be %d, but it was %d",i, id)
				return
			}
		})
	}
}


func TestFetchPage(t *testing.T) {
	store, err := NewDatabase(":memory:")
	if err != nil {
		t.Errorf("couldn't create the new database %s", err)
		return
	}
	content  := "hello world fetchPage"
	page1, _ := NewPage("https://github/lililatortue", &content)	
	page2, _ := NewPage("https://lol/usuck.com", &content)

	id1, _ := store.InsertPage(context.TODO(),page1);
	id2, _ := store.InsertPage(context.TODO(),page2);
	
	ids := []int64 {id1, id2}
	pages := store.FetchPage(context.TODO(), ids)	

	if len(pages) != 2 {
		t.Errorf("expected 2 pages got %d,  %d", len(pages), len(ids))
		return
	}

	if pages[0].URL != page1.URL {
		t.Errorf("expected URL %s but it is %s", page1.URL, pages[1].URL)
		return
	}
	if  pages[0].Checksum != page1.Checksum {
		t.Errorf("expected HASH %s but it is %s", page1.Checksum, pages[1].Checksum)
		return
	}
	if pages[1].URL != page2.URL {
		t.Errorf("expected URL %s but it is %s", page2.URL, pages[2].URL)
		return
	}
	if  pages[1].Checksum != page2.Checksum {
		t.Errorf("expected HASH %s but it is %s", page2.Checksum, pages[2].Checksum)
		return
	}
}


