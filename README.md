Generates a web page with maps based on The Relay KMZ/KML maps.

## Live version
https://bit.ly/ggrr1 or https://greenido.github.io/relaymaps/

## Build

You will need to install [Bazel](https://bazel.build/). Then:

    $ bazel build relaymaps:main

## Run

First, download a kmz file. On Google maps editor, select "Download KML", then
"Keep data up to date with network link KML". This will download a file with a
URL reference to the "live" map.

Then:

    $ bazel-bin/relaymaps/main --kmz=path/to/file.kmz --out=/tmp/relay.html

## Notes

A fair amount of stuff is hardcoded. For example, legs must be named "Leg XX",
with "Exchange XX", or "Exch. XX", and "Start" / "Finish" items.


[![Analytics](https://ga-beacon.appspot.com/UA-65622529-1/golden-gate-replay-run/main)](https://github.com/igrigorik/ga-beacon)

