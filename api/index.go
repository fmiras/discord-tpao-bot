package handler
 
import (
  "fmt"
	"os"
  "net/http"
	"github.com/bwmarrin/discordgo"
)
 
func Handler(w http.ResponseWriter, r *http.Request) {
	token, ok := os.LookupEnv("DISCORD_API_TOKEN")
	if !ok {
			fmt.Println("DISCORD_API_TOKEN environment variable not set")
			return
	}

	channel, ok := os.LookupEnv("DISCORD_CHANNEL_ID")
	if !ok {
			fmt.Println("DISCORD_CHANNEL_ID environment variable not set")
			return
	}

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
			fmt.Println("Error creating Discord session:", err)
			fmt.Fprintf(w, "Error sending message: %s", err)
			return
	}

	_, err = dg.ChannelMessageSend(channel, "Hello, world!")
	if err != nil {
			fmt.Println("Error sending message:", err)
			fmt.Fprintf(w, "Error sending message: %s", err)
			return
	}

	dg.Close()

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Message sent!")
}