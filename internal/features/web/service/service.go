package web_service

type WebService struct {
	webRepository WebRepository
}

type WebRepository interface {
	GetFile(FilePath string) ([]byte, error)
}

func NewWebService(w WebRepository) *WebService {
	return &WebService{webRepository: w}
}
