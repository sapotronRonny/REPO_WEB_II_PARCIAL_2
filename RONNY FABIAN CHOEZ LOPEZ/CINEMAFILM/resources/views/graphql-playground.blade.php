<!DOCTYPE html>
<html>
<head>
    <title>GraphQL Playground</title>
    <link rel="stylesheet" href="//cdn.jsdelivr.net/npm/graphql-playground-react@1.7.20/build/static/css/index.css" />
    <link rel="shortcut icon" href="//cdn.jsdelivr.net/npm/graphql-playground-react@1.7.20/build/favicon.png" />
    <script src="//cdn.jsdelivr.net/npm/graphql-playground-react@1.7.20/build/static/js/middleware.js"></script>
</head>
<body>
    <div id="root"></div>
    <script>window.addEventListener('load', function (event) {
        GraphQLPlayground.init(document.getElementById('root'), {
            endpoint: '/graphql'
        })
    })</script>
</body>
</html>