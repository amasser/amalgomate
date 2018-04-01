// Copyright (c) 2016 Palantir Technologies Inc. All rights reserved.
// Use of this source code is governed by the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Copyright 2016 Palantir Technologies, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"os"

	"github.com/palantir/godel/framework/pluginapi/v2/pluginapi"
	"github.com/palantir/pkg/cobracli"

	"github.com/palantir/amalgomate/godelplugin/cmd"
)

func main() {
	if ok := pluginapi.InfoCmd(os.Args, os.Stdout, cmd.PluginInfo); ok {
		return
	}
	os.Exit(cobracli.ExecuteWithDefaultParamsWithVersion(cmd.RootCmd, &cmd.DebugFlag, ""))
}