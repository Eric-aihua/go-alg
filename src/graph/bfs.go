package main

/**
广度优先算法
**/

func InitGraph()  map[string][]string{
	relationGraph  := make(map[string][]string)
	relationGraph["you"] = []string{"alice", "bob", "claire"}
	relationGraph["bob"] = []string{"anuj", "peggy"}
	relationGraph["alice"] = []string{"peggy"}
	relationGraph["claire"] = []string{"thom", "jonny"}
	relationGraph["anuj"] = []string{}
	relationGraph["peggy"] = []string{}
	relationGraph["thom"] = []string{}
	relationGraph["jonny"] = []string{}
	//fmt.Println(relationGraph)
	return relationGraph
}

//func BFSSearch() {
//	relationsGraph :=InitGraph()
//	searchedQueue := queue.NewQueue()
//	//将you添加到队列
//	searchedQueue.Offer(relationsGraph["you"])
//	for{
//		if ! searchedQueue.IsEmpty(){
//			return
//		}else{
//			currentPerson := searchedQueue.Poll()
//			if isSelectedMan(){
//
//			}
//		}
//	}
//
//
//}
//
//func isSelectedMan(currentPerson string) bool {
//	if
//
//}

func main() {
	//BFSSearch()
}

