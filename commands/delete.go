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
	//only accept the word and the messages count
	if len(msg) == 2 {
		//if the second element is number, proceed, otherwise alert the user
		delete_count, err := strconv.Atoi(msg[1])
		if err == nil {
			last_messages, _ := s.ChannelMessages(channel, 2, "", "", "")
			//Second to last
			secondtolastmessage := last_messages[1]
			lastmessage := last_messages[0] //is always the delete command
			slastmsg := secondtolastmessage.Content
			switch slastmsg {
			//if someone is oxyel
			case Delete_Success, Delete_FuckYou:
				if !IsEnderlord(authorID) {
					s.ChannelMessageSend(channel, Delete_FuckYou)
					return
				}
			}

			//After negotiation with Oleg, I came to number 79, no on vse ravno pidoras (po facts)
			if !IsEnderlord(authorID) && (delete_count > 79) {
				s.ChannelMessageSend(channel, "А не дохуя ли?")
				return
			}
			//message count + the delete command
			messages_to_delete, _ := s.ChannelMessages(channel, delete_count+1, "", "", "")
			var ids []string
			for _, m := range messages_to_delete {
				ids = append(ids, m.ID)
			}
			author := lastmessage.Author.Username
			fmt.Println(fmt.Sprint(author, " Has deleted ", strconv.Itoa(len(ids)-1), " messages:"))
			s.ChannelMessagesBulkDelete(channel, ids)
			s.ChannelMessageSend(channel, Delete_Success)
			//because printing takes a long time, put it after everything's deleted
			//I should look into threading or async processes for this
			count := 1
			for i := len(messages_to_delete) - 1; i >= 1; i-- {
				m := messages_to_delete[i]
				sb := strings.Builder{}
				sb.WriteString(fmt.Sprint("\nauthor: ", m.Author.Username, "\n"))
				sb.WriteString(fmt.Sprint("message ", count, ": ", m.Content))
				if len(m.Attachments) > 0 {
					sb.WriteString(m.Attachments[0].Filename)
				}
				fmt.Println(sb.String())
				count++
			}
			return

		}
		s.ChannelMessageSend(channel, Delete_Usage)
		return
	}
	s.ChannelMessageSend(channel, Delete_Usage)
}
