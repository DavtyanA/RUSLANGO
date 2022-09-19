package commands

func WhatToPlay(authorID string) string {
	var responses []string
	name := "сука"
	//maybe make IDS enums, but not sure how to handle multiple IDs bois (Руслан и Козлов идите нахуй)
	switch authorID {
	case "329339194029899776":
		return "Хеллоу ебать ты кто"
	case "333343403163123712":
		name = "Олег"
	case "395293715742064641":
		name = "Даня"
	case "434375071591563264", "911692496227139675", "796899505181032509", "395544758832988160", "313007272127102986":
		name = "Никита"
	case "328835966783717386":
		name = "Айс"
	case "395319132863856650":
		return "DOTA 2"
	case "310408199658405901":
		name = "Михей"
	case "431920924946858004":
		name = "Сандро"
	case "413729427227410442":
		name = "Давид"
	case "378566714310262785":
		name = "Эрик"
	case "516013742560116741":
		name = "Дима"
	case "359012016876421131":
		name = "Карпов"
	case "333341352404320267":
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
