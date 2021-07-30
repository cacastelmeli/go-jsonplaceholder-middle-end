package domain

type PostRepository interface {
	FetchAllPosts() ([]*Post, error)
}
