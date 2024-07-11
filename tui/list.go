package tui

import (
	"cff/types"
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	appStyle = lipgloss.NewStyle().Padding(1, 2)

	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#25A065")).
			Padding(0, 1)

	statusMessageStyle = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#04B575"}).
				Render
)

type item struct {
	title       string
	description string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.description }
func (i item) FilterValue() string { return i.title }

type listKeyMap struct {
	toggleSpinner    key.Binding
	toggleTitleBar   key.Binding
	toggleStatusBar  key.Binding
	togglePagination key.Binding
	toggleHelpMenu   key.Binding
	insertItem       key.Binding
}

func newListKeyMap() *listKeyMap {
	return &listKeyMap{
		insertItem: key.NewBinding(
			key.WithKeys("a"),
			key.WithHelp("a", "add item"),
		),
		toggleSpinner: key.NewBinding(
			key.WithKeys("s"),
			key.WithHelp("s", "toggle spinner"),
		),
		toggleTitleBar: key.NewBinding(
			key.WithKeys("T"),
			key.WithHelp("T", "toggle title"),
		),
		toggleStatusBar: key.NewBinding(
			key.WithKeys("S"),
			key.WithHelp("S", "toggle status"),
		),
		togglePagination: key.NewBinding(
			key.WithKeys("P"),
			key.WithHelp("P", "toggle pagination"),
		),
		toggleHelpMenu: key.NewBinding(
			key.WithKeys("H"),
			key.WithHelp("H", "toggle help"),
		),
	}
}

type TeaList struct {
	list         list.Model
	keys         *listKeyMap
	delegateKeys *delegateKeyMap
}

func NewList(horaires types.Horaires) TeaList {
	var (
		delegateKeys = newDelegateKeyMap()
		listKeys     = newListKeyMap()
	)

	// Make initial list of items
	items := make([]list.Item, len(horaires.Connections))
	for i, conn := range horaires.Connections {
		item := item{
			title:       conn.From.Station.ToString(),
			description: conn.From.Departure.Horaire(nil),
			//conn.To.Station.ToString(),
			//conn.To.Arrival.Horaire(),
			//description: user.Email + " [" + strconv.Itoa(user.Id) + "]",
			//description: user.Email + " [" + user.Id + "]",
		}
		items[i] = item
		//fmt.Println(user.ToString(nil))
	}

	// Setup list
	delegate := newItemDelegate(delegateKeys)
	groceryList := list.New(items, delegate, 0, 0)
	groceryList.Title = "Choisissez un utilisateur"
	groceryList.Styles.Title = titleStyle
	groceryList.AdditionalFullHelpKeys = func() []key.Binding {
		return []key.Binding{
			listKeys.toggleSpinner,
			listKeys.insertItem,
			listKeys.toggleTitleBar,
			listKeys.toggleStatusBar,
			listKeys.togglePagination,
			listKeys.toggleHelpMenu,
		}
	}
	fmt.Println(groceryList.SelectedItem())

	return TeaList{
		list:         groceryList,
		keys:         listKeys,
		delegateKeys: delegateKeys,
	}
}

func (l TeaList) Init() tea.Cmd {
	return tea.EnterAltScreen
}

func (l TeaList) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := appStyle.GetFrameSize()
		l.list.SetSize(msg.Width-h, msg.Height-v)

	case tea.KeyMsg:
		// Don't match any of the keys below if we're actively filtering.
		if l.list.FilterState() == list.Filtering {
			break
		}

		switch {
		case key.Matches(msg, l.keys.toggleSpinner):
			cmd := l.list.ToggleSpinner()
			return l, cmd

		case key.Matches(msg, l.keys.toggleTitleBar):
			v := !l.list.ShowTitle()
			l.list.SetShowTitle(v)
			l.list.SetShowFilter(v)
			l.list.SetFilteringEnabled(v)
			return l, nil

		case key.Matches(msg, l.keys.toggleStatusBar):
			l.list.SetShowStatusBar(!l.list.ShowStatusBar())
			return l, nil

		case key.Matches(msg, l.keys.togglePagination):
			l.list.SetShowPagination(!l.list.ShowPagination())
			return l, nil

		case key.Matches(msg, l.keys.toggleHelpMenu):
			l.list.SetShowHelp(!l.list.ShowHelp())
			return l, nil

		case key.Matches(msg, l.delegateKeys.choose):
			l.list.SetShowHelp(!l.list.ShowHelp())
			//Trace("quit")
			return l, tea.Quit

			/*case key.Matches(msg, m.keys.insertItem):
			m.delegateKeys.remove.SetEnabled(true)
			newItem := m.itemGenerator.next()
			insCmd := m.list.InsertItem(0, newItem)
			statusCmd := m.list.NewStatusMessage(statusMessageStyle("Added " + newItem.Title()))
			return m, tea.Batch(insCmd, statusCmd)*/
		}
	}

	// This will also call our delegate's update function.
	newListModel, cmd := l.list.Update(msg)
	l.list = newListModel
	cmds = append(cmds, cmd)

	return l, tea.Batch(cmds...)
}

func (l TeaList) View() string {
	return appStyle.Render(l.list.View())
}

func newItemDelegate(keys *delegateKeyMap) list.DefaultDelegate {
	d := list.NewDefaultDelegate()

	d.UpdateFunc = func(msg tea.Msg, m *list.Model) tea.Cmd {
		var title string

		if i, ok := m.SelectedItem().(item); ok {
			title = i.Title()
		} else {
			return nil
		}

		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, keys.choose):

				//tea.Quit()
				return m.NewStatusMessage(statusMessageStyle("Vous avez sélectionné " + title))

				/*case key.Matches(msg, keys.remove):
				index := m.Index()
				m.RemoveItem(index)
				if len(m.Items()) == 0 {
					keys.remove.SetEnabled(false)
				}
				return m.NewStatusMessage(statusMessageStyle("Deleted " + title))*/
			}
		}

		return nil
	}

	help := []key.Binding{keys.choose, keys.remove}

	d.ShortHelpFunc = func() []key.Binding {
		return help
	}

	d.FullHelpFunc = func() [][]key.Binding {
		return [][]key.Binding{help}
	}

	return d
}

type delegateKeyMap struct {
	choose key.Binding
	remove key.Binding
}

// Additional short help entries. This satisfies the help.KeyMap interface and
// is entirely optional.
func (d delegateKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		d.choose,
		d.remove,
	}
}

// Additional full help entries. This satisfies the help.KeyMap interface and
// is entirely optional.
func (d delegateKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			d.choose,
			d.remove,
		},
	}
}

func newDelegateKeyMap() *delegateKeyMap {
	return &delegateKeyMap{
		choose: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "choose"),
		),
		remove: key.NewBinding(
			key.WithKeys("x", "backspace"),
			key.WithHelp("x", "delete"),
		),
	}
}
