from django.db import models

class GeneroPelicula(models.Model):
    id_genero_pelicula = models.AutoField(primary_key=True)
    nombre = models.CharField(max_length=100)

    def __str__(self):
        return self.nombre

class ActorPelicula(models.Model):
    id_actor_pelicula = models.AutoField(primary_key=True)
    ESTADO_CHOICES = [
        ('Activo', 'Activo'),
        ('Retirado', 'Retirado'),
        ('Muerto', 'Muerto'),
    ]

    nombre = models.CharField(max_length=255)
    fecha_inicio_carrera = models.DateField()
    estado = models.CharField(max_length=10, choices=ESTADO_CHOICES)

    def __str__(self):
        return self.nombre

class Pelicula(models.Model):
    id_pelicula = models.AutoField(primary_key=True)
    titulo = models.CharField(max_length=255)
    sinopsis = models.TextField(null=True, blank=True)
    genero = models.ForeignKey(GeneroPelicula, on_delete=models.CASCADE)
    actor_principal = models.ForeignKey(ActorPelicula, on_delete=models.CASCADE)

    def __str__(self):
        return self.titulo
