# A Golang library for the Stratum Protocol

Stratum is how Bitcoin miners connect with mining pools.

See [here](https://docs.google.com/document/d/1ocEC8OdFYrvglyXbag1yi8WoskaZoYuR5HGhwf0hWAY)
for a description of Stratum and
[here](https://github.com/slushpool/stratumprotocol/blob/master/stratum-extensions.mediawiki)
for Stratum extensions. Extensions are necessary to support ASIC Boost.

No tests have been written and no methods are supported but the overall plan
for version 1 has been laid out.

## Supported Methods

(none)

## Methods to be supported for version 1 of this library.

* mining.configure
* mining.authorize
* mining.subscribe
* mining.set_difficulty
* mining.set_version_mask
* mining.notify
* mining.submit

## Methods with no plans to be supported for version 1.

* mining.set_extranonce
* mining.suggest_difficulty
* mining.suggest_target
* mining.get_transactions
* client.get_version
* client.reconnect
* client.show_message

## Method types

Some methods are client-to-server, others are server-to-client. Some methods
require a response, others do not.

### Client-to-server

| method | type |
|--|--|
| mining.configure          | request / response |
| mining.authorize          | request / response |
| mining.subscribe          | request / response |
| mining.submit             | request / response |
| mining.suggest_difficulty |       notification |
| mining.suggest_target     |       notification |
| mining.get_transactions   | request / response |

### Server-to-client

| method | type |
|--|--|
| mining.set_difficulty     |       notification |
| mining.set_version_mask   |       notification |
| mining.notify             |       notification |
| mining.set_extranonce     |       notification |
| client.get_version        | request / response |
| client.reconnect          |       notification |
| client.show_message       |       notification |

## Message Formats

Stratum uses json. There are three message types: notification, request, and response.

### Notification

Notification is for methods that don't require a response.

```
{
  method: string,        // one of the methods above
  params: [json...]      // array of json values
}
```

### Request

Request is for methods that require a response.

```
{
  "id": integer or string, // a unique id
  "method": string,        // one of the methods above
  "params": [json...]      // array of json values
}
```

### Response

Response is the response to requests.

```
{
  "id": integer or string,   // a unique id, must be the same as on the request
  "result": json,            // could be anything
  "error": null or [
    unsigned int,            // error code
    string                   // error message
  ]
}
```

## Methods

### mining.authorize

The first message that is sent in classic Stratum. For extended Stratum,
`mining.configure` comes first and then `mining.authorize`.

### mining.subscribe

Sent by the client after `mining.authorize`.

### mining.set_difficulty

Sent by the server after responding to `mining.subscribe` and every time
the difficulty changes.

### mining.notify

Sent by the server whenever the block is updated. This happens periodically
as the block is being built and when a new block is discovered.

### mining.configure

The first message that is sent in extended Stratum. Necessary for ASICBoost.
The client sends this to tell the server what extensions it supports.

### mining.set_version_mask

Sent by the server to notify of a change in version mask. Requires the
`version-rolling` extension.

### mining.submit

Sent by the client when a new share is mined. Modified by `version-rolling`. 
