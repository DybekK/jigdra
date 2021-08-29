package repository

type Repository interface {
	FindById(string)
	FindAll()
	Create()
}
