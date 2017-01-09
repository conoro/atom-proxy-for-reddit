# Reddit Proxy

This tool lets you request specific subreddits via RSS and generates its own RSS feed for them

Just request https://yourdomain.com/?r=name_of_subreddit to get a feed for https://www.reddit.com/r/name_of_subreddit

The tool exists to deal with Feedly being blocked for (presumably) excessive requests to Reddit

Next step is to get running on AWS Lambda

Note I have vendored-in the rss module so I can add a needed User-Agent for the requests.

A simple reverse proxy might do the same job but this was quicker for me as I had previous code

Note I run this on an EC2 instance behind a Caddy server which gave me a Let's Encrypt Cert for free

Copyright © 2017 Conor O'Neill, conor@conoroneill.com

License MIT