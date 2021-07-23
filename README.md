# Updater Demo

This is a working updater demo (currently windows only) using the 
"Launcher/App Interdependence Architecture".

That means the launcher is the starting point. The launcher the checks for updates for the app and
applies them, if available. The launcher then launches the app and exist.
The app checks for updates for the launcher and applies them, if available.

That way, the app is always up-to-date. The launcher might have additional features, like a repair
process or sending feedback, whatever your heart desires.