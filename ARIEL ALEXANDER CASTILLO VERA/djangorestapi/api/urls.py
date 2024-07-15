from django.urls import path,include
from rest_framework import routers
from api import views

routers = routers.DefaultRouter()
routers.register(r'pelicula1', views.PeliculaViewSet)
routers.register(r'genero1', views.GeneroPeliculaViewSet)
routers.register(r'actor1', views.ActorPeliculaViewSet)

urlpatterns = [
    path('', include(routers.urls)),
]


