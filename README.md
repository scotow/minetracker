# minetracker

### What?

*minetracker* is a [Minecraft](https://www.minecraft.net) tracker library that use the [RCON](https://wiki.vg/RCON) protocol implemented by the server to communicate, run command and parse the results.

### Why?

While playing a 1.14 lightly modded Skyblock game, we needed to know when the [Wandering Trader](https://minecraft.gamepedia.com/Wandering_Trader) did spawn.

An other thing we needed was to know when someone did connect.

### How?

This library is based on three interfaces:


#### Runner

##### Role

A `Runner` is a way to contact/run a RCON command on the server.

##### Implementations

The library already provides two ways to contact the server, but you can implement your own if you want. The first implementation uses an external RCON library to connect directly to the RCON server. The other implementation uses the `mcrcon` command and the `os/exec` package to build and run this command. So, the `mcrcon` needs to be installed first if you want to use this `Runner`.

#### Tracker

##### Role

The second interface, `Tracker`, is used to build the command string that will be run on the server, parse the result and decide if a notification should be send. The `Tracker` also has to provide an retry/re-run interval.

##### Implementations

The library already provides two trackers.

The first one tracks online players every *n* seconds, and uses the last result to check if someone did disconnect or connect.

The second tracker checks if entity is present in the game, or more precisely, checks if it has spawn (it will not trigger a notification if the entity is present two checks in a row).

#### Notifier

##### Role

The last interface, `Notifier`, is used to notify the finds of a tracker.

##### Implementations

The library provides a [notigo](https://github.com/scotow/notigo), an other library that I made, that use IFTTT Webhooks to send a Push Notification on mobile. For more information check the `notigo` repository.

A [Discord](https://discordapp.com) `Notifier` is also available. This tracker sends a message on Discord server channel. You will need to create, setup and get your [Discord Developer Credentials](https://discordapp.com/developers) for that.

There also some `Notifier` helpers, that allow you to chain `Notifier`s (called `MultiNotifier`). 

### Examples

The [cmd](https://github.com/scotow/minetracker/tree/master/cmd) directory contains some examples that I used on my server. It's mainly a combination of all the implemented `Tracker` and `Notifier` explained above.

For for, check the documentation.

### Disclaimer

*minetracker* provided by *Scotow* is for illustrative purposes only which provides customers with programming information regarding the products. This software is supplied "AS IS" without any warranties and support.

I assumes no responsibility or liability for the use of the software, conveys no license or title under any patent, copyright, or mask work right to the product.

***Enjoy tracking while playing!***