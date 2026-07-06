<!-- ![example](https://github.com/mrbrist/grandexcahneg-tracker/example.png "Logo Title Text 1") -->

# Grand Exchange Tracker

An item price history tracker for Old School Runescape's Grand Exchange

Written in Go + React

## How it works

This project is split into 3 destinct parts:

1. Frontend
2. Backend
3. Updater (cron)

The updater is the important part, without it the data will become stale it is also important for the updater to be seperate from the backend in order for them to run concurrently without causeing any issues and without producing messy code

## How to run it

1. Clone the project
2. Setup the .env files
3. Run `make dev -j` in the root folder and `make start-updater` in another terminal window
