package _interface

type ActionRepository interface {
	Insert(action ActionDO) (int, error)
	FindByID(actionID int) (*ActionDO, error)
}
