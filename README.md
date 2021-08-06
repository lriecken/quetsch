# Quetsch image size optimizer

Quetsch is german for squeeze.

Quetsch lets you reduce an image file to a maximum size. Quetsch will first reduce the resolution
until it hits a minimum resolution (conserving the aspect ratio of the original image. It will then
reduce the encoding quality.

Quetsch supports encoding to jpg, webp or avif. The latter is the new image format accompanying AV1. It is supported in some browsers.

This is so far only tested on MacOS.

You might need some additional system libraries for webp and avif.

You might have a problem if you are using Windows.
