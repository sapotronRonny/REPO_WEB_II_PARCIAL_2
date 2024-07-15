<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Documental extends Model
{
    use HasFactory;

    protected $fillable = ['titulo', 'id_genero', 'id_director', 'duracion', 'resumen'];

     // Especificar la clave primaria
     protected $primaryKey = 'id_documental';

     // Indicar que no use las convenciones de incrementación automática
     public $incrementing = true;
 
     // Indicar que la clave primaria no es de tipo entero
     protected $keyType = 'int';

     public function genero()
    {
        return $this->belongsTo(GeneroDocu::class, 'genero_id');
    }

    public function director()
    {
        return $this->belongsTo(Director::class, 'director_id');
    }
}
