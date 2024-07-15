<?php

namespace App\Http\Controllers;

use App\Models\Documental;
use Illuminate\Http\Request;

class DocumentalController extends Controller
{
    public function index()
    {
        return Documental::all();
    }

    public function store(Request $request)
    {
        $request->validate([
            'titulo' => 'required|string|max:100',
            'id_genero' => 'required|exists:genero_docus,id_genero',
            'id_director' => 'required|exists:directors,id_director',
            'duracion' => 'nullable|integer',
            'resumen' => 'nullable|string',
        ]);

        $documental = new Documental();
        $documental->titulo = $request->titulo;
        $documental->id_genero = $request->id_genero;
        $documental->id_director = $request->id_director;
        $documental->duracion = $request->duracion;
        $documental->resumen = $request->resumen;
        $documental->save();

        return $documental;
    }

    public function show($id_documental)
    {
        return Documental::findOrFail($id_documental);
    }

    public function update(Request $request, $id_documental)
    {
        $request->validate([
            'titulo' => 'required|string|max:100',
            'id_genero' => 'required|exists:genero_docus,id_genero',
            'id_director' => 'required|exists:directors,id_director',
            'duracion' => 'nullable|integer',
            'resumen' => 'nullable|string',
        ]);

        $documental = Documental::findOrFail($id_documental);
        $documental->titulo = $request->titulo;
        $documental->id_genero = $request->id_genero;
        $documental->id_director = $request->id_director;
        $documental->duracion = $request->duracion;
        $documental->resumen = $request->resumen;
        $documental->save();

        return $documental;
    }

    public function destroy($id_documental)
    {
        $documental = Documental::findOrFail($id_documental);
        $documental->delete();

        return response()->noContent();
    }
}