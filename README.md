# apt-grocery

The instructions below serve as reminders for us as we revisit this project periodically, rather than instructions for people unfamiliar with the project.

## Deploying to Heroku's Container Registry

1. `heroku container:push web --app apt-grocery`
1. `heroku container:release web`
1. `heroku logs -t` to see what's going on
