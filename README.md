# Gator 

An RSS feed aggregator in Go, called "Gator"

It's a CLI tool that allows users to:

*	Add RSS feeds from across the internet to be collected
*	Store the collected posts in a PostgreSQL database
*	Follow and unfollow RSS feeds that other users have added
*	View summaries of the aggregated posts in the terminal, with a link to the full post
*	RSS feeds are a way for websites to publish updates to their content. You can use this project to keep up with your favorite blogs, news sites, podcasts, and more!

This project was build with the intention of hands-on practice in GO, PostgreSQL, sqlc, goose and HTTP Requests and RSS feeds.

## Technologies Used

*   **GO**: The primary programming language.
*	**PostgreSQL**: The database tool

## Setup

To run this project, follow these steps:

    Copy the repo on your pc
    Install [GO](https://go.dev/doc/install)
    Install PostgreSQL:
        sudo apt update
        sudo apt install postgresql postgresql-contrib
        (Linux Only) sudo passwd postgres (This sets the password, use something easy like postgres)
    Create the Gator database:
        sudo service postgresql start
        sudo -u postgres psql
        CREATE DATABASE gator;
        (Linux Only) ALTER USER postgres PASSWORD 'postgres';
    Update the Gator database:
        install Goose: go install github.com/pressly/goose/v3/cmd/goose@latest
        cd sql/schema/
        goose postgres postgres://postgres:<PASSWORD>@localhost:5432/gator up
    Create your own gator config:
        nano ~/.gatorconfig.json
        Add the following json:
            {
                "db_url": "postgres://postgres:<PASSWORD>@localhost:5432/gator"
            }
    Lastely, we will build the actual program:
        go build . (while in the root of the gator project)

    Now your initial setup is complete!

## Usage

    Before using gator, make sure your postgres database is up and running:
        sudo service postgresql start
    Now you can use gator!
    See gator help for command information
    You will have to start with
        gator register <username>
        gator login <username>

# Credit
This project was completed as part of a guided course on [Boot.dev](https://www.boot.dev).
It was build following along with the [Build a Blog Aggregator in Go](https://www.boot.dev/courses/build-blog-aggregator-golang) course.