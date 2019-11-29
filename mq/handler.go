package mq

type HandlerIFace interface {
	HandleMessage(msg []byte)
}
