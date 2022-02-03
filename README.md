# Toggl
## A card service that can be used in various card games like poker and blackjack


✨ REST API to simulate a deck of cards.✨


## Tech

Toggl uses a number of open source projects to work properly:

- [Gorilla Mux](https://github.com/gorilla/mux) - A powerful HTTP router and URL matcher for building Go web servers with gorilla.

And of course Toggl itself is open source with a [public repository][robeng1]
 on GitHub.

## Design
The design follows the repository pattern or mvc 
- ```service``` : this folder houses the logic and models for the cards
- ```api```: this houses the handlers for the rest enpoints
- ```storage```: this house any storage implementations for this toggl
  Currently only in-memory implementation is provided, but others stores can be used
  simply by implementing the TogglRepository interface

You can run the tests yourself by running
```sh
  go test ./...
```


## Installation

Install the dependencies with go mod download and start the server.

```sh
go run .
```

Verify the deployment by navigating to your server address in
your preferred browser.

```sh
127.0.0.1:10000/
```

## Endpoints (Proper implemention will be to use Swagger/Open API to document the endpoints)
- ```/decks?cards=AS,KD,AC,2C,KH``` : post endpoint that creates a deck accepts an object that contains a boolean {shuffle: true}
  to indicate if the deck should be shuffled or not. Query params could be supplied to specify which cards to create by default it returns the standard 52
- ```/decks/{id}/``` : a get endpoint the returns the deck given the deck ID
- ```/decks/{id}/draw/{count}```: a get endpoint the draws a number of cards from the deck


## Development

Want to contribute? Great!



## License

MIT

**Free Software, As in Free Beer**