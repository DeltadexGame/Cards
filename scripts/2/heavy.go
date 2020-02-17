package scripts

import (
	"deltadex/gameplay"
	"deltadex/gameplay/events"
	"fmt"
)

func HandleMonsterDamageEvent(event events.Event) (events.Event, bool) {
	fmt.Println("running")
	fmt.Println(event.EventInfo)
	monster := event.EventInfo["monster"].(gameplay.Monster)
	if monster.Ability.AbilityID != 2 {
		return event, false
	}
	fmt.Println("got to halfing")
	event.EventInfo["damage"] = event.EventInfo["damage"].(int) / 2
	fmt.Println("Using heavy ability, damage was halved from " + string(event.EventInfo["damage"].(int)*2) + " to " + string(event.EventInfo["damage"].(int)) + ".")
	return event, true
}
