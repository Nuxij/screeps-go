package enum

type EnumItem struct {
	index int
	name  string
}

type Enum struct {
	items []EnumItem
}

func (enum Enum) Name(findIndex int) string {
	for _, item := range enum.items {
		if item.index == findIndex {
			return item.name
		}
	}
	return "ID not found"
}

func (enum Enum) Index(findName string) int {
	for _, item := range enum.items {
		if findName == item.name {
			return item.index
		}
	}
	return 0
}

func (enum Enum) Last() (int, string) {
	n := len(enum.items)
	return n - 1, enum.items[n-1].name
}
