using System.ComponentModel.DataAnnotations;

namespace WebApplication1.Models
{
    public class Serie
    {
        public int Id { get; set; }
        
        [Required(ErrorMessage = "El título es requerido.")]
        [StringLength(255, ErrorMessage = "El título debe tener como máximo {1} caracteres.")]
        public string Titulo { get; set; }
        
        public string Sinopsis { get; set; }
        
        [Display(Name = "Género")]
        public int IdGenero { get; set; }
        
        [Display(Name = "Actor Principal")]
        public int IdActorPrincipal { get; set; }
    }
}
