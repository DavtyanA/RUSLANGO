package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func WhatToPlay(authorID string) string {
	var responses []string
	name := "сука"
	//maybe make IDS enums, but not sure how to handle multiple IDs bois (Руслан и Козлов идите нахуй)
	switch authorID {
	case Enderlord_ID:
		return "Хеллоу ебать ты кто"
	case Giogis_ID:
		name = "Олег"
	case Mozart_ID:
		name = "Даня"
	case "434375071591563264", "911692496227139675", "796899505181032509", "395544758832988160", "313007272127102986":
		name = "Никита"
	case AYS_ID:
		name = "Айс"
	case Makich_ID:
		return "DOTA 2"
	case Mk7k_ID:
		name = "Михей"
	case Vnatureloh_ID:
		name = "Сандро"
	case David_ID:
		name = "Давид"
	case Zeeklik_ID:
		name = "Эрик"
	case Squirtana_ID:
		name = "Дима"
	case Neaus_ID:
		name = "Карпов"
	case Ducks_Fuhrer_ID:
		name = "Вадим"
	case Ruslan_ID:
		name = "Руслан"
		responses = append(responses, "Ты чё, долбоёб, сам с собой разговариваешь?")
	case "403451995601764352", "395256828075704322", "554784820924907570":
		responses = append(responses, `Пиздец ну давай я еще для девушек буду ответочки писать, в оригинале же 
													чтоб удобно было ответочка оканчивается на "заебал"`)
	}
	responses = append(responses, "DOTA 2", "Apex", "Овер", "Пубг", "Пираты", "Падающие Ребята", "Свояк", "Амогус", "Бля, "+name+", ты заебал. Ты сам решить не можешь?")
	response := GetRandomItem(responses)
	return response
}

func StoryTelling() string {
	responses := []string{"Бля, мне когда 5 лет было, у меня была серьезная болезнь, из-за которой " +
		"я теперь жирный. В смысле ты не веришь? ой блять долбаеб я жирный не из-за болезни, " +
		"а из-за операции, если бы мне не сделали операцию, я бы умер. Дебил хочешь у мамы моей спроси",

		"Я бы встретился с тобой и обсудил это в жизни, так как за мной тут следят, " +
			"но похуй, я сейчас попытаюсь. Все, получилось. Так вот, короче за мной теперь " +
			"охотится парень моей бывшей \\uD83D\\uDE02\\uD83D\\uDE02\\uD83D\\uDE02 Он в общем торгует " +
			"людьми и оружием \\uD83D\\uDE06 Я хуй знает, че мне с ним делать, он мне угрожает и " +
			"пытается через левых людей инфу на меня найти. Че думаешь мне делать?",

		"У меня брат уволился, сейчас работает на бирже Нью-Йорк Тайм Сквер, " +
			"знаешь такую? Ну это короче самая крутая биржа. Он взял меня к себе, я у него там " +
			"типо сись админ. бля, тупой т9, я хотел сказать сисадмин. Ну вот для меня это типо " +
			"стажировки + опыт работы",

		//need to change this one from svoyak
		"В следующий раз, когда я тебя увижу, Олег, я тебе разобью ебало, серьезно, " +
			"отвечай за слова свои. Вот докажи то, что ты ээ у тебя не получается. Давай, " +
			"alt+F10 ... Вот серьезно, ты такое ебло сидишь выебываешься, думаешь, " +
			"что тебе ничего не будет ... Ты выглядишь, как петух, говоришь, как петух блядь, " +
			"говоришь, как ебанный петух, почему нельзя быть как хоть немножечко иметь яйца," +
			" чтоб что-то говорить",

		"Я впервые в инете встретил педофила, причем такого явно, это как-то смешно " +
			"сидеть угарать над ним. ахахха. он типо написал моей девушке, я решил ради прикола " +
			"с ним поговорить, и он сейчас типо да ладно че такого, что у тебя есть парень, давай " +
			"встретимся, так далее. Я такой Ммм заебись. Бля может сказ может самому с этим челом" +
			" ради прикола встретиться"}
	return GetRandomItem(responses)
}

//should probably change nested ifs
func RandomAnek() string {
	var response string
	resp, err := http.Get("http://rzhunemogu.ru/RandJSON.aspx?CType=1")
	if err != nil {
		fmt.Println("error:", err)
		response = "Не удалось получить анек. Сервис поломался("
	} else {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("error:", err)
			fmt.Println("anek at the end", resp.Body)
			response = "Не удалось получить анек. Сервис поломался("
		} else {
			anekjson := anekJSON{}
			err = json.Unmarshal(body, &anekjson)
			if err != nil || strings.HasPrefix(anekjson.Content, "Ошибка обращения к БД") {
				fmt.Println("error:", err)
				fmt.Println("anek at the end", anekjson.Content)
				response = "Не удалось получить анек. Сервис поломался("
			} else {
				response = anekjson.Content
			}
		}
	}
	return response
}

type anekJSON struct {
	Content string `json:"content"`
}
