// Copyright 2018. Akamai Technologies, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package commands

import (
	"strings"

	"github.com/akamai/cli/pkg/config"
	"github.com/akamai/cli/pkg/terminal"

	"github.com/urfave/cli/v2"
)

func cmdConfigSet(c *cli.Context) error {
	section, key := parseConfigPath(c)

	value := strings.Join(c.Args().Tail(), " ")

	config.SetConfigValue(section, key, value)
	return config.SaveConfig()
}

func cmdConfigGet(c *cli.Context) error {
	section, key := parseConfigPath(c)

	terminal.
		Get(c.Context).
		Writeln(config.GetConfigValue(section, key))

	return nil
}

func cmdConfigUnset(c *cli.Context) error {
	section, key := parseConfigPath(c)

	config.UnsetConfigValue(section, key)
	return config.SaveConfig()
}

func cmdConfigList(c *cli.Context) error {
	config, err := config.OpenConfig()
	if err != nil {
		return err
	}

	term := terminal.Get(c.Context)

	if c.NArg() > 0 {
		sectionName := c.Args().First()
		section := config.Section(sectionName)
		for _, key := range section.Keys() {
			term.Printf("%s.%s = %s\n", sectionName, key.Name(), key.Value())
		}

		return nil
	}

	for _, section := range config.Sections() {
		for _, key := range section.Keys() {
			term.Printf("%s.%s = %s\n", section.Name(), key.Name(), key.Value())
		}
	}
	return nil
}

func parseConfigPath(c *cli.Context) (string, string) {
	path := strings.Split(c.Args().First(), ".")
	section := path[0]
	key := strings.Join(path[1:], "-")
	return section, key
}
