package models

type NumberGenerator interface {
	NumberGenerator(min int, max int) int
}