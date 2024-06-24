using System.ComponentModel.DataAnnotations;

namespace WebApplication1.Models
{
    public class Usuario
    {
        public int Id { get; set; }
        
        [Required(ErrorMessage = "El nombre es requerido.")]
        [StringLength(255, ErrorMessage = "El nombre debe tener como máximo {1} caracteres.")]
        public string Nombre { get; set; }
        
        [Required(ErrorMessage = "El email es requerido.")]
        [EmailAddress(ErrorMessage = "El email no es válido.")]
        [StringLength(255, ErrorMessage = "El email debe tener como máximo {1} caracteres.")]
        public string Email { get; set; }
        
        [Required(ErrorMessage = "La contraseña es requerida.")]
        [StringLength(255, ErrorMessage = "La contraseña debe tener como máximo {1} caracteres.")]
        public string Contraseña { get; set; }
    }
}
