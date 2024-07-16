package application

import (
	"context"
	l "pub-hf-product-p5/external/logger"
	ps "pub-hf-product-p5/external/strings"
	"pub-hf-product-p5/internal/core/domain/broker"
	"pub-hf-product-p5/internal/core/domain/entity/dto"
)

type Application interface {
	GetProductByID(msgID string, uuid string) error
	SaveProduct(msgID string, product dto.RequestProduct) error
	UpdateProductByID(msgID string, id string, product dto.RequestProduct) error
	GetProductByCategory(msgID string, category string) error
	DeleteProductByID(msgID string, id string) error
}

type application struct {
	ctx           context.Context
	productBroker broker.ProductBroker
}

func NewApplication(ctx context.Context, productBroker broker.ProductBroker) Application {
	return application{
		ctx:           ctx,
		productBroker: productBroker,
	}
}

func (app application) GetProductByID(msgID string, uuid string) error {
	app.setMessageIDCtx(msgID)
	l.Infof(msgID, "GetProductByIDApp: ", " | ", uuid)

	inputBroker := dto.ProductBroker{
		UUID:      uuid,
		MessageID: msgID,
	}

	if err := app.productBroker.GetProductByID(inputBroker); err != nil {
		l.Errorf(msgID, "GetProductByIDApp error: ", " | ", err)
		return err
	}

	l.Infof(msgID, "GetProductByIDApp output: ", " | ", "message sent with success!")
	return nil
}

func (app application) SaveProduct(msgID string, product dto.RequestProduct) error {
	app.setMessageIDCtx(msgID)
	l.Infof(msgID, "SaveProductApp: ", " | ", ps.MarshalString(product))

	inputBroker := dto.ProductBroker{
		UUID:          product.UUID,
		MessageID:     msgID,
		Name:          product.Name,
		Category:      product.Category,
		Image:         product.Image,
		Description:   product.Description,
		Price:         product.Price,
		CreatedAt:     product.CreatedAt,
		DeactivatedAt: product.DeactivatedAt,
	}

	if err := app.productBroker.SaveProduct(inputBroker); err != nil {
		l.Errorf(msgID, "SaveProductApp error: ", " | ", err)
		return err
	}

	l.Infof(msgID, "SaveProductApp output: ", " | ", "message sent with success!")
	return nil
}

func (app application) GetProductByCategory(msgID string, category string) error {
	app.setMessageIDCtx(msgID)
	l.Infof(msgID, "GetProductByCategoryApp: ", " | ", category)

	inputProker := dto.ProductBroker{
		MessageID: msgID,
		Category:  category,
	}

	if err := app.productBroker.GetProductByCategory(inputProker); err != nil {
		l.Errorf(msgID, "GetProductByCategoryApp error: ", " | ", err)
		return err
	}

	l.Infof(msgID, "GetProductByCategoryApp output: ", " | ", "message sent with success!")
	return nil
}

func (app application) UpdateProductByID(msgID string, id string, product dto.RequestProduct) error {
	app.setMessageIDCtx(msgID)
	l.Infof(msgID, "UpdateProductByIDApp: ", " | ", id, " | ", ps.MarshalString(product))

	inputBroker := dto.ProductBroker{
		UUID:          id,
		MessageID:     msgID,
		Name:          product.Name,
		Category:      product.Category,
		Image:         product.Image,
		Description:   product.Description,
		Price:         product.Price,
		CreatedAt:     product.CreatedAt,
		DeactivatedAt: product.DeactivatedAt,
	}

	if err := app.productBroker.UpdateProductByID(inputBroker); err != nil {
		l.Errorf(msgID, "UpdateProductByIDApp error: ", " | ", err)
		return err
	}

	l.Infof(msgID, "UpdateProductByIDApp output: ", " | ", "message sent with success!")
	return nil
}

func (app application) DeleteProductByID(msgID string, id string) error {
	app.setMessageIDCtx(msgID)
	l.Infof(msgID, "DeleteProductByIDApp: ", " | ", id)

	inputBroker := dto.ProductBroker{
		UUID:      id,
		MessageID: msgID,
	}
	if err := app.productBroker.DeleteProductByID(inputBroker); err != nil {
		l.Errorf(msgID, "DeleteProductByIDApp error: ", " | ", err)
	}

	l.Infof(msgID, "DeleteProductByIDApp output: ", " | ", "message sent with success!")
	return nil
}

func (app application) setMessageIDCtx(msgID string) {
	if app.ctx == nil {
		app.ctx = context.WithValue(context.Background(), l.MessageIDKey, msgID)
		return
	}
	app.ctx = context.WithValue(app.ctx, l.MessageIDKey, msgID)
}
