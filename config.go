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

/*
 * This file contains the configuration file parser
 */

package main

import (
    "fmt"
    "os"
    "reflect"
    "encoding/json"
)

type Config struct {
    ListenAddress string
    ListenPort uint16
    ExpiryDays int
    MaxPasteSizeBytes int
    ServerURL string
    FileStorePath string
}

type ConfigValidationError string
func (e ConfigValidationError) Error() string {
    return fmt.Sprintf("Failed to validate config field `%s`. Check your config.json", string(e))
}

func validateConfig(c Config) error {
    value := reflect.ValueOf(&c)
    elem := value.Elem()

    empty := Config{}
    e_value := reflect.ValueOf(&empty)
    e_elem := e_value.Elem()

    for i := 0; i < elem.NumField(); i++ {
        if elem.Field(i).Interface() == e_elem.Field(i).Interface() {
            return ConfigValidationError(reflect.TypeOf(c).Field(i).Name)
        }
    }

    return nil
}

func ParseConfig(path string) (Config, error) {
    file, err := os.Open(path)
    if err != nil {
        return Config{}, err
    }
    defer file.Close()

    // Parse config file
    config := Config{}
    jsonDecoder := json.NewDecoder(file)
    err = jsonDecoder.Decode(&config)
    if err != nil {
        return Config{}, err
    }

    // Validate config
    err = validateConfig(config)
    if err != nil {
        return Config{}, err
    }

    return config, nil
}
