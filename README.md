# apt-grocery

The instructions below serve as reminders for us as we revisit this project periodically, rather than instructions for people unfamiliar with the project.

## Deploying to Heroku's Container Registry

1. `heroku container:push web --app apt-grocery`
1. `heroku container:release web`
1. `heroku logs -t` to see what's going on

## Production Environment Configurations

apt-grocery's production environment needs to have the following environment variables set:

* `BOT_ID`: secret GroupMe bot ID
    * Given to you when you create a bot at [https://dev.groupme.com/bots/new](https://dev.groupme.com/bots/new)
* `RUN_ENV`: set to `prod`
    * This determines the path where `list.json` is expected to be

# To-Do List

## New Features

* ~~Store who adds which items to the apartment-wide list~~ ✅
    * ~~So whoever runs to the store can bill each person for what they asked for~~
* Store prices for commonly-requested items
    * Once this collection of prices gets large, we can make a rough calculation of the total grocery bill before anyone goes to the store
