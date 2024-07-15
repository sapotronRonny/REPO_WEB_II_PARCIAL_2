from rest_framework import viewsets
from .serializer import PeliculaSerializer, ActorPeliculaSerializer, GeneroPeliculaSerializer
from .models import Pelicula, ActorPelicula, GeneroPelicula

class PeliculaViewSet(viewsets.ModelViewSet):
    queryset = Pelicula.objects.all()
    serializer_class = PeliculaSerializer

class ActorPeliculaViewSet(viewsets.ModelViewSet):
    queryset = ActorPelicula.objects.all()
    serializer_class = ActorPeliculaSerializer

class GeneroPeliculaViewSet(viewsets.ModelViewSet):
    queryset = GeneroPelicula.objects.all()
    serializer_class = GeneroPeliculaSerializer
