// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "golang.org/x/mobile/app"

// This tiny program demonstrates https://github.com/golang/go/issues/10469.
// Run the app, rotate the device to reproduce the issue.
func main() {
	app.Run(app.Callbacks{})
}
