# [@TrendingGitlab](https://twitter.com/TrendingGitlab)

[![Build Status](https://travis-ci.org/whitef0x0/TrendingGitlab.svg?branch=master)](https://travis-ci.org/whitef0x0/TrendingGitlab)
[![GoDoc](https://godoc.org/github.com/whitef0x0/TrendingGitlab?status.svg)](https://godoc.org/github.com/whitef0x0/TrendingGitlab)
[![Coverage Status](https://coveralls.io/repos/whitef0x0/TrendingGitlab/badge.svg?branch=master&service=github)](https://coveralls.io/github/whitef0x0/TrendingGitlab?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/whitef0x0/TrendingGitlab)](https://goreportcard.com/report/github.com/whitef0x0/TrendingGitlab)

A twitter bot (**[@TrendingGitlab](https://twitter.com/TrendingGitlab)**) to tweet [trending repositories](https://gitlab.com/explore/projects/trending) from GitLab.

> Follow us at **[@TrendingGitlab](https://twitter.com/TrendingGitlab)**.

**Important:** This is not an official GitLab or Twitter product.

## Features

* Tweets trending projects every 30 minutes
* Refreshes the configuration of twitters URL shortener t.co every 24 hours
* Greylisting of repositories for 30 days (to avoid tweeting a project multiple times in a short timeframe)
* Maximum use of 140 chars per tweet to fill up with information
* Debug / development mode
* Multiple storage backends (currently [Redis](http://redis.io/) and in memory)

## Installation

1. Download the [latest release](https://github.com/whitef0x0/TrendingGitlab/releases/latest)
2. Extract the archive (zip / tar.gz)
3. Start the bot via `./TrendingGitlab -debug`

For linux this can look like:

```sh
curl -L  https://github.com/whitef0x0/TrendingGitlab/releases/download/v0.4.0/TrendingGitlab-v0.4.0-linux-amd64.tar.gz -o TrendingGitlab-v0.4.0-linux-amd64.tar.gz
tar xzvf TrendingGitlab-v0.4.0-linux-amd64.tar.gz
cd TrendingGitlab-v0.4.0-linux-amd64
./TrendingGitlab -debug
```

## Usage

```
$ ./TrendingGitlab -help
Usage of ./TrendingGitlab:
  -debug
    	Outputs the tweet instead of tweet it (useful for development). Env var: TrendingGitlab_DEBUG
  -expvar-port int
    	Port which will be used for the expvar TCP server. Env var: TrendingGitlab_EXPVAR_PORT (default 8123)
  -storage-auth string
    	Storage Auth (e.g. myPassword or <empty>). Env var: TrendingGitlab_STORAGE_AUTH
  -storage-url string
    	Storage URL (e.g. 1.2.3.4:6379 or :6379). Env var: TrendingGitlab_STORAGE_URL (default ":6379")
  -twitter-access-token string
    	Twitter-API: Access token. Env var: TrendingGitlab_TWITTER_ACCESS_TOKEN
  -twitter-access-token-secret string
    	Twitter-API: Access token secret. Env var: TrendingGitlab_TWITTER_ACCESS_TOKEN_SECRET
  -twitter-conf-refresh-time duration
    	Twitter: Time interval to refresh the configuration of twitter (e.g. char length for short url). Env var: TrendingGitlab_TWITTER_CONF_REFRESH_TIME (default 24h0m0s)
  -twitter-consumer-key string
    	Twitter-API: Consumer key. Env var: TrendingGitlab_TWITTER_CONSUMER_KEY
  -twitter-consumer-secret string
    	Twitter-API: Consumer secret. Env var: TrendingGitlab_TWITTER_CONSUMER_SECRET
  -twitter-follow-new-person
    	Twitter: Follows a friend of one of our followers. Env var: TrendingGitlab_TWITTER_FOLLOW_NEW_PERSON
  -twitter-follow-new-person-time duration
    	Growth hack: Time interval to search for a new person to follow. Env var: TrendingGitlab_TWITTER_FOLLOW_NEW_PERSON_TIME (default 45m0s)
  -twitter-tweet-time duration
    	Twitter: Time interval to search a new project and tweet it. Env var: TrendingGitlab_TWITTER_TWEET_TIME (default 30m0s)
  -version
    	Outputs the version number and exit. Env var: TrendingGitlab_VERSION
```

**Every parameter can be set by environment variable as well.**

**Twitter-API settings** (`twitter-access-token`, `twitter-access-token-secret`, `twitter-consumer-key` and `twitter-consumer-secret`) are necessary to use the Twitter API and to set up a tweet by your application.
You can get those settings by [Twitter's application management](https://apps.twitter.com/).

If you want to play around or develop this bot, use the `debug` setting.
It avoids using the Twitter API for tweet purposes and outputs the tweet on stdout.

The Redis url (`storage-url`)is the address of the Redis server in format *ip:port* (e.g. *192.168.0.12:6379*).
If your server is running on localhost you can use *:6379* as a shortcut.
`storage-auth` is the authentication string necessary for your Redis server if you use the [Authentication feature](http://redis.io/topics/security#authentication-feature).

## Storage backends

Why is a storage backend needed at all?

We are looking for popular projects in a regular interval.
To avoid tweeting a project or developer multiple times after another we add those records to a blacklist for a specific time.

At the moment there are two backends implemented:

* Memory (used in development)
* Redis (used in production)

## Growth hack

We implemented a small growth hack to get a few followers.
This hack was suggested by my colleague [@mre](https://github.com/mre).
It works like described:

* Get all followers from [@TrendingGitlab](https://twitter.com/TrendingGitlab)
* Choose a random one and get the followers of the choosen person
* Check if this person follows us already
* If yes, repeat
* If no, follow this person

This feature can be activated via the `twitter-follow-new-person` flag.

## Motivation

I love to discover new tools, new projects, new languages, new coding best practices, new exciting ideas and new people who share the same passion like me.
[I use twitter a lot](https://twitter.com/magic_cacti) and wanted to have a place to check [trending repositories](https://gitlab.com/explore/projects/trending) on GitLab as there seemed to be no bot or service currently doing this.

## License

This project is released under the terms of the [MIT license](http://en.wikipedia.org/wiki/MIT_License).
