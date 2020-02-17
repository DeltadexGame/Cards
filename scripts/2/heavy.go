package scripts

import (
	"deltadex/gameplay"
	"deltadex/gameplay/events"
)

func HandleMonsterDamageEvent(event *events.Event) {
	monster := event.EventInfo["monster"].(gameplay.Monster)
	if monster.Ability.AbilityID != 2 {
		return
	}
	event.EventInfo["damage"] = event.EventInfo["damage"].(int) / 2
}
