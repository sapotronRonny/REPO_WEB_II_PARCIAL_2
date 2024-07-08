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
    public class ComentariosSerieController : ControllerBase
    {
        private readonly ApplicationDbContext _context;

        public ComentariosSerieController(ApplicationDbContext context)
        {
            _context = context;
        }

        [HttpGet]
        public IActionResult GetComentariosSerie()
        {
            var comentarios = _context.ComentariosSerie
                .Include(c => c.Usuario)
                .Include(c => c.Serie)
                .ToList();
            
            return Ok(comentarios);
        }

        [HttpGet("{id}")]
        public IActionResult GetComentarioSerie(int id)
        {
            var comentario = _context.ComentariosSerie
                .Include(c => c.Usuario)
                .Include(c => c.Serie)
                .FirstOrDefault(c => c.Id == id);

            if (comentario == null)
            {
                return NotFound();
            }
            
            return Ok(comentario);
        }

        [HttpPost]
        public IActionResult PostComentarioSerie([FromBody] ComentarioSerie comentario)
        {
            if (ModelState.IsValid)
            {
                _context.ComentariosSerie.Add(comentario);
                _context.SaveChanges();
                return CreatedAtAction(nameof(GetComentarioSerie), new { id = comentario.Id }, comentario);
            }
            
            return BadRequest(ModelState);
        }

        [HttpPut("{id}")]
        public IActionResult PutComentarioSerie(int id, [FromBody] ComentarioSerie comentario)
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
        public IActionResult DeleteComentarioSerie(int id)
        {
            var comentario = _context.ComentariosSerie.Find(id);
            if (comentario == null)
            {
                return NotFound();
            }

            _context.ComentariosSerie.Remove(comentario);
            _context.SaveChanges();
            return NoContent();
        }
    }
}
