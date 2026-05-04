// Runs on every viewer request before CloudFront looks up the cache key.
// Two jobs:
//   1. Redirect the apex to the default landing page.
//   2. Rewrite extensionless URIs to their /index.html counterpart so
//      subpath requests (/brooklyn/therapy) actually find an S3 key.
function handler(event) {
  var request = event.request;
  var uri = request.uri;

  if (uri === '/' || uri === '') {
    return {
      statusCode: 302,
      statusDescription: 'Found',
      headers: { location: { value: '/brooklyn/therapy' } }
    };
  }

  if (uri.endsWith('/')) {
    request.uri = uri + 'index.html';
  } else if (!uri.split('/').pop().includes('.')) {
    request.uri = uri + '/index.html';
  }

  return request;
}
