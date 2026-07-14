# Grand Exchange Tracker

![example](https://github.com/mrbrist/grandexchange-tracker/blob/main/example.png "example image")

An item price history tracker for Old School Runescape's Grand Exchange

Written in Go + React (Typescript)

## Motivation
This is one of several projects on my journey to learning the Go + Typescript tech stack and has gone through a few itterations in how the projects are structured and I landed on this structure as it is the most modular and easiest to understand

This was a project that I wanted to do because I play OldSchool Runescape and it is easier to work on a complex project when you know what success looks like

## How it works
This project is split into 3 destinct parts:

1. Frontend
2. Backend
3. Updater (cron)

The updater is the important part, without it the data will become stale it is also important for the updater to be seperate from the backend in order for them to run concurrently without causeing any issues and without producing messy code

## Quickstart
1. Clone the project
2. Setup the .env files
3. Run `make dev -j` in the root folder and `make start-updater` in another terminal window

## Usage
1. Go to the website on the port that is printed into the terminal when you start the web app
2. Begin to type the name of an item and pick from the options that show up

## Contributing
If you'd like to contribute, please fork the repository and open a pull request to the `main` branch.
