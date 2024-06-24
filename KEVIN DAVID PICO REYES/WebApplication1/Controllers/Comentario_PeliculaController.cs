using Microsoft.AspNetCore.Mvc;
using WebApplication1.Models;
using WebApplication1.Data;
using Microsoft.AspNetCore.Authorization;
using Microsoft.EntityFrameworkCore;
using System.Linq;

namespace WebApplication1.Controllers
{
    [Route("api/[controller]")]
    [ApiController]
    [Authorize]
    public class ComentariosPeliculaController : ControllerBase
    {
        private readonly ApplicationDbContext _context;

        public ComentariosPeliculaController(ApplicationDbContext context)
        {
            _context = context;
        }

        [HttpGet]
        public IActionResult GetComentariosPelicula()
        {
            var comentarios = _context.ComentariosPelicula
                .Include(c => c.Usuario)
                .Include(c => c.Pelicula)
                .ToList();
            
            return Ok(comentarios);
        }

        [HttpGet("{id}")]
        public IActionResult GetComentarioPelicula(int id)
        {
            var comentario = _context.ComentariosPelicula
                .Include(c => c.Usuario)
                .Include(c => c.Pelicula)
                .FirstOrDefault(c => c.Id == id);

            if (comentario == null)
            {
                return NotFound();
            }
            
            return Ok(comentario);
        }

        [HttpPost]
        public IActionResult PostComentarioPelicula([FromBody] ComentarioPelicula comentario)
        {
            if (ModelState.IsValid)
            {
                _context.ComentariosPelicula.Add(comentario);
                _context.SaveChanges();
                return CreatedAtAction(nameof(GetComentarioPelicula), new { id = comentario.Id }, comentario);
            }
            
            return BadRequest(ModelState);
        }

        [HttpPut("{id}")]
        public IActionResult PutComentarioPelicula(int id, [FromBody] ComentarioPelicula comentario)
        {
            if (id != comentario.Id)
            {
                return BadRequest();
            }

            if (ModelState.IsValid)
            {
                _context.Entry(comentario).State = EntityState.Modified;
                _context.SaveChanges();
                return NoContent();
            }
            
            return BadRequest(ModelState);
        }

        [HttpDelete("{id}")]
        public IActionResult DeleteComentarioPelicula(int id)
        {
            var comentario = _context.ComentariosPelicula.Find(id);
            if (comentario == null)
            {
                return NotFound();
            }

            _context.ComentariosPelicula.Remove(comentario);
            _context.SaveChanges();
            return NoContent();
        }
    }
}
