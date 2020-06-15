package scaffold

import (
	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/input"
)

const versionShFile = "version"

type VersionSh struct {
	input.Input
}

func (v *VersionSh) GetInput() (input.Input, error) {
	if v.Path == "" {
		v.Path = versionShFile
	}
	v.TemplateBody = VersionShTmpl

	v.IsExec = true
	v.IfExistsAction = input.Skip
	return v.Input, nil
}

const VersionShTmpl = `NAME={{.ProjectName}}
VERSION=v0.0.0
PROJECT_ID="-"
REPO_URL="-"
ISSUES_URL="-"

[ "$(git status >/dev/null 2>&1 ; echo $?)" == 0 ] && IS_GIT_REPO="true"
IS_GIT_REPO="$IS_GIT_REPO"

[ "$IS_GIT_REPO" ] && GIT_COMMIT="$(git rev-list -1 HEAD)"
GIT_COMMIT="$GIT_COMMIT"

[ "$IS_GIT_REPO" ] && GIT_COMMIT_SHORT="$(git rev-parse --short -1 HEAD)"
GIT_COMMIT_SHORT="$GIT_COMMIT_SHORT"

[ "$IS_GIT_REPO" ] && GIT_BRANCH="$(git rev-parse --abbrev-ref HEAD | sed 's@[\/]@-@g')"
GIT_BRANCH="$GIT_BRANCH"

[ "$GIT_BRANCH" == "master" ] && IS_MASTER="true" || IS_MASTER="false"
IS_MASTER="$IS_MASTER"

## CI_VERSION is the version passed to the ci system. i.e. bamboo release version
[ "$IS_GIT_REPO" ] && CI_VERSION="${NAME}-${VERSION}_${GIT_COMMIT_SHORT}"
CI_VERSION="$CI_VERSION"
`
