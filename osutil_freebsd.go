//
// osutil_freebsd.go
//
// Copyright © 2012-2013 Damicon Kraa Oy
//
// This file is part of zfswatcher.
//
// Zfswatcher is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Zfswatcher is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with zfswatcher. If not, see <http://www.gnu.org/licenses/>.
//

package main

import (
	"time"
)

// Returns system uptime as time.Duration.
func getSystemUptime() (uptime time.Duration, err error) {
	return 0, nil
}

// Returns system load averages.
func getSystemLoadaverage() (la [3]float32, err error) {
	return la, nil
}

// Device lookup paths. (This list comes from lib/libzfs/libzfs_import.c)
var deviceLookupPaths = [...]string{
	"/dev",
}

// eof
