package tui

import (
	"cff/types"
	"fmt"
	"strconv"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type TravelList struct {
	Travels        []Travel
	actions        []Action
	selectedAction int
}

type Connection struct {
	Station string
	Time    types.Datetime
}

func getIcon(cat string) string {
	switch cat {
	case "B":
		return "🚌" //🚍"

	case "IC", "TGV":
		return "🚄"

	case "S", "R", "RE", "IR", "TER":
		return "🚆"
	}
	return "🚆"
}
func getLine(cat string, width int) string {
	gare := "○"
	switch cat {
	case "B":
		return LineBusStyle.Render(gare + strings.Repeat("┉", width) + gare)

	case "IC", "TGV", "ICE":
		return LineICStyle.Render(gare + strings.Repeat("═", width) + gare)

	case "S", "R", "RE", "IR", "TER":
		return LineTrainStyle.Render(gare + strings.Repeat("─", width) + gare)
	}
	return cat

}
func (c Connection) getName(widthStation int) string {

	if len(c.Station) <= widthStation {
		return c.Station
	} else {

		return string([]rune(c.Station)[0:widthStation]) // sinon problème d'index, Il faut donc transformer en rune[]

	}
}

type Section struct {
	Departure Connection
	Arrival   Connection
	Category  string
}

type Travel struct {
	Departure Connection
	Arrival   Connection
	Sections  []Section
	Duration  time.Duration
}

func (t Travel) details() string {
	var time string
	if t.Duration.Minutes() > 60 {

		hours := int(t.Duration.Hours())
		min := int(t.Duration.Minutes()) - (hours * 60)
		time = fmt.Sprintf("%d h %02d",
			hours,
			min)
	} else {
		time = fmt.Sprintf("%d min",
			int(t.Duration.Minutes()))
	}
	var changements string
	var attentes string
	for i, trans := range t.Sections {
		changements += getIcon(trans.Category)
		if i >= 1 {
			diff := t.Sections[i].Departure.Time.Sub(t.Sections[i-1].Arrival.Time)
			attentes += fmt.Sprintf("%d' ", int(diff.Minutes()))
		}
	}
	if len(attentes) > 0 {
		attentes = attentes[:len(attentes)-1]
	}

	switch len(t.Sections) {

	case 1:
		break
	case 2:
		//diff1 := t.Sections[1].Departure.Time.Sub(t.Sections[0].Arrival.Time)
		changements += fmt.Sprintf(" 1 changement (%v)", attentes)

	default:
		changements += fmt.Sprintf(" %d changements (%v))",
			len(t.Sections)-1,
			attentes)

	}
	return fmt.Sprintf("⏳ %v  | %v",
		time,
		changements)
}

func (t Travel) ToString() string {

	s := LightStyle.Render(t.details())
	switch len(t.Sections) {
	case 1:
		//fmt.Println(t)

		section := t.Sections[0]
		width := normalWidthStation
		s += fmt.Sprintf("\n%v %v %v %v %v\n",
			StationStyle.Width(width+2).Render(section.Departure.getName(width)),
			HoraireStyle.Render(section.Departure.Time.Horaire(&types.HoraireOptions{WithDay: false})),
			getLine(section.Category, 33),
			HoraireStyle.Render(section.Arrival.Time.Horaire(&types.HoraireOptions{WithDay: false})),
			StationStyle.Width(width+2).Render(section.Arrival.getName(width)),
		)

	case 2:
		section1 := t.Sections[0]
		section2 := t.Sections[1]
		width := shortWidthStation

		s += fmt.Sprintf("\n%v %v %v %v %v %v %v %v %v\n",
			StationStyle.Width(width+2).Render(section1.Departure.getName(width)),
			HoraireStyle.Render(section1.Departure.Time.Horaire(&types.HoraireOptions{WithDay: false})),
			getLine(section1.Category, 7),
			HoraireStyle.Render(section1.Arrival.Time.Horaire(&types.HoraireOptions{WithDay: false})),
			InterStationStyle.Width(width+2).Render(section1.Arrival.getName(width)),
			HoraireStyle.Render(section2.Departure.Time.Horaire(&types.HoraireOptions{WithDay: false})),
			getLine(section2.Category, 7),
			HoraireStyle.Render(section2.Arrival.Time.Horaire(&types.HoraireOptions{WithDay: false})),
			StationStyle.Width(width+2).Render(section2.Arrival.getName(width)),
		)

	case 3:
		section1 := t.Sections[0]
		section2 := t.Sections[1]
		section3 := t.Sections[2]
		width := veryShortWidthStation

		s += fmt.Sprintf("\n%v %v%v%v %v %v%v%v %v %v%v%v %v\n",

			StationStyle.Width(width+2).Render(section1.Departure.getName(width)),
			HoraireStyle.Render(section1.Departure.Time.Horaire(&types.HoraireOptions{WithDay: false})),
			getLine(section1.Category, 3),
			HoraireStyle.Render(section1.Arrival.Time.Horaire(&types.HoraireOptions{WithDay: false})),
			InterStationStyle.Width(width+2).Render(section1.Arrival.getName(width)),
			HoraireStyle.Render(section2.Departure.Time.Horaire(&types.HoraireOptions{WithDay: false})),
			getLine(section2.Category, 3),
			HoraireStyle.Render(section2.Arrival.Time.Horaire(&types.HoraireOptions{WithDay: false})),
			InterStationStyle.Width(width+2).Render(section2.Arrival.getName(width)),
			HoraireStyle.Render(section3.Departure.Time.Horaire(&types.HoraireOptions{WithDay: false})),
			getLine(section3.Category, 3),
			HoraireStyle.Render(section3.Arrival.Time.Horaire(&types.HoraireOptions{WithDay: false})),
			StationStyle.Width(width+2).Render(section3.Arrival.getName(width)),
		)

	case 4:
		section1 := t.Sections[0]
		section2 := t.Sections[1]
		section3 := t.Sections[2]
		section4 := t.Sections[3]
		width := veryShortWidthStation

		s += fmt.Sprintf("\n%v %v %v %v %v %v %v %v %v %v %v\n",

			StationStyle.Width(width+2).Render(section1.Departure.getName(width)),
			HoraireStyle.Render(section1.Departure.Time.Horaire(&types.HoraireOptions{WithDay: false})),
			getLine(section1.Category, 4),
			InterStationStyle.Width(width+2).Render(section1.Arrival.getName(width)),
			getLine(section2.Category, 4),
			InterStationStyle.Width(width+2).Render(section2.Arrival.getName(width)),
			getLine(section3.Category, 4),
			InterStationStyle.Width(width+2).Render(section3.Arrival.getName(width)),
			getLine(section4.Category, 4),
			HoraireStyle.Render(section4.Arrival.Time.Horaire(&types.HoraireOptions{WithDay: false})),
			StationStyle.Width(width+2).Render(section4.Arrival.getName(width)),
		)

	case 5:
		section1 := t.Sections[0]
		section2 := t.Sections[1]
		section3 := t.Sections[2]
		section4 := t.Sections[3]
		section5 := t.Sections[4]
		width := veryShortWidthStation

		s += fmt.Sprintf("\n%v %v %v%v%v%v%v%v%v%v%v %v %v\n",

			StationStyle.Width(width+2).Render(section1.Departure.getName(width)),
			HoraireStyle.Render(section1.Departure.Time.Horaire(&types.HoraireOptions{WithDay: false})),
			getLine(section1.Category, 2),
			InterStationStyle.Width(width+2).Render(section1.Arrival.getName(width)),
			getLine(section2.Category, 3),
			InterStationStyle.Width(width+2).Render(section2.Arrival.getName(width)),
			getLine(section3.Category, 3),
			InterStationStyle.Width(width+2).Render(section3.Arrival.getName(width)),
			getLine(section4.Category, 3),
			InterStationStyle.Width(width+2).Render(section4.Arrival.getName(width)),
			getLine(section5.Category, 2),
			HoraireStyle.Render(section5.Arrival.Time.Horaire(&types.HoraireOptions{WithDay: false})),
			StationStyle.Width(width+2).Render(section5.Arrival.getName(width)),
		)

	default:
		//fmt.Println(t)
		width := shortWidthStation
		s += fmt.Sprintf("%v %v %v %v %v\n",
			StationStyle.Width(width+2).Render(t.Departure.getName(width)),
			HoraireStyle.Render(t.Departure.Time.Horaire(&types.HoraireOptions{WithDay: false})),
			LineTrainStyle.Render("○───── "+strconv.Itoa(len(t.Sections))+" ─────○"),
			HoraireStyle.Render(t.Arrival.Time.Horaire(&types.HoraireOptions{WithDay: false})),
			StationStyle.Width(width+2).Render(t.Arrival.getName(width)),
		)
	}
	return s
}

func (tl TravelList) ToString() string {

	s := ""

	// Cherche le premier trajet pour avoir le nom complet des gares trouvées
	for i, h := range tl.Travels {

		if i == 0 {
			width := 22
			s += LargeTitleStyle.Render(fmt.Sprintf("\nTrajet %v%v%v\n",
				StationStyle.Bold(true).Width(width+2).Render(h.Departure.getName(width)),
				TitleStyle.Render("  ⇉  "),
				StationStyle.Bold(true).Width(width+2).Render(h.Arrival.getName(width)),
			))
		}

		s += "\n" + h.ToString()
	}

	return s
}

type Action struct {
	Label    string
	Shortcut string
	ObjectId int
}

var (
	normalWidthStation    = 13
	shortWidthStation     = 9
	veryShortWidthStation = 5
	rougeCffColor         = lipgloss.CompleteColor{TrueColor: "#EB0000", ANSI256: "160", ANSI: "9"}
	bleuCffColor          = lipgloss.CompleteColor{TrueColor: "#2D327D", ANSI256: "20", ANSI: "4"}
	bleuClairCffColor     = lipgloss.CompleteColor{TrueColor: "#005f5f", ANSI256: "37", ANSI: "14"}
	violetBusColor        = lipgloss.CompleteColor{TrueColor: "#875fff", ANSI256: "99", ANSI: "5"}

	HoraireStyle = lipgloss.NewStyle().Bold(true)
	StationStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("white")).
			Background(bleuCffColor).
			PaddingLeft(1).PaddingRight(1).
			Bold(false)

	InterStationStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("white")).
				Background(bleuClairCffColor).
				PaddingLeft(1).PaddingRight(1)

	LineTrainStyle = lipgloss.NewStyle().Foreground(rougeCffColor).Bold(true)
	LineICStyle    = lipgloss.NewStyle().Foreground(rougeCffColor).Bold(true)
	LineBusStyle   = lipgloss.NewStyle().Foreground(violetBusColor).Bold(true)

	TitleStyle      = lipgloss.NewStyle().Background(rougeCffColor).Bold(true)
	LargeTitleStyle = TitleStyle.Copy().Margin(1).PaddingLeft(2).PaddingRight(2)
)

func New(jsonHoraires types.Horaires) *TravelList {

	travels := make([]Travel, len(jsonHoraires.Connections))
	// loop over different travels
	for i, conn := range jsonHoraires.Connections {

		//sections := make([]Section, len(conn.Sections))
		var sections []Section
		// loop over sections in travels
		for _, trans := range conn.Sections {
			category := trans.Journey.Category
			// B = bus, R = régio, IC = IC, rien = pied
			if category != "" {

				section := Section{
					Departure: Connection{Station: trans.Departure.Station.Name,
						Time: trans.Departure.Departure},
					Arrival: Connection{Station: trans.Arrival.Station.Name,
						Time: trans.Arrival.Arrival},
					Category: trans.Journey.Category,
				}

				sections = append(sections, section)
				//sections[j] = section
			}
		}
		//fmt.Println("j'ai créé un trajet avec " + strconv.Itoa(len(sections)))
		diff := conn.To.Arrival.Sub(conn.From.Departure)
		travels[i] = Travel{Sections: sections,
			Departure: Connection{Station: conn.From.Station.Name,
				Time: conn.From.Departure},
			Arrival: Connection{Station: conn.To.Station.Name,
				Time: conn.To.Arrival},
			Duration: diff,
		}
	}

	return &TravelList{
		Travels: travels,
	}
}
func (m *TravelList) Init() tea.Cmd {

	var actions []Action
	actions = append(actions, Action{Label: "suivants", Shortcut: " ", ObjectId: -1})
	m.actions = actions
	return nil
}

func (m TravelList) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg: // startup and resize

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return &m, tea.Quit
		}
	}

	return &m, nil
}
func (m TravelList) View() string {
	var b strings.Builder

	fmt.Fprintf(&b, "%v ",
		m.ToString(),
	)
	fmt.Fprintf(&b, "%s ",
		"Actions : ",
	)

	// Iterate over our choices
	for i, choice := range m.actions {

		//fmt.Printf("%v-%v ", i, choice)
		cursor := " " // no cursor

		if m.selectedAction == i {
			cursor = "✔️" // cursor!
		}

		// Render the row
		fmt.Fprintf(&b, " %v %v [%v] ",
			cursor,
			choice.Label,
			choice.Shortcut,
		)
	}

	// Send the UI for rendering
	return b.String()
}
