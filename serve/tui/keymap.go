package tui

import "charm.land/bubbles/v2/key"

// KeyMap defines keyboard shortcuts for the serve TUI.
type KeyMap struct {
	Quit           key.Binding
	Rebuild        key.Binding
	CycleVerbosity key.Binding
	OpenBrowser    key.Binding
	Clear          key.Binding
	Help           key.Binding
	ScrollUp       key.Binding
	ScrollDown     key.Binding
	PageUp         key.Binding
	PageDown       key.Binding
	Home           key.Binding
	End            key.Binding
}

func DefaultKeyMap() KeyMap {
	return KeyMap{
		Quit:           key.NewBinding(key.WithKeys("q", "Q"), key.WithHelp("q", "quit")),
		Rebuild:        key.NewBinding(key.WithKeys("m", "M"), key.WithHelp("m", "rebuild manifest")),
		CycleVerbosity: key.NewBinding(key.WithKeys("v", "V"), key.WithHelp("v", "cycle log level")),
		OpenBrowser:    key.NewBinding(key.WithKeys("o", "O"), key.WithHelp("o", "open browser")),
		Clear:          key.NewBinding(key.WithKeys("c", "C"), key.WithHelp("c", "clear logs")),
		Help:           key.NewBinding(key.WithKeys("?"), key.WithHelp("?", "toggle help")),
		ScrollUp:       key.NewBinding(key.WithKeys("k", "up"), key.WithHelp("k/↑", "scroll up")),
		ScrollDown:     key.NewBinding(key.WithKeys("j", "down"), key.WithHelp("j/↓", "scroll down")),
		PageUp:         key.NewBinding(key.WithKeys("pgup"), key.WithHelp("pgup", "page up")),
		PageDown:       key.NewBinding(key.WithKeys("pgdown"), key.WithHelp("pgdn", "page down")),
		Home:           key.NewBinding(key.WithKeys("home", "g"), key.WithHelp("home", "top")),
		End:            key.NewBinding(key.WithKeys("end", "G"), key.WithHelp("end", "bottom")),
	}
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit, k.Rebuild, k.CycleVerbosity}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.ScrollUp, k.ScrollDown, k.PageUp, k.PageDown, k.Home, k.End},
		{k.Quit, k.Rebuild, k.CycleVerbosity, k.OpenBrowser, k.Clear, k.Help},
	}
}
