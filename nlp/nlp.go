package nlp

import (
	"github.com/gnarlyskier/wander/env"
	"regexp"
)

var cmdRegexp *regexp.Regexp = regexp.MustCompile(`^(\w+)(?:\s+(.+))?$`)

func ParsePlayerAction(player *env.Player, s string) *env.Action {
	res := cmdRegexp.FindStringSubmatch(s)
	if res != nil {
		cmd := res[1]
		for i := range player.Private {
			v := player.Private[i].WhatCanThisDo()
			for j := range v {
				aliases := v[j].CommandAliases
				for k := range aliases {
					if cmd == aliases[k] {
						var arg string
						if len(res) > 2 {
							arg = res[2]
						}
						return player.CreateAction(v[j], player.Private[i], nil, arg)
					}
				}
			}
		}
		player.Conn.Write <- "Unrecognized command \"" + cmd + "\"."
		return nil
	}
	player.Conn.Write <- "Invalid input."
	return nil
}
