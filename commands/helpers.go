package commands

import (
	"math/rand"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

//Get a random item from an array, works with dynamic types since Go 1.18! Zaebumba!
func GetRandomItem[T interface{}](inputarray []T) T {
	rand.Seed(time.Now().UnixNano())
	//not sure if I need (len - 1), examples don't do that
	randomItem := inputarray[rand.Intn(len(inputarray))]
	return randomItem
}

//String contains to lower case
func StringContains(S string, sub string) bool {
	return strings.Contains(strings.ToLower(S), strings.ToLower(sub))
}

//String starts with to lower case
func StringStartsWith(S string, sub string) bool {
	return strings.HasPrefix(strings.ToLower(S), strings.ToLower(sub))
}

//String contains but for array of substrings (for convenience)
func StringStartsWithArray(S string, subs []string) bool {
	for _, s := range subs {
		if StringStartsWith(S, s) {
			return true
		}
	}
	return false
}

//String contains but for array of substrings (for convenience)
func StringContainsArray(S string, subs []string) bool {
	for _, s := range subs {
		if StringContains(S, s) {
			return true
		}
	}
	return false
}

func MegaStory(s *discordgo.Session, channel string){

	SendFileFromS3(s, channel, Pictures_Folder_Other + "daiproidu.jpg")
	time.Sleep(1500 * time.Millisecond)
	SendFileFromS3(s, channel, Pictures_Folder_Other + "xuya.jpg")
	time.Sleep(1500 * time.Millisecond)
	SendFileFromS3(s, channel, Pictures_Folder_Other + "poebalu.jpg")
	time.Sleep(1500 * time.Millisecond)
	SendFileFromS3(s, channel, Pictures_Folder_Other + "choblyatb.jpg")
	time.Sleep(1500 * time.Millisecond)
	SendFileFromS3(s, channel, Pictures_Folder_Other + "razebu.jpg")
	time.Sleep(1500 * time.Millisecond)
	SendFileFromS3(s, channel, Pictures_Folder_Other + "willsee.jpg")
	time.Sleep(1500 * time.Millisecond)
	SendFileFromS3(s, channel, Pictures_Folder_Other + "taashaa.jpg")
	time.Sleep(1500 * time.Millisecond)
	SendFileFromS3(s, channel, Pictures_Folder_Other + "bilyateblo.jpg")
	time.Sleep(1500 * time.Millisecond)
	SendFileFromS3(s, channel, Pictures_Folder_Other + "che tam.jpg")
	time.Sleep(1500 * time.Millisecond)
	SendFileFromS3(s, channel, Pictures_Folder_Other + "blyatb.jpg")
}
