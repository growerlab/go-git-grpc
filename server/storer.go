package server

import (
	"context"
	"io"
	"time"

	"github.com/go-git/go-git/v5/plumbing/filemode"

	"github.com/go-git/go-git/v5/plumbing/format/index"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/growerlab/go-git-grpc/pb"
	"github.com/pkg/errors"
)

var _ pb.StorerServer = &Storer{}

type Storer struct {
	*pb.UnimplementedStorerServer

	root string // 仓库根目录
}

func (s *Storer) NewEncodedObject(ctx context.Context, none *pb.None) (*pb.EncodedObject, error) {
	var result *pb.EncodedObject
	err := repo(s.root, none.RepoPath, func(r *git.Repository) error {
		obj := r.Storer.NewEncodedObject()
		result = new(pb.EncodedObject)
		result.Type = obj.Type().String()
		result.Size = obj.Size()
		result.Hash = obj.Hash().String()
		return nil
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return result, nil
}

func (s *Storer) SetEncodedObjectType(ctx context.Context, i *pb.Int8) (*pb.None, error) {
	var result = &pb.None{RepoPath: i.RepoPath}
	var encodedObjectType plumbing.ObjectType

	if len(i.Value) > 0 {
		encodedObjectType = plumbing.ObjectType(i.Value[0])
	}

	err := repo(s.root, i.RepoPath, func(r *git.Repository) error {
		r.Storer.NewEncodedObject().SetType(encodedObjectType)
		return nil
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return result, nil
}

func (s *Storer) SetEncodedObjectSetSize(ctx context.Context, i *pb.Int64) (*pb.None, error) {
	var result = &pb.None{RepoPath: i.RepoPath}

	err := repo(s.root, i.RepoPath, func(r *git.Repository) error {
		r.Storer.NewEncodedObject().SetSize(i.Value)
		return nil
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return result, nil
}

func (s *Storer) EncodedObjectReader(none *pb.None, server pb.Storer_EncodedObjectReaderServer) error {
	var buf = make([]byte, 512)
	err := repo(s.root, none.RepoPath, func(r *git.Repository) error {
		reader, err := r.Storer.NewEncodedObject().Reader()
		if err != nil {
			return errors.WithStack(err)
		}
		defer reader.Close()

		for {
			buf := buf[:0]
			n, err := io.ReadFull(reader, buf)
			if n > 0 {
				err = server.Send(&pb.Bytes{
					Value: buf,
				})
			}
			if err != nil {
				return errors.WithStack(err)
			}
		}
	})
	return err
}

func (s *Storer) EncodedObjectWriter(server pb.Storer_EncodedObjectWriterServer) error {
	firstRecvBytes, err := server.Recv()
	if err != nil {
		return errors.WithStack(err)
	}

	err = repo(s.root, firstRecvBytes.RepoPath, func(r *git.Repository) error {
		writer, err := r.Storer.NewEncodedObject().Writer()
		if err != nil {
			return errors.WithStack(err)
		}
		defer writer.Close()

		_, err = writer.Write(firstRecvBytes.Value)
		if err != nil {
			return errors.WithStack(err)
		}

		for {
			recvMsg, err := server.Recv()
			if err != nil {
				return errors.WithStack(err)
			}
			_, err = writer.Write(recvMsg.Value)
			if err != nil {
				return errors.WithStack(err)
			}
		}
	})
	return err
}

func (s *Storer) SetReference(ctx context.Context, reference *pb.Reference) (*pb.None, error) {
	var result = &pb.None{}
	err := repo(s.root, reference.RepoPath, func(r *git.Repository) error {
		ref := plumbing.NewReferenceFromStrings(reference.N, reference.Target)
		err := r.Storer.SetReference(ref)
		return errors.WithStack(err)
	})
	return result, err
}

func (s *Storer) CheckAndSetReference(ctx context.Context, params *pb.SetReferenceParams) (*pb.None, error) {
	var result = &pb.None{}
	err := repo(s.root, params.RepoPath, func(r *git.Repository) error {
		newRef := plumbing.NewReferenceFromStrings(params.New.N, params.New.Target)
		oldRef := plumbing.NewReferenceFromStrings(params.Old.N, params.Old.Target)
		err := r.Storer.CheckAndSetReference(newRef, oldRef)
		return errors.WithStack(err)
	})
	return result, err
}

func (s *Storer) GetReference(ctx context.Context, name *pb.ReferenceName) (*pb.Reference, error) {
	var result *pb.Reference
	err := repo(s.root, name.RepoPath, func(r *git.Repository) error {
		ref, err := r.Storer.Reference(plumbing.ReferenceName(name.Name))
		if err != nil {
			return errors.WithStack(err)
		}

		result = buildRefToPbRef(ref)
		return nil
	})
	return result, err
}

func (s *Storer) GetReferences(ctx context.Context, none *pb.None) (*pb.References, error) {
	var result = new(pb.References)
	err := repo(s.root, none.RepoPath, func(r *git.Repository) error {
		iter, err := r.Storer.IterReferences()
		if err != nil {
			return errors.WithStack(err)
		}

		err = iter.ForEach(func(ref *plumbing.Reference) error {
			pbRef := buildRefToPbRef(ref)
			result.Refs = append(result.Refs, pbRef)
			return nil
		})
		return errors.WithStack(err)
	})
	return result, err
}

func (s *Storer) RemoveReference(ctx context.Context, name *pb.ReferenceName) (*pb.None, error) {
	err := repo(s.root, name.RepoPath, func(r *git.Repository) error {
		rn := plumbing.ReferenceName(name.Name)
		err := r.Storer.RemoveReference(rn)
		return errors.WithStack(err)
	})
	return &pb.None{}, err
}

func (s *Storer) CountLooseRefs(ctx context.Context, none *pb.None) (*pb.Int64, error) {
	var result = new(pb.Int64)
	err := repo(s.root, none.RepoPath, func(r *git.Repository) error {
		count, err := r.Storer.CountLooseRefs()
		if err != nil {
			return errors.WithStack(err)
		}
		result.Value = int64(count)
		return nil
	})
	return result, err
}

func (s *Storer) PackRefs(ctx context.Context, none *pb.None) (*pb.None, error) {
	var result = new(pb.None)
	err := repo(s.root, none.RepoPath, func(r *git.Repository) error {
		return errors.WithStack(r.Storer.PackRefs())
	})
	return result, err
}

func (s *Storer) SetShallow(ctx context.Context, hashs *pb.Hashs) (*pb.None, error) {
	var result = new(pb.None)
	err := repo(s.root, hashs.RepoPath, func(r *git.Repository) error {
		commits := make([]plumbing.Hash, len(hashs.Hash))
		for i := range hashs.Hash {
			commits[i] = plumbing.NewHash(hashs.Hash[i])
		}
		err := r.Storer.SetShallow(commits)
		return errors.WithStack(err)
	})
	return result, err
}

func (s *Storer) Shallow(ctx context.Context, none *pb.None) (*pb.Hashs, error) {
	var result = new(pb.Hashs)
	err := repo(s.root, none.RepoPath, func(r *git.Repository) error {
		commits, err := r.Storer.Shallow()
		if err != nil {
			return errors.WithStack(err)
		}
		for _, cmt := range commits {
			result.Hash = append(result.Hash, cmt.String())
		}
		return nil
	})
	return result, err
}

func (s *Storer) SetIndex(ctx context.Context, idx *pb.Index) (*pb.None, error) {
	var result = new(pb.None)
	err := repo(s.root, idx.RepoPath, func(r *git.Repository) error {
		var entries []*index.Entry
		var cache *index.Tree
		var trees []index.TreeEntry
		var resolveUndo *index.ResolveUndo
		var resolveUndoEntries []index.ResolveUndoEntry
		var endOfIndexEntry *index.EndOfIndexEntry

		if len(idx.Entries) > 0 {
			entries = make([]*index.Entry, 0, len(idx.Entries))
			for _, ent := range idx.Entries {
				entries = append(entries, &index.Entry{
					Hash:         plumbing.NewHash(ent.Hash),
					Name:         ent.Name,
					CreatedAt:    time.Unix(ent.CreatedAt, 0),
					ModifiedAt:   time.Unix(ent.ModifiedAt, 0),
					Dev:          ent.Dev,
					Inode:        ent.Inode,
					Mode:         filemode.FileMode(ent.Mode),
					UID:          ent.UID,
					GID:          ent.GID,
					Size:         ent.Size,
					Stage:        index.Stage(ent.Stage),
					SkipWorktree: ent.SkipWorktree,
					IntentToAdd:  ent.IntentToAdd,
				})
			}
		}

		if idx.Cache != nil && len(idx.Cache.Entries) > 0 {
			trees = make([]index.TreeEntry, 0, len(idx.Cache.Entries))
			for _, ent := range idx.Cache.Entries {
				trees = append(trees, index.TreeEntry{
					Path:    ent.Path,
					Entries: int(ent.Entries),
					Trees:   int(ent.Trees),
					Hash:    plumbing.NewHash(ent.Hash),
				})
			}
		}
		if trees != nil {
			cache = &index.Tree{
				Entries: trees,
			}
		}

		if idx.ResolveUndo != nil {
			for _, ent := range idx.ResolveUndo.Entries {
				stageSet := make(map[index.Stage]plumbing.Hash)
				for _, stg := range ent.Stages {
					stageSet[index.Stage(stg.Key)] = plumbing.NewHash(stg.Value)
				}
				resolveUndoEntries = append(resolveUndoEntries, index.ResolveUndoEntry{
					Path:   ent.Path,
					Stages: stageSet,
				})
			}
		}
		if len(resolveUndoEntries) > 0 {
			resolveUndo = &index.ResolveUndo{Entries: resolveUndoEntries}
		}
		if idx.EndOfIndexEntry != nil {
			endOfIndexEntry = &index.EndOfIndexEntry{
				Offset: idx.EndOfIndexEntry.Offset,
				Hash:   plumbing.NewHash(idx.EndOfIndexEntry.Hash),
			}
		}

		newIdx := &index.Index{
			Version:         idx.Version,
			Entries:         entries,
			Cache:           cache,
			ResolveUndo:     resolveUndo,
			EndOfIndexEntry: endOfIndexEntry,
		}
		err := r.Storer.SetIndex(newIdx)
		return errors.WithStack(err)
	})
	return result, err
}

func (s *Storer) GetIndex(ctx context.Context, none *pb.None) (*pb.Index, error) {
	panic("implement me")
}

func (s *Storer) GetConfig(ctx context.Context, none *pb.None) (*pb.Config, error) {
	panic("implement me")
}

func (s *Storer) SetConfig(ctx context.Context, c *pb.Config) (*pb.None, error) {
	panic("implement me")
}

func (s *Storer) Modules(ctx context.Context, none *pb.None) (*pb.ModuleNames, error) {
	panic("implement me")
}

func (s *Storer) mustEmbedUnimplementedStorerServer() {

}
