using System.ComponentModel.DataAnnotations;

namespace WebApplication1.Models
{
    public class ComentarioSerie
    {
        public int Id { get; set; }
        
        [Required(ErrorMessage = "El contenido del comentario es requerido.")]
        public string Contenido { get; set; }
        
        [Display(Name = "Usuario")]
        [Required(ErrorMessage = "El usuario es requerido.")]
        public int UsuarioId { get; set; }
        
        public Usuario Usuario { get; set; }
        
        [Display(Name = "Serie")]
        [Required(ErrorMessage = "La serie es requerida.")]
        public int SerieId { get; set; }
        
        public Serie Serie { get; set; }
    }
}
