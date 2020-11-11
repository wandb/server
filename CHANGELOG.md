## wandb/local:0.9.31 (November 10, 2020)

* SEO: run page noindex adjustment
* Make aliases copyable on Artifacts Overview (#6053)
* [WB-3621] Frontend for artifact collection deletion (#6032)
* Fixed Map Initialization (#6049)
* Upgrade lyft to 1.8 (#6042)
* Bump max cli version (#6048)
* Deduped Artifact Reference Limit Counter (#6046)
* [WB-3818] Artifacts overview - fix code copying, made version copyable (#6045)
* Set http-secure header in local (#6040)
* Remove tailon service from local (#6037)
* gorilla: lower max run config to 15 MB (#6044)
* gorilla: move away from linking to types directly within artifacts (#6039)
* Simple fix for suren (#6034)
* fix filtering (#6038)
* adding panel to panelbank section opens it
* update gallery webinar
* add panels to report: option to continue draft
* update comment notification email (#6031)
* fix crash when workspace has grouped runs with histogram-uberfile
* WB-3825 fix image panel step slider (#6029)
* Fixed Scrolling Overflow Issues for Artifact Tables (#6030)
* [WB-3824] update the links in UI (#6027)
* Bump max cli version (#6028)
* Fixes Dataviz Group UI (#6026)
* signup: fix error when user has duplicate invites (#6025)
* Check if a vega spec uses input bindings more thoroughly (#6024)
* gorilla: system metrics sampling returns keys in correct order (#6022)
* gorilla: disallow duplicate alerts per entity (#6021)
* gorilla: fix FetchArtifactDependenciesByArtifact query (#6018)
* Artifacts/tables and panels (#6020)
* moved time initialization and increased threshold (#6019)
* Lots of dataset & prediction viz UI updates (#6014)
* gorilla: fix run storage stats cleaner query (#6013)
* Modified Artifact Encoding Strategy (#6017)
* gorilla: fix FetchControllerRunNameBySweep (#6016)
* [CLI-476] Add runtime to parameter importance panel (#6004)
* also delete stars when deleting views
* Click image to show original size in modal (#5622)
* gorilla: autogenerate scaffolding for mysql fetch methods (#6007)
* Artifact Dependencies: Adding support for artifact dependencies as it relates to dataviz (#5998)
* gorilla: ignore permission errors in sentry (#6011)
* app: fix artifact query errors with nil project/entity (#6010)
* [WB-3794] Artifact Sidebar for cross comparisons (#6008)
* More spooky adjectives and nouns (#6005)
* gorilla: fix models query with nil entity / project (#6009)
* query test for license dashboard (#6006)
* gorilla: expose artifact tracked storage in the explorer (#5988)
* [WB-3766] Add button to insert text block in reports (#5986)
* actually add react-textarea-autosize dependency
* project page: add top padding to empty artifacts watermark
* Feature/tiny plot fixes (#6002)
* add to report from run page: preserve default colors (#6003)
* Local License Dashboard (#5924)
* Fix/skip vega resize under modal (#6001)
* Bounding Box labels and Collapsible controls (#5995)
* John/factor out infinite scroll (#5999)
* runs table: add hidden-by-default columns
* runs table: fix tags popping out of existence briefly
* remove hack that improved performance (#5997)
* Revert "WB-3535 gorilla: add metadata column on runs table (#5973)"
* panelbank: fix resize loop due to scrollbar
* gorilla: fix extraInfo sentry reporting (#5996)
* WB-3535 gorilla: add metadata column on runs table (#5973)
* WB-1833 batch select sweeps to delete (#5991)
* gallery event updates + community slack button
* Always force new local users to set a username. (#5992)
* WB-3657 DSVIZ filter (#5966)
* CLI-498: Performance improvements for heavy Vega charts. (#5993)
* revert gallery height change (#5994)
* preserve color/mark/title overrides when exporting panels from run page
* media panel: fix ULTRABONK resize loop issue
* [WB-3692]: Allow admins to change the number of seats allocate to an organization via web dashboard (#5967)
* SEO: don't index run pages (#5990)
* SEO: proper canonical url for benchmarks (#5989)
* gorilla: centralize sentry reporting (#5985)
* Fix parallel coordinates column renaming + reordering bugs (#5972)
* Fix/fix injecting colors (#5987)
* Feature/dataset tweaks (#5980)
* append viewer username as querystring to gallery submission typeform
* change text auto-appended to benchmark descriptions
* add panels to report: open report in new tab (#5976)
* profile page report showcase: update preview image (#5977)
* Bump max cli version (#5965)
* Fix tags sidebar scrollbar not hidden (#5982)
* remove useMemo around making modfiiedUserQuery (#5984)
* Hide legend scrollbars on all browsers (#5979)
* wandb/local now supports an external redis server with caching enabled, HA deployment, and a lock around migrations
* SEO: better titles and descriptions for benchmarks (#5975)
* logging around updateMediaBrowserSize
* disable invalid navigation links on /signup (#5974)
* prevent creating invite emails with leading/trailing spaces
* don't dump redux auth data to sentry to  solve import loop
* capture error when old-style queries fail
* useQuery: log additional data to sentry when data is null
* sagas saveLoop: more informative error messsage on view save failure
* runLogQuery: handle null project run
* Handle NodeList.forEach() being unsupported in old browsers
* Bounding boxes to parity with run bounding boxes (#5969)
* Memoize stuff so we don't recompute vega data (#5971)

## wandb/local:0.9.31 (November 10, 2020)

* Show API keys for service accounts (#6054)
* SEO: run page noindex adjustment
* Make aliases copyable on Artifacts Overview (#6053)
* [WB-3621] Frontend for artifact collection deletion (#6032)
* Fixed Map Initialization (#6049)
* Upgrade lyft to 1.8 (#6042)
* Bump max cli version (#6048)
* Deduped Artifact Reference Limit Counter (#6046)
* [WB-3818] Artifacts overview - fix code copying, made version copyable (#6045)
* Set http-secure header in local (#6040)
* Remove tailon service from local (#6037)
* gorilla: lower max run config to 15 MB (#6044)
* gorilla: move away from linking to types directly within artifacts (#6039)
* Simple fix for suren (#6034)
* fix filtering (#6038)
* adding panel to panelbank section opens it
* update gallery webinar
* add panels to report: option to continue draft
* update comment notification email (#6031)
* fix crash when workspace has grouped runs with histogram-uberfile
* WB-3825 fix image panel step slider (#6029)
* Fixed Scrolling Overflow Issues for Artifact Tables (#6030)
* [WB-3824] update the links in UI (#6027)
* Bump max cli version (#6028)
* Fixes Dataviz Group UI (#6026)
* signup: fix error when user has duplicate invites (#6025)
* Check if a vega spec uses input bindings more thoroughly (#6024)
* gorilla: system metrics sampling returns keys in correct order (#6022)
* gorilla: disallow duplicate alerts per entity (#6021)
* gorilla: fix FetchArtifactDependenciesByArtifact query (#6018)
* Artifacts/tables and panels (#6020)
* moved time initialization and increased threshold (#6019)
* Lots of dataset & prediction viz UI updates (#6014)
* gorilla: fix run storage stats cleaner query (#6013)
* Modified Artifact Encoding Strategy (#6017)
* gorilla: fix FetchControllerRunNameBySweep (#6016)
* [CLI-476] Add runtime to parameter importance panel (#6004)
* also delete stars when deleting views
* Click image to show original size in modal (#5622)
* gorilla: autogenerate scaffolding for mysql fetch methods (#6007)
* Artifact Dependencies: Adding support for artifact dependencies as it relates to dataviz (#5998)
* gorilla: ignore permission errors in sentry (#6011)
* app: fix artifact query errors with nil project/entity (#6010)
* [WB-3794] Artifact Sidebar for cross comparisons (#6008)
* More spooky adjectives and nouns (#6005)
* gorilla: fix models query with nil entity / project (#6009)
* query test for license dashboard (#6006)
* gorilla: expose artifact tracked storage in the explorer (#5988)
* [WB-3766] Add button to insert text block in reports (#5986)
* actually add react-textarea-autosize dependency
* project page: add top padding to empty artifacts watermark
* Feature/tiny plot fixes (#6002)
* add to report from run page: preserve default colors (#6003)
* Local License Dashboard (#5924)
* Fix/skip vega resize under modal (#6001)
* Bounding Box labels and Collapsible controls (#5995)
* John/factor out infinite scroll (#5999)
* runs table: add hidden-by-default columns
* runs table: fix tags popping out of existence briefly
* remove hack that improved performance (#5997)
* Revert "WB-3535 gorilla: add metadata column on runs table (#5973)"
* panelbank: fix resize loop due to scrollbar
* gorilla: fix extraInfo sentry reporting (#5996)
* WB-3535 gorilla: add metadata column on runs table (#5973)
* WB-1833 batch select sweeps to delete (#5991)
* gallery event updates + community slack button
* Always force new local users to set a username. (#5992)
* WB-3657 DSVIZ filter (#5966)
* CLI-498: Performance improvements for heavy Vega charts. (#5993)
* revert gallery height change (#5994)
* preserve color/mark/title overrides when exporting panels from run page
* media panel: fix ULTRABONK resize loop issue
* [WB-3692]: Allow admins to change the number of seats allocate to an organization via web dashboard (#5967)
* SEO: don't index run pages (#5990)
* SEO: proper canonical url for benchmarks (#5989)
* gorilla: centralize sentry reporting (#5985)
* Fix parallel coordinates column renaming + reordering bugs (#5972)
* Fix/fix injecting colors (#5987)
* Feature/dataset tweaks (#5980)
* append viewer username as querystring to gallery submission typeform
* change text auto-appended to benchmark descriptions
* add panels to report: open report in new tab (#5976)
* profile page report showcase: update preview image (#5977)
* Bump max cli version (#5965)
* Fix tags sidebar scrollbar not hidden (#5982)
* remove useMemo around making modfiiedUserQuery (#5984)
* Hide legend scrollbars on all browsers (#5979)
* wandb/local now supports an external redis server with caching enabled, HA deployment, and a lock around migrations
* SEO: better titles and descriptions for benchmarks (#5975)
* logging around updateMediaBrowserSize
* disable invalid navigation links on /signup (#5974)
* prevent creating invite emails with leading/trailing spaces
* don't dump redux auth data to sentry to  solve import loop
* capture error when old-style queries fail
* useQuery: log additional data to sentry when data is null
* sagas saveLoop: more informative error messsage on view save failure
* runLogQuery: handle null project run
* Handle NodeList.forEach() being unsupported in old browsers
* Bounding boxes to parity with run bounding boxes (#5969)
* Memoize stuff so we don't recompute vega data (#5971)

## wandb/local:0.9.31 (November 10, 2020)

* Show API keys for service accounts (#6054)
* SEO: run page noindex adjustment
* Make aliases copyable on Artifacts Overview (#6053)
* [WB-3621] Frontend for artifact collection deletion (#6032)
* Fixed Map Initialization (#6049)
* Upgrade lyft to 1.8 (#6042)
* Bump max cli version (#6048)
* Deduped Artifact Reference Limit Counter (#6046)
* [WB-3818] Artifacts overview - fix code copying, made version copyable (#6045)
* Set http-secure header in local (#6040)
* Remove tailon service from local (#6037)
* gorilla: lower max run config to 15 MB (#6044)
* gorilla: move away from linking to types directly within artifacts (#6039)
* Simple fix for suren (#6034)
* fix filtering (#6038)
* adding panel to panelbank section opens it
* update gallery webinar
* add panels to report: option to continue draft
* update comment notification email (#6031)
* fix crash when workspace has grouped runs with histogram-uberfile
* WB-3825 fix image panel step slider (#6029)
* Fixed Scrolling Overflow Issues for Artifact Tables (#6030)
* [WB-3824] update the links in UI (#6027)
* Bump max cli version (#6028)
* Fixes Dataviz Group UI (#6026)
* signup: fix error when user has duplicate invites (#6025)
* Check if a vega spec uses input bindings more thoroughly (#6024)
* gorilla: system metrics sampling returns keys in correct order (#6022)
* gorilla: disallow duplicate alerts per entity (#6021)
* gorilla: fix FetchArtifactDependenciesByArtifact query (#6018)
* Artifacts/tables and panels (#6020)
* moved time initialization and increased threshold (#6019)
* Lots of dataset & prediction viz UI updates (#6014)
* gorilla: fix run storage stats cleaner query (#6013)
* Modified Artifact Encoding Strategy (#6017)
* gorilla: fix FetchControllerRunNameBySweep (#6016)
* [CLI-476] Add runtime to parameter importance panel (#6004)
* also delete stars when deleting views
* Click image to show original size in modal (#5622)
* gorilla: autogenerate scaffolding for mysql fetch methods (#6007)
* Artifact Dependencies: Adding support for artifact dependencies as it relates to dataviz (#5998)
* gorilla: ignore permission errors in sentry (#6011)
* app: fix artifact query errors with nil project/entity (#6010)
* [WB-3794] Artifact Sidebar for cross comparisons (#6008)
* More spooky adjectives and nouns (#6005)
* gorilla: fix models query with nil entity / project (#6009)
* query test for license dashboard (#6006)
* gorilla: expose artifact tracked storage in the explorer (#5988)
* [WB-3766] Add button to insert text block in reports (#5986)
* actually add react-textarea-autosize dependency
* project page: add top padding to empty artifacts watermark
* Feature/tiny plot fixes (#6002)
* add to report from run page: preserve default colors (#6003)
* Local License Dashboard (#5924)
* Fix/skip vega resize under modal (#6001)
* Bounding Box labels and Collapsible controls (#5995)
* John/factor out infinite scroll (#5999)
* runs table: add hidden-by-default columns
* runs table: fix tags popping out of existence briefly
* remove hack that improved performance (#5997)
* Revert "WB-3535 gorilla: add metadata column on runs table (#5973)"
* panelbank: fix resize loop due to scrollbar
* gorilla: fix extraInfo sentry reporting (#5996)
* WB-3535 gorilla: add metadata column on runs table (#5973)
* WB-1833 batch select sweeps to delete (#5991)
* gallery event updates + community slack button
* Always force new local users to set a username. (#5992)
* WB-3657 DSVIZ filter (#5966)
* CLI-498: Performance improvements for heavy Vega charts. (#5993)
* revert gallery height change (#5994)
* preserve color/mark/title overrides when exporting panels from run page
* media panel: fix ULTRABONK resize loop issue
* [WB-3692]: Allow admins to change the number of seats allocate to an organization via web dashboard (#5967)
* SEO: don't index run pages (#5990)
* SEO: proper canonical url for benchmarks (#5989)
* gorilla: centralize sentry reporting (#5985)
* Fix parallel coordinates column renaming + reordering bugs (#5972)
* Fix/fix injecting colors (#5987)
* Feature/dataset tweaks (#5980)
* append viewer username as querystring to gallery submission typeform
* change text auto-appended to benchmark descriptions
* add panels to report: open report in new tab (#5976)
* profile page report showcase: update preview image (#5977)
* Bump max cli version (#5965)
* Fix tags sidebar scrollbar not hidden (#5982)
* remove useMemo around making modfiiedUserQuery (#5984)
* Hide legend scrollbars on all browsers (#5979)
* wandb/local now supports an external redis server with caching enabled, HA deployment, and a lock around migrations
* SEO: better titles and descriptions for benchmarks (#5975)
* logging around updateMediaBrowserSize
* disable invalid navigation links on /signup (#5974)
* prevent creating invite emails with leading/trailing spaces
* don't dump redux auth data to sentry to  solve import loop
* capture error when old-style queries fail
* useQuery: log additional data to sentry when data is null
* sagas saveLoop: more informative error messsage on view save failure
* runLogQuery: handle null project run
* Handle NodeList.forEach() being unsupported in old browsers
* Bounding boxes to parity with run bounding boxes (#5969)
* Memoize stuff so we don't recompute vega data (#5971)

## wandb/local:0.9.31 (November 10, 2020)

* Show API keys for service accounts (#6054)
* SEO: run page noindex adjustment
* Make aliases copyable on Artifacts Overview (#6053)
* [WB-3621] Frontend for artifact collection deletion (#6032)
* Fixed Map Initialization (#6049)
* Upgrade lyft to 1.8 (#6042)
* Bump max cli version (#6048)
* Deduped Artifact Reference Limit Counter (#6046)
* [WB-3818] Artifacts overview - fix code copying, made version copyable (#6045)
* Set http-secure header in local (#6040)
* Remove tailon service from local (#6037)
* gorilla: lower max run config to 15 MB (#6044)
* gorilla: move away from linking to types directly within artifacts (#6039)
* Simple fix for suren (#6034)
* fix filtering (#6038)
* adding panel to panelbank section opens it
* update gallery webinar
* add panels to report: option to continue draft
* update comment notification email (#6031)
* fix crash when workspace has grouped runs with histogram-uberfile
* WB-3825 fix image panel step slider (#6029)
* Fixed Scrolling Overflow Issues for Artifact Tables (#6030)
* [WB-3824] update the links in UI (#6027)
* Bump max cli version (#6028)
* Fixes Dataviz Group UI (#6026)
* signup: fix error when user has duplicate invites (#6025)
* Check if a vega spec uses input bindings more thoroughly (#6024)
* gorilla: system metrics sampling returns keys in correct order (#6022)
* gorilla: disallow duplicate alerts per entity (#6021)
* gorilla: fix FetchArtifactDependenciesByArtifact query (#6018)
* Artifacts/tables and panels (#6020)
* moved time initialization and increased threshold (#6019)
* Lots of dataset & prediction viz UI updates (#6014)
* gorilla: fix run storage stats cleaner query (#6013)
* Modified Artifact Encoding Strategy (#6017)
* gorilla: fix FetchControllerRunNameBySweep (#6016)
* [CLI-476] Add runtime to parameter importance panel (#6004)
* also delete stars when deleting views
* Click image to show original size in modal (#5622)
* gorilla: autogenerate scaffolding for mysql fetch methods (#6007)
* Artifact Dependencies: Adding support for artifact dependencies as it relates to dataviz (#5998)
* gorilla: ignore permission errors in sentry (#6011)
* app: fix artifact query errors with nil project/entity (#6010)
* [WB-3794] Artifact Sidebar for cross comparisons (#6008)
* More spooky adjectives and nouns (#6005)
* gorilla: fix models query with nil entity / project (#6009)
* query test for license dashboard (#6006)
* gorilla: expose artifact tracked storage in the explorer (#5988)
* [WB-3766] Add button to insert text block in reports (#5986)
* actually add react-textarea-autosize dependency
* project page: add top padding to empty artifacts watermark
* Feature/tiny plot fixes (#6002)
* add to report from run page: preserve default colors (#6003)
* Local License Dashboard (#5924)
* Fix/skip vega resize under modal (#6001)
* Bounding Box labels and Collapsible controls (#5995)
* John/factor out infinite scroll (#5999)
* runs table: add hidden-by-default columns
* runs table: fix tags popping out of existence briefly
* remove hack that improved performance (#5997)
* Revert "WB-3535 gorilla: add metadata column on runs table (#5973)"
* panelbank: fix resize loop due to scrollbar
* gorilla: fix extraInfo sentry reporting (#5996)
* WB-3535 gorilla: add metadata column on runs table (#5973)
* WB-1833 batch select sweeps to delete (#5991)
* gallery event updates + community slack button
* Always force new local users to set a username. (#5992)
* WB-3657 DSVIZ filter (#5966)
* CLI-498: Performance improvements for heavy Vega charts. (#5993)
* revert gallery height change (#5994)
* preserve color/mark/title overrides when exporting panels from run page
* media panel: fix ULTRABONK resize loop issue
* [WB-3692]: Allow admins to change the number of seats allocate to an organization via web dashboard (#5967)
* SEO: don't index run pages (#5990)
* SEO: proper canonical url for benchmarks (#5989)
* gorilla: centralize sentry reporting (#5985)
* Fix parallel coordinates column renaming + reordering bugs (#5972)
* Fix/fix injecting colors (#5987)
* Feature/dataset tweaks (#5980)
* append viewer username as querystring to gallery submission typeform
* change text auto-appended to benchmark descriptions
* add panels to report: open report in new tab (#5976)
* profile page report showcase: update preview image (#5977)
* Bump max cli version (#5965)
* Fix tags sidebar scrollbar not hidden (#5982)
* remove useMemo around making modfiiedUserQuery (#5984)
* Hide legend scrollbars on all browsers (#5979)
* wandb/local now supports an external redis server with caching enabled, HA deployment, and a lock around migrations
* SEO: better titles and descriptions for benchmarks (#5975)
* logging around updateMediaBrowserSize
* disable invalid navigation links on /signup (#5974)
* prevent creating invite emails with leading/trailing spaces
* don't dump redux auth data to sentry to  solve import loop
* capture error when old-style queries fail
* useQuery: log additional data to sentry when data is null
* sagas saveLoop: more informative error messsage on view save failure
* runLogQuery: handle null project run
* Handle NodeList.forEach() being unsupported in old browsers
* Bounding boxes to parity with run bounding boxes (#5969)
* Memoize stuff so we don't recompute vega data (#5971)

## wandb/local:0.9.31 (November 10, 2020)



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


