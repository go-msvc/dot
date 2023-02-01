package dot_test

import (
	"testing"

	"github.com/go-msvc/dot"
)

func TestSimple(t *testing.T) {
	o := map[string]interface{}{
		"a": 1,
		"b": true,
		"c": nil,
		"d": "d",
		"e": 3.141,
	}

	dot1 := dot.New(o)
	if a, err := dot1.Get(".a"); err != nil || a != 1 {
		t.Fatalf("a=(%T)%+v != 1", a, a)
	}
	if b, err := dot1.Get(".b"); err != nil || b != true {
		t.Fatalf("b=(%T)%+v != true", b, b)
	}
	if c, err := dot1.Get(".c"); err != nil || c != nil {
		t.Fatalf("c=(%T)%+v != nil", c, c)
	}
	if d, err := dot1.Get(".d"); err != nil || d != "d" {
		t.Fatalf("d=(%T)%+v != \"d\"", d, d)
	}
	if e, err := dot1.Get(".e"); err != nil || e != 3.141 {
		t.Fatalf("e=(%T)%+v != 3.141", e, e)
	}

	//set values
	if err := dot1.Set(".a", 2); err != nil {
		t.Fatal(err)
	}
	if err := dot1.Set(".b", false); err != nil {
		t.Fatal(err)
	}
	if err := dot1.Set(".c", &dot1); err != nil {
		t.Fatal(err)
	}
	if err := dot1.Set(".d", "e"); err != nil {
		t.Fatal(err)
	}
	if err := dot1.Set(".e", 5.678); err != nil {
		t.Fatal(err)
	}

	if a, err := dot1.Get(".a"); err != nil || a != 2 {
		t.Fatalf("a=(%T)%+v != 2", a, a)
	}
	if b, err := dot1.Get(".b"); err != nil || b != false {
		t.Fatalf("b=(%T)%+v != false", b, b)
	}
	if c, err := dot1.Get(".c"); err != nil || c != &dot1 {
		t.Fatalf("c=(%T)%+v != &dot1", c, c)
	}
	if d, err := dot1.Get(".d"); err != nil || d != "e" {
		t.Fatalf("d=(%T)%+v != \"e\"", d, d)
	}
	if e, err := dot1.Get(".e"); err != nil || e != 5.678 {
		t.Fatalf("e=(%T)%+v != 5.678", e, e)
	}
}

func TestSubs(t *testing.T) {
	s := map[string]interface{}{
		"a": 1,
		"b": true,
		"c": nil,
		"d": "d",
		"e": 3.141,
	}
	o := map[string]interface{}{
		"sub": s,
	}

	dot1 := dot.New(o)
	if a, err := dot1.Get(".sub.a"); err != nil || a != 1 {
		t.Fatalf("a=(%T)%+v != 1", a, a)
	}
	if b, err := dot1.Get(".sub.b"); err != nil || b != true {
		t.Fatalf("b=(%T)%+v != true", b, b)
	}
	if c, err := dot1.Get(".sub.c"); err != nil || c != nil {
		t.Fatalf("c=(%T)%+v != nil", c, c)
	}
	if d, err := dot1.Get(".sub.d"); err != nil || d != "d" {
		t.Fatalf("d=(%T)%+v != \"d\"", d, d)
	}
	if e, err := dot1.Get(".sub.e"); err != nil || e != 3.141 {
		t.Fatalf("e=(%T)%+v != 3.141", e, e)
	}

	//set values
	if err := dot1.Set(".sub.a", 2); err != nil {
		t.Fatalf("%+v", err)
	}
	if err := dot1.Set(".sub.b", false); err != nil {
		t.Fatalf("%+v", err)
	}
	if err := dot1.Set(".sub.c", &dot1); err != nil {
		t.Fatalf("%+v", err)
	}
	if err := dot1.Set(".sub.d", "e"); err != nil {
		t.Fatalf("%+v", err)
	}
	if err := dot1.Set(".sub.e", 5.678); err != nil {
		t.Fatalf("%+v", err)
	}

	if a, err := dot1.Get(".sub.a"); err != nil || a != 2 {
		t.Fatalf("sub.a=(%T)%+v != 2", a, a)
	}
	if b, err := dot1.Get(".sub.b"); err != nil || b != false {
		t.Fatalf("sub.b=(%T)%+v != false", b, b)
	}
	if c, err := dot1.Get(".sub.c"); err != nil || c != &dot1 {
		t.Fatalf("sub.c=(%T)%+v != &dot1", c, c)
	}
	if d, err := dot1.Get(".sub.d"); err != nil || d != "e" {
		t.Fatalf("sub.d=(%T)%+v != \"e\"", d, d)
	}
	if e, err := dot1.Get(".sub.e"); err != nil || e != 5.678 {
		t.Fatalf("sub.e=(%T)%+v != 5.678", e, e)
	}

}

func TestArrays(t *testing.T) {
	s := []interface{}{
		1,
		true,
		nil,
		"d",
		3.141,
	}
	o := map[string]interface{}{
		"sub": s,
	}

	dot1 := dot.New(o)
	if a, err := dot1.Get(".sub[0]"); err != nil || a != 1 {
		t.Fatalf("a=(%T)%+v != 1: err=%+v", a, a, err)
	}
	if b, err := dot1.Get(".sub[1]"); err != nil || b != true {
		t.Fatalf("b=(%T)%+v != true: err=%+v", b, b, err)
	}
	if c, err := dot1.Get(".sub[2]"); err != nil || c != nil {
		t.Fatalf("c=(%T)%+v != nil: err=%+v", c, c, err)
	}
	if d, err := dot1.Get(".sub[3]"); err != nil || d != "d" {
		t.Fatalf("d=(%T)%+v != \"d\": err=%+v", d, d, err)
	}
	if e, err := dot1.Get(".sub[4]"); err != nil || e != 3.141 {
		t.Fatalf("e=(%T)%+v != 3.141: err=%+v", e, e, err)
	}

	//set values
	if err := dot1.Set(".sub[0]", 2); err != nil {
		t.Fatalf("%+v", err)
	}
	if err := dot1.Set(".sub[1]", false); err != nil {
		t.Fatalf("%+v", err)
	}
	if err := dot1.Set(".sub[2]", &dot1); err != nil {
		t.Fatalf("%+v", err)
	}
	if err := dot1.Set(".sub[3]", "e"); err != nil {
		t.Fatalf("%+v", err)
	}
	if err := dot1.Set(".sub[4]", 5.678); err != nil {
		t.Fatalf("%+v", err)
	}

	if a, err := dot1.Get(".sub[0]"); err != nil || a != 2 {
		t.Fatalf("a=(%T)%+v != 2: err=%+v", a, a, err)
	}
	if b, err := dot1.Get(".sub[1]"); err != nil || b != false {
		t.Fatalf("b=(%T)%+v != false: err=%+v", b, b, err)
	}
	if c, err := dot1.Get(".sub[2]"); err != nil || c != &dot1 {
		t.Fatalf("c=(%T)%+v != &dot1: err=%+v", c, c, err)
	}
	if d, err := dot1.Get(".sub[3]"); err != nil || d != "e" {
		t.Fatalf("d=(%T)%+v != \"e\": err=%+v", d, d, err)
	}
	if e, err := dot1.Get(".sub[4]"); err != nil || e != 5.678 {
		t.Fatalf("e=(%T)%+v != 5.678: err=%+v", e, e, err)
	}
}
