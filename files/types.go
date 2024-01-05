package files

type Tree struct {
	Name     string
	IsDir    bool
	Children []*Tree
	Value    []byte
	Tracked bool // (may become handy in `gitpot status`)
}
