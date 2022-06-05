package main

import (
	"fmt"
	"github.com/SevereCloud/vksdk/v2/marusia"
	"net/http"
	"strings"
)

func main() {
	wh := marusia.NewWebhook()
	wh.EnableDebuging()

	question := 0
	category := Categories{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	wh.OnEvent(func(r marusia.Request) (resp marusia.Response) {
		if r.Request.Type == marusia.SimpleUtterance {
			fmt.Println(r.Request.Command)
			// switch r.Request.Command {
			if marusia.OnStart == r.Request.Command {
				resp.Text = "Как называется твоя команда?"
				resp.TTS = "Как называется твоя команда?"
			} else if strings.Contains(r.Request.Command, "вездекод") || strings.Contains(r.Request.Command, "вездеход") {
				resp.Text = "Привет вездекодерам!"
				resp.TTS = "Привет вездекодерам!"
				if question < 1 {
					resp.Text = "Привет вездекодерам! Отлично тогда я задам 8 вопросов чтобы помоч тебе с выбором категории." + " 1 Вопрос: кто спроектировал ядро линукс?"
					resp.TTS = "Привет вездекодерам!" + "Первый вопрос кто спроектировал ядро линукс?"
					resp.AddButton("не знаю", myPayload{Index: question})
					question++
					resp.Card = marusia.NewMiniApp("https://vk.com/app7543093")

				}
			} else if r.Request.Command == "посмотреть mini apps" {
				resp.Card = marusia.NewMiniApp("https://vk.com/app7543093")
			} else if message, err := Question(&question, r.Request.Command, &category); !err {
				if question <= 8 {
					fmt.Println(category, question)
					resp.Text = message
					resp.TTS = message
					resp.AddButton("не знаю", nil)
				} else if question > 8 {
					resp.Text = message + " Вам подходит категория: " + category.MaxCategory()
					marusia.SpeakerAudioVKID("-2000512004_456239026")
					fmt.Println(category.MaxCategory())
					cat := [9]struct {
						name string
						id   int
					}{
						{"Security", 457239021},
						{"Design", 457239020},
						{"Bot", 457239019},
						{"BackEnd", 457239022},
						{"JS", 457239017},
						{"XR", 457239023},
						{"Mobile", 457239025},
						{"Marusia", 457239018},
						{"Sport", 457239026},
					}
					var id int
					for _, val := range cat {
						if val.name == category.MaxCategory() {
							id = val.id
							break
						}
					}
					resp.Card = marusia.NewBigImage(category.MaxCategory(), "Привет", id)
					resp.AddButton("Посмотреть Mini Apps", nil)
				}

			} else if marusia.OnInterrupt == r.Request.Command {
				resp.Text = "Скилл закрыт"
				resp.TTS = "Пока"
				resp.EndSession = true
			} else {
				resp.Text = "Неизвестная команда"
				resp.TTS = "Я вас не поняла"
			}
		}

		return
	})

	http.HandleFunc("/", wh.HandleFunc)

	http.ListenAndServe(":8080", nil)
}

// case "картинка":
// 	resp.Card = marusia.NewBigImage("Заголовок", "Описание", 457239017)
// case "кнопки":
// 	resp.Text = "Держи кнопки"
// 	resp.TTS = "Жми на кнопки"
// 	resp.AddURL("ссылка", "https://vk.com")
// 	resp.AddButton("подсказка без нагрузки", nil)
// 	resp.AddButton("подсказка с нагрузкой", myPayload{
// 		Text: "test",
// 	})
