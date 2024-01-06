package status

type Tree struct {
	Name     string
	IsDir    bool
	Children []*Tree
	Value    []byte
	Tracked bool 
	Modified bool 
	Staged bool
}
