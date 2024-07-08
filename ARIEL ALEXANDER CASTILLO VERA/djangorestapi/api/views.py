from rest_framework import viewsets
from rest_framework.permissions import IsAuthenticated
from .serializers import PeliculaSerializer, GeneroSerializer, ActorSerializer
from .models import Pelicula, Genero, Actor  

class PeliculaViewSet(viewsets.ModelViewSet):
    permission_classes = [IsAuthenticated]
    queryset = Pelicula.objects.all()
    serializer_class = PeliculaSerializer

class GeneroViewSet(viewsets.ModelViewSet):
    permission_classes = [IsAuthenticated]
    queryset = Genero.objects.all()
    serializer_class = GeneroSerializer

class ActorViewSet(viewsets.ModelViewSet):
    permission_classes = [IsAuthenticated]
    queryset = Actor.objects.all()
    serializer_class = ActorSerializer
