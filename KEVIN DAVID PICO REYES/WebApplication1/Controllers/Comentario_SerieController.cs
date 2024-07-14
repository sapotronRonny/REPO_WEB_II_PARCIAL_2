using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;
using System.Linq;
using System.Threading.Tasks;
using WebApplication1.Data;
using WebApplication1.Models;

namespace WebApplication1.Controllers
{
    [ApiController]
    [Route("api/[controller]")]
    public class ComentarioSerieController : ControllerBase
    {
        private readonly ApplicationDbContext _context;

        public ComentarioSerieController(ApplicationDbContext context)
        {
            _context = context;
        }

        [HttpGet]
        public async Task<IActionResult> Get()
        {
            var comentarios = await _context.ComentariosSeries.ToListAsync();
            return Ok(comentarios);
        }

        [HttpGet("{id}")]
        public async Task<IActionResult> Get(int id)
        {
            var comentario = await _context.ComentariosSeries.FindAsync(id);
            if (comentario == null)
            {
                return NotFound();
            }
            return Ok(comentario);
        }

        [HttpPost]
        public async Task<IActionResult> Post([FromBody] ComentarioSerie comentario)
        {
            if (ModelState.IsValid)
            {
                _context.ComentariosSeries.Add(comentario);
                await _context.SaveChangesAsync();
                return CreatedAtAction(nameof(Get), new { id = comentario.Id }, comentario);
            }
            return BadRequest(ModelState);
        }

        [HttpPut("{id}")]
        public async Task<IActionResult> Put(int id, [FromBody] ComentarioSerie comentario)
        {
            if (id != comentario.Id)
            {
                return BadRequest();
            }

            _context.Entry(comentario).State = EntityState.Modified;

            try
            {
                await _context.SaveChangesAsync();
            }
            catch (DbUpdateConcurrencyException)
            {
                if (!_context.ComentariosSeries.Any(e => e.Id == id))
                {
                    return NotFound();
                }
                else
                {
                    throw;
                }
            }

            return NoContent();
        }

        [HttpDelete("{id}")]
        public async Task<IActionResult> Delete(int id)
        {
            var comentario = await _context.ComentariosSeries.FindAsync(id);
            if (comentario == null)
            {
                return NotFound();
            }

            _context.ComentariosSeries.Remove(comentario);
            await _context.SaveChangesAsync();

            return NoContent();
        }
    }
}
