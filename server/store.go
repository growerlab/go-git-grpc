package server

import (
	"context"
	"io"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/growerlab/go-git-grpc/pb"
	"github.com/pkg/errors"
)

var _ pb.StorerServer = &Store{}

type Store struct {
	*pb.UnimplementedStorerServer

	root string // 仓库根目录
}

func (s *Store) SetEncodedObject(ctx context.Context, o *pb.EncodedObject) (*pb.Hash, error) {
	var result *pb.Hash
	err := repo(s.root, o.RepoPath, func(r *git.Repository) error {
		t, err := plumbing.ParseObjectType(o.Type)
		if err != nil {
			return errors.WithStack(err)
		}

		obj := &MemoryObject{
			ctx:    ctx,
			server: s,
			t:      t,
			h:      plumbing.NewHash(o.Hash),
			cont:   []byte{},
			sz:     o.Size,
		}

		h, err := r.Storer.SetEncodedObject(obj)
		if err != nil {
			return errors.WithStack(err)
		}
		result = &pb.Hash{
			Value: h.String(),
		}
		return nil
	})
	return result, errors.WithStack(err)
}

func (s *Store) SetEncodedObjectType(ctx context.Context, i *pb.Int8) (*pb.None, error) {
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

func (s *Store) SetEncodedObjectSetSize(ctx context.Context, i *pb.Int64) (*pb.None, error) {
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

func (s *Store) EncodedObjectReader(none *pb.None, server pb.Storer_EncodedObjectReaderServer) error {
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

func (s *Store) EncodedObjectWriter(server pb.Storer_EncodedObjectWriterServer) error {
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

func (s *Store) SetReference(ctx context.Context, reference *pb.Reference) (*pb.None, error) {
	var result = &pb.None{}
	err := repo(s.root, reference.RepoPath, func(r *git.Repository) error {
		ref := plumbing.NewReferenceFromStrings(reference.N, reference.Target)
		err := r.Storer.SetReference(ref)
		return errors.WithStack(err)
	})
	return result, err
}

func (s *Store) CheckAndSetReference(ctx context.Context, params *pb.SetReferenceParams) (*pb.None, error) {
	var result = &pb.None{}
	err := repo(s.root, params.RepoPath, func(r *git.Repository) error {
		newRef := plumbing.NewReferenceFromStrings(params.New.N, params.New.Target)
		oldRef := plumbing.NewReferenceFromStrings(params.Old.N, params.Old.Target)
		err := r.Storer.CheckAndSetReference(newRef, oldRef)
		return errors.WithStack(err)
	})
	return result, err
}

func (s *Store) GetReference(ctx context.Context, name *pb.ReferenceName) (*pb.Reference, error) {
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

func (s *Store) GetReferences(ctx context.Context, none *pb.None) (*pb.References, error) {
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

func (s *Store) RemoveReference(ctx context.Context, name *pb.ReferenceName) (*pb.None, error) {
	err := repo(s.root, name.RepoPath, func(r *git.Repository) error {
		rn := plumbing.ReferenceName(name.Name)
		err := r.Storer.RemoveReference(rn)
		return errors.WithStack(err)
	})
	return &pb.None{}, err
}

func (s *Store) CountLooseRefs(ctx context.Context, none *pb.None) (*pb.Int64, error) {
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

func (s *Store) PackRefs(ctx context.Context, none *pb.None) (*pb.None, error) {
	var result = new(pb.None)
	err := repo(s.root, none.RepoPath, func(r *git.Repository) error {
		return errors.WithStack(r.Storer.PackRefs())
	})
	return result, err
}

func (s *Store) SetShallow(ctx context.Context, hashs *pb.Hashs) (*pb.None, error) {
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

func (s *Store) Shallow(ctx context.Context, none *pb.None) (*pb.Hashs, error) {
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

func (s *Store) SetIndex(ctx context.Context, idx *pb.Index) (*pb.None, error) {
	var result = new(pb.None)
	err := repo(s.root, idx.RepoPath, func(r *git.Repository) error {
		newIdx := buildPbRefToIndex(idx)
		err := r.Storer.SetIndex(newIdx)
		return errors.WithStack(err)
	})
	return result, err
}

func (s *Store) GetIndex(ctx context.Context, none *pb.None) (*pb.Index, error) {
	var result *pb.Index
	err := repo(s.root, none.RepoPath, func(r *git.Repository) error {
		idx, err := r.Storer.Index()
		if err != nil {
			return errors.WithStack(err)
		}
		result = buildIndexToPbRef(idx)
		return nil
	})
	return result, err
}

func (s *Store) GetConfig(ctx context.Context, none *pb.None) (*pb.Config, error) {
	var result = new(pb.Config)
	err := repo(s.root, none.RepoPath, func(r *git.Repository) error {
		cfg, err := r.Storer.Config()
		if err != nil {
			return errors.WithStack(err)
		}
		result = buildConfigFromPbConfig(cfg)
		return nil
	})
	return result, err
}

func (s *Store) SetConfig(ctx context.Context, c *pb.Config) (*pb.None, error) {
	var result = new(pb.None)
	err := repo(s.root, c.RepoPath, func(r *git.Repository) error {
		cfg := buildPbConfigFromConfig(c)
		err := r.Storer.SetConfig(cfg)
		return errors.WithStack(err)
	})
	return result, err
}

func (s *Store) Modules(ctx context.Context, none *pb.None) (*pb.ModuleNames, error) {
	var result = new(pb.ModuleNames)
	err := repo(s.root, none.RepoPath, func(r *git.Repository) error {
		cfg, err := r.Storer.Config()
		if err != nil {
			return errors.WithStack(err)
		}
		for _, submd := range cfg.Submodules {
			result.Names = append(result.Names, submd.Name)
		}
		return nil
	})
	return result, err
}

func (s *Store) mustEmbedUnimplementedStorerServer() {

}
