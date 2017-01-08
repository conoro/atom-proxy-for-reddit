# Reddit Proxy
This tool lets you request specific subreddits via RSS and generates its own RSS feeds for them

It exists to deal with Feedly being blocked for (presumably) excessive requests to Reddit

V1 is completely hard-coded for me but could easily be made configurable.

Next step is to get running on AWS Lambda

Note I have vendored-in gofeed so I can add a needed User-Agent for the requests.

A simple reverse proxy might do the same job but this was quicker for me as I had previous code

Copyright © 2017 Conor O'Neill, conor@conoroneill.com

License MIT