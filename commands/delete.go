package commands

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func Delete(s *discordgo.Session, channel string, messageobj *discordgo.MessageCreate) {
	message := messageobj.Content
	authorID := messageobj.Author.ID
	//To divide message into separate words
	msg := strings.Split(message, " ")
	//only accept the word and the count
	if len(msg) == 2 {
		//if the second element is number, proceed, otherwise alert the user
		number, err := strconv.Atoi(msg[1])
		if err == nil {
			// chann, _ := s.Channel(channel)
			msgs, _ := s.ChannelMessages(channel, 2, "", "", "")
			//Second to last
			lastmessage := msgs[1]
			lastmsg := lastmessage.Content
			switch lastmsg {
			//if someone is oxyel
			case Delete_Success, Delete_FuckYou:
				if !isEnderlord(authorID) {
					s.ChannelMessageSend(channel, Delete_FuckYou)
					return
				}
			}

			//20 is doxuya, less than 1 is in case someone is a pidoras (even though, the functions below will do nothing)
			if !isEnderlord(authorID) && (number > 79 || number < 1) {
				s.ChannelMessageSend(channel, "А не дохуя ли?")
				return
			} else {
				msgs, _ := s.ChannelMessages(channel, number+1, "", "", "")
				var ids []string
				fmt.Println(lastmessage.Author.Username + "Has deleted " + strconv.Itoa(len(ids)) + " messages: ")
				for _, m := range msgs {
					ids = append(ids, m.ID)
				}
				s.ChannelMessagesBulkDelete(channel, ids)
				s.ChannelMessageSend(channel, Delete_Success)
				//because printing takes a long time, put it after everything's deleted
				//I should look into threading or async processes for this
				for _, m := range msgs {
					fmt.Println("mesasge:", m.Content, m.Attachments[0].Filename, "\nauthor: ", m.Author.Username+"\n")
				}
				return
			}

		} else {
			s.ChannelMessageSend(channel, Delete_Usage)
			return
		}
		// s.ChannelMessageDelete()
	} else {
		s.ChannelMessageSend(channel, Delete_Usage)
		return
	}
}

func isEnderlord(id string) bool {
	return id == Enderlord_ID
}
