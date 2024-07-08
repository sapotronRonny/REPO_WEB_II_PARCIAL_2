from django.db import models

class Genero(models.Model):
    class Meta:
        db_table = 'Genero'  # Nombre exacto de la tabla en la base de datos
    nombre = models.CharField(max_length=255)

class Actor(models.Model):
    class Meta:
        db_table = 'Actor'  # Nombre exacto de la tabla en la base de datos
    nombre = models.CharField(max_length=255)

class Pelicula(models.Model):
    id_pelicula = models.AutoField(primary_key=True)
    titulo = models.CharField(max_length=255)
    sinopsis = models.TextField(blank=True, null=True)
    id_genero = models.ForeignKey('Genero', on_delete=models.CASCADE)
    id_actor_principal = models.ForeignKey('Actor', on_delete=models.CASCADE)
    fecha_creacion = models.DateField()
    duracion = models.TimeField()

    def __str__(self):
        return self.titulo
