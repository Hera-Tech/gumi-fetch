package types

type ShowStore interface {
	List() ([]Show, error)
	Create(Show) error
	Delete(int) error
}
