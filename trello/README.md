# Trello package

This package presents types and functions to work with lists and cards.

## Testing

Required flags for testing:

Option | Description
-----: | :----------
`-url` | url pointing to Trello API or other mock API, so that tests are run against it.
`-boardId` | board ID in Trello or mock API to get lists and cards from.

Example:

`go test -url "https://mymockapi.test/api/1"`

### Work with mock API

In mock API, you have to create next endpoints:

`/boards/:boardId/lists` - get lists from the given board
`/lists/:listId/cards` - get cards from the given list
