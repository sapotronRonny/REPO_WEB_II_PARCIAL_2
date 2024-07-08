from rest_framework import viewsets
from rest_framework.permissions import IsAuthenticated
from rest_framework.response import Response
from .serializer import PeliculaSerializer, ActorPeliculaSerializer, GeneroPeliculaSerializer
from .models import Pelicula, ActorPelicula, GeneroPelicula

# Create your views here.


class PeliculaViewSet(viewsets.ModelViewSet):
    permission_classes = [IsAuthenticated]
    queryset = Pelicula.objects.all()
    serializer_class = PeliculaSerializer

class ActorPeliculaViewSet(viewsets.ModelViewSet):
    permission_classes = [IsAuthenticated]
    queryset = ActorPelicula.objects.all()
    serializer_class = ActorPeliculaSerializer

class GeneroPeliculaViewSet(viewsets.ModelViewSet):
    permission_classes = [IsAuthenticated]
    queryset = GeneroPelicula.objects.all()
    serializer_class = GeneroPeliculaSerializer
# Path: api/urls.py