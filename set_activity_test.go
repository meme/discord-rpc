package discord_rpc

import (
	"log"
	"time"
	"testing"
)

const clientId = ""

func TestClient_SetActivity(t *testing.T) {
	client := Client{}
	err := client.Connect(clientId)

	if err != nil {
		t.Fatal(err)
	}

	err = client.SetActivity("Example", "Example", "default", "Example")

	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(10 * time.Second)

	err = client.ClearActivity()

	if err != nil {
		t.Fatal(err)
	}
}

func ExampleClient_SetActivity() {
	client := Client{}
	err := client.Connect(clientId)

	if err != nil {
		log.Fatal(err)
	}

	err = client.SetActivity("Example", "Example", "default", "Example")

	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(10 * time.Second)

	err = client.ClearActivity()

	if err != nil {
		log.Fatal(err)
	}
}