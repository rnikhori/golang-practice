package main

import "fmt"

type SparseSet struct {
	dense []int 
	sparse []int
	count int
	n int
} 

func NewSparseSet(n int) *SparseSet{
	return &SparseSet{
		dense: make([]int,n),
		sparse: make([]int, n),
		count:0,
		n: n,
	}
}

func(s* SparseSet) Contains(x int) bool{
	idx := s.sparse[x]
	return idx < s.count && s.dense[idx] == x 
}

func (s *SparseSet) Insert(x int){
	if s.Contains(x){
		return
	}
	s.dense[s.count] = x
	s.sparse[x] = s.count
	s.count++
}


func (s *SparseSet) Remove(x int){ /// [2,5,0] count =3 remove(0)
	if !s.Contains(x){
		return
	}

	idx := s.sparse[x] //idx =2
	s.dense[idx] = s.dense[s.count -1] //s.dense[2] = 0
	s.sparse[s.dense[idx]]= idx //
	s.count--
}

func (s *SparseSet) Clear(){
	s.count = 0
}

func (s *SparseSet) Iterate() []int{
	return s.dense[:s.count]
}

func main(){

	s := NewSparseSet(6)

	s.Insert(2)
	s.Insert(5)
	s.Insert(0)
	fmt.Println("after inserts:", s.Iterate())      // expect [2 5 0]

    fmt.Println("contains 5?", s.Contains(5))        // expect true
    fmt.Println("contains 3?", s.Contains(3)) 

	 s.Remove(5)
    fmt.Println("after remove(5):", s.Iterate())      // expect [2 0] (order may differ)
    fmt.Println("contains 5?", s.Contains(5)) 


    s.Insert(2)  // duplicate insert — should NOT double-count
    fmt.Println("after duplicate insert(2):", s.Iterate())  // still [2 0]

	 s.Clear()
    fmt.Println("after clear:", s.Iterate())          // expect [] (empty)
    fmt.Println("contains 2?", s.Contains(2))
	
	    s.Insert(4)
    fmt.Println("after insert(4) post-clear:", s.Iterate())  // expect [4]
}