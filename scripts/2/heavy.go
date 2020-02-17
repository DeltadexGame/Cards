package scripts

import (
	"deltadex/gameplay"
	"deltadex/gameplay/events"
)

func HandleMonsterDamageEvent(event events.Event) (events.Event, bool) {
	monster := event.EventInfo["monster"].(gameplay.Monster)
	if monster.Ability.AbilityID != 2 {
		return Event{}, false
	}
	event.EventInfo["damage"] = event.EventInfo["damage"].(int) / 2
	return event, true
}
