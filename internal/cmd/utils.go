package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/ghostrune/edge-tts-go/edge_tts"
)

func ListVoices() {
	voiceList, err := edge_tts.ListVoices(proxyURL)
	if err != nil {
		log.Fatal(err)
	}

	for _, voiceItem := range voiceList {
		fmt.Println("Name:", voiceItem.Name)
		fmt.Println("ShortName:", voiceItem.ShortName)
		fmt.Println("Gender:", voiceItem.Gender)
		fmt.Println("Locale:", voiceItem.Locale)
		fmt.Println("ContentCategories:", strings.Join(voiceItem.VoiceTag.ContentCategories, ","))
		fmt.Println("VoicePersonalities:", strings.Join(voiceItem.VoiceTag.VoicePersonalities, ","))
		fmt.Println()
	}
}
