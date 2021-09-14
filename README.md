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

| mining.configure          | request / response |
| mining.authorize          | request / response |
| mining.subscribe          | request / response |
| mining.submit             | request / response |
| mining.suggest_difficulty |       notification |
| mining.suggest_target     |       notification |
| mining.get_transactions   | request / response |

### Server-to-client

| mining.notify             |       notification |
| mining.set_difficulty     |       notification |
| mining.set_version_mask   |       notification |
| mining.set_extranonce     |       notification |
| client.get_version        | request / response |
| client.reconnect          |       notification |
| client.show_message       |       notification |

## Message Formats

### Notification

### Request

### Response

## Methods

### mining.configure

### mining.authorize

### mining.subscribe

### mining.set_difficulty

### mining.set_version_mask

### mining.notify

### mining.submit
