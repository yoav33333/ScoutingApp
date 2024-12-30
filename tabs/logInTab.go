package tabs

import (
	"context"
	"fmt"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"scouting_app/dataBase"
	"strconv"
)

func CreateLoginTab(db *dataBase.DB, mainTabs *container.AppTabs) *container.TabItem {
	code := widget.NewEntry()
	year := widget.NewEntry()
	userName := widget.NewEntry()
	code.SetPlaceHolder("event code")
	year.SetPlaceHolder("event year")

	userName.SetPlaceHolder("user name")
	return container.NewTabItemWithIcon("login tab", theme.LoginIcon(), container.NewVBox(code, year, userName, widget.NewButton("login", func() {
		Code := code.Text
		var err error = nil
		Year, err := strconv.Atoi(year.Text)
		if err != nil {
		}
		*db = *(dataBase.CreateDB(Year, Code, userName.Text))
		db.CreateUser(context.Background(), userName.Text)
		fmt.Printf("code: %v\n", code.Text)
		fmt.Printf("year: %v\n", year.Text)
		fmt.Printf("user name: %v\n", userName.Text)
	}),
		widget.NewButton("refresh", func() {
			for id, team := range db.GetUserSchedule(context.Background()) {
				fmt.Printf("%v: %v\n", id, team)
			}
			mainTabs.Remove(mainTabs.Items[1])
			mainTabs.Append(NewCreateScoutTab(db))
		}),
	))
}
