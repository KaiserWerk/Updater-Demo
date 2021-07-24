# Updater Demo

This is a working updater demo (currently windows only) using the 
"Launcher/App Interdependence Architecture" (great name, I know).

That means the launcher is the starting point. The launcher the checks for updates for the app and
applies them, if available. The launcher then launches the app and exits.
The app checks for updates for the launcher and applies them, if available.

That way, launcher and app are both always up-to-date. The launcher might have additional features, like a repair
process or sending feedback, whatever your heart desires.

# Setup

Execute the powershell scripts ``build-data.ps1`` to create app and launcher versions 1.0.0 through
1.0.3 (which are the files served by the update server) and ``build-demo.ps1`` to build the initial 
executables. You can place them in whatever directory you like, just make sure that the launcher and 
the app are placed in the same directory.

# Usage

1. Start up the update server.
1. Start the launcher. Not much should happen besides the launcher starting the app and then exiting.
1. After using the ``build-demo.ps1`` script, ``launcher.exe`` and ``app.exe`` are at version 1.0.0. Play
   around with the version strings in files in ``internal/assets/versionFiles``, run the launcher again and
   see the magic!
