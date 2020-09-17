// Copyright 2020 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import "context"

// Args provides plugin execution arguments.
type Args struct {
	Pipeline

	// Level defines the plugin log level.
	Level string `envconfig:"PLUGIN_LOG_LEVEL"`

	// TODO replace or remove
	Param1 string `envconfig:"PLUGIN_PARAM1"`
	Param2 string `envconfig:"PLUGIN_PARAM2"`
}

// Exec executes the plugin.
func Exec(ctx context.Context, args Args) error {
	// write code here
	return nil
}
