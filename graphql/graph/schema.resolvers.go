package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"graphql/graph/generated"
	"graphql/graph/model"
)

var authors = []*model.Author{
	{ID: 1, Name: "J. K. Rowling"},
	{ID: 2, Name: "J. R. R. Tolkien"},
	{ID: 3, Name: "Brent Weeks"},
}

var books = []*model.Book{
	{ID: 1, Name: "Harry Potter and the Chamber of Secrets", AuthorID: 1},
	{ID: 2, Name: "Harry Potter and the Prisoner of Azkaban", AuthorID: 1},
	{ID: 3, Name: "Harry Potter and the Goblet of Fire", AuthorID: 1},
	{ID: 4, Name: "The Fellowship of the Ring", AuthorID: 2},
	{ID: 5, Name: "The Two Towers", AuthorID: 2},
	{ID: 6, Name: "The Return of the King", AuthorID: 2},
	{ID: 7, Name: "The Way of Shadows", AuthorID: 3},
	{ID: 8, Name: "Beyond the Shadows", AuthorID: 3},
}

func (r *authorResolver) Books(ctx context.Context, author *model.Author) ([]*model.Book, error) {
	authorBooks := make([]*model.Book, 0, 0)
	for _, book := range books {
		if book.AuthorID == author.ID {
			authorBooks = append(authorBooks, book)
		}
	}

	return authorBooks, nil
}

func (r *bookResolver) Author(ctx context.Context, book *model.Book) (*model.Author, error) {
	for _, author := range authors {
		if author.ID == book.AuthorID {
			return author, nil
		}
	}

	return nil, nil
}

func (r *queryResolver) Books(ctx context.Context) ([]*model.Book, error) {
	return books, nil
}

func (r *queryResolver) Book(ctx context.Context, id *int) (*model.Book, error) {
	for _, book := range books {
		if book.ID == *id {
			return book, nil
		}
	}

	return nil, nil
}

func (r *queryResolver) Authors(ctx context.Context) ([]*model.Author, error) {
	return authors, nil
}

func (r *queryResolver) Author(ctx context.Context, id *int) (*model.Author, error) {
	for _, author := range authors {
		if author.ID == *id {
			return author, nil
		}
	}

	return nil, nil
}

// Author returns generated.AuthorResolver implementation.
func (r *Resolver) Author() generated.AuthorResolver { return &authorResolver{r} }

// Book returns generated.BookResolver implementation.
func (r *Resolver) Book() generated.BookResolver { return &bookResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type (
	authorResolver struct{ *Resolver }
	bookResolver   struct{ *Resolver }
	queryResolver  struct{ *Resolver }
)
