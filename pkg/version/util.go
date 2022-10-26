package version

// XXX Move this into geomyidia project ....

import "fmt"

const na string = "N/A"

// ProjectVersion data
type ProjectVersion struct {
	Semantic   string
	BuildDate  string
	GitCommit  string
	GitBranch  string
	GitSummary string
}

// BuildString ...
func BuildString(version *ProjectVersion) string {
	if version.GitCommit == "" {
		return na
	}
	return fmt.Sprintf("%s@%s, %s", version.GitBranch, version.GitCommit, version.BuildDate)
}

// String ...
func String(version *ProjectVersion) string {
	if version.Semantic == "" {
		return na
	}
	return version.Semantic
}

// VersionedBuildString ...
func VersionedBuildString(version *ProjectVersion) string {
	v := version.Semantic
	gc := version.GitCommit
	if v == "" {
		v = na
	}
	if gc == "" {
		gc = na
	}
	return fmt.Sprintf("%s, %s@%s, %s", v, version.GitBranch, gc, version.BuildDate)
}
