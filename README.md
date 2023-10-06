# IBLAPI

IBLAPI is a simple utility for bot/server developers (AKA *users*) of Infinity Bot List

For more information, try running `iblapi --help`

## Supported Features

### Webhook funneling

For anyone who wants to use the latest and greatest webhooks v2 (with support for review add events and more events to come (in the pipeline)) without dealing with the complex authentication and body decryption bits of it, webhook funneling is the solution.

You need `iblapi` downloaded first for your OS and CPU architecture. The executable will be named ``iblapi``

Then, run `iblapi funnel setup`, login when prompted by following the onscreen prompts and opening the OAuth2 URL when prompted. Next set a port that the webhook will listen on and the domain at which Infinity Bot List will POST to the webhook by following the onscreen instructions. 

Create the funnels for your bot and you're all set (for forwards to, provide a script that `kitescratch` should call with the webhook data in the DATA environment variable or a localhost (firewall'd if needed) server that iblcli should send the data to. Note that `kitescratch` will handle all the painful stuff like setting up the webhook config on the site using the Patch Webhook API and generating a random webhook secret so you don't have to!

Once you're ready, run `iblapi funnel start` and voila! Everything should just work!!!
