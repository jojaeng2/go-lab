package main

import (
	"fmt"
)

func main() {
	s := "example string"

	for i, c := range s {
		fmt.Printf("%c->%d ", c, i) // prints "e->0 x->1 a->2 m->3 p->4 l->5 e->6  ->7 s->8 t->9 r->10 i->11 n->12 g->13 "
	}
}
