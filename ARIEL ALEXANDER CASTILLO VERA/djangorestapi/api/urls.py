from django.urls import path, include
from rest_framework import routers
from api import views

router = routers.DefaultRouter()
router.register(r'peliculas', views.PeliculaViewSet)
router.register(r'generos', views.GeneroViewSet)
router.register(r'actores', views.ActorViewSet)

urlpatterns = [
    path('', include(router.urls)),
]
