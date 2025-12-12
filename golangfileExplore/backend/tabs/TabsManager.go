package tabs

import (
	"github.com/google/uuid"
)

type Manager struct{}

func NewTabsManager() *Manager {
	return &Manager{}
}

type Tab struct {
	ID      string
	Title   string
	History []string
	Index   int
	Files   []string
}

func (t *Manager) NewTab() *Tab {
	/* open dir and find files */
	u := uuid.New()
	tab := &Tab{
		ID:    u.String(),
		Title: "New Tab",
		History: []string{
			"home",
		},
		Index: 0,
		Files: []string{},
	}
	return tab
}

/* func (t *Manager) CloseTab(tabId string) {
	delete(TabsInstance.Tabs, tabId)
}

func (t *Manager) SelectTab(tabId string) (*Tab, bool) {
	tab, exists := TabsInstance.Tabs[tabId]
	println(tab.Title)
	return tab, exists
} */

func (t *Manager) HistoryBack(tab *Tab) {
	if tab.Index > 0 {
		tab.Index--
	}
}

func (t *Manager) HistoryForward(tab *Tab) {
	if tab.Index < len(tab.History)-1 {
		tab.Index++
	}
}
