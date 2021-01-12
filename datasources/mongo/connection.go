package mongo

type DbService struct {
	Connection string
}

func NewDbService(username, password, host, schema, connection string) *DbService {
	return &DbService{Connection: "mongo_connection"}
}
