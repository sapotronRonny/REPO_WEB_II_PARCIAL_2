using System.ComponentModel.DataAnnotations;

namespace WebApplication1.Models
{
    public class ComentarioPelicula
    {
        public int Id { get; set; }
        
        [Required(ErrorMessage = "El contenido del comentario es requerido.")]
        public string Contenido { get; set; }
        
        [Display(Name = "Usuario")]
        [Required(ErrorMessage = "El usuario es requerido.")]
        public int UsuarioId { get; set; }
        
        public Usuario Usuario { get; set; }
        
        [Display(Name = "Película")]
        [Required(ErrorMessage = "La película es requerida.")]
        public int PeliculaId { get; set; }
        
        public Pelicula Pelicula { get; set; }
    }
}
