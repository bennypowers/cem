package tui

import (
	"fmt"
	"strings"

	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"

	"bennypowers.dev/cem/internal/logging"
	itui "bennypowers.dev/cem/internal/tui"
)

const (
	maxTermLines    = 200
	separatorHeight = 1
	statusHeight    = 1
)

// Callbacks holds function pointers for side effects the Model dispatches.
// Avoids importing serve package (which would create a cycle).
// All fields are optional; nil callbacks are safely ignored.
type Callbacks struct {
	InitServer      func() tea.Msg
	RebuildManifest func() (int, error)
	OpenBrowser     func() error
	CycleVerbosity  func() string
	Shutdown        func()
}

// Model is the bubbletea model for the serve command TUI.
type Model struct {
	viewport  viewport.Model
	help      help.Model
	keyMap    KeyMap
	callbacks Callbacks

	logMsgs    []LogMsg
	lines      []string
	status     string
	width      int
	height     int
	showHelp   bool
	autoScroll bool
	ready      bool
	pending    []tea.Msg
}

// NewModel creates a serve TUI model with the given callbacks.
// pending contains messages buffered before the program started.
func NewModel(callbacks Callbacks, pending []tea.Msg) Model {
	vp := viewport.New()
	vp.KeyMap = viewport.KeyMap{}
	vp.SoftWrap = true
	vp.MouseWheelEnabled = true

	h := help.New()

	return Model{
		viewport:   vp,
		help:       h,
		keyMap:     DefaultKeyMap(),
		callbacks:  callbacks,
		lines:      make([]string, 0, maxTermLines),
		status:     "Initializing...",
		autoScroll: true,
		pending:    pending,
	}
}

// SetPending adds buffered messages to drain on Init. Must be called before Run.
func (m *Model) SetPending(msgs []tea.Msg) {
	m.pending = append(m.pending, msgs...)
}

func (m Model) Init() tea.Cmd {
	var cmds []tea.Cmd
	for _, msg := range m.pending {
		cmds = append(cmds, func() tea.Msg { return msg })
	}
	if m.callbacks.InitServer != nil {
		init := m.callbacks.InitServer
		cmds = append(cmds, func() tea.Msg { return init() })
	}
	return tea.Batch(cmds...)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.recalcLayout()
		m.reflowLines()
		return m, nil

	case LogMsg:
		m.appendLog(msg)
		return m, nil

	case StatusMsg:
		m.status = string(msg)
		return m, nil

	case ClearMsg:
		m.logMsgs = m.logMsgs[:0]
		m.lines = m.lines[:0]
		m.viewport.SetContent("")
		return m, nil

	case ServerDoneMsg:
		return m, tea.Quit

	case ServerReadyMsg:
		m.ready = true
		if msg.WatchDone != nil {
			done := msg.WatchDone
			return m, func() tea.Msg {
				<-done
				return ServerDoneMsg{}
			}
		}
		return m, nil

	case ServerInitErrorMsg:
		m.status = "Init failed: " + msg.Err.Error()
		return m, tea.Quit

	case RebuildResultMsg:
		return m, nil

	case OpenBrowserResultMsg:
		return m, nil

	case VerbosityCycledMsg:
		m.reflowLines()
		return m, nil

	case tea.KeyPressMsg:
		return m.handleKey(msg)

	case tea.InterruptMsg:
		return m, tea.Sequence(m.shutdownCmd(), tea.Quit)
	}

	return m, nil
}

func (m Model) handleKey(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	km := m.keyMap

	switch {
	case key.Matches(msg, km.Quit):
		return m, tea.Sequence(m.shutdownCmd(), tea.Quit)

	case key.Matches(msg, km.Rebuild):
		return m, m.rebuildCmd()

	case key.Matches(msg, km.CycleVerbosity):
		return m, m.cycleVerbosityCmd()

	case key.Matches(msg, km.OpenBrowser):
		return m, m.openBrowserCmd()

	case key.Matches(msg, km.Clear):
		return m, func() tea.Msg { return ClearMsg{} }

	case key.Matches(msg, km.Help):
		m.showHelp = !m.showHelp
		m.help.ShowAll = m.showHelp
		m.recalcLayout()
		return m, nil

	case key.Matches(msg, km.ScrollDown):
		m.viewport.ScrollDown(1)
		m.autoScroll = m.viewport.AtBottom()
		return m, nil

	case key.Matches(msg, km.ScrollUp):
		m.viewport.ScrollUp(1)
		m.autoScroll = false
		return m, nil

	case key.Matches(msg, km.PageDown):
		m.viewport.ScrollDown(m.viewport.Height())
		m.autoScroll = m.viewport.AtBottom()
		return m, nil

	case key.Matches(msg, km.PageUp):
		m.viewport.ScrollUp(m.viewport.Height())
		m.autoScroll = false
		return m, nil

	case key.Matches(msg, km.Home):
		m.viewport.GotoTop()
		m.autoScroll = false
		return m, nil

	case key.Matches(msg, km.End):
		m.viewport.GotoBottom()
		m.autoScroll = true
		return m, nil
	}

	return m, nil
}

func (m Model) View() tea.View {
	if m.width == 0 || m.height == 0 {
		return tea.NewView("")
	}

	var b strings.Builder

	b.WriteString(m.viewport.View())
	b.WriteByte('\n')

	b.WriteString(itui.SeparatorStyle.Render(strings.Repeat("─", m.width)))
	b.WriteByte('\n')

	b.WriteString(itui.StatusBulletStyle.Render("● "))
	b.WriteString(m.status)
	b.WriteByte('\n')

	b.WriteString(m.help.View(m.keyMap))

	v := tea.NewView(b.String())
	v.AltScreen = true
	return v
}

func (m *Model) appendLog(msg LogMsg) {
	m.logMsgs = append(m.logMsgs, msg)
	if len(m.logMsgs) > maxTermLines {
		m.logMsgs = m.logMsgs[len(m.logMsgs)-maxTermLines:]
	}
	if logging.ShouldDisplay(msg.Entry.Type) {
		m.lines = append(m.lines, formatLogLine(msg, m.width))
		if len(m.lines) > maxTermLines {
			m.lines = m.lines[len(m.lines)-maxTermLines:]
		}
		m.viewport.SetContent(strings.Join(m.lines, "\n"))
		if m.autoScroll {
			m.viewport.GotoBottom()
		}
	}
}

func (m *Model) reflowLines() {
	m.lines = m.lines[:0]
	for _, msg := range m.logMsgs {
		if logging.ShouldDisplay(msg.Entry.Type) {
			m.lines = append(m.lines, formatLogLine(msg, m.width))
		}
	}
	m.viewport.SetContent(strings.Join(m.lines, "\n"))
}


func (m *Model) recalcLayout() {
	helpHeight := 1
	if m.showHelp {
		helpHeight = 4
	}
	vpHeight := max(m.height-separatorHeight-statusHeight-helpHeight, 1)
	m.viewport.SetWidth(m.width)
	m.viewport.SetHeight(vpHeight)
	m.help.SetWidth(m.width)
}

func (m Model) rebuildCmd() tea.Cmd {
	if m.callbacks.RebuildManifest == nil {
		return nil
	}
	rebuild := m.callbacks.RebuildManifest
	return func() tea.Msg {
		size, err := rebuild()
		return RebuildResultMsg{Size: size, Err: err}
	}
}

func (m Model) openBrowserCmd() tea.Cmd {
	if m.callbacks.OpenBrowser == nil {
		return nil
	}
	open := m.callbacks.OpenBrowser
	return func() tea.Msg {
		return OpenBrowserResultMsg{Err: open()}
	}
}

func (m Model) cycleVerbosityCmd() tea.Cmd {
	if m.callbacks.CycleVerbosity == nil {
		return nil
	}
	cycle := m.callbacks.CycleVerbosity
	return func() tea.Msg {
		label := cycle()
		return VerbosityCycledMsg{Label: label}
	}
}

func (m Model) shutdownCmd() tea.Cmd {
	if m.callbacks.Shutdown == nil {
		return nil
	}
	shutdown := m.callbacks.Shutdown
	return func() tea.Msg {
		shutdown()
		return ShutdownMsg{}
	}
}

func styleMessage(levelType, message string) string {
	switch levelType {
	case "warning":
		return itui.WarnStyle.Render(message)
	case "error":
		return itui.ErrorStyle.Render(message)
	case "debug":
		return itui.DebugStyle.Render(message)
	case "trace":
		return itui.TraceStyle.Render(message)
	default:
		return message
	}
}

func levelPrefix(levelType string) string {
	switch levelType {
	case "info":
		return itui.InfoStyle.Render("INFO ")
	case "warning":
		return itui.WarnStyle.Render("WARN ")
	case "error":
		return itui.ErrorStyle.Render("ERROR")
	case "success":
		return itui.SuccessStyle.Render(" OK  ")
	case "debug":
		return itui.DebugStyle.Render("DEBUG")
	case "trace":
		return itui.TraceStyle.Render("TRACE")
	default:
		return fmt.Sprintf("%-5s", levelType)
	}
}

const (
	prefixCols    = 7
	timestampCols = 10
)

func formatLogLine(msg LogMsg, width int) string {
	if width == 0 {
		width = 80
	}

	timestamp := ""
	if len(msg.Entry.Date) >= 19 {
		timestamp = msg.Entry.Date[11:19]
	}

	prefix := levelPrefix(msg.Entry.Type)
	lines := strings.Split(msg.Entry.Message, "\n")
	indent := strings.Repeat(" ", prefixCols)

	var b strings.Builder
	for i, line := range lines {
		if i == 0 {
			fmt.Fprintf(&b, " %s %s", prefix, styleMessage(msg.Entry.Type, line))
		} else {
			b.WriteString(indent)
			b.WriteString(styleMessage(msg.Entry.Type, line))
		}

		if i == 0 {
			visualLen := prefixCols + lipgloss.Width(line) + timestampCols
			padding := max(width-visualLen, 1)
			b.WriteString(strings.Repeat(" ", padding))
			b.WriteString(itui.TimestampStyle.Render(timestamp))
		}
		if i < len(lines)-1 {
			b.WriteByte('\n')
		}
	}

	if d := msg.Entry.Data; d != nil && d.Kind == "durations" {
		b.WriteByte('\n')
		durationData := make([]itui.DurationData, len(d.Durations))
		for i, de := range d.Durations {
			durationData[i] = itui.DurationData{
				Name:     de.Name,
				Duration: de.Duration,
				Percent:  de.Percent,
			}
		}
		itui.FormatDurationBars(&b, durationData, indent, timestamp, width)
	}

	return b.String()
}


// Logs returns the formatted lines for testing.
func (m Model) Lines() []string {
	return m.lines
}

// Status returns the current status line for testing.
func (m Model) Status() string {
	return m.status
}

// AutoScroll returns whether auto-scroll is enabled for testing.
func (m Model) AutoScroll() bool {
	return m.autoScroll
}
