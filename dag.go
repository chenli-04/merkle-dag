package merkledag

//first
import "hash"

func Add(store KVStore, node Node, h hash.Hash) ([]byte, error) {

	h.Reset()

	if node.Type() == FILE {
		fileNode := node.(File)
		data := fileNode.Bytes()
		_, err := h.Write(data)
		if err != nil {
			return nil, err
		}
		return h.Sum(nil), nil
	}

	if node.Type() == DIR {
		dirNode := node.(Dir)
		it := dirNode.It()
		for it.Next() {
			file := it.Node().(File)
			data := file.Bytes()
			_, err := h.Write(data)
			if err != nil {
				return nil, err
			}
		}
		return h.Sum(nil), nil
	}

	return nil, nil
}
