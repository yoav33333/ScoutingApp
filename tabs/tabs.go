package tabs

import (
	"fyne.io/fyne/v2/container"
	"scouting_app/dataBase"
)

func CreateTabs(db *dataBase.DB) *container.AppTabs {
	ret := container.NewAppTabs()
	ret.Append(CreateLoginTab(db, ret))
	ret.Append(NewCreateScoutTab(db))
	ret.SetTabLocation(container.TabLocationBottom)
	return ret
}
