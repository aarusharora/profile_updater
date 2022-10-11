package utils

import (
	"fmt"
	"strings"
	"testing"
)

func TestMakeRequest(t *testing.T) {
	// https://jsonplaceholder.typicode.com/todos/1
	got, err := MakeGetRequest("https://mocki.io/v1/d4867d8b-b5d5-4a48-a4ab-79131b5809b8")
	if err != nil {
		fmt.Println(err)
	}
	want := `[{"name":"Harry Potter","city":"London"},{"name":"Don Quixote","city":"Madrid"},{"name":"Joan of Arc","city":"Paris"},{"name":"Rosa Park","city":"Alabama"}]`
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestAppendToMarkdown(t *testing.T) {
	AppendToMarkdown("../../input/profile_markdown.md", "../../output/README.md", "- Testing")
}

func TestUnpackStringSlice(t *testing.T) {
	var t1, t2, t3 string = "", "", ""
	UnpackStringSlice(strings.Split("City,State,Country", ","), &t1, &t2, &t3)
	// fmt.Println(t1, t2, t3)
	if t1 == "" || t2 == "" || t3 == "" {
		t.Errorf("got %s %s %s ", t1, t2, t3)
	}
	t1, t2, t3 = "", "", ""
	UnpackStringSlice(strings.Split("City,State", ","), &t1, &t2, &t3)
	// fmt.Println(t1, t2, t3)
	if t3 != "" {
		t.Errorf("got %s %s %s ", t1, t2, t3)
	}
	t1, t2, t3 = "", "", ""
	UnpackStringSlice(strings.Split("City", ","), &t1, &t2, &t3)
	// fmt.Println(t1, t2, t3)
	if t2 != "" && t3 != "" {
		t.Errorf("got %s %s %s ", t1, t2, t3)
	}
	t1, t2, t3 = "", "", ""
	UnpackStringSlice(strings.Split("", ","), &t1, &t2, &t3)
	// fmt.Println(t1, t2, t3)
	if t1 != "" && t2 != "" && t3 != "" {
		t.Errorf("got %s %s %s ", t1, t2, t3)
	}
	t1, t2, t3 = "", "", ""
	UnpackStringSlice(strings.Split("City,State,Country,Test", ","), &t1, &t2, &t3)
	// fmt.Println(t1, t2, t3)
	if t1 == "" || t2 == "" || t3 == "" {
		t.Errorf("got %s %s %s ", t1, t2, t3)
	}
}
