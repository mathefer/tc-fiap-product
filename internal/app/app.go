package app

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/fx"

	productController "github.com/mathefer/tc-fiap-product/internal/product/controller"
	productRepositories "github.com/mathefer/tc-fiap-product/internal/product/domain/repositories"
	productApiController "github.com/mathefer/tc-fiap-product/internal/product/infrastructure/api/controller"
	productPersistence "github.com/mathefer/tc-fiap-product/internal/product/infrastructure/persistence"
	productPresenter "github.com/mathefer/tc-fiap-product/internal/product/presenter"
	productUseCasesAdd "github.com/mathefer/tc-fiap-product/internal/product/usecase/addProduct"
	productUseCasesDelete "github.com/mathefer/tc-fiap-product/internal/product/usecase/deleteProduct"
	productUseCasesGet "github.com/mathefer/tc-fiap-product/internal/product/usecase/getProduct"
	productUseCasesUpdate "github.com/mathefer/tc-fiap-product/internal/product/usecase/updateProduct"

	"github.com/mathefer/tc-fiap-product/pkg/rest"
	"github.com/mathefer/tc-fiap-product/pkg/storage/postgres"
)

func InitializeApp() *fx.App {
	return fx.New(
		fx.Provide(
			postgres.NewPostgresDB,
			fx.Annotate(productPersistence.NewProductRepositoryImpl, fx.As(new(productRepositories.ProductRepository))),
			fx.Annotate(productController.NewProductControllerImpl, fx.As(new(productController.ProductController))),
			fx.Annotate(productPresenter.NewProductPresenterImpl, fx.As(new(productPresenter.ProductPresenter))),
			fx.Annotate(productUseCasesAdd.NewAddProductUseCaseImpl, fx.As(new(productUseCasesAdd.AddProductUseCase))),
			fx.Annotate(productUseCasesGet.NewGetProductUseCaseImpl, fx.As(new(productUseCasesGet.GetProductUseCase))),
			fx.Annotate(productUseCasesUpdate.NewUpdateProductUseCaseImpl, fx.As(new(productUseCasesUpdate.UpdateProductUseCase))),
			fx.Annotate(productUseCasesDelete.NewDeleteProductUseCaseImpl, fx.As(new(productUseCasesDelete.DeleteProductUseCase))),
			chi.NewRouter,
			func(
				productController productController.ProductController) []rest.Controller {
				return []rest.Controller{
					productApiController.NewProductController(productController),
				}
			},
		),
		fx.Invoke(registerRoutes),
		fx.Invoke(startHTTPServer),
	)
}

func registerRoutes(r *chi.Mux, controllers []rest.Controller) {
	r.Use(middleware.Logger)

	// Swagger UI
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), // The URL pointing to API definition
	))

	for _, controller := range controllers {
		controller.RegisterRoutes(r)
	}
}

func startHTTPServer(lc fx.Lifecycle, r *chi.Mux) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				log.Printf("Starting HTTP server on :%s", port)
				if err := http.ListenAndServe(":"+port, r); err != nil {
					log.Fatalf("Failed to start HTTP server: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Shutting down HTTP server gracefully")
			return nil
		},
	})
}

