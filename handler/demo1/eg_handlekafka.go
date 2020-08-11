package demo1

import (
	"context"
	"time"

	"github.com/Shopify/sarama"
	"github.com/gin-gonic/gin"
	"github.com/pquerna/ffjson/ffjson"

	"use-gin/config"
	cfg "use-gin/config"
	"use-gin/errcode"
	"use-gin/handler"
	"use-gin/logger"
	"use-gin/model/kafka"
)

const consumeTaskTimeout = 6 * time.Minute

type ProduceTask struct {
	ID       string
	StrDemo1 string `json:"str_demo1"`
	StrDemo2 string `json:"str_demo2"`
}

type ConsumeTask struct {
	ID       string
	StrDemo1 string `json:"str_demo1"`
	StrDemo2 string `json:"str_demo2"`
}

type consumeHandler struct {
	c         *gin.Context
	taskID    string
	ctxCancel func()
}

func (consumeHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (consumeHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h consumeHandler) ConsumeClaim(
	_ sarama.ConsumerGroupSession,
	claim sarama.ConsumerGroupClaim,
) error {
	for v := range claim.Messages() {
		// do not crash when received the invalid topic messages.
		defer func() {
			if err := recover(); err != nil {
				logger.Log.Errorf("meet up panic: %v", err)
			}
		}()

		var ct ConsumeTask
		if err := ffjson.Unmarshal(v.Value, &ct); err != nil {
			logger.Log.Errorf("kafka message unmarshal: %s, error: %v", v, err)
			continue
		}
		logger.Log.Infof(
			"received from %s: %+v",
			config.Config().Kafka.ConsumerTopic,
			ct,
		)

		if ct.ID == h.taskID {

			// your specific logic

			h.ctxCancel()
			return nil
		}
	}
	return nil
}

func HandleKafkaDemo(c *gin.Context) {
	// producer
	produceTask := ProduceTask{}
	if err := addTaskToTopic(produceTask); err != nil {
		err1 := errcode.New(errcode.InternalServerError, err)
		handler.SendResponse(c, err1, nil)
		return
	}
	handler.SendResponse(c, nil, nil)

	// consumer
	taskID := produceTask.ID
	consumeTopics := []string{cfg.Config().Kafka.ConsumerTopic}
	ctx, cancel := context.WithCancel(context.Background())
	handler := consumeHandler{c, taskID, cancel}

	go kafka.ConsumeTopics(
		ctx,
		cancel,
		consumeTaskTimeout,
		consumeTopics,
		handler,
	)
}

func addTaskToTopic(task ProduceTask) error {
	reqStr, err := ffjson.Marshal(task)
	if err != nil {
		logger.Log.Errorf("task serialization error: %v", err)
		return err
	}

	topic := cfg.Config().Kafka.ProducerTopic
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(reqStr),
	}
	logger.Log.Infof("put task into topic %s, task: %s", topic, reqStr)

	kafka.Producer.Input() <- msg
	return nil
}