using System;
using System.ComponentModel.DataAnnotations;

namespace WebApplication1.Models
{
    public class Pelicula
    {
        public Pelicula()
        {
            FechaCreacion = DateTime.Now; // Establecer la fecha actual como valor predeterminado
        }

        public int Id { get; set; }

        [Required(ErrorMessage = "El título es requerido.")]
        [StringLength(255, ErrorMessage = "El título debe tener como máximo {1} caracteres.")]
        public string Titulo { get; set; } = string.Empty;

        public string? Sinopsis { get; set; }  // Hacerlo anulable

        [Required(ErrorMessage = "La fecha de creación es requerida.")]
        [Display(Name = "Fecha de Creación")]
        public DateTime FechaCreacion { get; set; }  // Inicializar con la fecha actual

        [Required(ErrorMessage = "El género es requerido.")]
        [Display(Name = "Género")]
        public string Genero { get; set; } = string.Empty;  // Asegúrate de inicializar con un valor predeterminado

        [Required(ErrorMessage = "El actor principal es requerido.")]
        [Display(Name = "Actor Principal")]
        public string ActorPrincipal { get; set; } = string.Empty;  // Asegúrate de inicializar con un valor predeterminado

        [Required(ErrorMessage = "El país de origen es requerido.")]
        [Display(Name = "País de Origen")]
        public string PaisOrigen { get; set; } = string.Empty;  // Asegúrate de inicializar con un valor predeterminado

        [Required(ErrorMessage = "La duración en segundos es requerida.")]
        [Display(Name = "Duración (segundos)")]
        [Range(1, int.MaxValue, ErrorMessage = "La duración debe ser mayor a cero.")]
        public int DuracionSegundos { get; set; }

        [Required(ErrorMessage = "La fecha de estreno es requerida.")]
        [Display(Name = "Fecha de Estreno")]
        public DateTime FechaEstreno { get; set; }
    }
}
