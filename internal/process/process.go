// Copyright 2021 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// +build windows

package process

import (
	"context"
	"strings"

	"github.com/pkg/errors"
	"github.com/shirou/gopsutil/process"
	log "unknwon.dev/clog/v2"
)

// List returns all the windows processes.
func List(ctx context.Context) ([]*process.Process, error) {
	return process.ProcessesWithContext(ctx)
}

// KillByKeyword kills the windows processes whose name contains the given keyword.
func KillByKeyword(ctx context.Context, keyword string) error {
	processes, err := List(ctx)
	if err != nil {
		return errors.Wrap(err, "list process")
	}

	for _, p := range processes {
		name, err := p.Name()
		if err != nil {
			continue
		}
		if strings.Contains(name, keyword) {
			if err := p.KillWithContext(ctx); err != nil {
				return errors.Wrap(err, "kill")
			}
			log.Trace("Kill process: %v", name)
		}
	}
	return nil
}
