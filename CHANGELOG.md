## wandb/local:0.9.18 (June 12, 2020)

* New advanced Bar Chart options for std-dev, and error bars
* Restore deleted users via the user admin api
* Configure local instances programmatically using environment variables.
* Fix wandb.HTML rendering issues inside of wandb/local
* Fix flagging users as admin when reconfiguring the database
* Report UI improvements


## wandb/local:0.9.17 (June 5, 2020)

* Forgot password emails and password reset link generation
* Support for artifacts!


## wandb/local:0.9.16 (May 26, 2020)

* Emails now sent directly from local
* New frontend host settings to ensure links in emails are valid
* No longer depend on external google fonts
* Improved testing / QA process


## wandb/local:0.9.15 (May 12, 2020)

* Fixes initial admin user flagging
* Increases the default session length to 1 week.
* Code saving is enabled by default.


## wandb/local:0.9.14 (May 7, 2020)

* New user management UI
* admins only get access to management UI's
* Benchmarks enabled


## wandb/local:0.9.13 (April 21, 2020)

* The api properly restarts when settings change
* Improvements to report saving
* W&B now supports wandb.scikit and wandb.plots
* fixes an issues with local instances deployed to DNS with `app.` in the url


## wandb/local:0.9.12 (April 14, 2020)

* We now run the container as non-root by default
* OpenShift usage now supported.
* We've also enable the ability to record statsd information locally.


## wandb/local:0.9.11 (April 3, 2020)

* Fixes panics in Azure
* Makes images uploads work in azure
* New file metadata service


## wandb/local:0.9.10 (March 24, 2020)

* Fixes new user emails
* Adds caching to frontend assets
* Don't load Auth0 when not needed


## wandb/local:0.9.9 (March 20, 2020)

* Fixed user invites
* Fixed version check
* Pretty login page
* Test improvements
* The last 17 days of product improvements.


## wandb/local:0.9.8 (March 3, 2020)

* Ensures local works offline
* Fixes minio redis event streaming
* Disables teams with no license


## wandb/local:0.9.7 (February 26, 2020)

* Fixes cron jobs


## wandb/local:0.9.6 (February 24, 2020)

* Fixes sweep server
* UI Updates


## wandb/local:0.9.5 (February 20, 2020)

* New loading screen
* Better validation for auth0 setup
* New more robust startup script
* Fixed cron jobs
* Auto set tls=preferred for external mysql


## wandb/local:0.9.4 (February 13, 2020)

* First official release from the new build system
* Fixes auth0 integration
* Fixes azure mysql setup
* Doesn't bind to port 80

