package commands

import (
	"github.com/bwmarrin/discordgo"
)

func Delete(s *discordgo.Session, channel string, message string, number int) {
	return
	// msg := strings.Split(message, " ")

	// if len(msg) == 2 {
	// 	number, err := strconv.Atoi(msg[1])
	// 	if err == nil {
	// 		chann, _ := s.Channel(channel)
	// 		if number > 20 {
	// 			s.ChannelMessageSend(channel, "А не дохуя ли?")
	// 		} else if s.ChannelMessages()
	// 	} else {
	// 		s.ChannelMessageSend(channel, "Я не знаю хотел ты удалить сообщения или нет, но если хотел, нужно написать сколько. Например 'Удали 5'")
	// 	}
	// 	// s.ChannelMessageDelete()
	// } else {
	// 	s.ChannelMessageSend(channel, "Я не знаю хотел ты удалить сообщения или нет, но если хотел, нужно написать сколько. Например 'Удали 5'")
	// }
}
