from django.contrib import admin
from .models import Pelicula, GeneroPelicula, ActorPelicula
# Register your models here.


admin.site.register(Pelicula)
admin.site.register(GeneroPelicula)
admin.site.register(ActorPelicula)
