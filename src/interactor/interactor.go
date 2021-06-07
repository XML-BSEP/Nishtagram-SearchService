package interactor

import (
	"go.mongodb.org/mongo-driver/mongo"
	"search-service/http/handler"
	"search-service/repository"
	"search-service/usecase"
)

type interactor struct {
	db *mongo.Client
}

type Interactor interface {
	NewLocationRepository() repository.LocationRepo
	NewPostLocationRepository() repository.PostLocationRepo
	NewPostTagRepository() repository.PostTagRepo

	NewLocationUseCase() usecase.LocationUsecase
	NewPostLocationUseCase() usecase.PostLocationUsecase
	NewPostTagUseCase() usecase.PostTagUsecase

	NewAppHandler() AppHandler

	NewLocationHandler() handler.LocationHandler
	NewPostLocationHandler() handler.PostLocationHandler
	NewPostTagHandler() handler.PostTagHandler
	NewPostHandler() handler.PostHandler

}

func NewInteractor(db *mongo.Client) Interactor {
	return &interactor{db: db}
}

type appHandler struct {
	handler.LocationHandler
	handler.PostLocationHandler
	handler.PostTagHandler
	handler.PostHandler
}

type AppHandler interface {
	handler.LocationHandler
	handler.PostLocationHandler
	handler.PostTagHandler
	handler.PostHandler
}

func (i *interactor) NewLocationRepository() repository.LocationRepo {
	return repository.NewLocationRepo(i.db)
}

func (i *interactor) NewPostLocationRepository() repository.PostLocationRepo {
	return  repository.NewPostLocationRepo(i.db)
}

func (i *interactor) NewPostTagRepository() repository.PostTagRepo {
	return repository.NewPostTagRepo(i.db)
}

func (i *interactor) NewLocationUseCase() usecase.LocationUsecase {
	return usecase.NewLocationUsecase(i.NewLocationRepository())
}

func (i *interactor) NewPostLocationUseCase() usecase.PostLocationUsecase {
	return usecase.NewPostLocationUsecase(i.NewPostLocationRepository())
}

func (i *interactor) NewPostTagUseCase() usecase.PostTagUsecase {
	return usecase.NewPostTagUseCase(i.NewPostTagRepository())
}

func (i *interactor) NewLocationHandler() handler.LocationHandler {
	return handler.NewLocationHandler(i.NewLocationUseCase())
}

func (i *interactor) NewPostLocationHandler() handler.PostLocationHandler {
	return handler.NewPostLocationHandler(i.NewPostLocationUseCase())
}

func (i *interactor) NewPostTagHandler() handler.PostTagHandler {
	return handler.NewPostTagHandler(i.NewPostTagUseCase())
}

func (i *interactor) NewAppHandler() AppHandler{
	appHandler := &appHandler{}
	appHandler.LocationHandler = i.NewLocationHandler()
	appHandler.PostLocationHandler = i.NewPostLocationHandler()
	appHandler.PostTagHandler = i.NewPostTagHandler()
	appHandler.PostHandler = i.NewPostHandler()

	return appHandler
}

func (i *interactor) NewPostHandler() handler.PostHandler {
	return handler.NewPostHandler(i.NewPostTagUseCase(), i.NewPostLocationUseCase())
}




