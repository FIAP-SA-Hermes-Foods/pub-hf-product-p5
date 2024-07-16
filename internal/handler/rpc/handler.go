package rpc

import (
	"context"
	l "pub-hf-product-p5/external/logger"
	"pub-hf-product-p5/internal/core/application"
	"pub-hf-product-p5/internal/core/domain/entity/dto"
	cp "pub-hf-product-p5/product_pub_proto"
)

type HandlerGRPC interface {
	Handler() *handlerGRPC
}

type handlerGRPC struct {
	app application.Application
	cp.UnimplementedProductServer
}

func NewHandler(app application.Application) HandlerGRPC {
	return &handlerGRPC{app: app}
}

func (h *handlerGRPC) Handler() *handlerGRPC {
	return h
}

func (h *handlerGRPC) GetProductByID(ctx context.Context, req *cp.GetProductByIDRequest) (*cp.GetProductByIDResponse, error) {
	msgID := ctx.Value(l.MessageIDKey).(string)
	msgID = l.MessageID(msgID)

	if err := h.app.GetProductByID(msgID, req.Uuid); err != nil {
		return nil, err
	}

	return nil, nil
}

func (h *handlerGRPC) CreateProduct(ctx context.Context, req *cp.CreateProductRequest) (*cp.CreateProductResponse, error) {
	msgID := ctx.Value(l.MessageIDKey).(string)
	msgID = l.MessageID(msgID)

	input := dto.RequestProduct{
		Name:          req.Name,
		Category:      req.Category,
		Image:         req.Image,
		Description:   req.Description,
		Price:         float64(req.Price),
		CreatedAt:     req.CreatedAt,
		DeactivatedAt: req.DeactivatedAt,
	}

	if err := h.app.SaveProduct(msgID, input); err != nil {
		return nil, err
	}

	return nil, nil
}

func (h *handlerGRPC) GetProductByCategory(ctx context.Context, req *cp.GetProductByCategoryRequest) (*cp.GetProductByCategoryResponse, error) {
	msgID := ctx.Value(l.MessageIDKey).(string)
	msgID = l.MessageID(msgID)

	if err := h.app.GetProductByCategory(msgID, req.Category); err != nil {
		return nil, err
	}

	return nil, nil
}

func (h *handlerGRPC) UpdateProduct(ctx context.Context, req *cp.UpdateProductRequest) (*cp.UpdateProductResponse, error) {
	msgID := ctx.Value(l.MessageIDKey).(string)
	msgID = l.MessageID(msgID)

	input := dto.RequestProduct{
		Name:          req.Name,
		Category:      req.Category,
		Image:         req.Image,
		Description:   req.Description,
		Price:         float64(req.Price),
		CreatedAt:     req.CreatedAt,
		DeactivatedAt: req.DeactivatedAt,
	}

	if err := h.app.UpdateProductByID(msgID, req.Uuid, input); err != nil {
		return nil, err
	}

	return nil, nil
}

func (h *handlerGRPC) DeleteProductByID(ctx context.Context, req *cp.DeleteProductByIDRequest) (*cp.DeleteProductByIDResponse, error) {
	msgID := ctx.Value(l.MessageIDKey).(string)
	msgID = l.MessageID(msgID)

	if err := h.app.DeleteProductByID(msgID, req.Uuid); err != nil {
		return nil, err
	}

	return nil, nil
}
