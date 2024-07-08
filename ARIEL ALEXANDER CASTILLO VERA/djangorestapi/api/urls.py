from django.urls import path,include
from rest_framework import routers
from api import views

routers = routers.DefaultRouter()
routers.register(r'pelicula', views.PeliculaViewSet)
routers.register(r'genero', views.GeneroPeliculaViewSet)
routers.register(r'actor', views.ActorPeliculaViewSet)

urlpatterns = [
    path('', include(routers.urls)),
]


