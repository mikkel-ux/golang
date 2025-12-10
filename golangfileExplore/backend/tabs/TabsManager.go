package tabs

type Manager struct{}

func NewTabsManager() *Manager {
	return &Manager{}
}

type Tab struct {
	ID      int
	Title   string
	History []string
	Index   int
	Files   []string
}

type TabsManager struct {
	Tabs map[string]*Tab
}
