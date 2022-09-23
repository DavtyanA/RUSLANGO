package events

import (
	"RUSLANGO/commands"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func OnServerJoin(s *discordgo.Session, e *discordgo.GuildMemberAdd) {
	var responses []string
	user := e.User
	name := user.Mention()
	if user.Bot {
		responses = append(responses,
			"Ну и чё за "+name+"? А, новый бот... Ну-ну...",
			"Ну и нахрена нам "+name+"? Что он такого умеет, чего не умею я?",
			"Капец, у нас и так полно ботов, так ты еще и "+name+"позвал. Сервер не резиновый знаешь ли!",
			"Такс-такс, кто тут у нас... "+name+"? Пфф, да в этой мете его вообще никто на сервера не приглашает",
			"Ты чё наделал? Ты в курсе, что "+name+" Сервера взламывает и всех кикает? Убирай его отсюда, давай-давай",
			"О, пополнение в семействе ботов! "+name+", будешь сыночкой!")
	} else {
		responses = append(responses,
			"Ебать да к нам пожаловал "+name,
			"Явился, не запылился, "+name+". А голову ты дома не забыл?",
			name+" зашёл на хату. Мыло со стола или хлеб с параши?",
			"Пацаны к нам в рубку просится "+name+". Можно впустить?",
			"РУСТАМПИДОРАСРУСТАМПИДОРАСРУСТАМПИДОРАСРУСТАМПИДОРАС. Ах да, к нам зашёл "+name,
			"О, "+name+", поёдешь с нами на БТ?",
			"Аллах дал нам 86400 секунд сегодня, "+name+", используй хоть 1 чтобы сказать АЛЬХАМДУЛИЛЛЯХ",
			"А я всё думал, когда же ты появишься,  "+name,
			"Здорова, "+name+". Напиши Вадиму, он пришлёт тебе очень смешные биохимические мемы",
			"ЁПТА "+name+" ЗАШЁЛ, КИКАЙТЕ ЕГО НАХРЕН!!!!!!!!!!!!!!!!!!!!!!",
			"Roses are red, violets are blue, "+name+" Зашёл на сервак, а мне похую...",
			name+", Если ты мужчина, снимай штаны. Если женщина, попрошу уйти с сервера",
			"Бляяяяяяяя, опять Таня или Эрик своих друзей сюда зовут... "+name+", тебе тут не рады!!!",
			"ТЕСТИРУЕМ ПРИВЕТСТВИЯ! ПРИВЕТ, "+name+", НАПИШИ В ЧАТ 'ОЛЕГ ЕРМОЛАЕВ' ДЛЯ ДОСТУПА К СЕКРЕТНОМУ КАНАЛУ")
	}
	greeting := commands.GetRandomItem(responses)
	s.ChannelMessageSend(commands.General_Chat_ID, greeting)
}

func OnServerLeave(s *discordgo.Session, e *discordgo.GuildMemberRemove) {
	var responses []string
	user := e.User
	name := user.Mention()
	responses = append(responses,
		"Чё, сука, ушёл, "+name+"?",
		name+" съебался крыса!",
		"Уходишь, "+name+"? Да и не особо хотелось!",
		name+" может и не возвращаться, он не Саске",
		"От нас ушёл "+name+". А это кто вообще был?",
		"Как от Вадима ушёл Руслан, так и от нас ушёл "+name,
		"РУСТАМПИДОРАСРУСТАМПИДОРАСРУСТАМПИДОРАСРУСТАМПИДОРАС. Ах да, от нас ушёл "+name,
		"Как говорится, с корабля первым бежит "+name,
		"Олег, ты нахрена "+name+" забанил?")
	farewell := commands.GetRandomItem(responses)
	s.ChannelMessageSend(commands.General_Chat_ID, farewell)
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func OnMessage(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	channel := m.ChannelID
	var responses []string

	//MESSAGE IS EQUAL
	switch strings.ToLower(m.Content) {
	case "руслан":
		commands.SendFileFromS3(s, channel, commands.Pictures_Folder_Other+"ismail.jpg")
		s.ChannelMessageSend(channel, "Я НОВЫЙ УСОВЕРШЕНСТВОВАННЫЙ РУСЛАН, RUSLAN GO!")

	case "пошел нахуй", "пошёл нахуй", "иди нахуй", "пошол нахуй", "пишов нахуй", "пашел нахуй", "пашол нахуй":
		responses = append(responses, "Дорогу покажешь?",
			"Когда пошёл, тебя там нашёл!",
			"Нахуй твая жопа хороша",
			"нахуй не дорожка запомни мондовошка",
			"Нахуй не такси догани да отсоси!",
			"Ходить на хyй - твоя ежедневная работа. Не буду отнимать у тебя хлеб.",
			"Боюсь, что в твоих трусах я его не найду!",
			"Я та пойду нахуй мне на нём всю жизнь вертеться,а ті один раз сядишь и пидорасом станешь!!!",
			"Да мне всю жизнь на хую вертеться, а ты один раз и пидорас!",
			"Жопа не подушка,хуй не расскладушка",
			"По хуям прыгать твоя часть работы хуй в зубы канистру в очече и пошел вон отсюда , а то обоссу как того бедного пацана",
			"Кусай захуй")
		response := commands.GetRandomItem(responses)
		s.ChannelMessageSend(channel, response)

	// case "споки", "спокойной ночи", "сладких снов":
	// 	commands.SendFileFromS3(s, channel, commands.Pictures_Folder_Other+"isleep.gif")
	case "во что поиграть", "во что поиграть?":
		response := commands.WhatToPlay(m.Author.ID)
		s.ChannelMessageSend(channel, response)
	case "справедливо":
		commands.SendFileFromS3(s, channel, commands.Pictures_Folder_Other+"orehus-sticker.png")
		chann, _ := s.Channel(channel) //idk what name to give
		s.MessageReactionAdd(channel, chann.LastMessageID, ":orehus:400349897578250255")
	case "кто", "кто?":
		s.ChannelMessageSend(channel, "Дарцаев Исмаил Умарпашаевич 11 микрорайон космонавтов 54 приезжайте я чеченец таких пидорасов я буду разъебывать, и вас я буду разъебывать")
		commands.SendFileFromS3(s, channel, commands.Pictures_Folder_Other+"ismail.jpg")
	case "бот":
		s.ChannelMessageSend(channel, "Чё надо?")
	case "канал":
		s.ChannelMessageSend(channel, "https://www.youtube.com/c/AndreDavtyan\nПОДПИСЫВАЕМСЯ НА КАНАЛ\nСТАВИМ КОЛОКОЛЬЧИКИ")
	case "!soundcloud https://soundcloud.com/fendiglock/sets/true-religion":
		s.ChannelMessageSend(channel, "Наконец-то нормальную музыку врубили")
	case "брук":
		s.ChannelMessageSend(channel, "Да никогда!")
	case "порно комиксы", "порнокомиксы":
		s.ChannelMessageSend(channel, "https://vk.com/porno_comics_goog")
	case "михей":
		s.ChannelMessageSend(channel, "Хуйло твой рот")
	case "спорим?":
		s.ChannelMessageSend(channel, "Ты кончишь")
	case "йоу мадина":
		s.ChannelMessageSend(channel, "Я НА ГЕЛЕНТВАГЕНЕ")
	case "салам", "селем":
		s.ChannelMessageSend(channel, "Салам Алейкум!")
	case "салам алейкум", "ассалам алейкум", "ассаламу уалайкум", "ассаламу алейкум", "салам алайкум":
		s.ChannelMessageSend(channel, "Алейкум Ассалам!")
	case "олег":
		s.ChannelMessageSend(channel, "Ну а сразу сказать сложно сказать? Вот серьезно, ты такой ебло сидишь выёбываешься, думаешь тебе ничё не будет, "+
			"скажи, яйца иметь надо, ты выглядишь как петух блять, говоришь как ебаный петух! Почему нельзя быть хоть немножечко иметь яйца, что-то говорить?")
	case "руслан пидорас соси хуй":
		s.ChannelMessageSend(channel, "Сосу брат...")
	case "никита пидорас соси хуй":
		s.ChannelMessageSend(channel, "Big Facts...")
	case "олег пидорас соси хуй":
		commands.SendFileFromS3(s, channel, commands.Pictures_Folder_Other+"olegfuckyou.jpg")
	case "как дела", "как дела?":
		s.ChannelMessageSend(channel, "Пока не родила")
	case "нихуя":
		commands.SendFileFromS3(s, channel, commands.Pictures_Folder_Other+"nixuya.jpg")
	case "batman":
		commands.SendFileFromS3(s, channel, commands.Pictures_Folder_Other+"batman.jpg")
	case "блядь":
		commands.SendFileFromS3(s, channel, commands.Pictures_Folder_Other+"blyad.jpg")
	case "вадим":
		commands.SendFileFromS3(s, channel, commands.Pictures_Folder_Other+"vadim.jpg")
	case "лучший герой в доте":
		commands.SendFileFromS3(s, channel, commands.Pictures_Folder_Other+"pudge.jpg")
	case "хуя":
		commands.SendFileFromS3(s, channel, commands.Pictures_Folder_Other+"xuya.jpg")
	case "бля":
		commands.SendFileFromS3(s, channel, commands.Pictures_Folder_Other+"blep.png")
	case "блятб":
		commands.SendFileFromS3(s, channel, commands.Pictures_Folder_Other+"blyatb.jpg")
	case "zipper":
		commands.SendFileFromS3(s, channel, commands.Pictures_Folder_Other+"bruno.png")
	case "мясо":
		commands.SendFileFromS3(s, channel, commands.Pictures_Folder_Other+"myaso.jpg")
	case "усы":
		commands.SendFileFromS3(s, channel, commands.Pictures_Folder_Other+"moustache.jpg")
	case "сандро":
		commands.SendFileFromS3(s, channel, commands.Pictures_Folder_Other+"sandro.jpg")
	case "рамзан":
		commands.SendFileFromS3(s, channel, commands.Pictures_Folder_Other+"ramzan.jpg")
	case "мадина текст":
		s.ChannelMessageSend(channel, commands.Madina_Text)
	case "сам нюхни":
		commands.SendFileFromS3(s, channel, commands.Pictures_Folder_Other+"smell-bebra.gif")
	case "lf", "da", "да", "ну да":
		commands.SendFileFromS3(s, channel, commands.Pictures_Folder_Other+"pizda.gif")
	case "серьезно?":
		commands.SendFileFromS3(s, channel, commands.Pictures_Folder_Other+"pizdabol.gif")
	case "f":
		commands.SendRandomFileFromFolder(s, channel, "F") //need to remember the F below
	case "anek", "anec", "anecode", "анек", "анекдот", "юмор", "юмореска", "поржать", "ржать", "нуждик":
		anecdote := commands.Make_request("getRandItem")
		// s.ChannelFileSendWithMessage(channel, anecdote+"\n\n ДАННЫЙ АНЕКДОТ ПРОСПОНСИРОВАН ОЛЕГОМ ЕРМОЛАЕВЫМ", "oleg.jpg", )
		s.ChannelMessageSend(channel, anecdote+"\n ДАННЫЙ АНЕКДОТ ПРОСПОНСИРОВАН ОЛЕГОМ ЕРМОЛАЕВЫМ")
	case "расскажи историю":
		if commands.Roll(5) == "5" {
			s.ChannelMessageSend(channel, commands.StoryTelling())
		} else {
			commands.MegaStory(s, channel)
		}
	}

	switch {
	case commands.StringContainsArray(m.Content, []string{"бебр", "bebr"}):
		commands.SendRandomFileFromFolder(s, channel, "bebra")
	case commands.StringContainsArray(m.Content, []string{"бан", "ban"}):
		commands.SendRandomFileFromFolder(s, channel, "ban")
	case commands.StringContainsArray(m.Content, []string{"пиздец", "бля...", "жаль", "грустно", "хуево", "хуёво", "мде", "press f"}): //need to remember the F above
		commands.SendRandomFileFromFolder(s, channel, "F")
	case commands.StringContainsArray(m.Content, []string{"фото член дрочить", "дрочить", "порно"}):
		commands.SendRandomFileFromFolder(s, channel, "cock")
	case commands.StringContainsArray(m.Content, []string{"loss", "потеря"}):
		commands.SendRandomFileFromFolder(s, channel, "loss")
	case commands.StringContainsArray(m.Content, []string{"amogus", "амогус", "амонг", "amog", "а мог", "сус", "sus", "among us"}):
		commands.SendRandomFileFromFolder(s, channel, "amogus")
	case commands.StringStartsWith(m.Content, "споки"), commands.StringStartsWith(m.Content, "спокойной ночи"), commands.StringStartsWith(m.Content, "сладких снов"):
		commands.SendFileFromS3(s, channel, commands.Pictures_Folder_Other+"isleep.gif")
	case commands.StringContains(m.Content, "козлов"), commands.StringContains(m.Content, "кызлар"):
		commands.SendFileFromS3(s, channel, commands.Pictures_Folder_Other+"kozlov1.jpg")
		time.Sleep(1 * time.Second)
		commands.SendFileFromS3(s, channel, commands.Pictures_Folder_Other+"kozlov2.jpg")
	case commands.StringContains(m.Content, "русские вперед"), commands.StringContains(m.Content, "русские вперёд"):
		commands.SendFileFromS3(s, channel, commands.Pictures_Folder_Other+"russiansgo.jpg")
	case commands.StringContainsArray(m.Content, []string{"пошел нахер", "пошёл нахер", "пошёл ты нахер козёл", "пошёл ты нахер козел", "пошел ты нахер козёл", "пошел ты нахер козел"}):
		s.ChannelMessageSend(channel, "https://youtu.be/qks8SgT1B4M")
	case commands.StringStartsWithArray(m.Content, []string{"roll", "ролляй", "роляй", "ролл"}):
		message := strings.Split(m.Content, " ")
		var response string
		if len(message) > 1 {
			num, err := strconv.Atoi(message[1])
			if err != nil {
				response = "ептвою мать пиши нормально `ролл 5`, `ролляй 100`, `roll 228` нахуй мне твои буквы"
			} else {
				response = commands.Roll(num)
			}
		} else {
			response = commands.Roll(100)
		}
		s.ChannelMessageSend(channel, response)
	case commands.StringContains(m.Content, "посмотрим"), commands.StringContains(m.Content, "will see"), commands.StringContains(m.Content, "will see"):
		commands.SendFileFromS3(s, channel, commands.Pictures_Folder_Other+"WillSee.jpg")
	case commands.StringContainsArray(m.Content, []string{"кастом", "ресел", "кемп", "дроп", "км", "сток", "сникеры", "хайпбист", "оффер", "сайз", "кондей", "ритейл", "легит чек",
		"броук", "лейм", "шакал", "кук группа", "лоуболлер"}):
		commands.SendFileFromS3(s, channel, commands.Pictures_Folder_Other+"custom.jpg")
	case commands.StringContains(m.Content, "ты кто"), commands.StringContains(m.Content, "кто это"), commands.StringContains(m.Content, "это кто"):
		s.ChannelMessageSend(channel, "Я Дарцаев Исмаил Умарпашаевич 11 микрорайон космонавтов 54 приезжайте я чеченец таких пидорасов я буду разъебывать, и вас я буду разъебывать")
		commands.SendFileFromS3(s, channel, commands.Pictures_Folder_Other+"ismail.jpg")
	case commands.StringContains(m.Content, "сука"), commands.StringContains(m.Content, "cerf"):
		if m.Author.ID == "333341352404320267" {
			s.ChannelMessageSend(channel, "Ну тут сыглы")
		} else {
			s.ChannelMessageSend(channel, "Сам сука")
		}
	case commands.StringContains(m.Content, "что делать"),
		commands.StringContains(m.Content, "че делать"),
		commands.StringContains(m.Content, "че сделать"),
		commands.StringContains(m.Content, "что сделать"):
		commands.SendFileFromS3(s, channel, commands.Pictures_Folder_Other+"choosepudge.jpg")
	case commands.StringContains(m.Content, "мем"), commands.StringContains(m.Content, "meme"):
		commands.SendFileFromS3(s, channel, commands.Pictures_Folder_Other+"meme.jpg")
	case commands.StringStartsWith(m.Content, "удали"):
		commands.Delete(s, channel, m)
	case commands.StringContains(m.Content, "вадим"), m.Author.ID == commands.Ducks_Fuhrer_ID:
		s.ChannelMessageSend(channel, "А вы в курсе что Вадим натурал?")
		// discordgo.MessageEmbed
		// s.ChannelMessageSendEmbed()
	}
}

func OnBotReady(s *discordgo.Session, m *discordgo.Connect) {
	s.ChannelMessageSend(commands.Botchat_ID, "Я ТУТ ГАЙЗ")
}
