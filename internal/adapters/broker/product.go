package broker

import (
	l "pub-hf-product-p5/external/logger"
	ps "pub-hf-product-p5/external/strings"
	sqsBroker "pub-hf-product-p5/internal/core/broker"
	pBroker "pub-hf-product-p5/internal/core/domain/broker"
	"pub-hf-product-p5/internal/core/domain/entity/dto"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

var _ pBroker.ProductBroker = (*productBroker)(nil)

type productBroker struct {
	queueURL string
	broker   sqsBroker.SQSBroker
}

func NewProductBroker(broker sqsBroker.SQSBroker, queueURL string) *productBroker {
	return &productBroker{broker: broker, queueURL: queueURL}
}

func (p *productBroker) GetProductByID(input dto.ProductBroker) error {
	l.Infof(input.MessageID, "msg broker input: ", " | ", ps.MarshalString(input))

	msgBody := ps.MarshalString(input)

	inPub := &sqs.SendMessageInput{
		QueueUrl:    &p.queueURL,
		MessageBody: &msgBody,
	}

	if _, err := p.broker.Pub(inPub); err != nil {
		return err
	}

	l.Infof(input.MessageID, "Message published successfully: ", " | ", ps.MarshalString(input))
	return nil
}

func (p *productBroker) SaveProduct(input dto.ProductBroker) error {
	l.Infof(input.MessageID, "msg broker input: ", " | ", ps.MarshalString(input))

	msgBody := ps.MarshalString(input)

	inPub := &sqs.SendMessageInput{
		QueueUrl:    &p.queueURL,
		MessageBody: &msgBody,
	}

	if _, err := p.broker.Pub(inPub); err != nil {
		return err
	}

	l.Infof(input.MessageID, "Message published successfully: ", " | ", ps.MarshalString(input))
	return nil
}

func (p *productBroker) UpdateProductByID(input dto.ProductBroker) error {
	l.Infof(input.MessageID, "msg broker input: ", " | ", ps.MarshalString(input))

	msgBody := ps.MarshalString(input)

	inPub := &sqs.SendMessageInput{
		QueueUrl:    &p.queueURL,
		MessageBody: &msgBody,
	}

	if _, err := p.broker.Pub(inPub); err != nil {
		return err
	}

	l.Infof(input.MessageID, "Message published successfully: ", " | ", ps.MarshalString(input))
	return nil
}

func (p *productBroker) GetProductByCategory(input dto.ProductBroker) error {
	l.Infof(input.MessageID, "msg broker input: ", " | ", ps.MarshalString(input))

	msgBody := ps.MarshalString(input)

	inPub := &sqs.SendMessageInput{
		QueueUrl:    &p.queueURL,
		MessageBody: &msgBody,
	}

	if _, err := p.broker.Pub(inPub); err != nil {
		return err
	}

	l.Infof(input.MessageID, "Message published successfully: ", " | ", ps.MarshalString(input))
	return nil
}

func (p *productBroker) DeleteProductByID(input dto.ProductBroker) error {
	l.Infof(input.MessageID, "msg broker input: ", " | ", ps.MarshalString(input))

	msgBody := ps.MarshalString(input)

	inPub := &sqs.SendMessageInput{
		QueueUrl:    &p.queueURL,
		MessageBody: &msgBody,
	}

	if _, err := p.broker.Pub(inPub); err != nil {
		return err
	}

	l.Infof(input.MessageID, "Message published successfully: ", " | ", ps.MarshalString(input))
	return nil
}
