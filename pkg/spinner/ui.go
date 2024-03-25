package spinner

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type BubbleSpinner struct {
	spin spinner.Model
	tea  *tea.Program
}

func New() *BubbleSpinner {
	b := &BubbleSpinner{
		spin: spinner.New(),
	}
	b.tea = tea.NewProgram(b,
		tea.WithInput(os.Stdin),
		tea.WithOutput(os.Stdout),
	)
	return b
}

func (b *BubbleSpinner) Start() error {
	go func() {
		// Wait 2 seconds
		<-time.After(2 * time.Second)
		b.tea.Quit()
	}()
	if _, err := b.tea.Run(); err != nil {
		return err
	}
	fmt.Println("Done")
	return nil
}

func (b *BubbleSpinner) Init() tea.Cmd {
	return b.spin.Tick
}

func (b *BubbleSpinner) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m, c := b.spin.Update(msg)
	b.spin = m
	return b, c
}

func (b *BubbleSpinner) View() string {
	return fmt.Sprintf("Waiting 2 sec %s", b.spin.View())
}
