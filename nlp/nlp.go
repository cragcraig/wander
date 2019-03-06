package nlp

import (
	"github.com/cragcraig/wander/env"
	"regexp"
	"strings"
)

var cmdRegexp *regexp.Regexp = regexp.MustCompile(`^(\w+)(?:\s+(\S.*))?$`)

func ParsePlayerAction(player *env.Player, s string) *env.Action {
	res := cmdRegexp.FindStringSubmatch(strings.TrimSpace(s))
	if res != nil {
		cmd := res[1]
		for i := range player.Private {
			v := player.Private[i].WhatCanThisDo()
			for j := range v {
				aliases := v[j].CommandAliases
				for k := range aliases {
					if cmd == aliases[k] {
						var arg interface{}
						if len(res) > 2 && v[j].ArgParser != nil {
							arg = v[j].ArgParser(res[2])
						}
						return player.CreateAction(v[j], player.Private[i], nil, arg)
					}
				}
			}
		}
		player.Conn.Write <- "Unrecognized command \"" + cmd + "\"."
		return nil
	}
	player.Conn.Write <- "Invalid command."
	return nil
}
