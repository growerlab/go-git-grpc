package main

import (
	"context"
	"log"
	"os"
	"path"
	"time"

	"github.com/go-git/go-git/v5/plumbing/object"

	gggrpc "github.com/growerlab/go-git-grpc"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func main() {
	go func() {
		gitRoot := path.Join(os.Getenv("GOPATH"), "src/github.com/growerlab/go-git-grpc")
		err := gggrpc.NewServer(gitRoot, "localhost:8081")
		if err != nil {
			panic(err)
		}
	}()
	time.Sleep(1 * time.Second)

	clientCtx := context.Background()
	repoPath := "test/testrepo_bare"
	store, closeFn, err := gggrpc.NewClient(clientCtx, "localhost:8081", repoPath)
	if err != nil {
		panic(err)
	}
	defer closeFn.Close()

	repo, err := git.Open(store, nil)
	if err != nil {
		panic(err)
	}

	// testReferences(repo)
	// testTags(repo)
	testCommits(repo)
	// tag下的文件列表
	// testFileTreesInTag(repo, "v1.0")

	time.Sleep(500 * time.Millisecond)
}

func testFileTreesInTag(repo *git.Repository, tagName string) {
	iter, err := repo.TreeObjects()
	if err != nil {
		log.Fatalf("get trees was err:%+v", err)
	}
	iter.ForEach(func(t *object.Tree) error {
		return nil
	})
	// tag, err := repo.Tag(tagName)
	// if err != nil {
	// 	log.Fatalf("get tag '%s' was err: %+v", tagName, err)
	// }
	// tree, err := repo.TreeObject(tag.Hash())
	// if err != nil {
	// 	log.Fatalf("get tree '%s' was err: %+v", tagName, err)
	// }
	// n := 0
	// tree.Files().ForEach(func(file *object.File) error {
	// 	n++
	// 	log.Printf("file %s in tag '%s'", file.Name, tagName)
	// 	return nil
	// })
}

func testCommits(repo *git.Repository) {
	commitIter, err := repo.CommitObjects()
	if err != nil {
		log.Fatalf("get commits was err: %+v", err)
	}
	n := 0
	_ = commitIter.ForEach(func(c *object.Commit) error {
		n++
		log.Printf("commit committer: %s message: %s hash: %s\n",
			c.Committer.String(),
			c.Message,
			c.Hash.String(),
		)
		fileIter, err := c.Files()
		if err != nil {
			log.Fatalf("get files in %s was err: %+v", c.Hash.String(), err)
		}
		fileIter.ForEach(func(file *object.File) error {
			cts, err := file.Contents()
			if err != nil {
				log.Fatalf("get file %s in %s was err: %+v", file.Name, c.Hash.String(), err)
			}
			log.Printf("file '%s' content: %+v\n", file.Name, cts)
			return nil
		})
		return nil
	})
}

func testTags(repo *git.Repository) {
	refs, err := repo.Tags()
	if err != nil {
		log.Fatalf("get tags was err: %+v", err)
	}
	n := 0
	_ = refs.ForEach(func(r *plumbing.Reference) error {
		n++
		log.Printf("tag name: %s type: %d hash: %s string: %s target: %s\n",
			r.Name(),
			r.Type(),
			r.Hash().String(),
			r.String(),
			r.Target())
		return nil
	})
	if n == 0 {
		log.Fatalf("Not found tags")
	}
}

func testReferences(repo *git.Repository) {
	refs, err := repo.References()
	if err != nil {
		log.Fatalf("get branchs was err: %+v", err)
	}
	n := 0
	_ = refs.ForEach(func(r *plumbing.Reference) error {
		n++
		log.Printf("branch name: %s type: %d hash: %s string: %s target: %s\n",
			r.Name(),
			r.Type(),
			r.Hash().String(),
			r.String(),
			r.Target())
		return nil
	})
	if n == 0 {
		log.Fatalf("Not found branchs")
	}
}
