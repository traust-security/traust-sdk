// Code generated from traust-contracts bf4eb27b9eff357587fe1369db3b43e4a9534eba. DO NOT EDIT.

package types

import (
	"github.com/traust-security/traust-sdk/go/v1/enums"
)

type BelowRangePackets struct {
	Artifact        interface{}            `json:"artifact"`
	Candidates      int                    `json:"candidates"`
	CoverageGaps    []CoverageGap          `json:"coverage_gaps"`
	InstalledExact  int                    `json:"installed_exact"`
	Packets         []Packet               `json:"packets"`
	RepositoryRefId string                 `json:"repository_ref_id"`
	SkippedInstalls map[string]int         `json:"skipped_installs"`
	Status          enums.BelowRangeStatus `json:"status"`
}

type Advisory struct {
	Affected   Affected `json:"affected"`
	Aliases    []string `json:"aliases"`
	Details    string   `json:"details"`
	Identifier string   `json:"identifier"`
	Summary    string   `json:"summary"`
}

type Affected struct {
	Ecosystem enums.Ecosystem `json:"ecosystem"`
	Imports   []string        `json:"imports,omitempty"`
	Name      string          `json:"name"`
	Spans     []Span          `json:"spans,omitempty"`
	Symbols   []string        `json:"symbols,omitempty"`
	Versions  []string        `json:"versions,omitempty"`
}

type Candidate struct {
	Advisory      string                        `json:"advisory"`
	Aliases       []string                      `json:"aliases,omitempty"`
	Ecosystem     enums.Ecosystem               `json:"ecosystem"`
	Kind          enums.BelowRangeCandidateKind `json:"kind"`
	MajorDistance interface{}                   `json:"major_distance"`
	Package       string                        `json:"package"`
	StatedStart   string                        `json:"stated_start"`
	Version       string                        `json:"version"`
}

type Commit struct {
	Sha     string `json:"sha"`
	Subject string `json:"subject"`
}

type CoverageGap struct {
	Case      interface{}               `json:"case,omitempty"`
	Detail    interface{}               `json:"detail,omitempty"`
	Ecosystem *enums.Ecosystem          `json:"ecosystem,omitempty"`
	Package   interface{}               `json:"package,omitempty"`
	Reason    enums.BelowRangeGapReason `json:"reason"`
	Tried     []FixAttempt              `json:"tried,omitempty"`
	Version   interface{}               `json:"version,omitempty"`
}

type DeclaredSource struct {
	Commit     string `json:"commit"`
	Nuspec     string `json:"nuspec"`
	Repository string `json:"repository"`
}

type DependencyGap struct {
	Package interface{} `json:"package,omitempty"`
	Reason  string      `json:"reason"`
	Version interface{} `json:"version,omitempty"`
}

type DependencySource struct {
	Files      int      `json:"files"`
	ImportedBy []string `json:"imported_by"`
	Package    string   `json:"package"`
	Root       string   `json:"root"`
	Sha256     string   `json:"sha256"`
	Url        string   `json:"url"`
	Version    string   `json:"version"`
}

type FilePresence struct {
	Hunks         []HunkPresence `json:"hunks,omitempty"`
	InstalledPath interface{}    `json:"installed_path"`
	PatchPath     string         `json:"patch_path"`
}

type Fix struct {
	Alternatives  []FixAttempt   `json:"alternatives"`
	Chosen        FixAttempt     `json:"chosen"`
	FixedReleases []FixedRelease `json:"fixed_releases"`
}

type FixAttempt struct {
	CommitDate           interface{} `json:"commit_date,omitempty"`
	Commits              []Commit    `json:"commits,omitempty"`
	Error                interface{} `json:"error,omitempty"`
	Files                interface{} `json:"files,omitempty"`
	FilesInInstalledTree interface{} `json:"files_in_installed_tree,omitempty"`
	NamesAdvisory        interface{} `json:"names_advisory,omitempty"`
	OnlyAddsCode         interface{} `json:"only_adds_code,omitempty"`
	Rank                 int         `json:"rank"`
	ReleaseCommit        interface{} `json:"release_commit,omitempty"`
	Url                  string      `json:"url"`
}

type FixedRelease struct {
	Published interface{} `json:"published"`
	Version   string      `json:"version"`
}

type HunkPresence struct {
	AddedLines     int `json:"added_lines"`
	AddedPresent   int `json:"added_present"`
	RemovedLines   int `json:"removed_lines"`
	RemovedPresent int `json:"removed_present"`
}

type InstalledSource struct {
	DeclaredBy *DeclaredSource `json:"declared_by,omitempty"`
	Files      int             `json:"files"`
	Root       string          `json:"root"`
	Sha256     string          `json:"sha256"`
	Url        string          `json:"url"`
}

type Packet struct {
	Advisory          Advisory           `json:"advisory"`
	Candidate         Candidate          `json:"candidate"`
	Case              string             `json:"case"`
	DependencyGaps    []DependencyGap    `json:"dependency_gaps,omitempty"`
	DependencySources []DependencySource `json:"dependency_sources,omitempty"`
	Fix               Fix                `json:"fix"`
	HunkPresence      []FilePresence     `json:"hunk_presence"`
	InstalledSource   InstalledSource    `json:"installed_source"`
	Patch             string             `json:"patch"`
}

type Span struct {
	Fixed        interface{} `json:"fixed,omitempty"`
	Introduced   string      `json:"introduced"`
	LastAffected interface{} `json:"last_affected,omitempty"`
}
