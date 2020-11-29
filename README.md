# hangman

## It's hangman the game

Play hangman with your friends and enemies, all while socially distanced in your homes!

Feeling like you're going insane during this semester?? Relax with some hangman!

# How's it work?

Pretty simply.
I used Go to develop my application and also web server. I used [Gin](https://github.com/gin-gonic/gin) instead of the standard go routing. 
A MySQL databse is used to store data. Nothing crazy or mind blowing.

I am using docker-compose to deploy the application and database. An Nginx proxy sits in it's own container at the front of this application. all of this is contained in the `Dockerfile` and `docker-compose.yml`

It's assumed that this app will be running at the root of whatever you are running it on. If you need to change any proxy settings, adjust `reverse-proxy.conf` in the proxy directory. 

I don't know if anyone cares, but I deployed this on an EC2 isntance. I originally had it on a lightsail instance, but that was trash. 

# Setup

## Requirements
- Ubuntu 20.04 (This is what I developed on but in theory almost any major distro and MacOS will work)
- Docker and Docker-Compose

In order to run the hangman app, there a few things we need to do.
1. Create a `.env` file and setup [`Environment Variables`](#environment-variables)environment variables

## Environment Variables 
- **MYSQL_USER** - any username you want setup on your mysql container. Will get access  to the table you establish
- **MYSQL_PASSWORD** - set a password for the above user
- **MYSQL_DB** - name of the DB in use. See SQL for this project (feel free to change if desired)
- **MYSQL_ROOT_PASSWORD** - Change the root password
- **DOMAIN** - domain in use to use for cookie creation and registering certs. (you need an actual domain, using `localhost` or `127.0.0.1` will result in an error)
- **SALT** - a secret token for salting password hashes

# Run It

I have attempted to make running this as easy as possible. Ready?
1. `docker-compose up`

And that's it. That's hangman. Pretty good, pretty pretty good. 
