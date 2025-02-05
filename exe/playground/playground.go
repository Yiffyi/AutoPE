package main

import (
	"fmt"
	"path/filepath"

	"github.com/charmbracelet/lipgloss"
	"github.com/rs/zerolog/log"
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

func tryWMI() error {

	wmi.CoInitialize()
	defer wmi.CoUninitialize()

	locator, err := wmi.NewSWbemLocator()
	if err != nil {
		log.Error().Err(err).Msg("could not initialize WMI service locator")
		return err
		// return nil, err
	}

	svc, err := locator.ConnectServerDefault()
	if err != nil {
		log.Error().Err(err).Msg("could not connect to WMI service")
		return err
	}

	volumes, err := svc.ExecQuery("SELECT * FROM Win32_Volume")
	if err != nil {
		fmt.Println("ExecQuery:", err)
		return err
	}

	cnt, err := volumes.Count()
	fmt.Println("Count:", cnt, err)

	s, err := volumes.ToSlice()
	fmt.Println("ToSlice:", s, err)

	v := s[0]
	fmt.Println("volumes[0].Name", v.PropertyMustGetValue("Name").(string), err)
	fmt.Println(autope.SearchOfflineWindows(svc))
	return nil
}

func main() {

	autope.SetupDefaultLogger()
	// t, err := toml.Marshal(autope.Playbook{})
	// fmt.Println(string(t), err)

	// tryRichOutput()
	// trySpinner()
	// tryWMI()

	fmt.Println(filepath.Join("C:", "\\Windows"))

	// fmt.Println(autope.LoadHive(`C:\Windows\System32\config\SYSTEM`, "OfflineWindows"))
	fmt.Println(autope.PickupNetCfg(`SYSTEM\ControlSet001`))
}
