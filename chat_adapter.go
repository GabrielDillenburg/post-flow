package main

// ChatAdapter defines the standard interface for any chat service.
type ChatAdapter interface {
	GenerateResponse(input string) (string, error)
}
