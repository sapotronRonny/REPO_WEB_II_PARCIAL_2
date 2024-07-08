using System;
using System.ComponentModel.DataAnnotations;

namespace WebApplication1.Models
{
    public class ComentarioSerie
    {
        public ComentarioSerie()
        {
            FechaComentario = DateTime.Now;  // Establecer la fecha actual por defecto
        }

        public int Id { get; set; }

        [Required(ErrorMessage = "El contenido del comentario es requerido.")]
        public string Contenido { get; set; } = string.Empty;

        [Display(Name = "Usuario")]
        [Required(ErrorMessage = "El usuario es requerido.")]
        public int UsuarioId { get; set; }

        public Usuario Usuario { get; set; } = new Usuario();

        [Display(Name = "Serie")]
        [Required(ErrorMessage = "La serie es requerida.")]
        public int SerieId { get; set; }

        public Serie Serie { get; set; } = new Serie();

        [Required(ErrorMessage = "La fecha del comentario es requerida.")]
        [Display(Name = "Fecha del Comentario")]
        public DateTime FechaComentario { get; set; }  // Establecido por defecto en el constructor
    }
}