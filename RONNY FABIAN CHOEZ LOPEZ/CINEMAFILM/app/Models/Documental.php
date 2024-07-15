<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Documental extends Model
{
    use HasFactory;

    protected $fillable = ['nombre'];

     // Especificar la clave primaria
     protected $primaryKey = 'id_documental';

     // Indicar que no use las convenciones de incrementación automática
     public $incrementing = true;
 
     // Indicar que la clave primaria no es de tipo entero
     protected $keyType = 'int';
}
