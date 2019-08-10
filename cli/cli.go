package cli

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/staubichsauger/uno-cli/game"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Client struct {
	Url url.URL
	Id string
}

func (a *Client) Login() error {

	prompt := promptui.Prompt{
		Label:    "Your name",
	}

	result, err := prompt.Run()

	if err != nil {
		return errors.New("Prompt failed: " + err.Error())
	}

	fmt.Printf("You chose %s\n", result)

	me := game.Join{
		Name: result,
	}
	reqBytes, err := json.Marshal(&me)
	if err != nil {
		return err
	}

	res, err := http.Post(a.Url.String() + "/join", "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return errors.New("Got status: " + strconv.Itoa(res.StatusCode))
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	id := game.Id{}
	err = json.Unmarshal(body, &id)
	if err != nil {
		return err
	}

	a.Id = id.PlayerId

	fmt.Println("Successfully logged into the game.")
	return nil
}

func (a *Client) Play(stop chan error) {
	for _ = range time.Tick(time.Millisecond * 40) {
		res, err := http.Get(a.Url.String() + "/games?id=" + a.Id)
		if err != nil {
			stop <- err
			return
		}
		if res.StatusCode != http.StatusOK {
			stop <- errors.New("Error getting games endpoint, Status: " + strconv.Itoa(res.StatusCode))
			continue
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			stop <- err
			return
		}

		gs := game.Status{}
		err = json.Unmarshal(body, &gs)
		if err != nil {
			stop <- err
			return
		}

		if gs.MyTurn {
			turn := game.Turn{}

			fmt.Println("The top card is: " + gs.DiscardedCard.Value + "-" + gs.DiscardedCard.Color)

			prompt := promptui.Select{
				Label: "Select a card to play",
				Items: gs.GetCards(),
				Size: 10,
			}

			_, result, err := prompt.Run()

			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

			fmt.Printf("You choose %q\n", result)

			turn.PlayCard = gs.GetCard(result)

			if turn.PlayCard != nil && strings.Contains(turn.PlayCard.Value, "WILD") {
				prompt = promptui.Select{
					Label: "Select a color",
					Items: []string{"RED", "BLUE", "GREEN", "YELLOW"},
					Size: 4,
				}
				_, result, err = prompt.Run()

				turn.PlayCard.Color = result

				fmt.Printf("You choose %q\n", result)
			}

			bj, err := json.Marshal(&turn)
			if err != nil {
				stop <- err
				return
			}
			res, err := http.Post(a.Url.String() + "/games?id=" + a.Id, "application/json", bytes.NewBuffer(bj))
			if err != nil {
				stop <- err
				return
			}
			if res.StatusCode != http.StatusOK {
				//log.Print(*turn.PlayCard)
				body, err := ioutil.ReadAll(res.Body)
				if err != nil {
					stop <- err
					return
				}
				stop <- errors.New("Error posting turn: " + strconv.Itoa(res.StatusCode) + "-> " + string(body))
			}
		}
	}
}