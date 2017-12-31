package easyrouter

// import (
// 	"log"
// 	"strings"
// )

// type Trie struct {
// 	Nodes map[string]*Node
// }

// type Node struct {
// 	Value    string
// 	Children map[string]*Node
// 	Route    Route //is leaf
// }

// func (s *Server) makeTrie() *Trie {
// 	t := Trie{
// 		Nodes: make(map[string]*Node),
// 	}
// 	for _, route := range s.Routes {
// 		t.Nodes[strings.ToUpper(route.Method)] = &Node{
// 			Value: strings.ToUpper(route.Method),
// 		}
// 		strArray := strings.Split(route.Path, "/")
// 		t.assign(strArray, t.Nodes[strings.ToUpper(route.Method)], route)
// 	}
// 	return &t
// }

// func (s *Server) traverseTrie(t Trie) *Route {
// 	return nil
// }

// func (t *Trie) assign(strArray []string, node *Node, r Route) {
// 	log.Print(strArray)
// 	for i, s := range strArray {
// 		if s == "" {
// 			continue
// 		}
// 		n := &Node{
// 			Value: s,
// 		}
// 		if i == len(strArray)-1 {
// 			n.Route = r
// 		}
// 		if node.Children == nil {
// 			node.Children = make(map[string]*Node)
// 		}
// 		node.Children[s] = n
// 		t.assign(strArray[1:], n, r)
// 	}
// }
