package parser

import (
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

// Repository is a representation of the existing Git repo.
type Repository struct {
	repo *git.Repository
}

// NewRepository returns repository of given path.
// Note the path must point to where `.git` folder exists.
func NewRepository(path string) (*Repository, error) {

	dir, err := filepath.Abs(path)

	if err != nil {
		return nil, err
	}

	got, err := git.PlainOpen(dir)

	if err != nil {
		return nil, err
	}

	return &Repository{
		repo: got,
	}, nil
}

// TagNames returns git tags in representation.
func (c Repository) TagNames() ([]SemanticVersion, error) {

	iter, err := c.repo.Tags()

	if err != nil {
		return nil, err
	}

	var vers []SemanticVersion

	if err := iter.ForEach(func(ref *plumbing.Reference) error {
		got, err := c.repo.TagObject(ref.Hash())

		if err != nil {
			return err
		}

		v, err := NewSemanticVersion(got.Name)

		if err != nil {
			return err
		}

		vers = append(vers, *v)

		return nil

	}); err != nil {
		return nil, err
	}

	return SortVersions(vers), nil
}
