using System;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;

namespace WebApplication1.Models
{
    public class ComentarioPelicula
    {
        public ComentarioPelicula()
        {
            FechaComentario = DateTime.Now; // Establecer la fecha actual como valor predeterminado
        }

        public int Id { get; set; }

        [Required(ErrorMessage = "El contenido del comentario es requerido.")]
        public string Contenido { get; set; } = string.Empty;

        [Display(Name = "Usuario")]
        [Required(ErrorMessage = "El usuario es requerido.")]
        public int UsuarioId { get; set; }

        [ForeignKey("UsuarioId")]
        public Usuario Usuario { get; set; } = new Usuario();

        [Display(Name = "Película")]
        [Required(ErrorMessage = "La película es requerida.")]
        public int PeliculaId { get; set; }

        [ForeignKey("PeliculaId")]
        public Pelicula Pelicula { get; set; } = new Pelicula();

        [Required(ErrorMessage = "La fecha del comentario es requerida.")]
        [Display(Name = "Fecha del Comentario")]
        public DateTime FechaComentario { get; set; }
    }
}
