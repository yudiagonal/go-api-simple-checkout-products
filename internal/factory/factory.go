package factory

import (
	"database/sql"
	"net/http"
	"time"

	brandHandler "github.com/yudiagonal/go-api-simple-checkout-products/internal/app/brand/delivery/http"
	BrandRepository "github.com/yudiagonal/go-api-simple-checkout-products/internal/app/brand/repository"
	BrandService "github.com/yudiagonal/go-api-simple-checkout-products/internal/app/brand/service"

	productHandler "github.com/yudiagonal/go-api-simple-checkout-products/internal/app/product/delivery/http"
	ProductRepository "github.com/yudiagonal/go-api-simple-checkout-products/internal/app/product/repository"
	ProductService "github.com/yudiagonal/go-api-simple-checkout-products/internal/app/product/service"

	orderHandler "github.com/yudiagonal/go-api-simple-checkout-products/internal/app/order/delivery/http"
	OrderRepository "github.com/yudiagonal/go-api-simple-checkout-products/internal/app/order/repository"
	OrderService "github.com/yudiagonal/go-api-simple-checkout-products/internal/app/order/service"
)

func RegisterHandlers(mux *http.ServeMux, db *sql.DB) {

	const contextTimeout = 2 * time.Second

	brandRepository := BrandRepository.NewBrand(db)
	brandService := BrandService.NewBrandService(brandRepository, contextTimeout)
	brandHandler.NewBrandHandlers(mux, brandService)

	productRepository := ProductRepository.NewProduct(db)
	productService := ProductService.NewProductService(productRepository, brandService, contextTimeout)
	productHandler.NewProductHandler(mux, productService)

	orderRepository := OrderRepository.NewOrder(db)
	orderService := OrderService.NewOrderService(orderRepository, productService, contextTimeout)
	orderHandler.NewOrderHandler(mux, orderService)

}
