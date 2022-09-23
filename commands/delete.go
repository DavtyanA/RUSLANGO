package commands

import (
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
			lastmsg := msgs[1].Content
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
				for _, m := range msgs {
					ids = append(ids, m.ID)
				}
				s.ChannelMessagesBulkDelete(channel, ids)
				s.ChannelMessageSend(channel, Delete_Success)
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
