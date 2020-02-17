package scripts

import (
	"deltadex/gameplay"
	"deltadex/gameplay/events"
	"fmt"
	"strconv"
)

func HandleMonsterDamageEvent(event events.Event) (events.Event, bool) {
	fmt.Println("running")
	info := event.EventInfo
	monster := info["monster"].(gameplay.Monster)
	if monster.Ability.AbilityID != 2 {
		return event, false
	}
	damage := info["damage"].(int)
	info["damage"] = damage / 2
	fmt.Println("Using heavy ability, damage was halved from " + strconv.Itoa(damage*2) + " to " + strconv.Itoa(damage) + ".")
	event.EventInfo = info
	return event, true
}
