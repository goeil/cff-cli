package tui

// A simple program demonstrating the spinner component from the Bubbles
// component library.

import (
	"cff/api"
	"cff/types"
	"encoding/json"
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type errMsg error

var (
	lightColor = lipgloss.CompleteColor{TrueColor: "#888898", ANSI256: "160", ANSI: "9"}
	LightStyle = lipgloss.NewStyle().Foreground(lightColor)
)

type TravelLoader struct {
	spinner  spinner.Model
	quitting bool
	err      error
	status   string

	json     []byte
	horaires types.Horaires
}

func (m TravelLoader) GetHoraires() types.Horaires {
	return m.horaires
}

func InitialLoader() TravelLoader {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	return TravelLoader{spinner: s}
}

func (m *TravelLoader) Init() tea.Cmd {
	m.status = "Initialisation"
	return m.spinner.Tick
}

func (m *TravelLoader) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case FetchedApi:
		//fmt.Println(msg)

		m.quitting = true
		return m, tea.Quit

	case errorApi:
		//fmt.Println(msg)
		return m, tea.Quit

	case structBuilt:
		return m, tea.Quit

	case tea.QuitMsg:
		m.quitting = true
		return m, tea.Quit

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		default:
			return m, nil
		}

	case errMsg:
		m.err = msg
		return m, nil

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd

	default:
		return m, nil
	}
}

func (m TravelLoader) View() string {
	if m.err != nil {
		return m.err.Error()
	}
	icon := "✔️ "
	status := "Données bien récupérées."

	if !m.quitting {
		icon = m.spinner.View()
		status = m.status + "…"
	}
	str := fmt.Sprintf("\n   %s %s\n", icon, LightStyle.Render(status))

	return str

}

type FetchedApi string
type errorApi string
type structBuilt string

func (m *TravelLoader) GetClient(params api.RequestParameters) tea.Msg {
	m.status = "J'interroge l'API CFF"
	c := api.GetClient()
	myjson, err := c.Get(params)
	if err != nil {
		return errorApi("Error")
		//os.Exit(1)
	}
	m.json = myjson
	return FetchedApi("OK")
}

func (m *TravelLoader) Unmarshal() tea.Msg {
	var horaires types.Horaires
	m.status = "J'interprète les données"
	err := json.Unmarshal(m.json, &horaires)
	if err != nil {
		return errorApi("Error")
	}
	m.horaires = horaires
	return structBuilt("OK")
}

// loader.Update(tea.QuitMsg{})
//var horaires types.Horaires
