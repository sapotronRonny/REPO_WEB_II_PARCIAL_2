import graphene
from graphene_django import DjangoObjectType
from django.db.models import Count
from .models import GeneroPelicula, ActorPelicula, Pelicula

class GeneroPeliculaType(DjangoObjectType):
    class Meta:
        model = GeneroPelicula
        fields = "__all__"

class ActorPeliculaType(DjangoObjectType):
    class Meta:
        model = ActorPelicula
        fields = "__all__"

class PeliculaType(DjangoObjectType):
    class Meta:
        model = Pelicula
        fields = "__all__"

class ReportePeliculasPorGenero(graphene.ObjectType):
    genero = graphene.String()
    cantidad = graphene.Int()

class Query(graphene.ObjectType):
    all_generos = graphene.List(GeneroPeliculaType)
    all_actores = graphene.List(ActorPeliculaType)
    all_peliculas = graphene.List(PeliculaType)
    reporte_peliculas_por_genero = graphene.List(ReportePeliculasPorGenero)

    def resolve_all_generos(root, info):
        return GeneroPelicula.objects.all()

    def resolve_all_actores(root, info):
        return ActorPelicula.objects.all()

    def resolve_all_peliculas(root, info):
        return Pelicula.objects.all()

    def resolve_reporte_peliculas_por_genero(root, info, **kwargs):
        result = Pelicula.objects.values('genero__nombre').annotate(cantidad=Count('id_pelicula'))
        reporte = [
            {
                'genero': r['genero__nombre'],
                'cantidad': r['cantidad']
            }
            for r in result
        ]
        return reporte

schema = graphene.Schema(query=Query)