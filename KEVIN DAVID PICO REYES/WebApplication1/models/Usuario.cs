using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;

namespace WebApplication1.Models
{
    public class Usuario
    {
        public int Id { get; set; }

        [Required(ErrorMessage = "El nombre es requerido.")]
        [StringLength(255, ErrorMessage = "El nombre debe tener como máximo {1} caracteres.")]
        public string Nombre { get; set; } = string.Empty;

        [Required(ErrorMessage = "El email es requerido.")]
        [EmailAddress(ErrorMessage = "El email no es válido.")]
        [StringLength(255, ErrorMessage = "El email debe tener como máximo {1} caracteres.")]
        public string Email { get; set; } = string.Empty;

        [Required(ErrorMessage = "La contraseña es requerida.")]
        [StringLength(255, ErrorMessage = "La contraseña debe tener como máximo {1} caracteres.")]
        public string Contraseña { get; set; } = string.Empty;

        // Relación uno a muchos con ComentarioPelicula
        public ICollection<ComentarioPelicula> ComentariosPelicula { get; set; } = new List<ComentarioPelicula>();

        // Relación uno a muchos con ComentarioSerie
        public ICollection<ComentarioSerie> ComentariosSerie { get; set; } = new List<ComentarioSerie>();
    }
}
