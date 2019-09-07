package domain

import "errors"

type TreeInterface interface {
	GetParentId() string
	SaveChilds(TreeInterface)
	GetId() string
}

func MakeTree(rootId string, reader []TreeInterface) (TreeInterface, error) {

	var writer TreeInterface
	writer = nil

	m := make(map[string][]int)

	for key, value := range reader {
		if value.GetId() == rootId {
			writer = reader[key]
		}
		m[value.GetParentId()] = append(m[value.GetParentId()], key)
	}
	if writer == nil {
		return nil, errors.New("root id not found")
	}

	//очередь указателей на объект
	var queue []*TreeInterface
	queue = append(queue, &writer)

	for i := 0; i < len(queue); i++ {
		if childsIds, ok := m[(*queue[i]).GetId()]; ok {
			for _, childId := range childsIds {
				queue = append(queue, &reader[childId])
				(*queue[i]).SaveChilds(reader[childId])
			}
		}
	}
	return writer, nil
}