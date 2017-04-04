"use strict";

console.log('SW: executing the magic...');
var version = 'v1::';

/* These resources will be downloaded and cached by the service worker
   during the installation process. If any resource fails to be downloaded,
   then the service worker won't be installed either.
*/
var offlineFundamentals = [
  '/relaymaps/index-sw.html',
  'css/main.css',
  "https://fonts.googleapis.com/css?family=Roboto:regular,bold,italic,thin,light,bolditalic,black,medium&amp;lang=en",
  "https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-alpha.6/css/bootstrap.min.css",
  'relay_files/staticmap',
  'relay_files/staticmap(1)',
  "https://code.jquery.com/jquery-3.1.1.slim.min.js",
  "https://cdnjs.cloudflare.com/ajax/libs/tether/1.4.0/js/tether.min.js",
  "https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-alpha.6/js/bootstrap.min.js"
];

/* The install event fires when the service worker is first installed.
   You can use this event to prepare the service worker to be able to serve
   files while visitors are offline.
*/
self.addEventListener("install", function(event) {
  console.log('SW: install event in progress.');
  /* Using event.waitUntil(p) blocks the installation process on the provided
     promise. If the promise is rejected, the service worker won't be installed.
  */
  event.waitUntil(
    caches.open(version + 'fundamentals').then(function(cache) {
        /* After the cache is opened, we can fill it with the offline fundamentals.
           The method below will add all resources in `offlineFundamentals` to the
           cache, after making requests for them.
        */
        return cache.addAll(offlineFundamentals);
      })
      .then(function() {
        console.log('SW: install completed');
      })
  );
});

/* The fetch event fires whenever a page controlled by this service worker requests
   a resource. This isn't limited to `fetch` or even XMLHttpRequest. Instead, it
   comprehends even the request for the HTML page on first load, as well as JS and
   CSS resources, fonts, any images, etc.
*/
self.addEventListener("fetch", function(event) {
  console.log('SW: fetch event in progress.');

  /* We should only cache GET requests, and deal with the rest of method in the
     client-side, by handling failed POST,PUT,PATCH,etc. requests.
  */
  if (event.request.method !== 'GET') {
    /* If we don't block the event as shown below, then the request will go to
       the network as usual.
    */
    console.log('SW: fetch event ignored.', event.request.method, event.request.url);
    return;
  }
  
  event.respondWith(
    caches.match(event.request).then(function(cached) {
        /* Even if the response is in our cache, we go to the network as well.
           This pattern is known for producing "eventually fresh" responses,
           where we return cached responses immediately, and meanwhile pull
           a network response and store that in the cache.
        */
        var networked = fetch(event.request).then(fetchedFromNetwork, unableToResolve)
          // We should catch errors on the fetchedFromNetwork handler as well.
          .catch(unableToResolve);
        /* We return the cached response immediately if there is one, and fall
           back to waiting on the network as usual.
        */
        console.log('SW: fetch event', cached ? '(cached)' : '(network)', event.request.url);
        return cached || networked;

        function fetchedFromNetwork(response) {
          var cacheCopy = response.clone();
          console.log('SW: fetch response from network.', event.request.url);

          caches.open(version + 'pages').then(function add(cache) {
              cache.put(event.request, cacheCopy);
            })
            .then(function() {
              console.log('SW: fetch response stored in cache.', event.request.url);
            });

          // Return the response so that the promise is settled in fulfillment.
          return response;
        }

        /* When this method is called, it means we were unable to produce a response
           from either the cache or the network. This is our opportunity to produce
           a meaningful response even when all else fails. It's the last chance, so
           you probably want to display a "Service Unavailable" view or a generic
           error response.
        */
        function unableToResolve () {
          /* There's a couple of things we can do here.
             - Test the Accept header and then return one of the `offlineFundamentals`
               e.g: `return caches.match('/some/cached/image.png')`
             - You should also consider the origin. It's easier to decide what
               "unavailable" means for requests against your origins than for requests
               against a third party, such as an ad provider.
             - Generate a Response programmaticaly, as shown below, and return that.
          */

          console.log('SW: fetch request failed in both cache and network.');

          /* Here we're creating a response programmatically. The first parameter is the
             response body, and the second one defines the options for the response.
          */
          return new Response('<h1>Service Unavailable</h1>', {
            status: 503,
            statusText: 'Service Unavailable',
            headers: new Headers({
              'Content-Type': 'text/html'
            })
          });
        }
      })
  );
});

/* The activate event fires after a service worker has been successfully installed.
   It is most useful when phasing out an older version of a service worker, as at
   this point you know that the new worker was installed correctly. In this example,
   we delete old caches that don't match the version in the worker we just finished
   installing.
*/
self.addEventListener("activate", function(event) {
  /* Just like with the install event, event.waitUntil blocks activate on a promise.
     Activation will fail unless the promise is fulfilled.
  */
  console.log('SW: activate event in progress.');

  event.waitUntil(
    caches
      /* This method returns a promise which will resolve to an array of available
         cache keys.
      */
      .keys()
      .then(function (keys) {
        // We return a promise that settles when all outdated caches are deleted.
        return Promise.all(
          keys
            .filter(function (key) {
              // Filter by keys that don't start with the latest version prefix.
              return !key.startsWith(version);
            })
            .map(function (key) {
              /* Return a promise that's fulfilled
                 when each outdated cache is deleted.
              */
              return caches.delete(key);
            })
        );
      })
      .then(function() {
        console.log('SW: activate completed.');
      })
  );
});