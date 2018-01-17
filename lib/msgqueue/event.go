package msgqueue

// Interface definition for events that are emitted using an EventEmitter
type Event interface {
	EventName() string
	PartitionKey() string
}
