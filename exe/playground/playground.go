package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/pelletier/go-toml/v2"
	"github.com/yiffyi/autope"
	"github.com/yiffyi/autope/wmi"
)

func tryRichOutput() {
	var style1 = lipgloss.NewStyle().
		Bold(true).
		Italic(true).
		Faint(true).
		Blink(true).
		Strikethrough(true).
		Underline(true).
		Reverse(true)

	fmt.Println(style1.Render("Styled Text"))

	var style2 = lipgloss.NewStyle().
		Foreground(lipgloss.ANSIColor(5))
	fmt.Println(style2.Render("ANSI 5"))

	var style3 = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("63"))
	fmt.Println(style3.Render("I have border"))
}

func tryWMI() {
	wmi.CoInitialize()
	locator, err := wmi.NewSWbemLocator()
	if err != nil {
		fmt.Println("NewSWbemLocator:", err)
		return
	}

	service, err := locator.ConnectServerDefault()
	if err != nil {
		fmt.Println("ConnectServerDefault:", err)
		return
	}

	volumes, err := service.ExecQuery("SELECT * FROM Win32_Volume")
	if err != nil {
		fmt.Println("ExecQuery:", err)
		return
	}

	cnt, err := volumes.Count()
	fmt.Println("Count:", cnt, err)

	s, err := volumes.ToSlice()
	fmt.Println("ToSlice:", s, err)

	v := s[0]
	fmt.Println("volumes[0].Name", v.PropertyMustGetValue("Name").(string), err)
}

func main() {

	t, err := toml.Marshal(autope.Playbook{})
	fmt.Println(string(t), err)

	// tryRichOutput()
	// trySpinner()
	tryWMI()
}
