## wandb/local:0.9.34 (December 21, 2020)

* Fixed bug that caused certain metrics to not display in filters
* Fix regression in 0.9.33 that prevented instances using a mounted volume from starting up if they also specified an external storage bucket
* Files downloaded from artifacts now use their original name instead of their MD5 hash
* Fixed bug that showed duplicated graph in histogram tooltip
* Disable Google Translate, which was causing the app to crash
* Added point cloud support to dataset visualization
* Added the ability to select categorical variables in the parallel coordinates plot
* Added the ability to click on a node in the Artifact DAG to open the corresponding run or artifact page in a new tab
* Fixed bug that sometimes caused an incorrect layout when creating a new report
* Fixed tensorboard for runs where the project name has been changed.
* Fixed bug that sometimes prevented results in filterable dropdowns from updating when search text changed
* Allowed dots in config keys in the custom charts query editor
* Fixed bug where model query failed when missing an entity or project name
* Fixed bug where users who recently signed up sometimes wouldn't receive emails
* Fixed bug that prevented logging new versions of an artifact if its latest version was previously deleted
* Fixed crash in DataSignalViewer when row count is null
* Fixed crash from using forEach() in Internet Explorer
* Fixed bug that would sometimes cause Firefox to throw an error while measuring the size of rendered text
* Fixed bug preventing setting a preview image through report showcase on a user's profile
* Fixed fullscreen image sizing in the media panel
* Fixed bug where groups of runs in the runs table would randomly expand and contract
* Allow text wrapping in tagged refs in report comments
* Make the horizontal scrollbar on the media panel appear when media is wider than the panel
* Deprecate old dataframe table
* Add All and None toggle buttons to mask and bounding box controls
* Show an error message when visible column limit is reached in Runs table
* Show an error message in the browser when the container fails to start
* Show TPU usage in System Metrics tab if applicable
* Fixed bug where cursor jumped to end when editing panel bank section name


## wandb/local:0.9.33 (December 4, 2020)

This release contains fixes for bugs introduced in version 0.9.32.  Users should upgrade at their earliest convenience to ensure runs are being correctly marked as crashed and that early stopping works properly for sweeps.

* Improvements to the system settings UI
* Projects are now sorted by the time they were created by default
* The media panel now takes the current step into account when computing which columns to display
* Fixed bug that sometimes allowed users without appropriate permissions to create sweeps
* Service accounts in teams can now write runs
* Only panels included in the set of visible runs are now exported
* Added slack integration for local deployments
* Fixed bug in Artifacts that sometimes prevented deleted files from being garbage collected
* Fixed bug that sometimes caused workspaces to crash
* Improvements to the Artifacts Overview tab
* Scriptable run alerts are now enabled by default for new users
* Column controls in the Artifacts table are now in the header for each column
* Fixed bug that prevented users from creating a custom chart if they specified tableColumns
* Improvements to filter key search (keys now match on their display names)
* Added support for list/tuple values for categorical parameters in Sweeps
* Fixed bug preventing use of Artifacts with Azure storage
* Fixed bug where parallel coordinates charts wouldn't fill the screen
* Improvements to the system settings UI
* Fixed bug preventing the deletion of the newst version of an artifact
* Fixed bug where images would be cropped when opening in fullscreen view
* Report run tables with customized column configurations no longer automatically insert new columns that appear in the data
* Fixed bug where runs sometimes wouldn't highlight when hovering corresponding lines in area plots



## wandb/local:0.9.32 (November 25, 2020)

This release includes important security fixes, users should update to this release at their earliest convenience.

## Security

* Improved crypto library usage for more secure random strings
* Added HSTS headers and cookie flags to prevent CSRF attacks
* Removed local admin UI log viewer, logs can still be obtained via the debug bundle

## Features & Bugs

* Reports now support comments on specific charts or sections! https://docs.wandb.com/reports#comment-on-reports
* Beta release of Dataset visualiztion for artifacts, see: https://docs.wandb.com/datasets-and-predictions
* Custom charts are now exporable as SVG, PNG, or CSV
* New storage browser and bulk deletion UI accessible from settings pages
* New home page onboarding experience
* Disable comment @-mentions if there are no team members in the report entity
* Chart and run table CSV export improvements
* Fixed resizable run table snapping behaviour
* Hide Create Report button in jupyter notebooks
* File permissions in /vol are no longer changed for faster boot times
* Parallel Coordinates UI improvements
* Make sorting icons consistent in the runs table
* Custom chart performance improvements
* Custom charts now display all columns of all tables when several tables selected
* Fix issue where users can't load public artifact files in UI.
* Better error handling of artifacts error when users don't have write permission
* Improvements to keyset handling for projects with lots of keys.
* Stretch images when they're small, auto-join comparisons join on id.
* Better permission handling for run deletion on teams
* Better error messages for artifact type mis-matches
* Better searchable select widgets in various UI components


## wandb/local:0.9.31 (November 13, 2020)

* Added support for scriptable alerts (requires CLI v0.10.9)
* Make the minmax lines work in case where y value is zero
* Better read-only view for Vega custom charts editor
* Allow artifacts to be renamed and retyped
* API keys for service accounts are now visible
* Implemented a new Vega spec editor for custom charts
* Made aliases copyable on Artifacts Overview
* Added ability to delete artifacts by collection
* Fixed bug that prevented the copy button from working in the code tab in artifacts
* Security improvements for the local container environment (added HSTS header in SSL mode, removed admin log viewer)
* Adding a panel to a section now opens that section
* Added option to continue draft when adding panels to a report
* Fixed step slider not updating the image panel
* Fixed scrolling overflow issues for Artifact tables
* Fixed error when user has duplicate invites
* Added runtime to parameter importance panel
* Added sidebar for artifact comparisons
* Added button to insert text block in reports
* Added batch deletion for sweeps
* Forced all new local users to set a username.
* Added support for aggregation queries in the UI using the MongoDB query language
* Preserve color/mark/title overrides when exporting panels from run page
* Fixed several bugs in the parallel coordinates column
* Added initial support for high-availability deployments (multiple nodes)
* Improved UI around bounding boxes
* Improved performance for custom charts
* Add instance settings to the debug bundle
* Fixed a crash that occurred when deleting runs
* Fixed bug where panels would constantly re-render
* Fixed bug where changing tags would reset the tags' sort order
* Added the "Samples" grouping option to line plots, which displays the underlying lines combined to form a grouped line
* Fixed crash when attempting to show data from a private project in a report (an error now displays in the panel instead)


## wandb/local:0.9.30 (October 14, 2020)

* Enabled exporting panels to different report sections
* Fixed artifact header names
* New [custom charts](https://wandb.ai/wandb/posts/reports/The-W-B-Machine-Learning-Visualization-IDE--VmlldzoyNjk3Nzg)!
* We now show all x-axis labels in charts
* Fixed sporadic image not available errors in media panels when runs are in progress
* Better graphql type checking for api requests
* The table now supports sorting by multiple columns!
* Sweeps now have an early-stopping column when enabled
* The collapsed state of groups is now persisted across page loads
* Line plot CSV export now contains min and max values for grouped metrics
* Searching panels is more performant
* Added bounding box artifact visualization
* New storage stats table backend for exposing storage usage, UI coming in the next release
* Properly escape file paths when using a Minio storage backend
* Ensure the current user maintains admin status when connecting an external MySQL store


## wandb/local:0.9.29 (October 6, 2020)

* Fixed bug that prevented some users from logging in when using Auth0 authentication
* Fixed bug where elements in SVG exports would render twice
* Fixed various bugs in custom visualization system
* Hide popups in custom visualization editor when the page is scrolled
* Various frontend performance improvements
* Warn on slow queries in custom visualization system
* Fixed bug where some data would still show after unselecting all runs
* Performance improvements for reports
* Add ability to choose where in a report a panel will land when exporting from a workspace


## wandb/local:0.9.28 (October 1, 2020)

* In the browser, hide the full API key when authorizing a new client or viewing keys on settings page
* Fixed a rare bug causing inaccurate aggregations on config and summary values
* Added the ability to send notifications to an external service
* Fixed an infinite update loop that could consume excessive resources when using S3
* Improved UI performance when attaching notes to a runset
* Improved performance on some run queries
* Added support for Markdown links with arbitrary text
* Fixed a bug that caused some runs to be excluded from the parallel coordinates chart
* Added close button to fullscreen view for panels


## wandb/local:0.9.27 (September 24, 2020)

* In history panels, added ability to group runs over metrics.
* In history panels, improved usability axes.
* Fixed an issue where certain types of logged tables caused a page crash.
* Improved frontend performance in heavy workspaces and reports.
* Added backend support for custom panels logged via CLI (CLI support pending).
* Fixed a bug with artifact reference tracking.
* Improved run metric cache behavior.
* Fixed a bug with user management.
* Removed cookies from logs
* Fixed restart command and version upgrade logic in the settings app


## wandb/local:0.9.26 (September 17, 2020)

* Fixed panic during artifact uploads
* Fixed history scan panic
* Supports v2 artifacts which are created by 0.10.X python clients
* Support rendering dev and rc versions of python clients in the UI
* Custom chart UI improvments
* Support default region for artifacts
* Fixed artifact download links from python clients
* Added name to the report creation dialogue
* Improved frontend error message design
* General plot improvements
* Add artifact aliases from the UI
* Custom charts support run colors
* Remove chart options on the individual run page


## wandb/local:0.9.25 (September 11, 2020)

* Support for client version querying
* Redacted sensitive environment variables from logs
* Images render in Jupyter notebook previews
* Prevented metadata update loops for external file storage configurations
* Custom chart configuration improvements
* Support for dynamic hosts to serve the api and UI from different urls
* Artifact reference linking
* Move panels between grids in reports
* JobType can be added to run tables
* CSV Export contains only visible columns
* Artifacts support pagination
* Sweep config improvements
* Auth0 library updates


## wandb/local:0.9.24 (September 3, 2020)

* New Artifacts backend changes
* Private report link sharing
* Improved report comments
* Better support for ENV var settings
* Custom Panels with new graphql query UI's


## wandb/local:0.9.23 (August 4, 2020)

* Google cloud external buckets 
* Settings are not over written when clicking "Set frontend host"
* Frontend console logs in debug bundles
* Reporting comments and general improvements
* Private report sharing


## wandb/local:0.9.22 (July 27, 2020)

* Admin interface works properly in Safari
* Handle crashing expression editor
* Improved team / org management
* Report comments and images in markdown

> NOTE: There's a large data migration in this release that migrates files from an int to bigint primary key to allow for more file records.  This migration takes ~5 seconds per 1 million file records so could result in some downtime during the rollover.


## wandb/local:0.9.21 (July 7, 2020)

* s3 file store updates would loop infinitely from SQS events
* error logs weren't being written to disk
* Multiple agents caused sweeps to crash


## wandb/local:0.9.20 (July 6, 2020)

* This release primarily fixes an issue with users upgrading from v0.9.16.


## wandb/local:0.9.19 (July 1, 2020)

* Artifact UI improvements including deletion and multi-zoned buckets
* Artifact downloads now work properly in local
* No more onboarding UI
* Performance improvements for Vega rendering
* Report interface polish and improvements
* Charting improvements and polish


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


