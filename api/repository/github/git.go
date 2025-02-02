package github

import (
	"context"
	"strings"

	"github.com/tinyci/ci-agents/ci-gen/grpc/services/repository"
	"github.com/tinyci/ci-agents/errors"
	"github.com/tinyci/ci-agents/utils"
	"google.golang.org/grpc/codes"
)

// GetFileList finds all the files in the tree for the given repository
func (rs *RepositoryServer) GetFileList(ctx context.Context, rsp *repository.RepoSHAPair) (*repository.StringList, error) {
	owner, repo, err := utils.OwnerRepo(rsp.RepoName)
	if err != nil {
		return nil, err.ToGRPC(codes.FailedPrecondition)
	}

	gh, err := rs.getClientForRepo(ctx, rsp.RepoName)
	if err != nil {
		return nil, err.ToGRPC(codes.FailedPrecondition)
	}

	tree, _, eErr := gh.Git.GetTree(ctx, owner, repo, rsp.Sha, true)
	if eErr != nil {
		return nil, errors.New(err).Wrapf("obtaining tree for repo %v/%v", owner, repo).ToGRPC(codes.FailedPrecondition)
	}

	files := []string{}

	for _, entry := range tree.Entries {
		files = append(files, entry.GetPath())
	}

	return &repository.StringList{List: files}, nil
}

// GetSHA retrieves the SHA for the branch in the given repository
func (rs *RepositoryServer) GetSHA(ctx context.Context, rrp *repository.RepoRefPair) (*repository.String, error) {
	owner, repo, err := utils.OwnerRepo(rrp.RepoName)
	if err != nil {
		return nil, err.ToGRPC(codes.FailedPrecondition)
	}

	gh, err := rs.getClientForRepo(ctx, rrp.RepoName)
	if err != nil {
		return nil, err.ToGRPC(codes.FailedPrecondition)
	}

	ref, _, eErr := gh.Git.GetRef(ctx, owner, repo, rrp.RefName)
	if eErr != nil {
		return nil, errors.New(eErr).Wrapf("Obtaining ref for %v/%v", owner, repo).ToGRPC(codes.FailedPrecondition)
	}

	return &repository.String{Name: ref.GetObject().GetSHA()}, nil
}

// GetRefs gets the refs that match the given SHA. Only heads and tags are considered.
func (rs *RepositoryServer) GetRefs(ctx context.Context, rsp *repository.RepoSHAPair) (*repository.StringList, error) {
	owner, repo, err := utils.OwnerRepo(rsp.RepoName)
	if err != nil {
		return nil, err.ToGRPC(codes.FailedPrecondition)
	}

	gh, err := rs.getClientForRepo(ctx, rsp.RepoName)
	if err != nil {
		return nil, err.ToGRPC(codes.FailedPrecondition)
	}

	// FIXME pagination (sigh)
	refs, _, eErr := gh.Git.ListRefs(ctx, owner, repo, nil)
	if eErr != nil {
		return nil, errors.New(eErr).Wrapf("listing refs for %v/%v", owner, repo).ToGRPC(codes.FailedPrecondition)
	}

	list := []string{}

	for _, ref := range refs {
		if ref.GetObject().GetSHA() == rsp.Sha {
			list = append(list, strings.TrimPrefix(ref.GetRef(), "refs/"))
		}
	}

	return &repository.StringList{List: list}, nil
}

// GetFile returns the file in full as a byte array.
func (rs *RepositoryServer) GetFile(ctx context.Context, fr *repository.FileRequest) (*repository.Bytes, error) {
	owner, repo, err := utils.OwnerRepo(fr.RepoName)
	if err != nil {
		return nil, err.ToGRPC(codes.FailedPrecondition)
	}

	gh, err := rs.getClientForRepo(ctx, fr.RepoName)
	if err != nil {
		return nil, err.ToGRPC(codes.FailedPrecondition)
	}

	tree, _, eErr := gh.Git.GetTree(ctx, owner, repo, fr.Sha, true)
	if eErr != nil {
		return nil, errors.New(eErr).Wrapf("getting tree for repo %v/%v (sha: %v)", owner, repo, fr.Sha).ToGRPC(codes.FailedPrecondition)
	}

	for _, entry := range tree.Entries {
		if entry.GetPath() == fr.Filename {
			content, _, err := gh.Git.GetBlobRaw(context.Background(), owner, repo, entry.GetSHA())
			if err != nil {
				return nil, errors.New(err).ToGRPC(codes.FailedPrecondition)
			}

			return &repository.Bytes{Value: content}, nil
		}
	}

	return nil, errors.New("file not found").ToGRPC(codes.FailedPrecondition)
}

func (rs *RepositoryServer) getFileList(ctx context.Context, repoName, sha string) ([]string, *errors.Error) {
	owner, repo, err := utils.OwnerRepo(repoName)
	if err != nil {
		return nil, err
	}

	gh, err := rs.getClientForRepo(ctx, repoName)
	if err != nil {
		return nil, err
	}

	tree, _, eErr := gh.Git.GetTree(ctx, owner, repo, sha, true)
	if eErr != nil {
		return nil, errors.New(eErr)
	}

	files := []string{}

	for _, entry := range tree.Entries {
		files = append(files, entry.GetPath())
	}

	return files, nil
}

// GetDiffFiles retrieves the files present in the diff between the base and the head.
func (rs *RepositoryServer) GetDiffFiles(ctx context.Context, fdr *repository.FileDiffRequest) (*repository.StringList, error) {
	owner, repo, err := utils.OwnerRepo(fdr.RepoName)
	if err != nil {
		return nil, err.ToGRPC(codes.FailedPrecondition)
	}

	if fdr.Base == strings.Repeat("0", 40) {
		list, err := rs.getFileList(ctx, fdr.RepoName, fdr.Head)
		if err != nil {
			return nil, err.ToGRPC(codes.FailedPrecondition)
		}

		return &repository.StringList{List: list}, nil
	}

	if fdr.Head == strings.Repeat("0", 40) {
		return nil, errors.New("branch deleted").ToGRPC(codes.FailedPrecondition)
	}

	gh, err := rs.getClientForRepo(ctx, fdr.RepoName)
	if err != nil {
		return nil, err.ToGRPC(codes.FailedPrecondition)
	}

	commits, _, eErr := gh.Repositories.CompareCommits(ctx, owner, repo, fdr.Base, fdr.Head)
	if eErr != nil {
		return nil, errors.New(eErr).ToGRPC(codes.FailedPrecondition)
	}

	files := []string{}

	for _, file := range commits.Files {
		files = append(files, file.GetFilename())
	}

	return &repository.StringList{List: files}, nil
}
