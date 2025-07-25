package selector

import (
	"fmt"
	"remx/internal/ui"
	"remx/pkg/utility"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// RGB defines a simple RGB structure
type RGB struct {
	R, G, B int
}

// Options for customizing colors
type Options struct {
	ItemSelectedColour RGB
	ItemFocusedColour  RGB
	SelectorIcon       string
}

type model struct {
	title    string
	items    []string
	cursor   int
	selected bool
	options  Options
	choice   string
	styles   styleSet
}

type styleSet struct {
	focused  lipgloss.Style
	selected lipgloss.Style
	normal   lipgloss.Style
}

func Select(title string, items []string, opts Options) (string, error) {
	initialModel := model{
		title:   title,
		items:   items,
		options: opts,
		styles: styleSet{
			focused: lipgloss.NewStyle().Foreground(lipgloss.Color(utility.RGBToHex(opts.ItemFocusedColour.R, opts.ItemFocusedColour.G, opts.ItemFocusedColour.B))),
			selected: lipgloss.NewStyle().
				Foreground(lipgloss.Color(utility.RGBToHex(opts.ItemSelectedColour.R, opts.ItemSelectedColour.G, opts.ItemSelectedColour.B))).
				Bold(true),
			normal: lipgloss.NewStyle().Foreground(lipgloss.Color(ui.Colours["item:normal"].Hex)),
		},
	}

	p := tea.NewProgram(initialModel)
	finalModel, err := p.Run()
	if err != nil {
		return "", err
	}

	m := finalModel.(model)
	return m.choice, nil
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.items)-1 {
				m.cursor++
			}

		case "enter":
			m.selected = true
			m.choice = m.items[m.cursor]
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	if m.selected {
		return fmt.Sprintf("%s %s\n", m.title, m.styles.selected.Render(m.choice))
	}

	s := m.title + "\n"
	for i, item := range m.items {
		cursor := "  "
		if m.cursor == i {
			cursor = m.options.SelectorIcon
		}
		style := m.styles.normal
		if m.cursor == i {
			style = m.styles.focused
		}
		s += fmt.Sprintf("%s%s\n", cursor, style.Render(item))
	}
	return s
}
