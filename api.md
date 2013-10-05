API
===

Start Game
----------
```
GET /api/game/new/:game_id
```

Error response:

```
{
	"Error": "Error message."
}
```

Successs response:

```
{
	"GameId": "Game id",
	"PlayerPublicKey": "PublicKey",
	"PlayerPrivateKey": "PrivateKey"
}
```

Join Game
---------
```
POST /api/game/join/:game_id
```

Error response:

```
{
	"Error": "Error message."
}
```

Success response:

```
{
	"GameId": "Game id",
	"PlayerPublicKey": "PublicKey",
	"PlayerPrivateKey": "PrivateKey"
}
```

Play Game
---------
```
WS /api/game/play/:game_id/player/:public_key/:private_key
```