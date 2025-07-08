package version

import (
	"encoding/json"
	"fmt"
	"runtime/debug"
	"strings"

	"github.com/blang/semver/v4"
	log "github.com/sirupsen/logrus"

	"github.com/geomyidia/util/pkg/errors"
)

const na string = "N/A"

type Version struct {
	Binary         string         `json:"executable"`
	ShortName      string         `json:"short-name"`
	LongName       string         `json:"long-name"`
	Description    string         `json:"description"`
	Semantic       semver.Version `json:"version"`
	Go             string         `json:"go-version"`
	BuildPath      string         `json:"build-path"`
	CurrentBranch  string         `json:"current-branch"`
	CommitID       string         `json:"commid-id"`
	BuildTime      string         `json:"build-time"`
	GitSummary     string         `json:"git-summery"`
	DeployedFromIP string         `json:"deployed-from-ip"`
}

func New(
	version,
	shortName,
	longName,
	desc,
	branch,
	commitID,
	buildTime,
	gitSummary,
	deployedFromIP string,
) (*Version, error) {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		return nil, errors.Combine(errors.ErrVersionData, errors.ErrBuildInfo)
	}
	sem, err := semver.Parse(version)
	if err != nil {
		return nil, errors.Combine(errors.ErrVersionData, err)
	}

	short := strings.TrimSpace(shortName)
	v := &Version{
		Binary:        short,
		ShortName:     short,
		LongName:      strings.TrimSpace(longName),
		Description:   strings.TrimSpace(desc),
		Semantic:      sem,
		BuildPath:     bi.Path,
		CurrentBranch: strings.TrimSpace(branch),
		CommitID:      strings.TrimSpace(commitID),
		BuildTime:     strings.TrimSpace(buildTime),
		GitSummary:    strings.TrimSpace(gitSummary),
		Go:            bi.GoVersion,
	}
	if deployedFromIP != "" {
		v.DeployedFromIP = strings.TrimSpace(deployedFromIP)
	} else {
		v.DeployedFromIP = na
	}
	return v, nil
}

// BuildString ...
func (version *Version) BuildString() string {
	if version.CommitID == "" {
		return na
	}
	return fmt.Sprintf("%s@%s, %s", version.CurrentBranch, version.CommitID, version.BuildTime)
}

// String ...
func (version *Version) String() string {
	return version.Semantic.String()
}

func (version *Version) AsJSON() string {
	data, err := json.Marshal(version)
	if err != nil {
		log.Error(errors.Combine(errors.ErrVersionData, err))
	}
	return string(data)
}

func (version *Version) Major() int {
	return int(version.Semantic.Major)
}

func (version *Version) Minor() int {
	return int(version.Semantic.Minor)
}

func (version *Version) Patch() int {
	return int(version.Semantic.Patch)
}

func (version *Version) Micro() int {
	return version.Patch()
}

// VersionedBuildString ...
func (version *Version) VersionedBuildString() string {
	gc := version.CommitID
	if gc == "" {
		gc = na
	}
	return fmt.Sprintf("%s, %s@%s, %s", version.String(), version.CurrentBranch, gc, version.BuildTime)
}
