package commands

import (
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func Delete(s *discordgo.Session, channel string, message string) {
	//To divide message into separate words
	msg := strings.Split(message, " ")
	//only accept the word and the count
	if len(msg) == 2 {
		//if the second element is number, proceed, otherwise alert the user
		number, err := strconv.Atoi(msg[1])
		if err == nil {
			chann, _ := s.Channel(channel)
			lastmsg, _ := s.ChannelMessage(channel, chann.LastMessageID)
			switch lastmsg.Content {
			//if someone is oxyel
			case Delete_Success, Delete_FuckYou:
				s.ChannelMessageSend(channel, Delete_FuckYou)
				return
			}

			//20 is doxuya, less than 1 is in case someone is a pidoras (even though, the functions below will do nothing)
			if number > 20 || number < 1 {
				s.ChannelMessageSend(channel, "А не дохуя ли?")
				return
			} else {
				msgs, _ := s.ChannelMessages(channel, number, "", "", "")
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
