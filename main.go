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

package main

import (
    "log"
    "flag"
)

func main() {
    configFile := flag.String("config", "config.json", "Configuration file path")
    flag.Parse()

    config, err := ParseConfig(*configFile)
    if err != nil {
        log.Printf("Unable to open configuration file: %v\n", err)
        log.Fatalln("Please create it or specify a different file with --config.")
    }

    StartServer(config)
}
