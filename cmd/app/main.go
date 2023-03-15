package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"edmar.lima/edmarlima/product-api/internal/infra/akafka"
	"edmar.lima/edmarlima/product-api/internal/infra/repository"
	"edmar.lima/edmarlima/product-api/internal/infra/web"
	"edmar.lima/edmarlima/product-api/internal/usecase"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:3306)/products")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	repository := repository.NewProductRepositoryMysql(db)
	createProductUseCase := usecase.NewCreateProductUseCase(repository)
	listProducsUseCase := usecase.NewListProductsUseCase(repository)

	productsHandlers := web.NewProductsHandlers(createProductUseCase, listProducsUseCase)

	//Api web
	route := chi.NewRouter()

	route.Post("/products", productsHandlers.CreateProductHandler)
	route.Get("/products", productsHandlers.ListProductHandler)

	go http.ListenAndServe(":8000", route)

	//Kafka
	msgChan := make(chan *kafka.Message)
	go akafka.Consume([]string{"products"}, "host.docker.internal:9094", msgChan)

	for msg := range msgChan {
		dto := usecase.CreateProductInputDto{}

		err := json.Unmarshal(msg.Value, &dto)

		if err != nil {
			//logar erro
		}

		_, err = createProductUseCase.Execute(dto)
	}

}
