package main

import (
  "fmt"
)



// up -> i-1, j
// down -> i+1, j
// left -> i, j-1
// right -> i, j+1


func isValid(maze [][]int, visited [][]bool, i, j, x, y int) bool {
  
  if (i>=0 && j>=0 && i<=x && j<=y && !visited[i][j] && maze[i][j]==1 ) {
     fmt.Println("validating", i, j, true)
     return true
  }
    fmt.Println("validating", i, j, false)

  return false

}

func dequeue(queue []int) (int, []int) {
  fmt.Println("dequeuing", queue)
  q := queue[len(queue)-1]
  queue = queue[:len(queue)-1]
  fmt.Println("queue", q)
  return q, queue
}

func findPath(maze [][]int, x, y int) int {
  i,j := 0,0
  explored := 1
  visits := 0
  moves := 0
  foundPath := false
  queuei := make([]int, 0)
  queuej := make([]int, 0)
  visited := [][]bool{{ false, false, false, false},
                  { false, false, false, false},
                  { false, false, false, false},
                  { false, false, false, false}}

  visited[i][j] = true
  queuei = append(queuei, i)
  queuej = append(queuej, j)
  for (len(queuei) >0) {
    i,queuei=dequeue(queuei)
    j,queuej=dequeue(queuej)
    fmt.Println ("dequeue", i, j)
    
    fmt.Println("explored", explored)
    
    if ( i == x && j == y) {
      foundPath = true
      break
    }
    if isValid(maze, visited, i+1, j, x,y) {
      fmt.Println ("down")
      visited[i+1][j]=true
      queuei = append(queuei,i+1)
      queuej = append(queuej,j)
      fmt.Println ("down", i+1, j)
      //visits++
    }
    if isValid(maze, visited, i, j+1, x,y) {
      fmt.Println ("right")
      visited[i][j+1]=true
      queuei = append(queuei,i)
      queuej = append(queuej,j+1)
      fmt.Println ("right", i, j+1)
      //visits++
    }
    if isValid(maze, visited, i-1, j, x,y) {
      fmt.Println ("up")
      visited[i-1][j]=true
      queuei = append(queuei,i-1)
      queuej = append(queuej,j)
      fmt.Println ("up", i-1, j)
      //visits++
    }
    if isValid(maze, visited, i, j-1, x,y) {
      fmt.Println ("left")
      visited[i][j-1]=true
      queuei = append(queuei,i)
      queuej = append(queuej,j-1)
      fmt.Println ("left", i, j-1)
      //visits++
    }
    visits++
    explored--
    if explored == 0 {
      explored=visits
      visits=0
      moves++
      
    }
  }
 if foundPath{
   return moves
 }
 return 0

}

//    0 1 2 3
// 0{ 1 1 1 0}
// 1{ 1 1 0 1}
// 2{ 0 1 1 0}
// 3{ 0 1 1 1}


func main () {
  x,y := 3,3
  maze := [][]int{{ 0, 1, 1, 0},
                  { 1, 1, 0, 1},
                  { 0, 1, 1, 0},
                  { 0, 1, 1, 1}}
  moves := findPath(maze, x, y)
  fmt.Println("moves", moves)
  
}
