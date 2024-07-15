from django.contrib import admin
from django.urls import path, include
from rest_framework.documentation import include_docs_urls
from graphene_django.views import GraphQLView

urlpatterns = [
    path('admin/', admin.site.urls),
    path('api/v1/', include('api.urls')),
    path('docs/', include_docs_urls(title='Api Documentation')),
    path("graphql", GraphQLView.as_view(graphiql=True)),
]
