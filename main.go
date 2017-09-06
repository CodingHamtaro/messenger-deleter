package main

import (
	"log"
	"strconv"
	"time"

	"github.com/sclevine/agouti"
)

var (
	settings, settingsErr = ParseSettings()
	todayFolder           = time.Now().String()
)

func processConversation(page *agouti.Page, c *agouti.MultiSelection, convoCount int) (result string) {
	var removableConvo []*agouti.Selection
	log.Println(procFlag, "Excluding conversations based from settings.json...")
	for i := 0; i < convoCount; i++ {
		convoName, _ := c.At(i).Find("span._1ht6").Text()
		if !FindString(convoName, settings.Messages.Excluded) {
			removableConvo = append(removableConvo, c.At(i))
		}
	}

	if len(removableConvo) == 0 {
		return "No removable conversations found."
	}

	for _, c := range removableConvo {
		personName, _ := c.Find("span._1ht6").Text()
		// due to an unknown reason, we need to reverify the name from the excluded list
		if !FindString(personName, settings.Messages.Excluded) {
			log.Println(procFlag, "Deleting conversations with "+personName)

			// click first the conversation
			c.Find("span._1ht6").Click()

			// Clicking the gear icon (for options)
			page.SetImplicitWait(150)

			page.Find("div[data-testid=info_panel] div._5blh._4-0h").Click()

			// After clicking the gear icon, it will show the menu on the right side
			/**
			Normally (but the top two menu is always the same for 'Friends', 'Blocked', 'Group', 'Blocked From Messenger', 'Non-Friend'):
			- index:0 Archive
			- index:1 Delete
			- index:2 (Separator)
			- index:3 Mark As Unread
			- index:4 Mark As Spam
			- index:5 Report Spam or Abuse...
			- index:6 Block Message
			**/
			menu := page.All("div[class='uiContextualLayerPositioner uiLayer'] li.__MenuItem")
			err := menu.At(1).Click()
			if err != nil {
				log.Println(errFlag, "An error occured, the element might no longer attached to DOM. Continuing the process...")
			}

			page.Screenshot("./screenshots/" + todayFolder + "/c-" + personName + ".png")

			// Expected that the dialog box is open, then we will click the delete button
			dialogButtons := page.All("div[aria-label='Dialog content'] button")
			dialogButtons.At(2).Click()
		}
	}

	log.Println(procFlag, "Done with the current batch. Refreshing the page and checks the remaining available conversations...")

	page.Refresh()

	convoList := page.All("ul[aria-label='Conversation List'] li")

	_, err := convoList.Visible()
	if err != nil {
		log.Println(errFlag, err.Error())
		return "An error occured"
	}

	currentNoOfConvo, err := convoList.Count()
	if err != nil {
		log.Println(errFlag, err.Error())
		return "An error occured"
	}
	log.Println(procFlag, "Number of conversations: "+strconv.Itoa(currentNoOfConvo))
	log.Println(procFlag, "Going to process the unwanted conversations...")
	if currentNoOfConvo > 0 {
		processConversation(page, convoList, currentNoOfConvo)
	}

	return "No removable conversations found."
}

func main() {
	// Opening
	SigOpeming()

	CreateDirIfNotExist("./screenshots/" + todayFolder)
	// settings
	if settingsErr != nil {
		log.Println(errFlag, settingsErr.Error())
		return
	}

	// web driver1
	capabilities := agouti.NewCapabilities()
	capabilities = capabilities.Platform("WINDOWS").Browser("internet explorer")
	capabilities["phantomjs.page.settings.XSSAuditingEnabled"] = true
	capabilitesOptions := agouti.Desired(capabilities)
	driver := agouti.PhantomJS(capabilitesOptions)
	err := driver.Start()
	if err != nil {
		log.Println(errFlag, err.Error())
		return
	}

	// navigation
	page, err := driver.NewPage()
	if err != nil {
		log.Println(errFlag, err.Error())
		driver.Stop()
		return
	}
	page.Size(960, 700)

	// home page and fill login form
	err = page.Navigate("https://www.facebook.com/")
	if err != nil {
		log.Println(errFlag, err.Error())
		driver.Stop()
		return
	}
	title, _ := page.Title()
	log.Println(pageFlag, title)
	page.Screenshot("./screenshots/" + todayFolder + "/login.png")

	page.AllByName("email").Fill(settings.Account.Username)
	page.AllByName("pass").Fill(settings.Account.Password)

	page.Screenshot("./screenshots/" + todayFolder + "/login-filled.png")

	page.FindByID("login_form").Submit()
	page.Refresh()
	log.Println(procFlag, "Logged In!")
	page.Screenshot("./screenshots/" + todayFolder + "/loggedin.png")

	// go to messenger
	page.Navigate("https://www.facebook.com/messages/t")

	title, _ = page.Title()
	log.Println(pageFlag, title)
	page.Screenshot("./screenshots/" + todayFolder + "/messenger.png")
	convoList := page.All("ul[aria-label='Conversation List'] li")

	_, err = convoList.Visible()
	if err != nil {
		log.Println(errFlag, err.Error())
		driver.Stop()
		return
	}

	currentNoOfConvo, err := convoList.Count()
	if err != nil {
		log.Println(errFlag, err.Error())
		driver.Stop()
		return
	}
	log.Println(procFlag, "Number of conversations: "+strconv.Itoa(currentNoOfConvo))
	log.Println(procFlag, "Going to process the unwanted conversations...")

	res := processConversation(page, convoList, currentNoOfConvo)
	log.Println(procFlag, res)

	driver.Stop()
}
