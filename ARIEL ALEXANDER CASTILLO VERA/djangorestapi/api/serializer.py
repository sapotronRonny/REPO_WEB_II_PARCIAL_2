from rest_framework import serializers
from .models import Pelicula, GeneroPelicula, ActorPelicula



class PeliculaSerializer(serializers.ModelSerializer):
    class Meta:
        model = Pelicula
        fields = '__all__'

class ActorPeliculaSerializer(serializers.ModelSerializer):
    class Meta:
        model = ActorPelicula
        fields = '__all__'

class GeneroPeliculaSerializer(serializers.ModelSerializer):
    class Meta:
        model = GeneroPelicula
        fields = '__all__'