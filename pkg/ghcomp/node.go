package ghcomp

type node struct {
	value    byte
	children map[byte]*node
	parent   *node
}

func (n *node) entree(remaining []byte) {
	if len(remaining) == 0 {
		return
	}

	if n.children[remaining[0]] == nil {
		n.children[remaining[0]] = &node{
			value:    remaining[0],
			children: make(map[byte]*node),
			parent:   n,
		}
	}

	n.children[remaining[0]].entree(remaining[1:])
}
