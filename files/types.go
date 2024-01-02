package files

type Tree struct {
	Name     string
	IsDir    bool
	Children []*Tree
	Value    []byte
}
