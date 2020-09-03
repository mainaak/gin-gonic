package newsfeed

import "testing"

func TestAdd(t *testing.T){
	n := New()
	n.Add(Item{
		Title: "Title of struct Item",
		Post:  "Post of struct Item",
	})

	if condition := len(n.Items) != 1; condition {
		t.Error("Was unable to add the item to the Slice of Items")
	}
}

func TestGetAll(t *testing.T){
	n := New()
	n.Add(Item{
		Title: "Title of struct Item",
		Post:  "Post of struct Item",
	})
	get := n.GetAll()

	title := get[0].Title == "Title of struct Item"
	post := get[0].Post == "Post of struct Item"

	if !post || !title {
		t.Error("Incorrect title or post received")
	}
}
