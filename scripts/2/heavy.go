package scripts

import (
	"deltadex/gameplay"
	"deltadex/gameplay/events"
	"fmt"
)

func HandleMonsterDamageEvent(event events.Event) (events.Event, bool) {
	returnEvent := event
	fmt.Println("running")
	info := event.EventInfo
	monster := info["monster"].(gameplay.Monster)
	if monster.Ability.AbilityID != 2 {
		return event, false
	}
	damage := info["damage"].(int)
	returnEvent.EventInfo["damage"] = damage / 2
	fmt.Println("Using heavy ability, damage was halved from " + string(damage*2) + " to " + string(damage) + ".")
	return returnEvent, true
}
