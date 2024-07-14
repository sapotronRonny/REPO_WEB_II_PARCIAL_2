﻿// <auto-generated />
using System;
using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Infrastructure;
using Microsoft.EntityFrameworkCore.Storage.ValueConversion;
using Npgsql.EntityFrameworkCore.PostgreSQL.Metadata;
using WebApplication1.Data;

#nullable disable

namespace WebApplication1.Migrations
{
    [DbContext(typeof(ApplicationDbContext))]
    partial class ApplicationDbContextModelSnapshot : ModelSnapshot
    {
        protected override void BuildModel(ModelBuilder modelBuilder)
        {
#pragma warning disable 612, 618
            modelBuilder
                .HasAnnotation("ProductVersion", "8.0.5")
                .HasAnnotation("Relational:MaxIdentifierLength", 63);

            NpgsqlModelBuilderExtensions.UseIdentityByDefaultColumns(modelBuilder);

            modelBuilder.Entity("WebApplication1.Models.ComentarioPelicula", b =>
                {
                    b.Property<int>("Id")
                        .ValueGeneratedOnAdd()
                        .HasColumnType("integer");

                    NpgsqlPropertyBuilderExtensions.UseIdentityByDefaultColumn(b.Property<int>("Id"));

                    b.Property<string>("Contenido")
                        .IsRequired()
                        .HasColumnType("text");

                    b.Property<DateTime>("FechaComentario")
                        .HasColumnType("timestamp with time zone");

                    b.Property<int>("PeliculaId")
                        .HasColumnType("integer");

                    b.Property<int>("UsuarioId")
                        .HasColumnType("integer");

                    b.HasKey("Id");

                    b.HasIndex("PeliculaId");

                    b.HasIndex("UsuarioId");

                    b.ToTable("ComentariosPeliculas");
                });

            modelBuilder.Entity("WebApplication1.Models.ComentarioSerie", b =>
                {
                    b.Property<int>("Id")
                        .ValueGeneratedOnAdd()
                        .HasColumnType("integer");

                    NpgsqlPropertyBuilderExtensions.UseIdentityByDefaultColumn(b.Property<int>("Id"));

                    b.Property<string>("Contenido")
                        .IsRequired()
                        .HasColumnType("text");

                    b.Property<DateTime>("FechaComentario")
                        .HasColumnType("timestamp with time zone");

                    b.Property<int>("SerieId")
                        .HasColumnType("integer");

                    b.Property<int>("UsuarioId")
                        .HasColumnType("integer");

                    b.HasKey("Id");

                    b.HasIndex("SerieId");

                    b.HasIndex("UsuarioId");

                    b.ToTable("ComentariosSeries");
                });

            modelBuilder.Entity("WebApplication1.Models.Pelicula", b =>
                {
                    b.Property<int>("Id")
                        .ValueGeneratedOnAdd()
                        .HasColumnType("integer");

                    NpgsqlPropertyBuilderExtensions.UseIdentityByDefaultColumn(b.Property<int>("Id"));

                    b.Property<string>("ActorPrincipal")
                        .IsRequired()
                        .HasColumnType("text");

                    b.Property<int>("DuracionSegundos")
                        .HasColumnType("integer");

                    b.Property<DateTime>("FechaCreacion")
                        .HasColumnType("timestamp with time zone");

                    b.Property<DateTime>("FechaEstreno")
                        .HasColumnType("timestamp with time zone");

                    b.Property<string>("Genero")
                        .IsRequired()
                        .HasColumnType("text");

                    b.Property<string>("PaisOrigen")
                        .IsRequired()
                        .HasColumnType("text");

                    b.Property<string>("Sinopsis")
                        .HasColumnType("text");

                    b.Property<string>("Titulo")
                        .IsRequired()
                        .HasMaxLength(255)
                        .HasColumnType("character varying(255)");

                    b.HasKey("Id");

                    b.ToTable("Peliculas");
                });

            modelBuilder.Entity("WebApplication1.Models.Serie", b =>
                {
                    b.Property<int>("Id")
                        .ValueGeneratedOnAdd()
                        .HasColumnType("integer");

                    NpgsqlPropertyBuilderExtensions.UseIdentityByDefaultColumn(b.Property<int>("Id"));

                    b.Property<string>("ActorPrincipal")
                        .IsRequired()
                        .HasColumnType("text");

                    b.Property<int>("CantidadEpisodios")
                        .HasColumnType("integer");

                    b.Property<int>("DuracionSegundos")
                        .HasColumnType("integer");

                    b.Property<DateTime>("FechaCreacion")
                        .HasColumnType("timestamp with time zone");

                    b.Property<string>("Genero")
                        .IsRequired()
                        .HasColumnType("text");

                    b.Property<string>("IdiomaOriginal")
                        .IsRequired()
                        .HasColumnType("text");

                    b.Property<string>("PaisOrigen")
                        .IsRequired()
                        .HasColumnType("text");

                    b.Property<string>("Sinopsis")
                        .HasColumnType("text");

                    b.Property<string>("Titulo")
                        .IsRequired()
                        .HasMaxLength(255)
                        .HasColumnType("character varying(255)");

                    b.HasKey("Id");

                    b.ToTable("Series");
                });

            modelBuilder.Entity("WebApplication1.Models.Usuario", b =>
                {
                    b.Property<int>("Id")
                        .ValueGeneratedOnAdd()
                        .HasColumnType("integer");

                    NpgsqlPropertyBuilderExtensions.UseIdentityByDefaultColumn(b.Property<int>("Id"));

                    b.Property<string>("Contraseña")
                        .IsRequired()
                        .HasMaxLength(255)
                        .HasColumnType("character varying(255)");

                    b.Property<string>("Email")
                        .IsRequired()
                        .HasMaxLength(255)
                        .HasColumnType("character varying(255)");

                    b.Property<string>("Nombre")
                        .IsRequired()
                        .HasMaxLength(255)
                        .HasColumnType("character varying(255)");

                    b.HasKey("Id");

                    b.ToTable("Usuarios");
                });

            modelBuilder.Entity("WebApplication1.Models.ComentarioPelicula", b =>
                {
                    b.HasOne("WebApplication1.Models.Pelicula", "Pelicula")
                        .WithMany()
                        .HasForeignKey("PeliculaId")
                        .OnDelete(DeleteBehavior.Cascade)
                        .IsRequired();

                    b.HasOne("WebApplication1.Models.Usuario", "Usuario")
                        .WithMany("ComentariosPelicula")
                        .HasForeignKey("UsuarioId")
                        .OnDelete(DeleteBehavior.Cascade)
                        .IsRequired();

                    b.Navigation("Pelicula");

                    b.Navigation("Usuario");
                });

            modelBuilder.Entity("WebApplication1.Models.ComentarioSerie", b =>
                {
                    b.HasOne("WebApplication1.Models.Serie", "Serie")
                        .WithMany()
                        .HasForeignKey("SerieId")
                        .OnDelete(DeleteBehavior.Cascade)
                        .IsRequired();

                    b.HasOne("WebApplication1.Models.Usuario", "Usuario")
                        .WithMany("ComentariosSerie")
                        .HasForeignKey("UsuarioId")
                        .OnDelete(DeleteBehavior.Cascade)
                        .IsRequired();

                    b.Navigation("Serie");

                    b.Navigation("Usuario");
                });

            modelBuilder.Entity("WebApplication1.Models.Usuario", b =>
                {
                    b.Navigation("ComentariosPelicula");

                    b.Navigation("ComentariosSerie");
                });
#pragma warning restore 612, 618
        }
    }
}
