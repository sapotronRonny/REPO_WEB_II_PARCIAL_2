from rest_framework import viewsets
from rest_framework.permissions import IsAuthenticated
from rest_framework.response import Response
from .serializer import RegistroSerializer, UsuarioSerializer, ComentarioSerializer, PeliculaSerializer, IdiomaSerializer, PeliculaGeneroSerializer, GeneroSerializer, ActorPeliculaSerializer, ActorSerializer
from .models import Registro, Usuario, Comentario, Pelicula, Idioma, PeliculaGenero, Genero, ActorPelicula, Actor  

# Create your views here.

class PeliculaViewSet(viewsets.ModelViewSet):
    permission_classes = [IsAuthenticated]
    queryset = Pelicula.objects.all()
    serializer_class = PeliculaSerializer

class generoViewSet(viewsets.ModelViewSet):
    permission_classes = [IsAuthenticated]
    queryset = Genero.objects.all()
    serializer_class = GeneroSerializer

class actoresViewSet(viewsets.ModelViewSet):
    permission_classes = [IsAuthenticated]
    queryset = Actor.objects.all()
    serializer_class = ActorSerializer


# Path: api/urls.py