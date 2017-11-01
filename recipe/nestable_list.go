package recipe

// TODO: Evaluate need of source func.

// NestableList ...
type NestableList []NestableItem

// NestableItem ...
type NestableItem struct {
	Item string
	List NestableList
}
