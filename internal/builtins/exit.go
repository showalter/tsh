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
	"errors"
	"os"

	. "github.com/showalter/tsh/internal/env"
)

func Exit(_ []string, _ Environment) error {
	os.Exit(0)

	// This is unreachable, but it is necessary for this function to return
	// an error to be considered a BuiltInFunction. The exit builtin is a
	// special case in this regard.
	return errors.New("")
}
