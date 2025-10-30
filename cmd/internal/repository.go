package internal

type MongoClient interface {
}

type Repository struct {
	MongoClient MongoClient
}

func NewRepository(mongoCLient MongoClient) *Repository {
	return &Repository{MongoClient: mongoCLient}
}
