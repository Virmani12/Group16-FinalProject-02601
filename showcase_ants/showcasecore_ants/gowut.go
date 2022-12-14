package showcasecore_ants

import (
	"fmt"
	"log"

	"github.com/icza/gowut/gwu"
)

// plural returns an empty string if i is equal to 1,
// "s" otherwise.
func plural(i int) string {
	if i == 1 {
		return ""
	}
	return "s"
}

func buildHomeDemo(event gwu.Event) gwu.Comp {
	p := gwu.NewPanel()

	content := []string{
		"This app is written in and showcases Gowut " + gwu.GowutVersion + ".",
		"Everything you see here is modeled and represented in Go (server side). Everything is created with Go code only (no HTML).",
		"Select components on the right side to see them in action",
	}

	for _, s := range content {
		p.Add(gwu.NewLabel(s))
		p.AddVSpace(20)
	}

	return p
}

func buildTextBoxDemo(event gwu.Event) gwu.Comp {
	// need to make a panel to access these fuinctions for the user
	p := gwu.NewPanel()

	p.Add(gwu.NewLabel("Enter alpha value (max 15 characters):"))
	row := gwu.NewHorizontalPanel()
	alphatb := gwu.NewTextBox("")
	alphatb.SetMaxLength(15)
	alphatb.AddSyncOnETypes(gwu.ETypeKeyUp)
	length := gwu.NewLabel("")
	length.Style().SetFontSize("80%").SetFontStyle(gwu.FontStyleItalic)
	alphatb.AddEHandlerFunc(func(e gwu.Event) {
		rem := 15 - len(alphatb.Text())
		length.SetText(fmt.Sprintf("(%d character%s left...)", rem, plural(rem)))
		e.MarkDirty(length)
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row.Add(alphatb)
	row.Add(length)
	p.Add(row)
	// space between last text box and next
	p.AddVSpace(10)

	/// new entry for beta paramter

	p.Add(gwu.NewLabel("Enter beta value (max 15 characters):"))
	row2 := gwu.NewHorizontalPanel()
	betatb := gwu.NewTextBox("")
	betatb.SetMaxLength(15)
	betatb.AddSyncOnETypes(gwu.ETypeKeyUp)
	length2 := gwu.NewLabel("")
	length.Style().SetFontSize("80%").SetFontStyle(gwu.FontStyleItalic)
	betatb.AddEHandlerFunc(func(e gwu.Event) {
		rem := 15 - len(betatb.Text())
		length2.SetText(fmt.Sprintf("(%d character%s left...)", rem, plural(rem)))
		e.MarkDirty(length)
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row2.Add(betatb)
	row2.Add(length)
	p.Add(row2)
	// space between last text box and next
	p.AddVSpace(10)
	/// new entry for rho new paramter

	p.Add(gwu.NewLabel("Enter beta value (max 15 characters):"))
	row3 := gwu.NewHorizontalPanel()
	rhotb := gwu.NewTextBox("")
	rhotb.SetMaxLength(15)
	rhotb.AddSyncOnETypes(gwu.ETypeKeyUp)
	length3 := gwu.NewLabel("")
	length3.Style().SetFontSize("80%").SetFontStyle(gwu.FontStyleItalic)
	rhotb.AddEHandlerFunc(func(e gwu.Event) {
		rem := 15 - len(rhotb.Text())
		length3.SetText(fmt.Sprintf("(%d character%s left...)", rem, plural(rem)))
		e.MarkDirty(length)
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row3.Add(rhotb)
	row3.Add(length)
	p.Add(row3)
	// space between last text box and next
	p.AddVSpace(10)

	// parameter entry for Q

	p.Add(gwu.NewLabel("Enter beta value (max 15 characters):"))
	row4 := gwu.NewHorizontalPanel()
	qtb := gwu.NewTextBox("")
	qtb.SetMaxLength(15)
	qtb.AddSyncOnETypes(gwu.ETypeKeyUp)
	length4 := gwu.NewLabel("")
	length4.Style().SetFontSize("80%").SetFontStyle(gwu.FontStyleItalic)
	qtb.AddEHandlerFunc(func(e gwu.Event) {
		rem := 15 - len(qtb.Text())
		length4.SetText(fmt.Sprintf("(%d character%s left...)", rem, plural(rem)))
		e.MarkDirty(length)
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row4.Add(rhotb)
	row4.Add(length)
	p.Add(row4)
	// space between last text box and next
	p.AddVSpace(10)

	// parameter entry for numTowns
	p.Add(gwu.NewLabel("Enter beta value (max 15 characters):"))
	row5 := gwu.NewHorizontalPanel()
	initialinttb := gwu.NewTextBox("")
	initialinttb.SetMaxLength(15)
	initialinttb.AddSyncOnETypes(gwu.ETypeKeyUp)
	length5 := gwu.NewLabel("")
	length5.Style().SetFontSize("80%").SetFontStyle(gwu.FontStyleItalic)
	initialinttb.AddEHandlerFunc(func(e gwu.Event) {
		rem := 15 - len(initialinttb.Text())
		length5.SetText(fmt.Sprintf("(%d character%s left...)", rem, plural(rem)))
		e.MarkDirty(length)
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row5.Add(rhotb)
	row5.Add(length)
	p.Add(row5)
	// space between last text box and next
	p.AddVSpace(10)

	// parameter entry for numTowns
	p.Add(gwu.NewLabel("Enter beta value (max 15 characters):"))
	row6 := gwu.NewHorizontalPanel()
	townstb := gwu.NewTextBox("")
	townstb.SetMaxLength(15)
	townstb.AddSyncOnETypes(gwu.ETypeKeyUp)
	length6 := gwu.NewLabel("")
	length6.Style().SetFontSize("80%").SetFontStyle(gwu.FontStyleItalic)
	townstb.AddEHandlerFunc(func(e gwu.Event) {
		rem := 15 - len(townstb.Text())
		length6.SetText(fmt.Sprintf("(%d character%s left...)", rem, plural(rem)))
		e.MarkDirty(length)
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row6.Add(rhotb)
	row6.Add(length)
	p.Add(row6)
	// space between last text box and next
	p.AddVSpace(10)

	// parameter entry for numAnts
	p.Add(gwu.NewLabel("Enter beta value (max 15 characters):"))
	row7 := gwu.NewHorizontalPanel()
	antstb := gwu.NewTextBox("")
	antstb.SetMaxLength(15)
	antstb.AddSyncOnETypes(gwu.ETypeKeyUp)
	length7 := gwu.NewLabel("")
	length7.Style().SetFontSize("80%").SetFontStyle(gwu.FontStyleItalic)
	antstb.AddEHandlerFunc(func(e gwu.Event) {
		rem := 15 - len(antstb.Text())
		length7.SetText(fmt.Sprintf("(%d character%s left...)", rem, plural(rem)))
		e.MarkDirty(length)
	}, gwu.ETypeChange, gwu.ETypeKeyUp)
	row7.Add(rhotb)
	row7.Add(length)
	p.Add(row7)
	// space between last text box and next
	p.AddVSpace(10)

	return p
}

// will need to make a button that will act as an initialization starter
// will need a button for random towns or Oliver30 placement
// will enter values from repsective text box; i.e. firstarg := gwu.HasText.Text(alpatb)

type demo struct {
	link      gwu.Label
	buildFunc func(gwu.Event) gwu.Comp
	comp      gwu.Comp // Lazily initialized demo comp
}
type pdemo *demo

var extraHeadHTMLs []string

func buildShowcaseWin(sess gwu.Session) {
	win := gwu.NewWindow("show", "Showcase of Features - Gowut")
	for _, headHTML := range extraHeadHTMLs {
		win.AddHeadHTML(headHTML)
	}

	win.Style().SetFullSize()
	win.AddEHandlerFunc(func(e gwu.Event) {
		switch e.Type() {
		case gwu.ETypeWinLoad:
			log.Println("LOADING window:", e.Src().ID())
		case gwu.ETypeWinUnload:
			log.Println("UNLOADING window:", e.Src().ID())
		}
	}, gwu.ETypeWinLoad, gwu.ETypeWinUnload)

	hiddenPan := gwu.NewNaturalPanel()
	sess.SetAttr("hiddenPan", hiddenPan)

	header := gwu.NewHorizontalPanel()
	header.Style().SetFullWidth().SetBorderBottom2(2, gwu.BrdStyleSolid, "#cccccc")
	title := gwu.NewLink("Gowut - Showcase of Features", win.Name())
	title.SetTarget("")
	title.Style().SetColor(gwu.ClrBlue).SetFontWeight(gwu.FontWeightBold).SetFontSize("120%").Set("text-decoration", "none")
	header.Add(title)
	header.AddHConsumer()
	header.Add(gwu.NewLabel("Session timeout:"))
	header.Add(gwu.NewSessMonitor())
	header.AddHSpace(10)
	header.Add(gwu.NewLabel("Theme:"))
	themes := gwu.NewListBox([]string{"default", "debug"})
	themes.AddEHandlerFunc(func(e gwu.Event) {
		win.SetTheme(themes.SelectedValue())
		e.ReloadWin("show")
	}, gwu.ETypeChange)
	header.Add(themes)
	header.AddHSpace(10)
	reset := gwu.NewLink("Reset", "#")
	reset.Style().SetColor(gwu.ClrBlue)
	reset.SetTarget("")
	reset.AddEHandlerFunc(func(e gwu.Event) {
		e.RemoveSess()
		e.ReloadWin("show")
	}, gwu.ETypeClick)
	header.Add(reset)
	setNoWrap(header)
	win.Add(header)

	content := gwu.NewHorizontalPanel()
	content.SetCellPadding(1)
	content.SetVAlign(gwu.VATop)
	content.Style().SetFullSize()

	demoWrapper := gwu.NewPanel()
	demoWrapper.Style().SetPaddingLeftPx(5)
	demoWrapper.AddVSpace(10)
	demoTitle := gwu.NewLabel("")
	demoTitle.Style().SetFontWeight(gwu.FontWeightBold).SetFontSize("110%")
	demoWrapper.Add(demoTitle)
	demoWrapper.AddVSpace(10)

	links := gwu.NewPanel()
	links.SetCellPadding(1)
	links.Style().SetPaddingRightPx(5)

	demos := make(map[string]pdemo)
	var selDemo pdemo

	selectDemo := func(d pdemo, e gwu.Event) {
		if selDemo != nil {
			selDemo.link.Style().SetBackground("")
			if e != nil {
				e.MarkDirty(selDemo.link)
			}
			demoWrapper.Remove(selDemo.comp)
		}
		selDemo = d
		d.link.Style().SetBackground("#88ff88")
		demoTitle.SetText(d.link.Text())
		if d.comp == nil {
			d.comp = d.buildFunc(e)
		}
		demoWrapper.Add(d.comp)
		if e != nil {
			e.MarkDirty(d.link, demoWrapper)
		}
	}

	createDemo := func(name string, buildFunc func(gwu.Event) gwu.Comp) pdemo {
		link := gwu.NewLabel(name)
		link.Style().SetFullWidth().SetCursor(gwu.CursorPointer).SetDisplay(gwu.DisplayBlock).SetColor(gwu.ClrBlue)
		demo := &demo{link: link, buildFunc: buildFunc}
		link.AddEHandlerFunc(func(e gwu.Event) {
			selectDemo(demo, e)
		}, gwu.ETypeClick)
		links.Add(link)
		demos[name] = demo
		return demo
	}

	links.Style().SetFullHeight().SetBorderRight2(2, gwu.BrdStyleSolid, "#cccccc")
	links.AddVSpace(5)
	homeDemo := createDemo("Home", buildHomeDemo)
	selectDemo(homeDemo, nil)
	links.AddVSpace(5)
	l := gwu.NewLabel("Component Palette")
	l.Style().SetFontWeight(gwu.FontWeightBold).SetFontSize("110%")
	links.Add(l)
	links.AddVSpace(5)
	l = gwu.NewLabel("Containers")
	l.Style().SetFontWeight(gwu.FontWeightBold)
	links.Add(l)
	/*
		createDemo("Expander", buildExpanderDemo)
		createDemo("Link (as Container)", buildLinkContainerDemo)
		createDemo("Panel", buildPanelDemo)
		createDemo("Table", buildTableDemo)
		createDemo("TabPanel", buildTabPanelDemo)
		createDemo("Window", buildWindowDemo)
	*/
	links.AddVSpace(5)
	// can call l something else here
	l = gwu.NewLabel("Input components")
	l.Style().SetFontWeight(gwu.FontWeightBold).SetDisplay(gwu.DisplayBlock)
	links.Add(l)
	createDemo("TextBox", buildTextBoxDemo)
	/*
		createDemo("CheckBox", buildCheckBoxDemo)
		createDemo("ListBox", buildListBoxDemo)
		createDemo("PasswBox", buildPasswBoxDemo)
		createDemo("RadioButton", buildRadioButtonDemo)
		createDemo("SwitchButton", buildSwitchButtonDemo)
	*/
	//links.AddVSpace(5)
	/*
		// can make l = maps and graphs
		l = gwu.NewLabel("Other components")
		l.Style().SetFontWeight(gwu.FontWeightBold)
		links.Add(l)

			createDemo("Button", buildButtonDemo)
			createDemo("HTML", buildHTMLDemo)
			createDemo("Image", buildImageDemo)
			createDemo("Label", buildLabelDemo)
			createDemo("Link", buildLinkDemo)
			createDemo("SessMonitor", buildSessMonitorDemo)
			createDemo("Timer", buildTimerDemo)
	*/
	links.AddVConsumer()
	setNoWrap(links)
	content.Add(links)
	content.Add(demoWrapper)
	content.CellFmt(demoWrapper).Style().SetFullWidth()

	win.Add(content)
	win.CellFmt(content).Style().SetFullSize()

	footer := gwu.NewHorizontalPanel()
	footer.Style().SetFullWidth().SetBorderTop2(2, gwu.BrdStyleSolid, "#cccccc")
	footer.Add(hiddenPan)
	footer.AddHConsumer()
	l = gwu.NewLabel("Copyright © 2013-2018 András Belicza. All rights reserved.")
	l.Style().SetFontStyle(gwu.FontStyleItalic).SetFontSize("95%")
	footer.Add(l)
	footer.AddHSpace(10)
	link := gwu.NewLink("Visit the Gowut Wiki", "https://github.com/icza/gowut/wiki")
	link.Style().SetFontStyle(gwu.FontStyleItalic).SetFontSize("95%")
	footer.Add(link)
	setNoWrap(footer)
	win.Add(footer)

	sess.AddWin(win)
}

// setNoWrap sets WhiteSpaceNowrap to all children of the specified panel.
func setNoWrap(panel gwu.Panel) {
	count := panel.CompsCount()
	for i := count - 1; i >= 0; i-- {
		panel.CompAt(i).Style().SetWhiteSpace(gwu.WhiteSpaceNowrap)
	}
}

// SessHandler is our session handler to build the showcases window.
type sessHandler struct{}

func (h sessHandler) Created(s gwu.Session) {
	buildShowcaseWin(s)
}

func (h sessHandler) Removed(s gwu.Session) {}

// StartServer creates and starts the Gowut GUI server.
func StartServer(appName, addr string, autoOpen bool) {
	// Create GUI server
	server := gwu.NewServer(appName, addr)
	for _, headHTML := range extraHeadHTMLs {
		server.AddRootHeadHTML(headHTML)
	}
	server.AddStaticDir("/asdf", "w:/")
	server.SetText("Gowut - Showcase of Features")

	server.AddSessCreatorName("show", "Showcase of Features - Gowut")
	server.AddSHandler(sessHandler{})
	// Just for the demo: Add an extra "Gowut-Server" header to all responses holding the Gowut version
	server.SetHeaders(map[string][]string{
		"Gowut-Server": {gwu.GowutVersion},
	})

	// Start GUI server
	var openWins []string
	if autoOpen {
		openWins = []string{"show"}
	}
	if err := server.Start(openWins...); err != nil {
		log.Println("Error: Cound not start GUI server:", err)
		return
	}
}
