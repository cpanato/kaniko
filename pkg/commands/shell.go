/*
Copyright 2018 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package commands

import (
	"github.com/chainguard-dev/kaniko/pkg/dockerfile"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/moby/buildkit/frontend/dockerfile/instructions"
)

type ShellCommand struct {
	BaseCommand
	cmd *instructions.ShellCommand
}

// ExecuteCommand handles command processing similar to CMD and RUN,
func (s *ShellCommand) ExecuteCommand(config *v1.Config, buildArgs *dockerfile.BuildArgs) error {
	config.Shell = s.cmd.Shell
	return nil
}

// String returns some information about the command for the image config history
func (s *ShellCommand) String() string {
	return s.cmd.String()
}
