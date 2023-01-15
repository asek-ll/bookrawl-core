package dao

import (
	"github.com/asek-ll/bookrawl-core/pkg/dao/abooks"
	"github.com/asek-ll/bookrawl-core/pkg/dao/authors"
	"github.com/asek-ll/bookrawl-core/pkg/dao/books"
	"github.com/asek-ll/bookrawl-core/pkg/dao/userbookstates"
	"github.com/asek-ll/bookrawl-core/pkg/dao/users"

	"go.mongodb.org/mongo-driver/mongo"
)

type DaoHolder struct {
	client             *mongo.Client
	abookStore         *abooks.AbookStore
	authorStore        *authors.Store
	userStore          *users.Store
	bookStore          *books.Store
	userBookStateState *userbookstates.Store
}

func NewDaoHolder(client *mongo.Client) *DaoHolder {
	return &DaoHolder{
		client: client,
		abookStore: &abooks.AbookStore{
			Collection: client.Database("bookrawl").Collection("abooks"),
		},
		authorStore: &authors.Store{
			Collection: client.Database("bookrawl").Collection("authors"),
		},
		userStore: &users.Store{
			Collection: client.Database("bookrawl").Collection("users"),
		},
		bookStore: &books.Store{
			Collection: client.Database("bookrawl").Collection("books"),
		},
		userBookStateState: &userbookstates.Store{
			Collection: client.Database("bookrawl").Collection("userBookStates"),
		},
	}
}

func (dh *DaoHolder) GetABookStore() *abooks.AbookStore {
	return dh.abookStore
}

func (dh *DaoHolder) GetAuthorsStore() *authors.Store {
	return dh.authorStore
}

func (dh *DaoHolder) GetUsersStore() *users.Store {
	return dh.userStore
}

func (dh *DaoHolder) GetBookStore() *books.Store {
	return dh.bookStore
}

func (dh *DaoHolder) GetUserBookStateStore() *userbookstates.Store {
	return dh.userBookStateState
}
