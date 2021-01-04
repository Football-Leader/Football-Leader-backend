package players

type Player struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// TODO: Change this when connect to database
var DefaultPlayers = []*Player{
	&Player{Id: "0", Name: "Миша Блэк Бразе"},
	&Player{Id: "1", Name: "Миксер"},
	&Player{Id: "2", Name: "Темп"},
	&Player{Id: "3", Name: "Серёга Юниор Худой"},
	&Player{Id: "4", Name: "Витя Маркс"},
	&Player{Id: "5", Name: "Костя"},
	&Player{Id: "6", Name: "Виталя"},
	&Player{Id: "7", Name: "Железный"},
	&Player{Id: "8", Name: "Хомячок"},
	&Player{Id: "9", Name: "Волшебник"},
	&Player{Id: "10", Name: "Пиф Паф"},
	&Player{Id: "11", Name: "Валера"},
	&Player{Id: "12", Name: "Вовати"},
	&Player{Id: "13", Name: "Вова Энгельс"},
	&Player{Id: "14", Name: "Серёга Юниор Полный"},
	&Player{Id: "15", Name: "Омар"},
	&Player{Id: "16", Name: "Буффон"},
	&Player{Id: "17", Name: "Лёха Соренто"},
	&Player{Id: "18", Name: "Юрка"},
	&Player{Id: "19", Name: "Рома"},
	&Player{Id: "20", Name: "Колян"},
	&Player{Id: "21", Name: "Динамо"},
	&Player{Id: "22", Name: "Шнур"},
	&Player{Id: "23", Name: "Дима от Миши"},
	&Player{Id: "24", Name: "Саня Орёл"},
	&Player{Id: "25", Name: "Андрей Воротчик"},
	&Player{Id: "26", Name: "Евгений Анатольевич"},
	&Player{Id: "27", Name: "Амарок"},
	&Player{Id: "28", Name: "Артём"},
	&Player{Id: "29", Name: "Саня Малыш"},
	&Player{Id: "30", Name: "Сосед"},
	&Player{Id: "31", Name: "Диман от Волшебника"},
	&Player{Id: "32", Name: "Илюха Вязкий"},
}
