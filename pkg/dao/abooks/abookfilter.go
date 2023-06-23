package abooks

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FindBooksFilter struct {
	AfterDate  *time.Time
	BeforeDate *time.Time
	Author     *primitive.ObjectID
	NoAuthor   *bool
}

type BooksFilterBuilder struct {
	filter *FindBooksFilter
}

func NewBooksFilterBuilder() *BooksFilterBuilder {
	return &BooksFilterBuilder{
		filter: &FindBooksFilter{},
	}
}

func (fb *BooksFilterBuilder) SetAfterDate(t *time.Time) *BooksFilterBuilder {
	fb.filter.AfterDate = t
	return fb
}

func (fb *BooksFilterBuilder) SetBeforeDate(t *time.Time) *BooksFilterBuilder {
	fb.filter.BeforeDate = t
	return fb
}

func (fb *BooksFilterBuilder) SetAuthor(author *primitive.ObjectID) *BooksFilterBuilder {
	fb.filter.Author = author
	return fb
}

func (fb *BooksFilterBuilder) NoAuthor() *BooksFilterBuilder {
	fb.filter.Author = nil
	t := true
	fb.filter.NoAuthor = &t
	return fb
}

func (fb *BooksFilterBuilder) Build() *FindBooksFilter {
	return fb.filter
}
