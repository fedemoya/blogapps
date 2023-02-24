package main

type EventsDeserializer[Visitor any] interface {
    Deserialize(message Message) (Event[Visitor], error)
}
