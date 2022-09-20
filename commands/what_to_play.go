package commands

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
	case mk7k_ID:
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
