package cmd

import (
	"fmt"
	"github.com/SachaCR/tfb/internals/music"
	"github.com/SachaCR/tfb/internals/neck"
	"github.com/SachaCR/tfb/internals/render"
	"github.com/SachaCR/tfb/internals/ui/modeList"
	key "github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"os"
	"sort"
	"strings"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

func init() {
	rootCmd.AddCommand(chordCmd)
	rootCmd.AddCommand(scaleCmd)

	rootCmd.PersistentFlags().StringVarP(&instrument, "inst", "i", "G", "Set the instrument type, G for guitar, B for Bass, U for Ukulele")
	rootCmd.PersistentFlags().StringVarP(&root, "root", "r", "C", "Set the root of your chord")
}

var rootCmd = &cobra.Command{
	Use:   "tfb",
	Short: "tfb displays musical mode rendering a guitar neck diagram",
	Long: `tfb displays musical mode on the guitar neck. By default tfb ask you to select the mode in the list but you can also pass it in argument directly.
	It also support bass and ukulele and potentially any instruments with strings and frets.`,
	Args: cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {

		var modeName string

		if len(args) > 0 {
			modeName = args[0]
		}

		// If no mode name is passed in arguments we allow the user to search in a predefined list
		if modeName == "" {

			list := list.New(buildChoices(), modeList.ItemDelegate{}, 0, 0)

			list.SetFilteringEnabled(true)

			list.Title = "Modes"

			list.KeyMap.CursorUp.SetHelp("↑/k", "Up")
			list.KeyMap.CursorDown.SetHelp("↓/j", "Down")
			list.KeyMap.NextPage.SetHelp("←/→ h/l ", "← Root →")

			list.KeyMap.PrevPage.SetEnabled(false)
			list.KeyMap.GoToEnd.SetEnabled(false)
			list.KeyMap.GoToStart.SetEnabled(false)

			instrumentBindings :=
				key.NewBinding(
					key.WithKeys("g"),
					key.WithKeys("u"),
					key.WithKeys("b"),
					key.WithHelp("u/g/b", "Change instrument"),
				)

			list.AdditionalShortHelpKeys = func() []key.Binding {
				return []key.Binding{
					list.KeyMap.NextPage,
					instrumentBindings,
				}
			}

			list.AdditionalFullHelpKeys = func() []key.Binding {
				return []key.Binding{
					instrumentBindings,
				}
			}

			model := &model{
				list:       list,
				root:       music.NoteFromString[root],
				instrument: instrument,
			}

			p := tea.NewProgram(model, tea.WithAltScreen())

			if _, err := p.Run(); err != nil {
				fmt.Printf("Alas, there's been an error: %v", err)
				os.Exit(1)
			}

			os.Exit(0)
		}

		mode := music.FindMode(modeName)

		if mode == nil {
			fmt.Println("Unknow Mode")
			return
		}

		scale := mode.ToScale(music.NoteFromString[root])
		fmt.Println(render.RenderScale(neck.New(instrument), scale, 1, 12, "circle"))

		fmt.Println("Intervals: ", strings.Join(mode.Intervals(), " "))
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

type model struct {
	cursor     int
	list       list.Model
	choice     string
	root       music.Note
	instrument string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "h", "left":
			m.root = m.root.AddTone(-1)
			return m, nil

		case "l", "right":
			m.root = m.root.AddTone(1)
			return m, nil

		case "U", "u":
			m.instrument = "U"
			return m, nil

		case "G", "g":
			m.instrument = "G"
			return m, nil

		case "B", "b":
			m.instrument = "B"
			return m, nil
		}

	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v-10)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)

	return m, cmd
}

func (m model) View() string {
	var modeName = "ionian"

	i, ok := m.list.SelectedItem().(modeList.Item)

	if ok {
		modeName = string(i.Title())
	}

	mode := music.FindMode(modeName)
	scale := mode.ToScale(music.NoteFromString[m.root.String()])
	neck := neck.New(m.instrument)
	diagramString := render.RenderScale(neck, scale, 1, 12, "circle")

	return "Instrument: " + neck.Instrument() + "\n" + diagramString + "\nIntervals: " + strings.Join(mode.Intervals(), " ") + "\n" + m.list.View()
}

func buildChoices() []list.Item {
	var choices []list.Item

	for _, mode := range music.ModeMap {
		choices = append(choices, modeList.NewItem(mode.Name()))
	}

	// Order choice alphabeticaly
	sort.Slice(choices, func(i, j int) bool {
		return choices[i].FilterValue() < choices[j].FilterValue()
	})

	return choices
}
