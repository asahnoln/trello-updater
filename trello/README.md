# Trello package

This package presents types and functions to work with lists and cards.

## Testing

Required flags for testing:

Option | Type | Description
-----: | :--- | :----------
`-url` | string | url pointing to Trello API or other mock API, so that tests are run against it.
`-boardId` | string | board ID in Trello or mock API to get lists and cards from. It should be string as trello ids are string.

Example:

`go test -url "https://mymockapi.test/api/1" -boardId "1"`

### Work with mock API

In mock API, you have to create next endpoints:

Endpoint | Description
-------: | :----------
`/boards/:boardId/lists` | Get lists from the given board
`/lists/:listId/cards` | Get cards from the given list
