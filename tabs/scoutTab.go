package tabs

import (
	"context"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"scouting_app/consts"
	"scouting_app/dataBase"
	"slices"
	"strconv"
)

type Counter struct {
	object  fyne.CanvasObject
	counter *int
}

type FullCounter struct {
	baseCounter *Counter
	refCounter  *Counter
}

type Results struct {
	ascent          string
	netZone         string
	observationZone string
	highBasket      string
	lowBasket       string
	highChamber     string
	lowChamber      string
}

func NewCounter() *FullCounter {
	var baseCounter = new(int)
	var refCounter = new(int)
	*baseCounter = 0
	*refCounter = 0

	labelB := widget.NewLabel(strconv.Itoa(*baseCounter))
	labelR := widget.NewLabel(strconv.Itoa(*refCounter))
	addB := widget.NewButton("add", func() {
		*baseCounter++
		*refCounter++
		labelB.SetText(strconv.Itoa(*baseCounter))
		labelR.SetText(strconv.Itoa(*refCounter))
	})
	subB := widget.NewButton("sub", func() {
		if *baseCounter <= 0 {
			return
		}
		*baseCounter--
		*refCounter--
		labelB.SetText(strconv.Itoa(*baseCounter))
		labelR.SetText(strconv.Itoa(*refCounter))
	})
	addR := widget.NewButton("add", func() {
		*refCounter++
		labelR.SetText(strconv.Itoa(*refCounter))
	})
	subR := widget.NewButton("sub", func() {
		if *baseCounter <= 0 {
			return
		}
		*refCounter--
		labelR.SetText(strconv.Itoa(*refCounter))
	})

	base := &Counter{container.NewHBox(addB, subB, labelB), baseCounter}
	ref := &Counter{container.NewHBox(addR, subR, labelR), refCounter}
	return &FullCounter{base, ref}
}

func createMatchTab(matchNum string, teamNum int, db *dataBase.DB) *container.TabItem {
	temp := make(map[string]*FullCounter)

	formAuto := widget.NewForm()
	formAuto.Append("team: ", widget.NewLabel(strconv.Itoa(teamNum)))
	formAuto.Append("auto", widget.NewLabel(""))
	for _, s := range consts.GetGameTemp() {
		temp[s] = NewCounter()
		formAuto.Append(s, temp[s].baseCounter.object)
	}
	formTele := widget.NewForm()

	formTele.Append("teleop", widget.NewLabel(""))
	for _, s := range consts.GetGameTemp() {
		formTele.Append(s, temp[s].refCounter.object)
	}

	submit := widget.NewButton("Submit", func() {
		ret := make(map[string]int)
		for _, i := range consts.GetGameTemp() {
			ret[i] = *temp[i].baseCounter.counter + *temp[i].refCounter.counter
			fmt.Println(ret[i])
		}
		err := db.SubmitData(context.Background(), ret, matchNum, strconv.Itoa(teamNum))
		if err != nil {
			return
		}
	})
	next := widget.NewButton("Next", func() {})
	prev := widget.NewButton("Prev", func() {})

	return container.NewTabItem(matchNum, container.NewVScroll(container.NewVBox(formAuto, formTele, container.NewHBox(prev, submit, next))))
}

func newCreateMatchTab(matchNum string, teamNum int, db *dataBase.DB, otherPages []*container.Scroll, index int, base *fyne.Container) *container.Scroll {
	temp := make(map[string]*FullCounter)

	formAuto := widget.NewForm()
	formAuto.Append("match number: ", widget.NewLabel(matchNum))
	formAuto.Append("team: ", widget.NewLabel(strconv.Itoa(teamNum)))
	formAuto.Append("auto", widget.NewLabel(""))
	for _, s := range consts.GetGameTemp() {
		temp[s] = NewCounter()
		formAuto.Append(s, temp[s].baseCounter.object)
	}
	formTele := widget.NewForm()

	formTele.Append("teleop", widget.NewLabel(""))
	for _, s := range consts.GetGameTemp() {
		formTele.Append(s, temp[s].refCounter.object)
	}

	submit := widget.NewButton("Submit", func() {
		ret := make(map[string]int)
		for _, i := range consts.GetGameTemp() {
			ret[i] = *temp[i].baseCounter.counter + *temp[i].refCounter.counter
			fmt.Println(ret[i])
		}
		err := db.SubmitData(context.Background(), ret, matchNum, strconv.Itoa(teamNum))
		if err != nil {
			return
		}
	})

	next := widget.NewButton("Next", func() {
		base.RemoveAll()
		if len(otherPages)-1 > index {
			base.Add(otherPages[index+1])
			base.Refresh()

			return
		}
		base.Add(otherPages[0])
		base.Refresh()

	})
	prev := widget.NewButton("Prev", func() {
		base.RemoveAll()
		if 0 != index {
			base.Add(otherPages[index-1])
			base.Refresh()
			return
		}
		base.Add(otherPages[len(otherPages)-1])
		base.Refresh()

	})

	return container.NewVScroll(container.NewVBox(formAuto, formTele, container.NewHBox(prev, submit, next)))
}

func NewCreateScoutTab(b *dataBase.DB) *container.TabItem {
	numLists := make([]int, 0)
	for match := range b.Schedule {
		t, err := strconv.Atoi(match)
		if err != nil {
			fmt.Println(err)
		}
		numLists = append(numLists, t)
	}
	slices.Sort(numLists)
	pages := make([]*container.Scroll, len(numLists))
	base := container.NewMax()
	for i, num := range numLists {
		pages[i] = newCreateMatchTab(strconv.Itoa(num+1), b.Schedule[strconv.Itoa(num)], b, pages, i, base)
	}
	if len(pages) == 0 {
		return container.NewTabItem("scout tab", base)
	}
	base.Add(pages[0])
	base.Refresh()

	return container.NewTabItem("scout tab", base)
}

func CreateScoutTab(b *dataBase.DB) *container.TabItem {
	tabs := container.NewAppTabs()
	numLists := make([]int, 0)
	for match := range b.Schedule {
		t, err := strconv.Atoi(match)
		if err != nil {
			fmt.Println(err)
		}
		numLists = append(numLists, t)
	}
	slices.Sort(numLists)
	for _, num := range numLists {
		tabs.Append(createMatchTab(strconv.Itoa(num+1), b.Schedule[strconv.Itoa(num)], b))
	}
	tabs.SetTabLocation(container.TabLocationTop)

	return container.NewTabItem("scout tab", tabs)
}
