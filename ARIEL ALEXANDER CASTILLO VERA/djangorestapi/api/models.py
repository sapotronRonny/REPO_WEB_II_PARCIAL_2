from django.db import models

class genero(models.Model):
    class Meta:
        db_table = 'genero'  # Nombre exacto de la tabla en la base de datos
    nombre = models.CharField(max_length=255)

class actores(models.Model):
    class Meta:
        db_table = 'actores'  # Nombre exacto de la tabla en la base de datos
    nombre = models.CharField(max_length=255)


class Pelicula(models.Model):
    id_pelicula = models.AutoField(primary_key=True)
    titulo = models.CharField(max_length=255)
    sinopsis = models.TextField(blank=True, null=True)
    id_genero = models.ForeignKey('genero', on_delete=models.CASCADE)
    id_actor_principal = models.ForeignKey('actor', on_delete=models.CASCADE)
    fecha_creacion = models.DateField()
    duracion = models.TimeField()
    def __str__(self):
        return self.titulo
