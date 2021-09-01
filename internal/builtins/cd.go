/*
 * Copyright 2021 Ryan Showalter
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License. You may obtain a copy
 * of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package builtins

import (
	"os"

	. "github.com/showalter/tsh/internal/env"
	_ "github.com/showalter/tsh/internal/errors"
)

var homeDirectory string = os.Getenv("HOME")

func Cd(args []string, env Environment) error {
	// Handle going to $HOME if no arguments are given.
	if len(args) == 1 {
		return os.Chdir(homeDirectory)
	}

	var newDirectory string

	switch args[1] {
	case "~": // Equivalent to $HOME
		newDirectory = homeDirectory
	case "-":
		newDirectory = env.Get("OLDPWD")
	default:
		newDirectory = args[1]
	}

	// Go ahead and change directory
	err := os.Chdir(newDirectory)

	// If that worked out, update environment variables
	if err == nil {
		env.Put("OLDPWD", env.Get("PWD"))

		pwd, getwdErr := os.Getwd()
		if getwdErr == nil {
			env.Put("PWD", pwd)
		}
	}

	return err
}
