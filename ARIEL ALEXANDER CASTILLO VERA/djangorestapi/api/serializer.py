from rest_framework import serializers
from .models import Pelicula, Genero, Actor

class PeliculaSerializer(serializers.ModelSerializer):
    class Meta:
        model = Pelicula
        fields = '__all__'

class generoSerializer(serializers.ModelSerializer):
    class Meta:
        model = genero
        fields = '__all__'

class ActoresSerializer(serializers.ModelSerializer):
    class Meta:
        model = actores
        fields = '__all__'
