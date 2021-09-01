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

package shell

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	. "github.com/showalter/tsh/internal/builtins"
	. "github.com/showalter/tsh/internal/env"
)

type Shell struct {
	builtins map[string]BuiltInFunction
	env      Environment
}

func (shell Shell) RunInteractiveShell() {
	shell.builtins = make(map[string]BuiltInFunction)
	shell.env = make(map[string]string)

	reader := bufio.NewReader(os.Stdin)
	shell.env.Put("PWD", os.Getenv("PWD"))
	shell.env.Put("HOME", os.Getenv("HOME"))

	for {
		fmt.Print("$ ")

		// Read keyboard input
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if err = shell.execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func (shell Shell) loadBuiltins() {
	shell.builtins["cd"] = Cd
	shell.builtins["exit"] = Exit
	shell.builtins["env"] = Env
}

func (shell Shell) execInput(input string) error {
	shell.loadBuiltins()

	// Remove newline
	input = strings.TrimSuffix(input, "\n")

	args := strings.Split(input, " ")

	// Check if the program is a built in program, and run it if it is
	if val, ok := shell.builtins[args[0]]; ok {
		return val(args, shell.env)
	}

	// Otherwise, exec the program
	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
