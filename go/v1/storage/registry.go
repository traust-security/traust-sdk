package storage

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/google/uuid"
)

// ProductRepo is a repo as a product ships it. Ref "" is the default branch.
type ProductRepo struct {
	ProductID    string
	RepoID       string
	Ref          string
	SubService   *string
	ResourceType *string
}

// ProductRepoVersion records that a product version ships a product_repo.
type ProductRepoVersion struct {
	ProductRepoID    string
	Version          string
	Category         *string
	ClusterOperators []string
	Images           []string
}

// RepoOwner records a team that owns a product_repo.
type RepoOwner struct {
	ProductRepoID string
	Team          string
	Manager       *string
	Individuals   []string
	Source        *string
	JiraProject   *string
	JiraComponent *string
}

// RegisterProduct registers a product by slug, idempotently, and returns its
// product_id. Segment is descriptive: re-registering sets it.
func (c *Client) RegisterProduct(ctx context.Context, slug string, segment *string) (string, error) {
	if err := registryText(slug, true, segment); err != nil {
		return "", err
	}
	id := newRegistryID()
	return c.register(ctx, func(conn *sql.Conn, s *sqlStore, at string) (string, error) {
		if err := s.queries.productUpsert(ctx, conn, productUpsertParams{
			productId: id, slug: slug, segment: segment, registeredAt: at,
		}); err != nil {
			return "", err
		}
		return storedID(s.queries.productGet(ctx, conn, productGetParams{slug: slug}))
	})
}

// RegisterRepo registers a repo by exact URL, idempotently, and returns its repo_id.
func (c *Client) RegisterRepo(ctx context.Context, repoURL string) (string, error) {
	if err := registryText(repoURL, true); err != nil {
		return "", err
	}
	id := newRegistryID()
	return c.register(ctx, func(conn *sql.Conn, s *sqlStore, at string) (string, error) {
		if err := s.queries.repoUpsert(ctx, conn, repoUpsertParams{
			repoId: id, repoUrl: repoURL, registeredAt: at,
		}); err != nil {
			return "", err
		}
		return storedID(s.queries.repoGet(ctx, conn, repoGetParams{repoUrl: repoURL}))
	})
}

// RegisterProductRepo registers a product's repo at a ref, idempotently, and
// returns its product_repo_id. Product and repo must already be registered.
func (c *Client) RegisterProductRepo(ctx context.Context, pr ProductRepo) (string, error) {
	if err := registryText(pr.ProductID, true, pr.SubService, pr.ResourceType); err != nil {
		return "", err
	}
	if err := registryText(pr.RepoID, true); err != nil {
		return "", err
	}
	if err := registryText(pr.Ref, false); err != nil {
		return "", err
	}
	id := newRegistryID()
	return c.register(ctx, func(conn *sql.Conn, s *sqlStore, at string) (string, error) {
		if err := s.queries.productRepoUpsert(ctx, conn, productRepoUpsertParams{
			productRepoId: id,
			productId:     pr.ProductID,
			repoId:        pr.RepoID,
			ref:           pr.Ref,
			subService:    pr.SubService,
			resourceType:  pr.ResourceType,
			registeredAt:  at,
		}); err != nil {
			return "", err
		}
		return storedID(s.queries.productRepoGet(ctx, conn, productRepoGetParams{
			productId: pr.ProductID, repoId: pr.RepoID, ref: pr.Ref,
		}))
	})
}

// RegisterProductRepoVersion records a version of a product_repo; re-registering refreshes it.
func (c *Client) RegisterProductRepoVersion(ctx context.Context, v ProductRepoVersion) error {
	if err := registryText(v.ProductRepoID, true, v.Category); err != nil {
		return err
	}
	if err := registryText(v.Version, true); err != nil {
		return err
	}
	operators, err := registryList(v.ClusterOperators)
	if err != nil {
		return err
	}
	images, err := registryList(v.Images)
	if err != nil {
		return err
	}
	_, err = c.register(ctx, func(conn *sql.Conn, s *sqlStore, at string) (string, error) {
		return "", s.queries.productRepoVersionUpsert(ctx, conn, productRepoVersionUpsertParams{
			productRepoId:    v.ProductRepoID,
			version:          v.Version,
			category:         v.Category,
			clusterOperators: operators,
			images:           images,
			registeredAt:     at,
		})
	})
	return err
}

// RegisterRepoOwner records a team that owns a product_repo; re-registering refreshes it.
func (c *Client) RegisterRepoOwner(ctx context.Context, o RepoOwner) error {
	if err := registryText(o.ProductRepoID, true, o.Manager, o.Source, o.JiraProject, o.JiraComponent); err != nil {
		return err
	}
	if err := registryText(o.Team, true); err != nil {
		return err
	}
	individuals, err := registryList(o.Individuals)
	if err != nil {
		return err
	}
	_, err = c.register(ctx, func(conn *sql.Conn, s *sqlStore, at string) (string, error) {
		return "", s.queries.repoOwnerUpsert(ctx, conn, repoOwnerUpsertParams{
			productRepoId: o.ProductRepoID,
			team:          o.Team,
			manager:       o.Manager,
			individuals:   individuals,
			source:        o.Source,
			jiraProject:   o.JiraProject,
			jiraComponent: o.JiraComponent,
			registeredAt:  at,
		})
	})
	return err
}

// FindProductRepo returns the product_repo_id for (product slug, repo URL, ref)
// without creating anything; found is false when the registry has no such row.
func (c *Client) FindProductRepo(ctx context.Context, slug, repoURL, ref string) (id string, found bool, err error) {
	if c == nil || c.store == nil {
		return "", false, wrap(OperationRead, PhaseInput, ErrNilDatabase)
	}
	conn, err := c.store.db.Conn(ctx)
	if err != nil {
		return "", false, wrap(OperationRead, PhaseConnect, err)
	}
	defer func() { _ = conn.Close() }()
	err = c.store.queries.productRepoFind(ctx, conn, productRepoFindParams{
		slug: slug, repoUrl: repoURL, ref: ref,
	}).Scan(&id)
	if errors.Is(err, sql.ErrNoRows) {
		return "", false, nil
	}
	if err != nil {
		return "", false, wrap(OperationRead, PhaseRead, err)
	}
	return id, true, nil
}

// RegistryProduct is a registered product as enumeration returns it.
// RegisteredAt is the RFC 3339 timestamp of the first registration.
type RegistryProduct struct {
	ProductID    string
	Slug         string
	Segment      *string
	RegisteredAt string
}

// RegistryProductRepo is a registered product_repo joined with its product and
// repo, as enumeration returns it. Ref "" is the default branch.
type RegistryProductRepo struct {
	ProductRepoID string
	ProductSlug   string
	Segment       *string
	RepoURL       string
	Ref           string
	SubService    *string
	ResourceType  *string
	RegisteredAt  string
}

// ProductScope is the product anchoring of one artifact binding: the
// product_repo the binding is anchored to, joined with its product. Bindings
// without a product_repo anchor are absent from the result.
type ProductScope struct {
	BindingID     string
	ProductRepoID string
	ProductID     string
	ProductSlug   string
	Ref           string
	Segment       string
}

// GetProductScopes resolves the product anchoring of every given artifact
// binding in one read. Bindings that are not anchored to a product_repo are
// absent from the returned map.
func (c *Client) GetProductScopes(ctx context.Context, bindingIDs []string) (map[string]ProductScope, error) {
	if c == nil || c.store == nil {
		return nil, wrap(OperationRead, PhaseInput, ErrNilDatabase)
	}
	ids, err := scopeValue(bindingIDs)
	if err != nil {
		return nil, wrap(OperationRead, PhaseInput, err)
	}
	conn, err := c.store.db.Conn(ctx)
	if err != nil {
		return nil, wrap(OperationRead, PhaseConnect, err)
	}
	defer func() { _ = conn.Close() }()
	rows, err := c.store.queries.productScopeGet(ctx, conn, productScopeGetParams{bindingIds: ids})
	if err != nil {
		return nil, wrap(OperationRead, PhaseRead, err)
	}
	defer func() { _ = rows.Close() }()
	scopes := make(map[string]ProductScope, len(bindingIDs))
	for rows.Next() {
		var scope ProductScope
		if err := rows.Scan(
			&scope.BindingID,
			&scope.ProductRepoID,
			&scope.ProductID,
			&scope.ProductSlug,
			&scope.Ref,
			&scope.Segment,
		); err != nil {
			return nil, wrap(OperationRead, PhaseRead, err)
		}
		scopes[scope.BindingID] = scope
	}
	if err := rows.Err(); err != nil {
		return nil, wrap(OperationRead, PhaseRead, err)
	}
	return scopes, nil
}

// RegistryFilter selects registry rows for enumeration. A nil ProductSlug or
// RepoURLPattern matches everything; RepoURLPattern uses SQL LIKE syntax
// (percent wildcards) and matches case-insensitively. Limit must be positive.
type RegistryFilter struct {
	ProductSlug    *string
	RepoURLPattern *string
	Limit          int
	Offset         int
}

func (f RegistryFilter) validate() error {
	if f.Limit <= 0 {
		return wrap(OperationRead, PhaseInput, ErrRegistryLimitInvalid)
	}
	if f.Offset < 0 {
		return wrap(OperationRead, PhaseInput, ErrRegistryLimitInvalid)
	}
	if f.ProductSlug != nil {
		if err := registryText(*f.ProductSlug, true); err != nil {
			return wrap(OperationRead, PhaseInput, err)
		}
	}
	if f.RepoURLPattern != nil {
		if err := registryText(*f.RepoURLPattern, true); err != nil {
			return wrap(OperationRead, PhaseInput, err)
		}
	}
	return nil
}

// filterPair expands one optional filter into the (flag, value) placeholder
// pair the search queries expect: an empty flag disables the condition.
func filterPair(value *string) (flag, actual string) {
	if value == nil {
		return "", ""
	}
	return *value, *value
}

// QueryProducts enumerates registered products ordered by slug.
func (c *Client) QueryProducts(ctx context.Context, filter RegistryFilter) ([]RegistryProduct, error) {
	if c == nil || c.store == nil {
		return nil, wrap(OperationRead, PhaseInput, ErrNilDatabase)
	}
	if err := filter.validate(); err != nil {
		return nil, err
	}
	conn, err := c.store.db.Conn(ctx)
	if err != nil {
		return nil, wrap(OperationRead, PhaseConnect, err)
	}
	defer func() { _ = conn.Close() }()
	slugFlag, slug := filterPair(filter.ProductSlug)
	rows, err := c.store.queries.productSearch(ctx, conn, productSearchParams{
		slugFilter: slugFlag, slug: slug,
		limit: int64(filter.Limit), offset: int64(filter.Offset),
	})
	if err != nil {
		return nil, wrap(OperationRead, PhaseRead, err)
	}
	defer func() { _ = rows.Close() }()
	out := []RegistryProduct{}
	for rows.Next() {
		var p RegistryProduct
		if err := rows.Scan(&p.ProductID, &p.Slug, &p.Segment, &p.RegisteredAt); err != nil {
			return nil, wrap(OperationRead, PhaseRead, err)
		}
		out = append(out, p)
	}
	if err := rows.Err(); err != nil {
		return nil, wrap(OperationRead, PhaseRead, err)
	}
	return out, nil
}

// QueryProductRepos enumerates registered product_repos joined with their
// product and repo, ordered by product slug, repo URL, then ref.
func (c *Client) QueryProductRepos(ctx context.Context, filter RegistryFilter) ([]RegistryProductRepo, error) {
	if c == nil || c.store == nil {
		return nil, wrap(OperationRead, PhaseInput, ErrNilDatabase)
	}
	if err := filter.validate(); err != nil {
		return nil, err
	}
	conn, err := c.store.db.Conn(ctx)
	if err != nil {
		return nil, wrap(OperationRead, PhaseConnect, err)
	}
	defer func() { _ = conn.Close() }()
	slugFlag, slug := filterPair(filter.ProductSlug)
	patternFlag, pattern := filterPair(filter.RepoURLPattern)
	rows, err := c.store.queries.productRepoSearch(ctx, conn, productRepoSearchParams{
		slugFilter: slugFlag, slug: slug,
		repoUrlFilter: patternFlag, repoUrlPattern: pattern,
		limit: int64(filter.Limit), offset: int64(filter.Offset),
	})
	if err != nil {
		return nil, wrap(OperationRead, PhaseRead, err)
	}
	defer func() { _ = rows.Close() }()
	out := []RegistryProductRepo{}
	for rows.Next() {
		var row RegistryProductRepo
		if err := rows.Scan(
			&row.ProductRepoID,
			&row.ProductSlug,
			&row.Segment,
			&row.RepoURL,
			&row.Ref,
			&row.SubService,
			&row.ResourceType,
			&row.RegisteredAt,
		); err != nil {
			return nil, wrap(OperationRead, PhaseRead, err)
		}
		out = append(out, row)
	}
	if err := rows.Err(); err != nil {
		return nil, wrap(OperationRead, PhaseRead, err)
	}
	return out, nil
}

type registryWrite func(conn *sql.Conn, store *sqlStore, registeredAt string) (string, error)

func (c *Client) register(ctx context.Context, write registryWrite) (string, error) {
	if c == nil || c.store == nil {
		return "", wrap(OperationRegister, PhaseInput, ErrNilDatabase)
	}
	conn, err := c.store.db.Conn(ctx)
	if err != nil {
		return "", wrap(OperationRegister, PhaseConnect, err)
	}
	defer func() { _ = conn.Close() }()
	if err := c.store.prepareConnection(ctx, conn); err != nil {
		return "", wrap(OperationRegister, PhaseConnect, err)
	}
	if err := begin(ctx, conn, c.store.dialect); err != nil {
		return "", wrap(OperationRegister, PhaseBegin, err)
	}
	committed := false
	defer rollbackUnlessCommitted(ctx, conn, &committed)

	id, err := write(conn, c.store, nowUTC())
	if err != nil {
		return "", wrap(OperationRegister, PhaseRegistry, err)
	}
	if _, err := conn.ExecContext(ctx, "COMMIT"); err != nil {
		return "", wrap(OperationRegister, PhaseCommit, err)
	}
	committed = true
	return id, nil
}

func storedID(row *sql.Row) (string, error) {
	var id string
	if err := row.Scan(&id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", ErrRegistryRowMissing
		}
		return "", err
	}
	return id, nil
}

// registryText validates a key value (required: non-empty) and optional
// attributes: NUL-free UTF-8, matching contracts' text rules.
func registryText(value string, required bool, optional ...*string) error {
	if required && value == "" {
		return wrap(OperationRegister, PhaseInput, ErrRegistryValueMissing)
	}
	if !validBindingText(value) {
		return wrap(OperationRegister, PhaseInput, ErrInvalidIdentifier)
	}
	for _, attribute := range optional {
		if attribute != nil && !validBindingText(*attribute) {
			return wrap(OperationRegister, PhaseInput, ErrInvalidIdentifier)
		}
	}
	return nil
}

// registryList encodes a list attribute as a compact JSON array; empty stays NULL.
func registryList(values []string) (*string, error) {
	if len(values) == 0 {
		return nil, nil
	}
	for _, value := range values {
		if err := registryText(value, true); err != nil {
			return nil, err
		}
	}
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(values); err != nil {
		return nil, wrap(OperationRegister, PhaseInput, err)
	}
	text := string(bytes.TrimRight(buf.Bytes(), "\n"))
	return &text, nil
}

func newRegistryID() string {
	return uuid.NewString()
}
