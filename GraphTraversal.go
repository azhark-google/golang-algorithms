package main

import (
  "fmt"
)

type graph struct {
  vertices []*Vertex
  edges map[*Vertex][]*Vertex
}

type Vertex struct {
  item int
}

func (g *graph) insertNode(v *Vertex) {
  if g.vertices == nil {
    g.vertices = make([]*Vertex, 0)
  }
    g.vertices = append(g.vertices, v)

}

func (g *graph) insertEdges(v1 *Vertex, v2 *Vertex) {
  if g.edges == nil {
    g.edges = make(map[*Vertex][]*Vertex)
  }
  g.edges[v1] = append(g.edges[v1], v2)
  g.edges[v2] = append(g.edges[v2], v1)
}

func traverse(g *graph) map[*Vertex]bool {
  queue := make ([]*Vertex, 0)
  visited := make (map[*Vertex]bool)
  queue = append(queue,g.vertices[0])
  for (len (queue) > 0) {
    v := queue[len(queue)-1]
    queue = queue[:len(queue)-1]
    for _, node := range g.edges[v] {
      if _, ok :=  visited[node]; !ok {
        visited[node] = true
        queue = append(queue, node)
      }
    }
  }
  return visited
}


func main () {
  g := &graph{ }
  v1 := &Vertex{item: 5 }
  v2 := &Vertex{item: 10 }
  g.insertNode(v1)
  g.insertNode(v2)
  g.insertEdges(v1,v2)
  fmt.Println(g)
  v := traverse(g)
  for node, _ := range v {
    fmt.Println(node.item)
  }

}
