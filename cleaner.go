/**
 * Copyright 2019 Shawn Anastasio
 *
 * This file is part of upaste.
 *
 * upaste is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * upaste is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with upaste.  If not, see <https://www.gnu.org/licenses/>.
 */

/**
 * This file contains routines to remove out-of-date pastes
 * from the filesystem.
 */

package main

import (
    "os"
    "log"
    "time"
    "path/filepath"
)

func StartCleaner(c Config) {
    // Iterate through all files in the store
    err := filepath.Walk(c.FileStorePath, func(path string, info os.FileInfo, err error) error {
        // Skip the store itself
        if path == c.FileStorePath {
            return nil
        }

        expire := info.ModTime().AddDate(0, 0, c.ExpiryDays)
        if expire.Before(time.Now()) {
            log.Printf("Deleting expired file: %s\n", info.Name())

            if err := os.Remove(path); err != nil {
                log.Printf("Failed to remove expired file %s: %v\n", path, err)
                return err
            }

        }

        return nil
    })

    if err != nil {
        log.Printf("Error encountered while cleaning store: %v\n", err)
    }

    // Sleep for 1 hour before cleaning again
    const secToNs = 1000000000
    time.Sleep(1 * 60 * 60 * secToNs)
}
