using System;
using System.ComponentModel.DataAnnotations;

namespace WebApplication1.Models
{
    public class Serie
    {
        public Serie()
        {
            FechaCreacion = DateTime.Now;  // Establecer la fecha actual por defecto
        }

        public int Id { get; set; }

        [Required(ErrorMessage = "El título es requerido.")]
        [StringLength(255, ErrorMessage = "El título debe tener como máximo {1} caracteres.")]
        public string Titulo { get; set; } = string.Empty;  // Asegúrate de inicializar con un valor predeterminado

        public string? Sinopsis { get; set; }  // Hacerlo anulable

        [Required(ErrorMessage = "La fecha de creación es requerida.")]
        [Display(Name = "Fecha de Creación")]
        public DateTime FechaCreacion { get; set; }  // Establecido por defecto en el constructor

        [Required(ErrorMessage = "La duración es requerida.")]
        [Range(1, int.MaxValue, ErrorMessage = "La duración debe ser mayor a cero.")]
        [Display(Name = "Duración (segundos)")]
        public int DuracionSegundos { get; set; }

        [Required(ErrorMessage = "El género es requerido.")]
        [Display(Name = "Género")]
        public string Genero { get; set; } = string.Empty;  // Asegúrate de inicializar con un valor predeterminado

        [Required(ErrorMessage = "El actor principal es requerido.")]
        [Display(Name = "Actor Principal")]
        public string ActorPrincipal { get; set; } = string.Empty;  // Asegúrate de inicializar con un valor predeterminado

        [Required(ErrorMessage = "El país de origen es requerido.")]
        [Display(Name = "País de Origen")]
        public string PaisOrigen { get; set; } = string.Empty;  // Asegúrate de inicializar con un valor predeterminado

        [Required(ErrorMessage = "La cantidad de episodios es requerida.")]
        [Range(1, int.MaxValue, ErrorMessage = "La cantidad de episodios debe ser mayor a cero.")]
        [Display(Name = "Cantidad de Episodios")]
        public int CantidadEpisodios { get; set; }

        [Required(ErrorMessage = "El idioma original es requerido.")]
        [Display(Name = "Idioma Original")]
        public string IdiomaOriginal { get; set; } = string.Empty;  // Asegúrate de inicializar con un valor predeterminado
    }
}