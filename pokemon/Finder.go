package pokemon

var N int = len(Effectiveness)

func contains(list *[]string, a string) bool {
	for _, b := range *list {
		if b == a {
			return true
		}
	}
	return false
}

func isSame(team1 *[]string, team2 *[]string) bool {
	for _,v:=range *team1 {
		if !contains(team2, v) {
			return false
		}
	}
	return true
}


func teamExisted(teams *[][]string, t *[]string) bool {
	for _,v:=range *teams {
		if isSame(&v, t) {
			return true
		}
	}
	return false
}



func Find() [][]string {
	var optimizedTeams [][]string
	bestCount := N + 1

	queue := make(NodeHeap, 0)

	//init the queue with all the attacker which cover the type 0
	for i := 0; i < N; i++ {
		if Effectiveness[i][0] == 20 {
			queue.Push(NewNode(nil, i))
		}
	}

	for len(queue) > 0 {
		currentNode := queue.Pop()

		firstUncoveredType := currentNode.computeFistUncoveredType()
		if firstUncoveredType == -1 { //every type is covered, we found a complete team, so print it
			completeTeam := currentNode.Export()
			countMembers := len(completeTeam)
			if countMembers < bestCount { //the complete team beat the old recorded best teams
				bestCount = countMembers
				//reset the optimzed teams record, remove every old team

				optimizedTeams = make([][]string, 1)
				optimizedTeams[0] = completeTeam
			} else if countMembers == bestCount { //found another optimize team
				//check if this team was already added
				if !teamExisted(&optimizedTeams, &completeTeam) {
					optimizedTeams = append(optimizedTeams, completeTeam)
				}
			}

		} else {
			//find all the node which cover this type, add to the heap as a child node of the current node
			for i := 0; i < N; i++ {
				if Effectiveness[i][firstUncoveredType] == 20 {
					queue.Push(NewNode(currentNode, i))
				}
			}
		}
	}
	return optimizedTeams
}
