package response

import (
	"qbills/models/domain"
	"qbills/models/schema"
	"qbills/models/web"
)

func ProductSchemaToProductDomain(product *schema.Product) *domain.Product {
	return &domain.Product{
		ID:            product.ID,
		ProductTypeID: product.ProductTypeID,
		AdminID:       product.AdminID,
		Name:          product.Name,
		Description:   product.Description,
		Price:         product.Price,
		TotalStock:    product.TotalStock,
		Size:          product.Size,
		Image:         product.Image,
	}
}

func ProductDomainToProductResponse(product *domain.Product) web.ProductResponse {
	return web.ProductResponse{
		ID:            product.ID,
		ProductTypeID: product.ProductTypeID,
		ProductType:   product.ProductType,
		AdminID:       product.AdminID,
		Admin:         product.Admin,
		Name:          product.Name,
		Description:   product.Description,
		Price:         product.Price,
		TotalStock:    product.TotalStock,
		Size:          product.Size,
		Image:         product.Image,
	}
}

func ProductDomainToProductCreateResponse(product *domain.Product) web.ProductCreateResponse {
	return web.ProductCreateResponse{
		ID:            product.ID,
		ProductTypeID: product.ProductTypeID,
		AdminID:       product.AdminID,
		Name:          product.Name,
		Description:   product.Description,
		Price:         product.Price,
		TotalStock:    product.TotalStock,
		Size:          product.Size,
		Image:         product.Image,
	}
}

func ProductDomainToProductUpdateResponse(product *domain.Product) web.ProductUpdateResponse {
	return web.ProductUpdateResponse{
		ID:            product.ID,
		ProductTypeID: product.ProductTypeID,
		Name:          product.Name,
		Description:   product.Description,
		Price:         product.Price,
		TotalStock:    product.TotalStock,
		Size:          product.Size,
		Image:         product.Image,
	}
}

func ConvertProductResponse(products []domain.Product) []web.ProductResponseCustom {
	var results []web.ProductResponseCustom

	for _, product := range products {

		Admins := AdminDomainToAdminDomainResponse(product.Admin)
		productResponse := web.ProductResponseCustom{
			ID:            product.ID,
			ProductTypeID: product.ProductTypeID,
			ProductType:   product.ProductType,
			AdminID:       product.AdminID,
			Admin:         Admins,
			Name:          product.Name,
			Description:   product.Description,
			Price:         product.Price,
			TotalStock:    product.TotalStock,
			Size:          product.Size,
			Image:         product.Image,
		}
		results = append(results, productResponse)
	}
	return results
}

func ProductResponseToProductCostumResponse(product web.ProductResponse) web.ProductResponseCustom {

	admin := AdminDomainToAdminDomainResponse(product.Admin)

	return web.ProductResponseCustom{
		ID:            product.ID,
		ProductTypeID: product.ProductTypeID,
		ProductType:   product.ProductType,
		AdminID:       product.AdminID,
		Admin:         admin,
		Name:          product.Name,
		Description:   product.Description,
		Price:         product.Price,
		TotalStock:    product.TotalStock,
		Size:          product.Size,
		Image:         product.Image,
	}
}

func ProductResponseToProductsCostumResponse(product web.ProductResponse) web.ProductResponseCustom {

	admin := AdminDomainToAdminDomainResponse(product.Admin)

	return web.ProductResponseCustom{
		ID:            product.ID,
		ProductTypeID: product.ProductTypeID,
		ProductType:   product.ProductType,
		AdminID:       product.AdminID,
		Admin:         admin,
		Name:          product.Name,
		Description:   product.Description,
		Price:         product.Price,
		TotalStock:    product.TotalStock,
		Size:          product.Size,
		Image:         product.Size,
	}
}

func ProductResponseToProductPreloadResponse(response domain.Product) domain.ProductPreloadResponse {
	return domain.ProductPreloadResponse{
		ID:            response.ID,
		ProductTypeID: response.ProductTypeID,
		AdminID:       response.AdminID,
		Name:          response.Name,
		Description:   response.Description,
		Price:         response.Price,
		TotalStock:    response.TotalStock,
		Size:          response.Size,
		Image:         response.Image,
	}
}

func DomainProductToDomainResponseProduct(response domain.Product) domain.ProductPreloadResponse {
	return domain.ProductPreloadResponse{
		ID:            response.ID,
		ProductTypeID: response.ProductTypeID,
		AdminID:       response.AdminID,
		Name:          response.Name,
		Description:   response.Description,
		Price:         response.Price,
		TotalStock:    response.TotalStock,
		Size:          response.Size,
		Image:         response.Image,
	}
}
