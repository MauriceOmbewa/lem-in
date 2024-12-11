package utils

type Ant struct {
	Id          int
	Path        []*Room
	PathIndex   int
	CurrentRoom *Room
	HasReached  bool
}

type Room struct {
	Name    string
	IsStart bool
	IsEnd   bool
}

type AntFarm struct {
	Ants  []*Ant
	Move   int
	Rooms map[string]*Room
	Start *Room
	End   *Room
}

type Path struct {
	Rooms []*Room
}
