const cacheName = 'mavtc-v1';
const appShellFiles = [
    '/public/assets/app-192x.png',
    '/public/assets/app-512x.png',
    '/public/assets/mnr.css',
    '/public/assets/MNR2007.cd059e2cbf346028f917.woff2',
    '/public/assets/stations.json',
    '/public/assets/train_types.svg',
    '/public/assets/trainclasses.svg',
];
self.addEventListener('install', (e) => {
    console.log('[Service Worker] Install');
    e.waitUntil((async () => {
        const cache = await caches.open(cacheName);
        console.log('[Service Worker] Caching all: app shell and content');
        await cache.addAll(appShellFiles);
    })());
});
self.addEventListener('fetch', (e) => {
    e.respondWith((async () => {
        const r = await caches.match(e.request);
        console.log(`[Service Worker] Fetching resource: ${e.request.url}`);
        if (r) { return r; }
        const response = await fetch(e.request);
        const cache = await caches.open(cacheName);
        console.log(`[Service Worker] Caching new resource: ${e.request.url}`);
        cache.put(e.request, response.clone());
        return response;
    })());
});