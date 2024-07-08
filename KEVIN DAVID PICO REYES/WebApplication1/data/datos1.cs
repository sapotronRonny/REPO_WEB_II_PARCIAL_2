using Microsoft.EntityFrameworkCore;
using WebApplication1.Models;

namespace WebApplication1.Data
{
    public class ApplicationDbContext : DbContext
    {
        public ApplicationDbContext(DbContextOptions<ApplicationDbContext> options)
            : base(options)
        {
        }

        public DbSet<Usuario> Usuarios { get; set; }
        public DbSet<ComentarioPelicula> ComentariosPelicula { get; set; }
        public DbSet<ComentarioSerie> ComentariosSerie { get; set; }

        // Otros DbSet para otras entidades si las tienes
        public DbSet<Pelicula> Peliculas { get; set; }
        public DbSet<Serie> Series { get; set; }

        protected override void OnModelCreating(ModelBuilder modelBuilder)
        {
            // Configuraciones adicionales de modelos si es necesario
        }
    }
}
