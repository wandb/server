## wandb/local:0.9.44 (September 28, 2021)

* Fixed bug where an external S3 object store wouldn't connect over tls
* Local supports being run under a sub-path of an existing domain, i.e. https://mycompany.com/wandb
* Local supports custom headers for authentication when running behind a proxy
* Fix a UI crash on organization dashboards.
* Adds new shortcuts for commands in Report editing mode
* Move default entity into its own category in home page sidebar
* Fix issue where some runqueues would not show up in launch config modal
* Media panel now shows a more informative error message when handling non-media values.
* Add option to filter runs by artifact
* Prevent non-admins from deleting projects
* Fixes a bug where users' sweeps may have duplicate runs
* Return error if run names already exist in destination project when moving runs
* Allow entering arbitrary string literals in the Weave expression editor.
* Responsive styles for the controls image overlay
* Render image placeholder when image in media panel changes
* Fixes an issue where grid search sweeps terminated prematurely
* Fixes a bug where resuming a grid search sweep would cause it to restart instead of picking up where it left off
* Fixes a bug where the media panel would endlessly send requests to the server.
* Fixes a race condition in preemptible sweeps that sometimes caused preempted runs to not be retried
* Users can now view the PyTorch Profiler tab inside Tensorboard in W&B UI, if they use the client library in conjunction with PyTorch to profile their code.
* Weave expressions now support `argmin` and `argmax`
* Fixed crash with image mask controls for masks with numeric labels.
* Fixed an issue where requests could fail for artifacts under deleted projects
* Weave Plot infers colors more intelligently
* Fix the issue with the onboarding survey where users can skip it with the browser's "back" button.
* Update the run's text colors (e.g. tooltip, legend) to be more accessible.
* Table of Contents can include different levels of headers in the Reports.
* Improves performance of sweeps by implementing them with run queues
* Fixed a bug that sometimes prevented Weave panels from loading.
* Improved iframe rendering when viewing runs in Jupyter
* Allow reports to be embedded as iframes in tools like Confluence
* Better error messages when access is denied to a resource in iframes
* Upgraded our LaTeX rendering engine to address a possible command injection vulnerability.
* Improves accuracy and utility of validation/suggestions in Monaco sweep editor in the UI
* Improve resource usage of artifact overview page
* Fixes the Code Comparer panel when used in a cross-project run set.
* Users without usernames can be deleted from the admin dashboard
* Update to WBButton component (only visible in storybook)
* Enable reordering in the runs grouping selector.
* Update the run's color to be more accessible.
* Enable support for anonymous functions in Weave expression editor
* W&B Tables can now leverage web workers to speed up computations!
* Fixes issue where fullscreen charts overflow the bottom of the page
* New improved sweep engine
* Users with `@` in their artifact type names won't have broken links anymore
* Fixes the "add visualization" preview not showing proper data, sometimes showing nothing.
* Fix bug in Weave expression editor where pressing space after a number replaced the entire expression with that number.
* Made the artifact sidebar resizable
* Changes made only affect new WB components, no existing ones.
* Update launch tab icon
* Improve run queue tab UI
* Improve performance of merging partial manifests
* Artifact Graph view now has the option to show/hide automatically generated Artifacts.
* Instance Activity Dashboard is now available! Navigate to `https://<your-wandb-host>/admin/usage` to try it out! (requires admin access and license upgrade)
* Users can programmatically fetch reports using API with pagination.
* On a user's profile page, hide the "Likes" tab if the user hasn't liked any reports yet.
* Fixed performance issues when loading the storage explorer with a large number of artifacts.
* Add public API examples to the project overview page.
* In Team settings, disable adding an existing team member.
* In Reports, don't show the "Write caption..." section when dragging the image in the view mode.
* In Reports, editors can change the visibility of the run sets in the panel grid for the saved reports.
* Storage explorer url changes from `/storage/entityName` to `/usage/entityName`
* Fixes a bug in grid search where arrays of yaml objects were not properly handled
* Prevent all pages except run overview from being embedded in an iframe
* Fixes an issue in sweeps were constant parameters were not handled in grid search
* Disabled image resizing in the reports.
* Weave Panel in Reports no longer automatically grabs focus and scroll page
* TracerPanel no longer automatically grabs focus and scroll page
* Removed TracerPanel's unnecessary scroll bar
* Weave ops that accept functions can now use inner ops that consume tags.
* Images/Videos/Media will no longer break after the direct URLs expire
* Fixes line plot legend positioning with east/west configuration
* Fixes safari issues with line plots
* Users can now use `.run` and `.project` in Tables and Weave to retrieve the source object.
* Tables which are the concatenation of tables from multiple runs will now automatically show the run name
* Fixes style bugs affecting icons.
* Grouped line plots will have a more subtle transition to highlighted state - increases legibility.
* Adds back the legends for box plot and violin plot
* When navigating to an Artifact file, users would previously see "no preview" before the file successfully loads. Now, a loading indicator is presented.
* Fixed the link for "used by" runs in the Artifacts page.
* Add validation to entity and team names


## wandb/local:0.9.43 (August 13, 2021)

Changelog:
* Prevent non-whitelisted domains from making CORS requests
* Fixed bug preventing users from logging in
* In Reports, fixed a bug that caused adding expressions to take a long time
* In Reports, the copy link icons show up when hovering over the headers
* Fixed the issue that the user can't undelete runs after deleting all the runs
* Fixed an issue where pages could crash if a run didn't have a `wandb-metadata.json` file
* In Reports, users can copy a hyperlink to a specific section on the view mode.
* Flips the default order of values and name in the line plot legend
* Makes the line plot legend smarter about truncating long legends
* Removes run highlighting on the single run page
* The slider for the video panels is now set to visible by default without the users having to click the setting button
* Artifact Graph now automatically hides source code assets
* Fixed image rendering pixelation in firefox
* Updated sweep validation warnings
* Added "Horizontal rule" to the slash menu in Reports
* In Reports, users can add a table of contents that contains a list of links to headings
* Fixed inserting images to panel grid markdown in Reports
* Fixed an issue that caused StringCompare panels to be very slow on large inputs.
* Fixed Netron Viewer in Artifacts / Tables
* Made direct OIDC authentication possible in local servers
* Added a session length configuration option for local servers
* Moved the existing OIDC callback url from "/login" to "/oidc/callback" (existing Auth0 tenants will need to update the allowed callback urls)
* Allow users to sort panels alphabetically in sections
* Added createdBy, metadata, aliases, and link properties to artifact versions in Weave
* Added verbose S3 logging, enabled by the `S3_LOG_VERBOSE` env var
* Tables: Added "remove all right" and "remove all left" column menu options
* Tables: Added ability to access a project's runs via `project.runs`
* Added the ability to configure SSL and path style when using an S3-compatible bucket store
* Fix display of runtime value when syncing from offline mode
* Officially dropped support for IE11
* Fixed artifacts issues when using third party S3-compatible object stores
* Typing a bracket now completes bracket expressions in the Weave expression editor
* The left/right arrows can now skip back and forth between nodes in the Weave expression editor
* Tables now show run colors on all bar plots
* In Reports, users can now add a checklist from the slash dropdown menu
* Fixed cursor behavior when adding a checklist in Report
* Users can now properly compare files using Weave Query.
* Fixed bug causing Tables in full screen mode to overflow
* In Reports, error messages are now shown in the specific panel grid where the error was triggered, instead of the whole report page
* Fixed bug preventing run tables logged with older clients from being queried from Weave
* Joined tables now properly show bar chart colors when the column is a list
* Tables now support Molecule media type
* Fixed issue with Artifact Creation when client loses the request connection
* Fixed an issue with the line plot where the config can get into an invalid state
* Fixed an issue with the run files explorer which can cause viewing files to fail
* In Reports, added the ability to link to a specific section
* Users can now drag and drop a panel in the Report's panel grid without moving up the panels below
* Added "gallery blocks" in the report editor, where each card in the gallery will link to another report
* Fixed an issue that could cause undeleting runs to sometimes fail
* GIFs are now allowed as report preview images
* Fixed bug causing excessive network requests when rendering Vega plots
* Default Table View for Run Workspace fixed to concat rather than potentially join.
* PanelWeave now loads runs from all active Runsets
* Tables now support concat on Grouped Joins
* In Reports, a user can add a callout box by typing `>>> ` or `/callout`.
* Fixed an issue that could cause distributed artifacts with retrying writers to fail during the commit flow.
* Fixed bug that caused invalid URLs to be highlighted in tables
* Fixed bug that caused Audio Wave Forms to render behind sticky columns and headers
* Increase maximum metrics payload size per `wandb.log` call from 4MB to 10MB
* Fixed sizing bug on Table cells showing a list of objects
* Fixed bug in Tables where some string values did not appear as possible suggestions
* Table Panels rendered with Weave now filter out empty runs.
* Upgraded Ray to 4.1 to address CVE-2021-33503
* Pinned urllib3>=1.26.5 to address CVE-2021-33503
* In Reports, snake_case no longer gets converted to italics. (e.g. 1_2_3 stays the same without italicizing 2)
* In line plots with a single line, the line will no longer thicken when its run is highlighted
* Improved styling of highlighted runs in the runs table
* For artifacts with no metadata, a link to the docs explaining how to set metadata is now shown on the Metadata tab
* In artifact page, the Used By table is now hidden if the artifact is not consumed by any runs
* Fixed bug that caused some artifacts to appear to take up more disk space than they actually did
* Small plots now hide the run legend for better readability
* Fixed bug that made artifacts inaccessible on a project after the project was moved to a different user or organization
* Run selector now shows a warning when user attempts to change visibility of a run that's outside the current bounds of a parallel coordinate chart or scatter plot
* Fixed bug that caused some Tables to incorrectly render as empty
* Fixed bug that caused panels to change from line plots to bar charts when timestamp values were entered as min/max
* Projects a user does not own and has not contributed runs to will no longer appear on their profile page
* In multi-run line plots, the legend is now displayed by default
* Fixed bug that made sweep configs containing certain parameter names unparseable
* Fixed bug that caused the app to crash when copy/pasting from Apple Notes
* Extended artifacts JSON parser to allow for non-standard Infinity value in Python's generated JSON
* Fixed bug that caused errors when interacting with some dropdowns
* Fixed header indentation in Reports
* Fixed bug that caused runs to rener in parallel coordinates chart even if disabled in the runset
* Fixed a bug that caused pinned columns in tables to always appear to the left of other columns
* Improved run grouping functionality



## wandb/local:0.9.42 (June 30, 2021)

* Added support for authentication via Google Identity-Aware Proxy
* Add "Import panel" feature to reports
* Fixed bugs resulting from spaces in user emails
* Added the ability to view PyTorch Kineto traces in artifacts
* Added popup to clarify restricted characters in artifact aliases
* Security updates for wandb local
* BarChart, Histogram, String Histrogram, MultiHistrogram, and MultiStringHistorgram all now color their data with run colors
* Disabled the Delete Run button for runs that contain undeleted Artifacts
* Added download dropdown to the runs table, with export API and CSV export options
* Added the ability to edit Weave expressions in plain text, and to copy/paste them
* Fixed rendering errors affecting box plots
* Improve UX when viewing multiple Tables side-by-side
* Clarify supported regex syntax for the runs table in a tooltip
* Fixed sweeps bug where preemptible / preempting runs in would crash a Bayes search
* Fixed bug where accessing media panels would sometimes erroneously trigger user rate limits
* Wrap long names in the projects table and on project cards
* Fixed a bug where panel settings changes in workspaces sometimes wouldn't be applied
* Added documentation link to Workspace menu
* New report-creation experience
* Better UX for moving/duplicating/deleting blocks in reports
* Made the smoothing input for panels support values greater than the slider max
* Fixed issue where long media keys could overflow the media panel
* Fixed issue in Tables expression editor where you couldn't check expressions equality to None
* Hid code tab in UI when code-saving is disabled
* Various bug fixes and improvements for Tables and Artifacts
* Fixed bugs where special characters in group names could break the grouping UI
* Support for more intuitive logging syntax for Joined and Partitioned tables



## wandb/local:0.9.41 (May 27, 2021)

* Added a few optimizations for frontend and backend performance.
* Added validation for sweep configs.
* Added support for preemptible runs and sweeps.
* Fixed several small issues with reports.
* Added new "Weave" panel type to workspaces and reports.
* Fixed an issue where the backend might return 500 at the end of a grid search sweep.
* Made several improvements to scatter plots.
* Made links in table cells clickable.


## wandb/local:0.9.40 (April 29, 2021)

* Added option to set a different set of options in each panel section within a workspace.
* Added option to change default panel arrangement in workspaces based on metric prefixes.
* Added toggle for "freezing" a set of runs in the project run table.
* Made a few usability improvements for custom charts.
* Optimized for a few expensive backend queries.
* Added support for writing incremental artifacts.
* Added backend support for define_metric.
* Made a few smaller UI element improvements.
* Made improvements to artifact deletion.
* Implemented live-updating comments within reports.
* Fixed a bug with our authentication layer.


## wandb/local:0.9.39 (March 26, 2021)

* Added more options for increased parallelism in processing uploaded data.
* Fixed a few report stying regressions.
* Added more diagnostic info for sweeps.
* Increased size of 3d points in 3d object panel.
* Added more robust support for S3 buckets using custom KMS keys.
* Improved drag/drop behavior of panels in workspaces.
* Added checkboxes to reports.
* Reconfigured default limits for local installs.
* UI polish and usability improvements for dataset & prediction visualization tables.
* Added warning when local install is out of date.
* Added support for incremental artifact commits.
* Improved filter validation for run table queries.
* Added ability to create reports in empty projects.
* Improved jupyter notebook renderer and added ability to save code as an artifact.


## wandb/local:0.9.38 (February 17, 2021)

- Fixed an issue with histograms rendering colors in a non-deterministic order
- Fixed issues where deleting projects or benchmarks could crash runs
- Added a LOGGING_ENABLED flag to send server logs to stddout+stderr
- WYSIWYG reports
- Fixed an issue where file uploads fail on buckets using KMS 
- Security fixes
- Fixed an issue where bulk deleting projects or artifacts could fail


## wandb/local:0.9.37 (February 2, 2021)

* UX improvements for artifact aliases and run tag editing.
* Extended deadline and implemented retry for loading system settings from a bucket.
* Performance improvements for sweep suggestions.
* Added support for parallel writers in artifacts.
* UX improvements for dataset visualization.
* Improved debug information if container fails to start


## wandb/local:0.9.36 (January 26, 2021)

* Fixed an issue with parsing unicode characters in history.
* Improved report stability.
* Fixed an issue with artifact file downloads.
* Improved handling of nested key structures.
* Fixed an issue that may cause file notification queue backups.
* Added support for expressions in big scalar panel.
* General bug fixes and stability improvements.


## wandb/local:0.9.35 (January 14, 2021)

* Upgraded the base image to Ubuntu Bionic
* Added support for specifying the AWS_S3_KMS_ID env variable to encrypt S3 Objects with a custom key
* Fixed crash in Artifacts UI when viewing media in a join table
* Fixed the bottom of the custom chart editor getting cut off
* Added ability to filter/sort/group runs by hostname
* Fixed bug when linking to runs from outside their project
* Fixed an error state in the custom chart editor that prevented the data pane from listing data in some cases
* Added the ability to view a custom chart spec and copy it to a new report
* Fixed an infinite loop when using Google Cloud Storage
* Redacted emails when resetting passwords
* Updated modal flow when creating report
* Enable iframe embedding (e.g. Streamlit apps) in HTML media objects
* Added support for video files, HTML media objects, Bokeh plots, and boolean values to Artifacts
* More informative error state when uploading payloads that are too large
* Improved UI flow for deleting artifacts with aliases
* Updated icons in the report editor
* Created tables showing input and output artifacts used for a given run
* Added links to tables and joined tables in the Artifacts overview
* Added the ability to select empty string columns in the parallel coordinates chart
* Fixed "no runs selected" error when making selections on the parallel coordinates chart
* Made various improvements to the output of the LaTeX export
* Fix server error when getting a list of projects from an access token when some projects have been deleted
* Fix server error on project lookups missing a project or entity namee


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


