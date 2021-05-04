package git

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"strings"
)

var (
	ErrRepositoryNotFound = errors.New("Repository not found.")
)

// DiscoverRepositoryPath finds the path of the repository from a given
// directory. It returns the path to the repository, or an an empty string and
// a non-nil error if it can't find the repository.
func DiscoverRepositoryPath(dir string) (string, error) {
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return "", ErrRepositoryNotFound
	}
	dir = filepath.Join(dir, ".git")
	for dir != "/.git" {
		fi, err := os.Stat(dir)
		if err == nil && fi.IsDir() {
			return dir, nil
		}
		dir = filepath.Join(dir, "..", "..", ".git")
	}
	return "", ErrRepositoryNotFound
}

// Repository represents a git repository.
type Repository struct {
	path string
}

// OpenRepository opens a repository by its filepath. You can use
// DiscoverRepositoryPath to discover the repository from any directory, and
// use the result of this call as parameter for OpenRepository.
//
// OpenRepository will return an error if the given path does not appear to be
// a git repository.
func OpenRepository(p string) (*Repository, error) {
	if !strings.HasSuffix(p, ".git") && !strings.HasSuffix(p, ".git/") {
		p = filepath.Join(p, ".git")
	}
	p = strings.TrimRight(p, "/")
	fi, err := os.Stat(filepath.Join(p, "config"))
	if err == nil && !fi.IsDir() {
		return &Repository{path: p}, nil
	}
	return nil, ErrRepositoryNotFound
}

// RemoteURL returns the URL of a remote by its name. Or an error, if the
// remote is not declared.
func (r *Repository) RemoteURL(name string) (string, error) {
	config, err := os.Open(filepath.Join(r.path, "config"))
	if err != nil {
		return "", err
	}
	defer config.Close()
	line := fmt.Sprintf("[remote %q]", name)
	scanner := bufio.NewScanner(config)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		if scanner.Text() == line {
			scanner.Scan()
			return strings.Split(scanner.Text(), " = ")[1], nil
		}
	}
	return "", errRemoteNotFound{name}
}

type errRemoteNotFound struct {
	name string
}

func (e errRemoteNotFound) Error() string {
	return fmt.Sprintf("Remote %q not found.", e.name)
}
