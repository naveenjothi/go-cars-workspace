package models

type PagingModel struct {
	Limit  int16
	Offset int16
}

func NewPagingModel() *PagingModel {
	return &PagingModel{
		Limit:  10,
		Offset: 0,
	}
}
