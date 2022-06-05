package main

func Question(q *int, c string, category *Categories) (message string, err bool) {
	switch *q {
	case 1:
		if c == "линус торвальдс" {
			category.BackEnd++
			category.Security++
			*q++
			return "Отлично! Вопрос №2: язык программирования на логотипе которого чашка кофе", false
		} else if c == "не знаю" {
			*q++
			return "Ответ линус торвальдс Вопрос №2: язык программирования на логотипе которого чашка кофе", false
		} else {
			return "Не верно!", false
		}
	case 2:
		if c == "java" {
			category.BackEnd++
			category.Mobile++
			category.Sport++
			*q++
			return "Верно! Вопрос №3: в какой компании изобрели язык Go/Golang?", false
		} else if c == "не знаю" {
			*q++
			return "Ответ java Вопрос №3: в какой компании изобрели язык Go/Golang?", false
		} else {
			return "Не верно!", false
		}
	case 3:
		if c == "google" {
			category.Bot++
			category.JS++
			*q++
			return "Верно! Вопрос №4: как можно сократить программный интерфейс приложения (3 латинские буквы)?", false
		} else if c == "не знаю" {
			*q++
			return "Ответ google Вопрос №4: как можно сократить программный интерфейс приложения (3 латинские буквы)?", false
		} else {
			return "Не верно!", false
		}
	case 4:
		if c == "api" {
			category.BackEnd++
			category.Bot++
			category.JS++
			category.Marusia++
			category.Mobile++
			category.Security++
			*q++
			return "Верно! Вопрос №5: аналог Figma у Adobe", false
		} else if c == "не знаю" {
			*q++
			return "Ответ api Вопрос №5: аналог Figma у Adobe", false
		} else {
			return "Не верно!", false
		}
	case 5:
		if c == "adobe xd" {
			category.Design += 3
			category.Marusia++
			*q++
			return "Верно! Вопрос №6: какой индекс у первого элемента в массиве Python", false
		} else if c == "не знаю" {
			*q++
			return "Ответ adobe xd Вопрос №6: какой индекс у первого элемента в массиве Python", false
		} else {
			return "Не верно!", false
		}
	case 6:
		if c == "0" {
			category.BackEnd++
			category.XR++
			category.Sport += 2
			*q++
			return "Верно! Вопрос №7: поисковая система разработанная в России?", false
		} else if c == "не знаю" {
			*q++
			return "Ответ 0 Вопрос №7: поисковая система разработанная в России?", false
		} else {
			return "Не верно!", false
		}
	case 7:
		if c == "яндекс" {
			category.XR++
			category.Design++
			*q++
			return "Верно! Вопрос №8: как звали первого программиста в истории?", false
		} else if c == "не знаю" {
			*q++
			return "Ответ яндекс Вопрос №8: как звали первого программиста в истории?", false
		} else {
			return "Не верно!", false
		}
	case 8:
		if c == "ада лавлейс" {
			category.XR++
			category.Security++
			category.BackEnd++
			category.Bot++
			*q++
			return "Верно! Тест окончен!", false
		} else if c == "не знаю" {
			*q++
			return "Ответ ада лавлейс Тест окончен!", false
		} else {
			return "Не верно!", false
		}
	default:
		return "Что-то пошло не так", true
	}

}
