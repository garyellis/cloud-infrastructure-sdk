package tflive

import (
	"path/filepath"

	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/input"
)

const metadataYamlShFile = "metadata.yaml.sh"

type MetadataYamlSh struct {
	input.Input
}

func (m *MetadataYamlSh) GetInput() (input.Input, error) {
	if m.Path == "" {
		m.Path = filepath.Join(LiveBaseDir, metadataYamlShFile)
	}
	m.TemplateBody = MetadataYamlShTmpl

	m.IsExec = true

	return m.Input, nil
}

const MetadataYamlShTmpl = `#!/bin/bash

DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
source $DIR/../../../version

cat <<EOF > $DIR/metadata.yaml
---
tags:
  project-id: "${PROJECT_ID}"
  project-name: "${NAME}"
  repo-url: ${REPO_URL}
  version: ${VERSION}
  commit: ${GIT_COMMIT_SHORT}
EOF
`
