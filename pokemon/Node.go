package pokemon

var PokemonTypes = [18]string{"Normal", "Fire", "Water", "Electric", "Grass", "Ice", "Fighting", "Poison", "Ground", "Flying", "Psychic", "Bug", "Rock", "Ghost", "Dragon", "Dark", "Steel", "Fairy"}

/*
Effectiveness[attack][defense]
*/
var Effectiveness = [18][18]int{
	{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 05, 00, 10, 10, 05, 10},
	{10, 05, 05, 10, 20, 20, 10, 10, 10, 10, 10, 20, 05, 10, 05, 10, 20, 10},
	{10, 20, 05, 10, 05, 10, 10, 10, 20, 10, 10, 10, 20, 10, 05, 10, 10, 10},
	{10, 10, 20, 05, 05, 10, 10, 10, 00, 20, 10, 10, 10, 10, 05, 10, 10, 10},
	{10, 05, 20, 10, 05, 10, 10, 05, 20, 05, 10, 05, 20, 10, 05, 10, 05, 10},
	{10, 05, 05, 10, 20, 05, 10, 10, 20, 20, 10, 10, 10, 10, 20, 10, 05, 10},
	{20, 10, 10, 10, 10, 20, 10, 05, 10, 05, 05, 05, 20, 00, 10, 20, 20, 05},
	{10, 10, 10, 10, 20, 10, 10, 05, 05, 10, 10, 10, 05, 05, 10, 10, 00, 20},
	{10, 20, 10, 20, 05, 10, 10, 20, 10, 00, 10, 05, 20, 10, 10, 10, 20, 10},
	{10, 10, 10, 05, 20, 10, 20, 10, 10, 10, 10, 20, 05, 10, 10, 10, 05, 10},
	{10, 10, 10, 10, 10, 10, 20, 20, 10, 10, 05, 10, 10, 10, 10, 00, 05, 10},
	{10, 05, 10, 10, 20, 10, 05, 05, 10, 05, 20, 10, 10, 05, 10, 20, 05, 05},
	{10, 20, 10, 10, 10, 20, 05, 10, 05, 20, 10, 20, 10, 10, 10, 10, 05, 10},
	{00, 10, 10, 10, 10, 10, 10, 10, 10, 10, 20, 10, 10, 20, 10, 05, 10, 10},
	{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 20, 10, 05, 00},
	{10, 10, 10, 10, 10, 10, 05, 10, 10, 10, 20, 10, 10, 20, 10, 05, 10, 05},
	{10, 05, 05, 05, 10, 20, 10, 10, 10, 10, 10, 10, 20, 10, 10, 10, 05, 20},
	{10, 05, 10, 10, 10, 10, 20, 05, 10, 10, 10, 10, 10, 10, 20, 20, 05, 10},
}

type Node struct {
	Parent            *Node
	Type               int
	FirstUncoveredType int //not important
}

func NewNode(parent *Node, pokemonType int) *Node {
	n := new(Node)
	n.Parent = parent
	n.Type = pokemonType
	n.FirstUncoveredType = -1
	return n
}

//return true if this node or one of its ancestor cover the pokemonType
func (o Node) Cover(pokemonType int) bool {
	if Effectiveness[o.Type][pokemonType] == 20 {
		return true
	}
	for o.Parent != nil {
		o = *o.Parent
		if Effectiveness[o.Type][pokemonType] == 20 {
			return true
		}
	}

	return false;
}

//set the FirstUncoveredType
func (o* Node) computeFistUncoveredType() int {
	for i:=0; i<len(PokemonTypes); i++ {
		if !o.Cover(i) {
			o.FirstUncoveredType = i
			return i
		}
	}
	o.FirstUncoveredType = -1
	return -1;
}

//export the node and all its ancestor to an array
func (o *Node) Export() []string {
	resu := make([]string, 1);
	resu[0] = PokemonTypes[o.Type]

	for ;o.Parent!=nil; {
		o = o.Parent
		resu = append(resu, PokemonTypes[o.Type])
	}

	return resu
}


